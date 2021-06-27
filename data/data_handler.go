package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cass-dlcm/splatstatsuploader/enums"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cass-dlcm/splatstatsuploader/iksm"
	"github.com/cass-dlcm/splatstatsuploader/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Monitor monitors JSON for changes/new battles/shifts and uploads them.
func Monitor(m int, s bool, salmon bool, apiKey string, version string, appHead http.Header, client *http.Client) {
	if salmon {
		GetSplatnetSalmon(s, apiKey, version, appHead, client)
	} else {
		GetSplatnetBattle(s, apiKey, version, appHead, client)
	}

	for {
		timer := time.NewTimer(time.Duration(m) * time.Second)
		<-timer.C

		if salmon {
			uploadLatestSalmon(s, apiKey, appHead, client)
		} else {
			uploadLatestBattle(s, apiKey, appHead, client)
		}
	}
}

func uploadLatestBattle(s bool, apiKey string, appHead http.Header, client *http.Client) {
	if _, err := fmt.Println("Pulling data from online..."); err != nil { // grab data from SplatNet 2
		panic(err)
	}

	url := "https://app.splatoon2.nintendo.net/api/results"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = appHead

	req.AddCookie(&http.Cookie{Name: "iksm_session", Value: viper.GetString("cookie")})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	var data types.BattleList
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	uploadSingleBattle(s, apiKey, appHead, data.Results[0].BattleNumber, client)
}

func uploadSingleBattle(s bool, apiKey string, appHead http.Header, battleNumber string, client *http.Client) {
	url := "https://app.splatoon2.nintendo.net/api/results/" + battleNumber
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = appHead

	req.AddCookie(&http.Cookie{Name: "iksm_session", Value: viper.GetString("cookie")})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var battle types.BattleSplatnet
	if err := json.NewDecoder(resp.Body).Decode(&battle); err != nil {
		panic(errors.Wrap(err, resp.Body.Close().Error()))
	}

	if err := resp.Body.Close(); err != nil {
		panic(err)
	}

	battleUpload := setBattlePayload(&battle)
	UploadBattle(&battleUpload, apiKey, client)

	if s {
		file, err := json.MarshalIndent(battle, "", " ")
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile("two_battle/"+battleNumber+".json", file, 0600); err != nil {
			panic(err)
		}
	}
}

func uploadSingleSalmon(s bool, apiKey string, appHead http.Header, jobId int64, client *http.Client) {
	url := "https://app.splatoon2.nintendo.net/api/coop_results/" + fmt.Sprint(jobId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = appHead

	req.AddCookie(&http.Cookie{Name: "iksm_session", Value: viper.GetString("cookie")})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var shift types.ShiftSplatnet

	if s {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(errors.Wrap(err, resp.Body.Close().Error()))
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile("two_salmon/"+fmt.Sprint(jobId)+".json", bodyBytes, 0600); err != nil {
			panic(err)
		}

		if err := json.Unmarshal(bodyBytes, &shift); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&shift); err != nil {
			panic(errors.Wrap(err, resp.Body.Close().Error()))
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}

	shiftUpload := setSalmonPayload(&shift)
	UploadSalmon(&shiftUpload, apiKey, client)
}

func uploadLatestSalmon(s bool, apiKey string, appHead http.Header, client *http.Client) {
	if _, err := fmt.Println("Pulling Salmon Run data from online..."); err != nil {
		panic(err)
	}

	url := "https://app.splatoon2.nintendo.net/api/coop_results"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = appHead

	req.AddCookie(&http.Cookie{Name: "iksm_session", Value: viper.GetString("cookie")})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	var data types.ShiftList

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	uploadSingleSalmon(s, apiKey, appHead, data.Results[0].JobId, client)
}

// File looks in a directory and uploads all the result data found in files in the directory.
func File(salmon bool, apiKey string, client *http.Client) {
	var files []fs.FileInfo

	var err error

	if salmon {
		files, err = ioutil.ReadDir("./two_salmon/")
		if err != nil {
			panic(err)
		}
	} else {
		files, err = ioutil.ReadDir("./two_battle/")
		if err != nil {
			panic(err)
		}
	}

	for _, file := range files {
		var jsonFile *os.File

		var shift types.ShiftSplatnet

		var battle types.BattleSplatnet

		if salmon {
			jsonFile, err = os.Open("./two_salmon/" + file.Name())
		} else {
			jsonFile, err = os.Open("./two_battle/" + file.Name())
		}

		if err != nil {
			panic(err)
		}

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			panic(errors.Wrap(err, jsonFile.Close().Error()))
		}

		if err := jsonFile.Close(); err != nil {
			panic(err)
		}

		if salmon {
			if err := json.Unmarshal(byteValue, &shift); err != nil {
				panic(err)
			}

			shiftUpload := setSalmonPayload(&shift)
			UploadSalmon(&shiftUpload, apiKey, client)
		} else {
			if err := json.Unmarshal(byteValue, &battle); err != nil {
				panic(err)
			}

			battleUpload := setBattlePayload(&battle)
			UploadBattle(&battleUpload, apiKey, client)
		}
	}
}

// GetSplatnetBattle retrieves the battles from SplatNet and uploads them all.
func GetSplatnetBattle(s bool, apiKey string, version string, appHead http.Header, client *http.Client) {
	if _, err := fmt.Println("Pulling data from online..."); err != nil { // grab data from SplatNet 2
		panic(err)
	}

	url := "https://app.splatoon2.nintendo.net/api/results"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = appHead

	req.AddCookie(&http.Cookie{Name: "iksm_session", Value: viper.GetString("cookie")})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	var data types.BattleList
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	if data.Code != nil {
		iksm.GenNewCookie("auth", version, client)
		GetSplatnetBattle(s, apiKey, version, appHead, client)

		return
	}

	for i := range data.Results {
		uploadSingleBattle(s, apiKey, appHead, data.Results[i].BattleNumber, client)
	}
}

// GetSplatnetSalmon retrieves the shifts from SplatNet and uploads them all.
func GetSplatnetSalmon(s bool, apiKey string, version string, appHead http.Header, client *http.Client) {
	if viper.GetString("cookie") == "" {
		iksm.GenNewCookie("blank", version, client)
	}

	if _, err := fmt.Println("Pulling Salmon Run data from online..."); err != nil {
		panic(err)
	}

	url := "https://app.splatoon2.nintendo.net/api/coop_results"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = appHead

	req.AddCookie(&http.Cookie{Name: "iksm_session", Value: viper.GetString("cookie")})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	var data types.ShiftList
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(errors.Wrap(err, resp.Body.Close().Error()))
	}

	if err := resp.Body.Close(); err != nil {
		panic(err)
	}

	if data.Code != nil {
		iksm.GenNewCookie("auth", version, client)
		GetSplatnetSalmon(s, apiKey, version, appHead, client)

		return
	}

	for i := range data.Results {
		uploadSingleSalmon(s, apiKey, appHead, data.Results[i].JobId, client)
	}
}

func setSalmonPayload(shift *types.ShiftSplatnet) types.Shift {
	shiftUpload := types.Shift{}
	shiftUpload.SplatnetJson = shift
	shiftUpload.SplatnetUpload = true
	shiftUpload.StatInkUpload = false
	shiftUpload.DangerRate = (*shift).DangerRate
	shiftUpload.DrizzlerCount = (*shift).BossCounts.Drizzler.Count
	shiftUpload.FailureWave = (*shift).JobResult.FailureWave
	shiftUpload.JobScore = (*shift).JobScore
	shiftUpload.FlyfishCount = (*shift).BossCounts.Flyfish.Count
	shiftUpload.GoldieCount = (*shift).BossCounts.Goldie.Count
	shiftUpload.GradePoint = (*shift).GradePoint
	shiftUpload.GradePointDelta = (*shift).GradePointDelta
	shiftUpload.GrillerCount = (*shift).BossCounts.Griller.Count
	shiftUpload.IsClear = (*shift).JobResult.IsClear
	shiftUpload.JobFailureReason = (*shift).JobResult.FailureReason
	shiftUpload.JobId = (*shift).JobId
	shiftUpload.MawsCount = (*shift).BossCounts.Maws.Count
	shiftUpload.PlayerDeathCount = (*shift).MyResult.DeadCount
	shiftUpload.PlayerDrizzlerKills = (*shift).MyResult.BossKillCounts.Drizzler.Count
	shiftUpload.PlayerFlyfishKills = (*shift).MyResult.BossKillCounts.Flyfish.Count
	shiftUpload.PlayerGender = (*shift).MyResult.PlayerType.Gender
	shiftUpload.PlayerGoldenEggs = (*shift).MyResult.GoldenEggs
	shiftUpload.PlayerGoldieKills = (*shift).MyResult.BossKillCounts.Goldie.Count
	shiftUpload.PlayerGrillerKills = (*shift).MyResult.BossKillCounts.Griller.Count
	shiftUpload.PlayerSplatnetId = (*shift).MyResult.Pid
	shiftUpload.PlayerMawsKills = (*shift).MyResult.BossKillCounts.Maws.Count
	shiftUpload.PlayerName = (*shift).MyResult.Name
	shiftUpload.PlayerPowerEggs = (*shift).MyResult.PowerEggs
	shiftUpload.PlayerReviveCount = (*shift).MyResult.HelpCount
	shiftUpload.PlayerScrapperKills = (*shift).MyResult.BossKillCounts.Scrapper.Count
	shiftUpload.PlayerSpecial = (*shift).MyResult.Special.Id
	shiftUpload.PlayerSpecies = (*shift).MyResult.PlayerType.Species
	shiftUpload.PlayerSteelEelKills = (*shift).MyResult.BossKillCounts.SteelEel.Count
	shiftUpload.PlayerSteelheadKills = (*shift).MyResult.BossKillCounts.Steelhead.Count
	shiftUpload.PlayerStingerKills = (*shift).MyResult.BossKillCounts.Stinger.Count
	shiftUpload.PlayerTitle = (*shift).Grade.Id
	salmonPlayerWeaponSpecials(shift, &shiftUpload)
	shiftSetTimes(shift, &shiftUpload)
	if (*shift).Schedule.Weapons[0].CoopSpecialWeapon != nil {
		shiftUpload.ScheduleWeapon0 = (*enums.SalmonWeaponScheduleEnum)(&(*shift).Schedule.Weapons[0].CoopSpecialWeapon.Name)
	} else {
		shiftUpload.ScheduleWeapon0 = &(*shift).Schedule.Weapons[0].Weapon.Name
	}
	if (*shift).Schedule.Weapons[1].CoopSpecialWeapon != nil {
		shiftUpload.ScheduleWeapon1 = (*enums.SalmonWeaponScheduleEnum)(&(*shift).Schedule.Weapons[1].CoopSpecialWeapon.Name)
	} else {
		shiftUpload.ScheduleWeapon1 = &(*shift).Schedule.Weapons[1].Weapon.Name
	}
	if (*shift).Schedule.Weapons[2].CoopSpecialWeapon != nil {
		shiftUpload.ScheduleWeapon2 = (*enums.SalmonWeaponScheduleEnum)(&(*shift).Schedule.Weapons[2].CoopSpecialWeapon.Name)
	} else {
		shiftUpload.ScheduleWeapon2 = &(*shift).Schedule.Weapons[2].Weapon.Name
	}
	if (*shift).Schedule.Weapons[3].CoopSpecialWeapon != nil {
		shiftUpload.ScheduleWeapon3 = (*enums.SalmonWeaponScheduleEnum)(&(*shift).Schedule.Weapons[3].CoopSpecialWeapon.Name)
	} else {
		shiftUpload.ScheduleWeapon3 = &(*shift).Schedule.Weapons[3].Weapon.Name
	}
	shiftUpload.ScrapperCount = (*shift).BossCounts.Scrapper.Count
	shiftUpload.Stage = (*shift).Schedule.Stage.Name
	shiftUpload.SteelEelCount = (*shift).BossCounts.SteelEel.Count
	shiftUpload.SteelheadCount = (*shift).BossCounts.Steelhead.Count
	shiftUpload.StingerCount = (*shift).BossCounts.Stinger.Count
	shiftSetTeammate0(shift, &shiftUpload)
	shiftSetTeammate1(shift, &shiftUpload)
	shiftSetTeammate2(shift, &shiftUpload)
	shiftUpload.Wave1EventType = (*shift).WaveDetails[0].EventType.Key
	shiftUpload.Wave1GoldenAppear = (*shift).WaveDetails[0].GoldenAppear
	shiftUpload.Wave1GoldenDelivered = (*shift).WaveDetails[0].GoldenEggs
	shiftUpload.Wave1PowerEggs = (*shift).WaveDetails[0].PowerEggs
	shiftUpload.Wave1Quota = (*shift).WaveDetails[0].QuotaNum
	shiftUpload.Wave1WaterLevel = (*shift).WaveDetails[0].WaterLevel.Key

	if len((*shift).WaveDetails) > 1 {
		shiftUpload.Wave2EventType = &(*shift).WaveDetails[1].EventType.Key
		shiftUpload.Wave2GoldenAppear = &(*shift).WaveDetails[1].GoldenAppear
		shiftUpload.Wave2GoldenDelivered = &(*shift).WaveDetails[1].GoldenEggs
		shiftUpload.Wave2PowerEggs = &(*shift).WaveDetails[1].PowerEggs
		shiftUpload.Wave2Quota = &(*shift).WaveDetails[1].QuotaNum
		shiftUpload.Wave2WaterLevel = &(*shift).WaveDetails[1].WaterLevel.Key

		if len((*shift).WaveDetails) > 2 {
			shiftUpload.Wave3EventType = &(*shift).WaveDetails[2].EventType.Key
			shiftUpload.Wave3GoldenAppear = &(*shift).WaveDetails[2].GoldenAppear
			shiftUpload.Wave3GoldenDelivered = &(*shift).WaveDetails[2].GoldenEggs
			shiftUpload.Wave3PowerEggs = &(*shift).WaveDetails[2].PowerEggs
			shiftUpload.Wave3Quota = &(*shift).WaveDetails[2].QuotaNum
			shiftUpload.Wave3WaterLevel = &(*shift).WaveDetails[2].WaterLevel.Key
		}
	}

	return shiftUpload
}

func UploadSalmon(shiftUpload *types.Shift, apiKey string, client *http.Client) {
	if (*shiftUpload).PlayerSplatnetId == "" || (*shiftUpload).JobId < 1 {
		log.Println("Skipping shift due to missing data.")
		return
	}

	bodyMarshalled, err := json.Marshal(shiftUpload)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "https://splatstats.cass-dlcm.dev/api/two_salmon/", bytes.NewReader(bodyMarshalled))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"session_token": []string{apiKey},
		"Content-Type":  []string{"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if resp.StatusCode == 302 {
		log.Printf("Shift #%d already uploaded to %s\n", (*shiftUpload).JobId, resp.Header.Get("location"))
	} else if resp.StatusCode == 201 {
		log.Printf("Shift #%d uploaded to %s\n", (*shiftUpload).JobId, resp.Header.Get("location"))
	} else {
		log.Println(resp.Status)
		log.Println(string(bodyMarshalled))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		panic(nil)
	}
}

func salmonPlayerWeaponSpecials(shift *types.ShiftSplatnet, shiftUpload *types.Shift) {
	(*shiftUpload).PlayerW1Weapon = (*shift).MyResult.WeaponList[0].Weapon.Name
	(*shiftUpload).PlayerW1Specials = (*shift).MyResult.SpecialCounts[0]

	if len((*shift).MyResult.WeaponList) > 1 {
		(*shiftUpload).PlayerW2Weapon = &(*shift).MyResult.WeaponList[1].Weapon.Name

		if len((*shift).MyResult.WeaponList) > 2 {
			(*shiftUpload).PlayerW3Weapon = &(*shift).MyResult.WeaponList[2].Weapon.Name
		}
	}

	if len((*shift).MyResult.SpecialCounts) > 1 {
		(*shiftUpload).PlayerW2Specials = &(*shift).MyResult.SpecialCounts[1]

		if len((*shift).MyResult.SpecialCounts) > 2 {
			(*shiftUpload).PlayerW3Specials = &(*shift).MyResult.SpecialCounts[2]
		}
	}
}

func shiftSetTimes(shift *types.ShiftSplatnet, shiftUpload *types.Shift) {
	(*shiftUpload).PlayTime = (*shift).PlayTime
	(*shiftUpload).EndTime = (*shift).EndTime
	(*shiftUpload).ScheduleEndTime = &(*shift).Schedule.EndTime
	(*shiftUpload).ScheduleStartTime = (*shift).Schedule.StartTime
	(*shiftUpload).StartTime = (*shift).StartTime
}

func shiftSetTeammate0(shift *types.ShiftSplatnet, shiftUpload *types.Shift) {
	if len((*shift).OtherResults) < 1 {
		return
	}
	(*shiftUpload).Teammate0DeathCount = &(*shift).OtherResults[0].DeadCount
	(*shiftUpload).Teammate0DrizzlerKills = &(*shift).OtherResults[0].BossKillCounts.Drizzler.Count
	(*shiftUpload).Teammate0FlyfishKills = &(*shift).OtherResults[0].BossKillCounts.Flyfish.Count
	(*shiftUpload).Teammate0Gender = &(*shift).OtherResults[0].PlayerType.Gender
	(*shiftUpload).Teammate0GoldenEggs = &(*shift).OtherResults[0].GoldenEggs
	(*shiftUpload).Teammate0GoldieKills = &(*shift).OtherResults[0].BossKillCounts.Goldie.Count
	(*shiftUpload).Teammate0GrillerKills = &(*shift).OtherResults[0].BossKillCounts.Griller.Count
	(*shiftUpload).Teammate0SplatnetId = &(*shift).OtherResults[0].Pid
	(*shiftUpload).Teammate0MawsKills = &(*shift).OtherResults[0].BossKillCounts.Maws.Count
	(*shiftUpload).Teammate0Name = &(*shift).OtherResults[0].Name
	(*shiftUpload).Teammate0PowerEggs = &(*shift).OtherResults[0].PowerEggs
	(*shiftUpload).Teammate0ReviveCount = &(*shift).OtherResults[0].HelpCount
	(*shiftUpload).Teammate0ScrapperKills = &(*shift).OtherResults[0].BossKillCounts.Scrapper.Count
	(*shiftUpload).Teammate0Special = &(*shift).OtherResults[0].Special.Id
	(*shiftUpload).Teammate0Species = &(*shift).OtherResults[0].PlayerType.Species
	(*shiftUpload).Teammate0SteelEelKills = &(*shift).OtherResults[0].BossKillCounts.SteelEel.Count
	(*shiftUpload).Teammate0SteelheadKills = &(*shift).OtherResults[0].BossKillCounts.Steelhead.Count
	(*shiftUpload).Teammate0StingerKills = &(*shift).OtherResults[0].BossKillCounts.Stinger.Count
	(*shiftUpload).Teammate0W1Specials = &(*shift).OtherResults[0].SpecialCounts[0]

	if len((*shift).OtherResults[0].WeaponList) > 0 {
		(*shiftUpload).Teammate0W1Weapon = &(*shift).OtherResults[0].WeaponList[0].Weapon.Name

		if len((*shift).OtherResults[0].WeaponList) > 1 {
			(*shiftUpload).Teammate0W2Weapon = &(*shift).OtherResults[0].WeaponList[1].Weapon.Name

			if len((*shift).OtherResults[0].WeaponList) > 2 {
				(*shiftUpload).Teammate0W3Weapon = &(*shift).OtherResults[0].WeaponList[2].Weapon.Name
			}
		}
	}

	if len((*shift).OtherResults[0].SpecialCounts) > 1 {
		(*shiftUpload).Teammate0W2Specials = &(*shift).OtherResults[0].SpecialCounts[1]

		if len((*shift).OtherResults[0].SpecialCounts) > 2 {
			(*shiftUpload).Teammate0W3Specials = &(*shift).OtherResults[0].SpecialCounts[2]
		}
	}
}

func shiftSetTeammate1(shift *types.ShiftSplatnet, shiftUpload *types.Shift) {
	if len((*shift).OtherResults) < 2 {
		return
	}
	(*shiftUpload).Teammate1DeathCount = &(*shift).OtherResults[1].DeadCount
	(*shiftUpload).Teammate1DrizzlerKills = &(*shift).OtherResults[1].BossKillCounts.Drizzler.Count
	(*shiftUpload).Teammate1FlyfishKills = &(*shift).OtherResults[1].BossKillCounts.Flyfish.Count
	(*shiftUpload).Teammate1Gender = &(*shift).OtherResults[1].PlayerType.Gender
	(*shiftUpload).Teammate1GoldenEggs = &(*shift).OtherResults[1].GoldenEggs
	(*shiftUpload).Teammate1GoldieKills = &(*shift).OtherResults[1].BossKillCounts.Goldie.Count
	(*shiftUpload).Teammate1GrillerKills = &(*shift).OtherResults[1].BossKillCounts.Griller.Count
	(*shiftUpload).Teammate1SplatnetId = &(*shift).OtherResults[1].Pid
	(*shiftUpload).Teammate1MawsKills = &(*shift).OtherResults[1].BossKillCounts.Maws.Count
	(*shiftUpload).Teammate1Name = &(*shift).OtherResults[1].Name
	(*shiftUpload).Teammate1PowerEggs = &(*shift).OtherResults[1].PowerEggs
	(*shiftUpload).Teammate1ReviveCount = &(*shift).OtherResults[1].HelpCount
	(*shiftUpload).Teammate1ScrapperKills = &(*shift).OtherResults[1].BossKillCounts.Scrapper.Count
	(*shiftUpload).Teammate1Special = &(*shift).OtherResults[1].Special.Id
	(*shiftUpload).Teammate1Species = &(*shift).OtherResults[1].PlayerType.Species
	(*shiftUpload).Teammate1SteelEelKills = &(*shift).OtherResults[1].BossKillCounts.SteelEel.Count
	(*shiftUpload).Teammate1SteelheadKills = &(*shift).OtherResults[1].BossKillCounts.Steelhead.Count
	(*shiftUpload).Teammate1StingerKills = &(*shift).OtherResults[1].BossKillCounts.Stinger.Count
	(*shiftUpload).Teammate1W1Specials = &(*shift).OtherResults[1].SpecialCounts[0]

	if len((*shift).OtherResults[1].SpecialCounts) > 1 {
		(*shiftUpload).Teammate1W2Specials = &(*shift).OtherResults[1].SpecialCounts[1]

		if len((*shift).OtherResults[1].SpecialCounts) > 2 {
			shiftUpload.Teammate1W3Specials = &(*shift).OtherResults[1].SpecialCounts[2]
		}
	}

	if len((*shift).OtherResults[1].WeaponList) > 0 {
		(*shiftUpload).Teammate1W1Weapon = &(*shift).OtherResults[1].WeaponList[0].Weapon.Name

		if len((*shift).OtherResults[1].WeaponList) > 1 {
			(*shiftUpload).Teammate1W2Weapon = &(*shift).OtherResults[1].WeaponList[1].Weapon.Name

			if len((*shift).OtherResults[1].WeaponList) > 2 {
				(*shiftUpload).Teammate1W3Weapon = &(*shift).OtherResults[1].WeaponList[2].Weapon.Name
			}
		}
	}
}

func shiftSetTeammate2(shift *types.ShiftSplatnet, shiftUpload *types.Shift) {
	if len((*shift).OtherResults) > 2 {
		(*shiftUpload).Teammate2DeathCount = &(*shift).OtherResults[2].DeadCount
		(*shiftUpload).Teammate2DrizzlerKills = &(*shift).OtherResults[2].BossKillCounts.Drizzler.Count
		(*shiftUpload).Teammate2FlyfishKills = &(*shift).OtherResults[2].BossKillCounts.Flyfish.Count
		(*shiftUpload).Teammate2Gender = &(*shift).OtherResults[2].PlayerType.Gender
		(*shiftUpload).Teammate2GoldenEggs = &(*shift).OtherResults[2].GoldenEggs
		(*shiftUpload).Teammate2GoldieKills = &(*shift).OtherResults[2].BossKillCounts.Goldie.Count
		(*shiftUpload).Teammate2GrillerKills = &(*shift).OtherResults[2].BossKillCounts.Griller.Count
		(*shiftUpload).Teammate2SplatnetId = &(*shift).OtherResults[2].Pid
		(*shiftUpload).Teammate2MawsKills = &(*shift).OtherResults[2].BossKillCounts.Maws.Count
		(*shiftUpload).Teammate2Name = &(*shift).OtherResults[2].Name
		(*shiftUpload).Teammate2PowerEggs = &(*shift).OtherResults[2].PowerEggs
		(*shiftUpload).Teammate2ReviveCount = &(*shift).OtherResults[2].HelpCount
		(*shiftUpload).Teammate2ScrapperKills = &(*shift).OtherResults[2].BossKillCounts.Scrapper.Count
		(*shiftUpload).Teammate2Special = &(*shift).OtherResults[2].Special.Id
		(*shiftUpload).Teammate2Species = &(*shift).OtherResults[2].PlayerType.Species
		(*shiftUpload).Teammate2SteelEelKills = &(*shift).OtherResults[2].BossKillCounts.SteelEel.Count
		(*shiftUpload).Teammate2SteelheadKills = &(*shift).OtherResults[2].BossKillCounts.Steelhead.Count
		(*shiftUpload).Teammate2StingerKills = &(*shift).OtherResults[2].BossKillCounts.Stinger.Count
		(*shiftUpload).Teammate2W1Specials = &(*shift).OtherResults[2].SpecialCounts[0]

		if len((*shift).OtherResults[2].SpecialCounts) > 1 {
			(*shiftUpload).Teammate2W2Specials = &(*shift).OtherResults[2].SpecialCounts[1]

			if len((*shift).OtherResults[2].SpecialCounts) > 2 {
				(*shiftUpload).Teammate2W3Specials = &(*shift).OtherResults[2].SpecialCounts[2]
			}
		}

		if len((*shift).OtherResults[2].WeaponList) > 0 {
			(*shiftUpload).Teammate2W1Weapon = &(*shift).OtherResults[2].WeaponList[0].Weapon.Name

			if len((*shift).OtherResults[2].WeaponList) > 1 {
				(*shiftUpload).Teammate2W2Weapon = &(*shift).OtherResults[2].WeaponList[1].Weapon.Name

				if len((*shift).OtherResults[2].WeaponList) > 2 {
					(*shiftUpload).Teammate2W3Weapon = &(*shift).OtherResults[2].WeaponList[2].Weapon.Name
				}
			}
		}
	}
}

func setBattlePayload(battle *types.BattleSplatnet) types.Battle {
	battleUpload := types.Battle{}
	battleUpload.SplatnetJson = battle
	battleUpload.SplatnetUpload = true
	battleUpload.StatInkUpload = false
	var err error
	battleUpload.BattleNumber, err = strconv.Atoi((*battle).BattleNumber)
	if err != nil {
		panic(err)
	}
	battleUpload.Rule = (*battle).Rule.Key
	battleUpload.MatchType = (*battle).GameMode.Key
	battleUpload.Stage = (*battle).Stage.Id
	battleUpload.Win = (*battle).MyTeamResult.Key == "victory"
	battleSetHasDc(battle, &battleUpload)
	battleUpload.Time = (*battle).StartTime
	battleUpload.WinMeter = (*battle).WinMeter
	battleSetScoreTime(battle, &battleUpload)
	battleUpload.TagId = (*battle).TagId
	battleUpload.LeaguePoint = (*battle).LeaguePoint

	battleSetPlayer(battle, &battleUpload)
	battleSetTeammate0(battle, &battleUpload)
	battleSetTeammate1(battle, &battleUpload)
	battleSetTeammate2(battle, &battleUpload)

	battleSetOpponent0(battle, &battleUpload)
	battleSetOpponent1(battle, &battleUpload)
	battleSetOpponent2(battle, &battleUpload)
	battleSetOpponent3(battle, &battleUpload)

	return battleUpload
}

func UploadBattle(battleUpload *types.Battle, apiKey string, client *http.Client) {
	if (*battleUpload).PlayerSplatnetId == "" || (*battleUpload).BattleNumber == 0 {
		log.Println("Skipping battle due to missing data.")
		return
	}

	bodyMarshalled, err := json.Marshal(battleUpload)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "https://splatstats.cass-dlcm.dev/api/two_battles/", bytes.NewReader(bodyMarshalled))
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"session_token": []string{apiKey},
		"Content-Type":  []string{"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if resp.StatusCode == 302 {
		log.Printf("Battle #%d already uploaded to %s\n", (*battleUpload).BattleNumber, resp.Header.Get("location"))
	} else if resp.StatusCode == 201 {
		log.Printf("Battle #%d uploaded to %s\n", (*battleUpload).BattleNumber, resp.Header.Get("location"))
	} else {
		log.Println(resp.Status)
		log.Println(string(bodyMarshalled))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		panic(nil)
	}
}

func battleSetHasDc(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	(*battleUpload).HasDisconnectedPlayer = false

	for i := range (*battle).MyTeamMembers {
		(*battleUpload).HasDisconnectedPlayer = (*battleUpload).HasDisconnectedPlayer || ((*battle).MyTeamMembers[i].GamePaintPoint == 0 && (*battle).MyTeamMembers[i].KillCount == 0 && (*battle).MyTeamMembers[i].SpecialCount == 0 && (*battle).MyTeamMembers[i].DeathCount == 0 && (*battle).MyTeamMembers[i].AssistCount == 0)
	}

	for i := range (*battle).OtherTeamMembers {
		(*battleUpload).HasDisconnectedPlayer = (*battleUpload).HasDisconnectedPlayer || ((*battle).OtherTeamMembers[i].GamePaintPoint == 0 && (*battle).OtherTeamMembers[i].KillCount == 0 && (*battle).OtherTeamMembers[i].SpecialCount == 0 && (*battle).OtherTeamMembers[i].DeathCount == 0 && (*battle).OtherTeamMembers[i].AssistCount == 0)
	}
}

func battleSetScoreTime(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	if (*battle).MyTeamCount != nil {
		(*battleUpload).MyTeamCount = float64(*(*battle).MyTeamCount)
	} else {
		(*battleUpload).MyTeamCount = *(*battle).MyTeamPercentage
	}

	if (*battle).OtherTeamCount != nil {
		(*battleUpload).OtherTeamCount = float64(*(*battle).OtherTeamCount)
	} else {
		(*battleUpload).OtherTeamCount = *(*battle).OtherTeamPercentage
	}

	if battleUpload.Rule == "turf_war" {
		(*battleUpload).ElapsedTime = 180
	} else {
		(*battleUpload).ElapsedTime = *(*battle).ElapsedTime
	}
}

func battleSetPlayer(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	(*battleUpload).PlayerSplatnetId = (*battle).PlayerResult.Player.PrincipalId
	(*battleUpload).PlayerName = (*battle).PlayerResult.Player.Nickname
	(*battleUpload).PlayerWeapon = (*battle).PlayerResult.Player.Weapon.Id
	if (*battle).Udemae != nil {
		(*battleUpload).PlayerRank = &(*battle).Udemae.Name
	}
	(*battleUpload).PlayerLevelStar = (*battle).StarRank
	(*battleUpload).PlayerLevel = (*battle).PlayerRank
	(*battleUpload).PlayerKills = (*battle).PlayerResult.KillCount
	(*battleUpload).PlayerDeaths = (*battle).PlayerResult.DeathCount
	(*battleUpload).PlayerAssists = (*battle).PlayerResult.AssistCount
	(*battleUpload).PlayerSpecials = (*battle).PlayerResult.SpecialCount
	(*battleUpload).PlayerGamePaintPoint = (*battle).PlayerResult.GamePaintPoint
	(*battleUpload).PlayerGender = (*battle).PlayerResult.Player.PlayerType.Gender
	(*battleUpload).PlayerSpecies = (*battle).PlayerResult.Player.PlayerType.Species
	(*battleUpload).PlayerHeadgear = (*battle).PlayerResult.Player.Head.Id
	(*battleUpload).PlayerHeadgearMain = (*battle).PlayerResult.Player.HeadSkills.Main.Id

	if len((*battle).PlayerResult.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).PlayerHeadgearSub0 = (*battle).PlayerResult.Player.HeadSkills.Subs[0].Id

		if len((*battle).PlayerResult.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).PlayerHeadgearSub1 = (*battle).PlayerResult.Player.HeadSkills.Subs[1].Id

			if len((*battle).PlayerResult.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).PlayerHeadgearSub2 = (*battle).PlayerResult.Player.HeadSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).PlayerClothes = (*battle).PlayerResult.Player.Clothes.Id
	(*battleUpload).PlayerClothesMain = (*battle).PlayerResult.Player.ClothesSkills.Main.Id

	if len((*battle).PlayerResult.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).PlayerClothesSub0 = (*battle).PlayerResult.Player.ClothesSkills.Subs[0].Id

		if len((*battle).PlayerResult.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).PlayerClothesSub1 = (*battle).PlayerResult.Player.ClothesSkills.Subs[1].Id

			if len((*battle).PlayerResult.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).PlayerClothesSub2 = (*battle).PlayerResult.Player.ClothesSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).PlayerShoes = (*battle).PlayerResult.Player.Shoes.Id
	(*battleUpload).PlayerShoesMain = (*battle).PlayerResult.Player.ShoesSkills.Main.Id

	if len((*battle).PlayerResult.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).PlayerShoesSub0 = (*battle).PlayerResult.Player.ShoesSkills.Subs[0].Id

		if len((*battle).PlayerResult.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).PlayerShoesSub1 = (*battle).PlayerResult.Player.ShoesSkills.Subs[1].Id

			if len((*battle).PlayerResult.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).PlayerShoesSub2 = (*battle).PlayerResult.Player.ShoesSkills.Subs[2].Id
			}
		}
	}
}

func battleSetTeammate0(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	if len((*battle).MyTeamMembers) < 1 {
		return
	}
	teammate0 := (*battle).MyTeamMembers[0]
	(*battleUpload).Teammate0SplatnetId = &teammate0.Player.PrincipalId
	(*battleUpload).Teammate0Name = &teammate0.Player.Nickname
	(*battleUpload).Teammate0LevelStar = &teammate0.Player.StarRank
	(*battleUpload).Teammate0Level = &teammate0.Player.PlayerRank
	if teammate0.Player.Udemae != nil {
		(*battleUpload).Teammate0Rank = &teammate0.Player.Udemae.Name
	}
	(*battleUpload).Teammate0Weapon = &teammate0.Player.Weapon.Id
	(*battleUpload).Teammate0Gender = &teammate0.Player.PlayerType.Gender
	(*battleUpload).Teammate0Species = &teammate0.Player.PlayerType.Species
	(*battleUpload).Teammate0Kills = &teammate0.KillCount
	(*battleUpload).Teammate0Deaths = &teammate0.DeathCount
	(*battleUpload).Teammate0Assists = &teammate0.AssistCount
	(*battleUpload).Teammate0GamePaintPoint = &teammate0.GamePaintPoint
	(*battleUpload).Teammate0Specials = &teammate0.SpecialCount
	(*battleUpload).Teammate0Headgear = &teammate0.Player.Head.Id
	(*battleUpload).Teammate0HeadgearMain = &teammate0.Player.HeadSkills.Main.Id

	if len(teammate0.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).Teammate0HeadgearSub0 = &teammate0.Player.HeadSkills.Subs[0].Id

		if len(teammate0.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).Teammate0HeadgearSub1 = &teammate0.Player.HeadSkills.Subs[1].Id

			if len(teammate0.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).Teammate0HeadgearSub2 = &teammate0.Player.HeadSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Teammate0Clothes = &teammate0.Player.Clothes.Id
	(*battleUpload).Teammate0ClothesMain = &teammate0.Player.ClothesSkills.Main.Id

	if len(teammate0.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).Teammate0ClothesSub0 = &teammate0.Player.ClothesSkills.Subs[0].Id

		if len(teammate0.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).Teammate0ClothesSub1 = &teammate0.Player.ClothesSkills.Subs[1].Id

			if len(teammate0.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).Teammate0ClothesSub2 = &teammate0.Player.ClothesSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Teammate0Shoes = &teammate0.Player.Shoes.Id
	(*battleUpload).Teammate0ShoesMain = &teammate0.Player.ShoesSkills.Main.Id

	if len(teammate0.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).Teammate0ShoesSub0 = &teammate0.Player.ShoesSkills.Subs[0].Id

		if len(teammate0.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).Teammate0ShoesSub1 = &teammate0.Player.ShoesSkills.Subs[1].Id

			if len(teammate0.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).Teammate0ShoesSub2 = &teammate0.Player.ShoesSkills.Subs[2].Id
			}
		}
	}
}

func battleSetTeammate1(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	if len((*battle).MyTeamMembers) < 2 {
		return
	}
	teammate1 := (*battle).MyTeamMembers[1]
	(*battleUpload).Teammate1SplatnetId = &teammate1.Player.PrincipalId
	(*battleUpload).Teammate1Name = &teammate1.Player.Nickname
	(*battleUpload).Teammate1LevelStar = &teammate1.Player.StarRank
	(*battleUpload).Teammate1Level = &teammate1.Player.PlayerRank
	if teammate1.Player.Udemae != nil {
		(*battleUpload).Teammate1Rank = &teammate1.Player.Udemae.Name
	}
	(*battleUpload).Teammate1Weapon = &teammate1.Player.Weapon.Id
	(*battleUpload).Teammate1Gender = &teammate1.Player.PlayerType.Gender
	(*battleUpload).Teammate1Species = &teammate1.Player.PlayerType.Species
	(*battleUpload).Teammate1Kills = &teammate1.KillCount
	(*battleUpload).Teammate1Deaths = &teammate1.DeathCount
	(*battleUpload).Teammate1Assists = &teammate1.AssistCount
	(*battleUpload).Teammate1GamePaintPoint = &teammate1.GamePaintPoint
	(*battleUpload).Teammate1Specials = &teammate1.SpecialCount
	(*battleUpload).Teammate1Headgear = &teammate1.Player.Head.Id
	(*battleUpload).Teammate1HeadgearMain = &teammate1.Player.HeadSkills.Main.Id

	if len(teammate1.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).Teammate1HeadgearSub0 = &teammate1.Player.HeadSkills.Subs[0].Id

		if len(teammate1.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).Teammate1HeadgearSub1 = &teammate1.Player.HeadSkills.Subs[1].Id

			if len(teammate1.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).Teammate1HeadgearSub2 = &teammate1.Player.HeadSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Teammate1Clothes = &teammate1.Player.Clothes.Id
	(*battleUpload).Teammate1ClothesMain = &teammate1.Player.ClothesSkills.Main.Id

	if len(teammate1.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).Teammate1ClothesSub0 = &teammate1.Player.ClothesSkills.Subs[0].Id

		if len(teammate1.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).Teammate1ClothesSub1 = &teammate1.Player.ClothesSkills.Subs[1].Id

			if len(teammate1.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).Teammate1ClothesSub2 = &teammate1.Player.ClothesSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Teammate1Shoes = &teammate1.Player.Shoes.Id
	(*battleUpload).Teammate1ShoesMain = &teammate1.Player.ShoesSkills.Main.Id

	if len(teammate1.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).Teammate1ShoesSub0 = &teammate1.Player.ShoesSkills.Subs[0].Id

		if len(teammate1.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).Teammate1ShoesSub1 = &teammate1.Player.ShoesSkills.Subs[1].Id

			if len(teammate1.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).Teammate1ShoesSub2 = &teammate1.Player.ShoesSkills.Subs[2].Id
			}
		}
	}
}

func battleSetTeammate2(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	if len((*battle).MyTeamMembers) < 3 {
		return
	}
	teammate2 := (*battle).MyTeamMembers[2]
	(*battleUpload).Teammate2SplatnetId = &teammate2.Player.PrincipalId
	(*battleUpload).Teammate2Name = &teammate2.Player.Nickname
	(*battleUpload).Teammate2LevelStar = &teammate2.Player.StarRank
	(*battleUpload).Teammate2Level = &teammate2.Player.PlayerRank
	if teammate2.Player.Udemae != nil {
		(*battleUpload).Teammate2Rank = &teammate2.Player.Udemae.Name
	}
	(*battleUpload).Teammate2Weapon = &teammate2.Player.Weapon.Id
	(*battleUpload).Teammate2Gender = &teammate2.Player.PlayerType.Gender
	(*battleUpload).Teammate2Species = &teammate2.Player.PlayerType.Species
	(*battleUpload).Teammate2Kills = &teammate2.KillCount
	(*battleUpload).Teammate2Deaths = &teammate2.DeathCount
	(*battleUpload).Teammate2Assists = &teammate2.AssistCount
	(*battleUpload).Teammate2GamePaintPoint = &teammate2.GamePaintPoint
	(*battleUpload).Teammate2Specials = &teammate2.SpecialCount
	(*battleUpload).Teammate2Headgear = &teammate2.Player.Head.Id
	(*battleUpload).Teammate2HeadgearMain = &teammate2.Player.HeadSkills.Main.Id

	if len(teammate2.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).Teammate2HeadgearSub0 = &teammate2.Player.HeadSkills.Subs[0].Id

		if len(teammate2.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).Teammate2HeadgearSub1 = &teammate2.Player.HeadSkills.Subs[1].Id

			if len(teammate2.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).Teammate2HeadgearSub2 = &teammate2.Player.HeadSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Teammate2Clothes = &teammate2.Player.Clothes.Id
	(*battleUpload).Teammate2ClothesMain = &teammate2.Player.ClothesSkills.Main.Id

	if len(teammate2.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).Teammate2ClothesSub0 = &teammate2.Player.ClothesSkills.Subs[0].Id

		if len(teammate2.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).Teammate2ClothesSub1 = &teammate2.Player.ClothesSkills.Subs[1].Id

			if len(teammate2.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).Teammate2ClothesSub2 = &teammate2.Player.ClothesSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Teammate2Shoes = &teammate2.Player.Shoes.Id
	(*battleUpload).Teammate2ShoesMain = &teammate2.Player.ShoesSkills.Main.Id

	if len(teammate2.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).Teammate2ShoesSub0 = &teammate2.Player.ShoesSkills.Subs[0].Id

		if len(teammate2.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).Teammate2ShoesSub1 = &teammate2.Player.ShoesSkills.Subs[1].Id

			if len(teammate2.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).Teammate2ShoesSub2 = &teammate2.Player.ShoesSkills.Subs[2].Id
			}
		}
	}
}

func battleSetOpponent0(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	if len((*battle).OtherTeamMembers) < 1 {
		return
	}
	opponent0 := (*battle).OtherTeamMembers[0]
	(*battleUpload).Opponent0SplatnetId = &opponent0.Player.PrincipalId
	(*battleUpload).Opponent0Name = &opponent0.Player.Nickname
	(*battleUpload).Opponent0LevelStar = &opponent0.Player.StarRank
	(*battleUpload).Opponent0Level = &opponent0.Player.PlayerRank
	if opponent0.Player.Udemae != nil {
		(*battleUpload).Opponent0Rank = &opponent0.Player.Udemae.Name
	}
	(*battleUpload).Opponent0Weapon = &opponent0.Player.Weapon.Id
	(*battleUpload).Opponent0Gender = &opponent0.Player.PlayerType.Gender
	(*battleUpload).Opponent0Species = &opponent0.Player.PlayerType.Species
	(*battleUpload).Opponent0Kills = &opponent0.KillCount
	(*battleUpload).Opponent0Deaths = &opponent0.DeathCount
	(*battleUpload).Opponent0Assists = &opponent0.AssistCount
	(*battleUpload).Opponent0GamePaintPoint = &opponent0.GamePaintPoint
	(*battleUpload).Opponent0Specials = &opponent0.SpecialCount
	(*battleUpload).Opponent0Headgear = &opponent0.Player.Head.Id
	(*battleUpload).Opponent0HeadgearMain = &opponent0.Player.HeadSkills.Main.Id

	if len(opponent0.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).Opponent0HeadgearSub0 = &opponent0.Player.HeadSkills.Subs[0].Id

		if len(opponent0.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).Opponent0HeadgearSub1 = &opponent0.Player.HeadSkills.Subs[1].Id

			if len(opponent0.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).Opponent0HeadgearSub2 = &opponent0.Player.HeadSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Opponent0Clothes = &opponent0.Player.Clothes.Id
	(*battleUpload).Opponent0ClothesMain = &opponent0.Player.ClothesSkills.Main.Id

	if len(opponent0.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).Opponent0ClothesSub0 = &opponent0.Player.ClothesSkills.Subs[0].Id

		if len(opponent0.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).Opponent0ClothesSub1 = &opponent0.Player.ClothesSkills.Subs[1].Id

			if len(opponent0.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).Opponent0ClothesSub2 = &opponent0.Player.ClothesSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Opponent0Shoes = &opponent0.Player.Shoes.Id
	(*battleUpload).Opponent0ShoesMain = &opponent0.Player.ShoesSkills.Main.Id

	if len(opponent0.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).Opponent0ShoesSub0 = &opponent0.Player.ShoesSkills.Subs[0].Id

		if len(opponent0.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).Opponent0ShoesSub1 = &opponent0.Player.ShoesSkills.Subs[1].Id

			if len(opponent0.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).Opponent0ShoesSub2 = &opponent0.Player.ShoesSkills.Subs[2].Id
			}
		}
	}
}

func battleSetOpponent1(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	if len((*battle).OtherTeamMembers) < 2 {
		return
	}
	opponent1 := (*battle).OtherTeamMembers[1]
	(*battleUpload).Opponent1SplatnetId = &opponent1.Player.PrincipalId
	(*battleUpload).Opponent1Name = &opponent1.Player.Nickname
	(*battleUpload).Opponent1LevelStar = &opponent1.Player.StarRank
	(*battleUpload).Opponent1Level = &opponent1.Player.PlayerRank
	if opponent1.Player.Udemae != nil {
		(*battleUpload).Opponent1Rank = &opponent1.Player.Udemae.Name
	}
	(*battleUpload).Opponent1Weapon = &opponent1.Player.Weapon.Id
	(*battleUpload).Opponent1Gender = &opponent1.Player.PlayerType.Gender
	(*battleUpload).Opponent1Species = &opponent1.Player.PlayerType.Species
	(*battleUpload).Opponent1Kills = &opponent1.KillCount
	(*battleUpload).Opponent1Deaths = &opponent1.DeathCount
	(*battleUpload).Opponent1Assists = &opponent1.AssistCount
	(*battleUpload).Opponent1GamePaintPoint = &opponent1.GamePaintPoint
	(*battleUpload).Opponent1Specials = &opponent1.SpecialCount
	(*battleUpload).Opponent1Headgear = &opponent1.Player.Head.Id
	(*battleUpload).Opponent1HeadgearMain = &opponent1.Player.HeadSkills.Main.Id

	if len(opponent1.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).Opponent1HeadgearSub0 = &opponent1.Player.HeadSkills.Subs[0].Id

		if len(opponent1.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).Opponent1HeadgearSub1 = &opponent1.Player.HeadSkills.Subs[1].Id

			if len(opponent1.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).Opponent1HeadgearSub2 = &opponent1.Player.HeadSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Opponent1Clothes = &opponent1.Player.Clothes.Id
	(*battleUpload).Opponent1ClothesMain = &opponent1.Player.ClothesSkills.Main.Id

	if len(opponent1.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).Opponent1ClothesSub0 = &opponent1.Player.ClothesSkills.Subs[0].Id

		if len(opponent1.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).Opponent1ClothesSub1 = &opponent1.Player.ClothesSkills.Subs[1].Id

			if len(opponent1.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).Opponent1ClothesSub2 = &opponent1.Player.ClothesSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Opponent1Shoes = &opponent1.Player.Shoes.Id
	(*battleUpload).Opponent1ShoesMain = &opponent1.Player.ShoesSkills.Main.Id

	if len(opponent1.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).Opponent1ShoesSub0 = &opponent1.Player.ShoesSkills.Subs[0].Id

		if len(opponent1.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).Opponent1ShoesSub1 = &opponent1.Player.ShoesSkills.Subs[1].Id

			if len(opponent1.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).Opponent1ShoesSub2 = &opponent1.Player.ShoesSkills.Subs[2].Id
			}
		}
	}
}

func battleSetOpponent2(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	if len((*battle).OtherTeamMembers) < 3 {
		return
	}
	opponent2 := (*battle).OtherTeamMembers[2]
	(*battleUpload).Opponent2SplatnetId = &opponent2.Player.PrincipalId
	(*battleUpload).Opponent2Name = &opponent2.Player.Nickname
	(*battleUpload).Opponent2LevelStar = &opponent2.Player.StarRank
	(*battleUpload).Opponent2Level = &opponent2.Player.PlayerRank
	if opponent2.Player.Udemae != nil {
		(*battleUpload).Opponent2Rank = &opponent2.Player.Udemae.Name
	}
	(*battleUpload).Opponent2Weapon = &opponent2.Player.Weapon.Id
	(*battleUpload).Opponent2Gender = &opponent2.Player.PlayerType.Gender
	(*battleUpload).Opponent2Species = &opponent2.Player.PlayerType.Species
	(*battleUpload).Opponent2Kills = &opponent2.KillCount
	(*battleUpload).Opponent2Deaths = &opponent2.DeathCount
	(*battleUpload).Opponent2Assists = &opponent2.AssistCount
	(*battleUpload).Opponent2GamePaintPoint = &opponent2.GamePaintPoint
	(*battleUpload).Opponent2Specials = &opponent2.SpecialCount
	(*battleUpload).Opponent2Headgear = &opponent2.Player.Head.Id
	(*battleUpload).Opponent2HeadgearMain = &opponent2.Player.HeadSkills.Main.Id

	if len(opponent2.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).Opponent2HeadgearSub0 = &opponent2.Player.HeadSkills.Subs[0].Id

		if len(opponent2.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).Opponent2HeadgearSub1 = &opponent2.Player.HeadSkills.Subs[1].Id

			if len(opponent2.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).Opponent2HeadgearSub2 = &opponent2.Player.HeadSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Opponent2Clothes = &opponent2.Player.Clothes.Id
	(*battleUpload).Opponent2ClothesMain = &opponent2.Player.ClothesSkills.Main.Id

	if len(opponent2.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).Opponent2ClothesSub0 = &opponent2.Player.ClothesSkills.Subs[0].Id

		if len(opponent2.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).Opponent2ClothesSub1 = &opponent2.Player.ClothesSkills.Subs[1].Id

			if len(opponent2.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).Opponent2ClothesSub2 = &opponent2.Player.ClothesSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Opponent2Shoes = &opponent2.Player.Shoes.Id
	(*battleUpload).Opponent2ShoesMain = &opponent2.Player.ShoesSkills.Main.Id

	if len(opponent2.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).Opponent2ShoesSub0 = &opponent2.Player.ShoesSkills.Subs[0].Id

		if len(opponent2.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).Opponent2ShoesSub1 = &opponent2.Player.ShoesSkills.Subs[1].Id

			if len(opponent2.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).Opponent2ShoesSub2 = &opponent2.Player.ShoesSkills.Subs[2].Id
			}
		}
	}
}

func battleSetOpponent3(battle *types.BattleSplatnet, battleUpload *types.Battle) {
	if len((*battle).OtherTeamMembers) < 4 {
		return
	}
	opponent3 := (*battle).OtherTeamMembers[3]
	(*battleUpload).Opponent3SplatnetId = &opponent3.Player.PrincipalId
	(*battleUpload).Opponent3Name = &opponent3.Player.Nickname
	(*battleUpload).Opponent3LevelStar = &opponent3.Player.StarRank
	(*battleUpload).Opponent3Level = &opponent3.Player.PlayerRank
	if opponent3.Player.Udemae != nil {
		(*battleUpload).Opponent0Rank = &opponent3.Player.Udemae.Name
	}
	(*battleUpload).Opponent3Weapon = &opponent3.Player.Weapon.Id
	(*battleUpload).Opponent3Gender = &opponent3.Player.PlayerType.Gender
	(*battleUpload).Opponent3Species = &opponent3.Player.PlayerType.Species
	(*battleUpload).Opponent3Kills = &opponent3.KillCount
	(*battleUpload).Opponent3Deaths = &opponent3.DeathCount
	(*battleUpload).Opponent3Assists = &opponent3.AssistCount
	(*battleUpload).Opponent3GamePaintPoint = &opponent3.GamePaintPoint
	(*battleUpload).Opponent3Specials = &opponent3.SpecialCount
	(*battleUpload).Opponent3Headgear = &opponent3.Player.Head.Id
	(*battleUpload).Opponent3HeadgearMain = &opponent3.Player.HeadSkills.Main.Id

	if len(opponent3.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).Opponent3HeadgearSub0 = &opponent3.Player.HeadSkills.Subs[0].Id

		if len(opponent3.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).Opponent3HeadgearSub1 = &opponent3.Player.HeadSkills.Subs[1].Id

			if len(opponent3.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).Opponent3HeadgearSub2 = &opponent3.Player.HeadSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Opponent3Clothes = &opponent3.Player.Clothes.Id
	(*battleUpload).Opponent3ClothesMain = &opponent3.Player.ClothesSkills.Main.Id

	if len(opponent3.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).Opponent3ClothesSub0 = &opponent3.Player.ClothesSkills.Subs[0].Id

		if len(opponent3.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).Opponent3ClothesSub1 = &opponent3.Player.ClothesSkills.Subs[1].Id

			if len(opponent3.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).Opponent3ClothesSub2 = &opponent3.Player.ClothesSkills.Subs[2].Id
			}
		}
	}

	(*battleUpload).Opponent3Shoes = &opponent3.Player.Shoes.Id
	(*battleUpload).Opponent3ShoesMain = &opponent3.Player.ShoesSkills.Main.Id

	if len(opponent3.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).Opponent3ShoesSub0 = &opponent3.Player.ShoesSkills.Subs[0].Id

		if len(opponent3.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).Opponent3ShoesSub1 = &opponent3.Player.ShoesSkills.Subs[1].Id

			if len(opponent3.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).Opponent3ShoesSub2 = &opponent3.Player.ShoesSkills.Subs[2].Id
			}
		}
	}
}
