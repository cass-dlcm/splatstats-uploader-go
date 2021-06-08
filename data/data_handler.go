package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/cass-dlcm/splatstatsuploader/iksm"
	"github.com/cass-dlcm/splatstatsuploader/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Monitor monitors JSON for changes/new battles/shifts and uploads them.
func Monitor(m int, s bool, salmon bool, apiKey string, version string, appHead map[string]string, client *http.Client) {
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

func uploadLatestBattle(s bool, apiKey string, appHead map[string]string, client *http.Client) {
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

	for key, element := range appHead {
		req.Header.Set(key, element)
	}

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

func uploadSingleBattle(s bool, apiKey string, appHead map[string]string, battleNumber string, client *http.Client) {
	url := "https://app.splatoon2.nintendo.net/api/results/" + battleNumber
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	for key, element := range appHead {
		req.Header.Set(key, element)
	}

	req.AddCookie(&http.Cookie{Name: "iksm_session", Value: viper.GetString("cookie")})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var battle types.Battle
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

func uploadSingleSalmon(s bool, apiKey string, appHead map[string]string, jobId *int, client *http.Client) {
	url := "https://app.splatoon2.nintendo.net/api/coop_results/" + fmt.Sprint(*jobId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	for key, element := range appHead {
		req.Header.Set(key, element)
	}

	req.AddCookie(&http.Cookie{Name: "iksm_session", Value: viper.GetString("cookie")})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var shift types.Shift

	if s {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(errors.Wrap(err, resp.Body.Close().Error()))
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile("two_salmon/"+fmt.Sprint(*jobId)+".json", bodyBytes, 0600); err != nil {
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

func uploadLatestSalmon(s bool, apiKey string, appHead map[string]string, client *http.Client) {
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

	for key, element := range appHead {
		req.Header.Set(key, element)
	}

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

	uploadSingleSalmon(s, apiKey, appHead, data.Results[0].JobID, client)
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

		var shift types.Shift

		var battle types.Battle

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
func GetSplatnetBattle(s bool, apiKey string, version string, appHead map[string]string, client *http.Client) {
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

	for key, element := range appHead {
		req.Header.Set(key, element)
	}

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
func GetSplatnetSalmon(s bool, apiKey string, version string, appHead map[string]string, client *http.Client) {
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

	for key, element := range appHead {
		req.Header.Set(key, element)
	}

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
		uploadSingleSalmon(s, apiKey, appHead, data.Results[i].JobID, client)
	}
}

func setSalmonPayload(shift *types.Shift) types.ShiftUpload {
	shiftUpload := types.ShiftUpload{}
	shiftUpload.SplatnetJSON = *shift
	shiftUpload.SplatnetUpload = true
	shiftUpload.StatInkUpload = false
	shiftUpload.DangerRate = (*shift).DangerRate
	shiftUpload.DrizzlerCount = (*shift).BossCounts.Drizzler.Count
	shiftUpload.FailureWave = (*shift).JobResult.FailureWave
	shiftUpload.JobScore = (*shift).JobScore
	shiftUpload.FlyfishCount = (*shift).BossCounts.Flyfish.Count
	shiftUpload.GoldieCount = (*shift).BossCounts.Goldie.Count
	shiftUpload.GradePoint = (*shift).GradePoint
	shiftUpload.GradePointDelta = *(*shift).GradePointDelta
	shiftUpload.GrillerCount = (*shift).BossCounts.Griller.Count
	shiftUpload.IsClear = (*shift).JobResult.IsClear
	shiftUpload.JobFailureReason = (*shift).JobResult.FailureReason
	shiftUpload.JobID = (*shift).JobID
	shiftUpload.MawsCount = (*shift).BossCounts.Maws.Count
	shiftUpload.PlayerDeathCount = (*shift).MyResult.DeadCount
	shiftUpload.PlayerDrizzlerKills = (*shift).MyResult.BossKillCounts.Drizzler.Count
	shiftUpload.PlayerFlyfishKills = (*shift).MyResult.BossKillCounts.Flyfish.Count
	shiftUpload.PlayerGender = (*shift).MyResult.PlayerType.Gender
	shiftUpload.PlayerGoldenEggs = (*shift).MyResult.GoldenEggs
	shiftUpload.PlayerGoldieKills = (*shift).MyResult.BossKillCounts.Goldie.Count
	shiftUpload.PlayerGrillerKills = (*shift).MyResult.BossKillCounts.Griller.Count
	shiftUpload.PlayerID = (*shift).MyResult.Pid
	shiftUpload.PlayerMawsKills = (*shift).MyResult.BossKillCounts.Maws.Count
	shiftUpload.PlayerName = (*shift).MyResult.Name
	shiftUpload.PlayerPowerEggs = (*shift).MyResult.PowerEggs
	shiftUpload.PlayerReviveCount = (*shift).MyResult.HelpCount
	shiftUpload.PlayerScrapperKills = (*shift).MyResult.BossKillCounts.Scrapper.Count
	shiftUpload.PlayerSpecial = *(*shift).MyResult.Special.ID
	shiftUpload.PlayerSpecies = (*shift).MyResult.PlayerType.Species
	shiftUpload.PlayerSteelEelKills = (*shift).MyResult.BossKillCounts.SteelEel.Count
	shiftUpload.PlayerSteelheadKills = (*shift).MyResult.BossKillCounts.Steelhead.Count
	shiftUpload.PlayerStingerKills = (*shift).MyResult.BossKillCounts.Stinger.Count
	shiftUpload.PlayerTitle = *(*shift).Grade.ID
	salmonPlayerWeaponSpecials(shift, &shiftUpload)
	shiftSetTimes(shift, &shiftUpload)
	shiftUpload.ScheduleWeapon0 = (*shift).Schedule.Weapons[0].ID
	shiftUpload.ScheduleWeapon1 = (*shift).Schedule.Weapons[1].ID
	shiftUpload.ScheduleWeapon2 = (*shift).Schedule.Weapons[2].ID
	shiftUpload.ScheduleWeapon3 = (*shift).Schedule.Weapons[3].ID
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
		shiftUpload.Wave2EventType = (*shift).WaveDetails[1].EventType.Key
		shiftUpload.Wave2GoldenAppear = (*shift).WaveDetails[1].GoldenAppear
		shiftUpload.Wave2GoldenDelivered = (*shift).WaveDetails[1].GoldenEggs
		shiftUpload.Wave2PowerEggs = (*shift).WaveDetails[1].PowerEggs
		shiftUpload.Wave2Quota = (*shift).WaveDetails[1].QuotaNum
		shiftUpload.Wave2WaterLevel = (*shift).WaveDetails[1].WaterLevel.Key

		if len((*shift).WaveDetails) > 2 {
			shiftUpload.Wave3EventType = (*shift).WaveDetails[2].EventType.Key
			shiftUpload.Wave3GoldenAppear = (*shift).WaveDetails[2].GoldenAppear
			shiftUpload.Wave3GoldenDelivered = (*shift).WaveDetails[2].GoldenEggs
			shiftUpload.Wave3PowerEggs = (*shift).WaveDetails[2].PowerEggs
			shiftUpload.Wave3Quota = (*shift).WaveDetails[2].QuotaNum
			shiftUpload.Wave3WaterLevel = (*shift).WaveDetails[2].WaterLevel.Key
		}
	}

	return shiftUpload
}

func UploadSalmon(shiftUpload *types.ShiftUpload, apiKey string, client *http.Client) {
	url := "https://splatstats.cass-dlcm.dev/two_salmon/api/shifts/"
	auth := map[string]string{
		"Authorization": "Token " + apiKey,
		"Content-Type":  "application/json",
	}

	bodyMarshalled, err := json.Marshal(*shiftUpload)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(bodyMarshalled))
	if err != nil {
		panic(err)
	}

	for key, element := range auth {
		req.Header.Add(key, element)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(errors.Wrap(err, resp.Body.Close().Error()))
	}

	if err := resp.Body.Close(); err != nil {
		panic(err)
	}

	bodyString := string(body)
	if resp.StatusCode == 400 && bodyString == "{\"non_field_errors\":[\"The fields player_id, job_id must make a unique set.\"]}" {
		if _, err := fmt.Printf("Shift #%d already uploaded\n", *(*shiftUpload).JobID); err != nil {
			panic(err)
		}
	} else if resp.StatusCode == 201 {
		if _, err := fmt.Printf("Shift #%d uploaded to %s\n", *(*shiftUpload).JobID, resp.Header.Get("location")[0:44]+resp.Header.Get("location")[55:]); err != nil {
			panic(err)
		}
	} else {
		if _, err := fmt.Println(resp.Status); err != nil {
			panic(err)
		}

		if _, err := fmt.Println(bodyString); err != nil {
			panic(err)
		}

		panic(nil)
	}
}

func salmonPlayerWeaponSpecials(shift *types.Shift, shiftUpload *types.ShiftUpload) {
	if len((*shift).MyResult.WeaponList) > 0 {
		(*shiftUpload).PlayerWeaponW1 = *(*shift).MyResult.WeaponList[0].ID

		if len((*shift).MyResult.WeaponList) > 1 {
			(*shiftUpload).PlayerWeaponW2 = *(*shift).MyResult.WeaponList[1].ID

			if len((*shift).MyResult.WeaponList) > 2 {
				(*shiftUpload).PlayerWeaponW3 = *(*shift).MyResult.WeaponList[2].ID
			}
		}
	}

	if len((*shift).MyResult.SpecialCounts) > 0 {
		(*shiftUpload).PlayerW1Specials = (*shift).MyResult.SpecialCounts[0]

		if len((*shift).MyResult.SpecialCounts) > 1 {
			(*shiftUpload).PlayerW2Specials = (*shift).MyResult.SpecialCounts[1]

			if len((*shift).MyResult.SpecialCounts) > 2 {
				(*shiftUpload).PlayerW3Specials = (*shift).MyResult.SpecialCounts[2]
			}
		}
	}
}

func shiftSetTimes(shift *types.Shift, shiftUpload *types.ShiftUpload) {
	if (*shift).PlayTime != nil {
		(*shiftUpload).Playtime = time.Unix(int64(*(*shift).PlayTime), 0).Format("2006-01-02 15:04:05")
	}

	if (*shift).EndTime != nil {
		(*shiftUpload).Endtime = time.Unix(int64(*(*shift).EndTime), 0).Format("2006-01-02 15:04:05")
	}

	if (*shift).Schedule.EndTime != nil {
		(*shiftUpload).ScheduleEndtime = time.Unix(int64(*(*shift).Schedule.EndTime), 0).Format("2006-01-02 15:04:05")
	}

	if (*shift).Schedule.StartTime != nil {
		(*shiftUpload).ScheduleStarttime = time.Unix(int64(*(*shift).Schedule.StartTime), 0).Format("2006-01-02 15:04:05")
	}

	if (*shift).StartTime != nil {
		(*shiftUpload).Starttime = time.Unix(int64(*(*shift).StartTime), 0).Format("2006-01-02 15:04:05")
	}
}

func shiftSetTeammate0(shift *types.Shift, shiftUpload *types.ShiftUpload) {
	if len((*shift).OtherResults) > 0 {
		(*shiftUpload).Teammate0DeathCount = (*shift).OtherResults[0].DeadCount
		(*shiftUpload).Teammate0DrizzlerKills = (*shift).OtherResults[0].BossKillCounts.Drizzler.Count
		(*shiftUpload).Teammate0FlyfishKills = (*shift).OtherResults[0].BossKillCounts.Flyfish.Count
		(*shiftUpload).Teammate0Gender = (*shift).OtherResults[0].PlayerType.Gender
		(*shiftUpload).Teammate0GoldenEggs = (*shift).OtherResults[0].GoldenEggs
		(*shiftUpload).Teammate0GoldieKills = (*shift).OtherResults[0].BossKillCounts.Goldie.Count
		(*shiftUpload).Teammate0GrillerKills = (*shift).OtherResults[0].BossKillCounts.Griller.Count
		(*shiftUpload).Teammate0ID = (*shift).OtherResults[0].Pid
		(*shiftUpload).Teammate0MawsKills = (*shift).OtherResults[0].BossKillCounts.Maws.Count
		(*shiftUpload).Teammate0Name = (*shift).OtherResults[0].Name
		(*shiftUpload).Teammate0PowerEggs = (*shift).OtherResults[0].PowerEggs
		(*shiftUpload).Teammate0ReviveCount = (*shift).OtherResults[0].HelpCount
		(*shiftUpload).Teammate0ScrapperKills = (*shift).OtherResults[0].BossKillCounts.Scrapper.Count
		(*shiftUpload).Teammate0Special = *(*shift).OtherResults[0].Special.ID
		(*shiftUpload).Teammate0Species = *(*shift).OtherResults[0].PlayerType.Species
		(*shiftUpload).Teammate0SteelEelKills = (*shift).OtherResults[0].BossKillCounts.SteelEel.Count
		(*shiftUpload).Teammate0SteelheadKills = (*shift).OtherResults[0].BossKillCounts.Steelhead.Count
		(*shiftUpload).Teammate0StingerKills = (*shift).OtherResults[0].BossKillCounts.Stinger.Count
		(*shiftUpload).Teammate0W1Specials = (*shift).OtherResults[0].SpecialCounts[0]

		if len((*shift).OtherResults[0].WeaponList) > 0 {
			(*shiftUpload).Teammate0WeaponW1 = (*shift).OtherResults[0].WeaponList[0].ID

			if len((*shift).OtherResults[0].WeaponList) > 1 {
				(*shiftUpload).Teammate0WeaponW2 = (*shift).OtherResults[0].WeaponList[1].ID

				if len((*shift).OtherResults[0].WeaponList) > 2 {
					shiftUpload.Teammate0WeaponW3 = (*shift).OtherResults[0].WeaponList[2].ID
				}
			}
		}

		if len((*shift).OtherResults[0].SpecialCounts) > 1 {
			(*shiftUpload).Teammate0W2Specials = (*shift).OtherResults[0].SpecialCounts[1]

			if len((*shift).OtherResults[0].SpecialCounts) > 2 {
				(*shiftUpload).Teammate0W3Specials = (*shift).OtherResults[0].SpecialCounts[2]
			}
		}
	}
}

func shiftSetTeammate1(shift *types.Shift, shiftUpload *types.ShiftUpload) {
	if len((*shift).OtherResults) > 1 {
		(*shiftUpload).Teammate1DeathCount = (*shift).OtherResults[1].DeadCount
		(*shiftUpload).Teammate1DrizzlerKills = (*shift).OtherResults[1].BossKillCounts.Drizzler.Count
		(*shiftUpload).Teammate1FlyfishKills = (*shift).OtherResults[1].BossKillCounts.Flyfish.Count
		(*shiftUpload).Teammate1Gender = (*shift).OtherResults[1].PlayerType.Gender
		(*shiftUpload).Teammate1GoldenEggs = (*shift).OtherResults[1].GoldenEggs
		(*shiftUpload).Teammate1GoldieKills = (*shift).OtherResults[1].BossKillCounts.Goldie.Count
		(*shiftUpload).Teammate1GrillerKills = (*shift).OtherResults[1].BossKillCounts.Griller.Count
		(*shiftUpload).Teammate1ID = (*shift).OtherResults[1].Pid
		(*shiftUpload).Teammate1MawsKills = (*shift).OtherResults[1].BossKillCounts.Maws.Count
		(*shiftUpload).Teammate1Name = (*shift).OtherResults[1].Name
		(*shiftUpload).Teammate1PowerEggs = (*shift).OtherResults[1].PowerEggs
		(*shiftUpload).Teammate1ReviveCount = (*shift).OtherResults[1].HelpCount
		(*shiftUpload).Teammate1ScrapperKills = (*shift).OtherResults[1].BossKillCounts.Scrapper.Count
		(*shiftUpload).Teammate1Special = *(*shift).OtherResults[1].Special.ID
		(*shiftUpload).Teammate1Species = *(*shift).OtherResults[1].PlayerType.Species
		(*shiftUpload).Teammate1SteelEelKills = (*shift).OtherResults[1].BossKillCounts.SteelEel.Count
		(*shiftUpload).Teammate1SteelheadKills = (*shift).OtherResults[1].BossKillCounts.Steelhead.Count
		(*shiftUpload).Teammate1StingerKills = (*shift).OtherResults[1].BossKillCounts.Stinger.Count
		(*shiftUpload).Teammate1W1Specials = (*shift).OtherResults[1].SpecialCounts[0]

		if len((*shift).OtherResults[1].SpecialCounts) > 1 {
			(*shiftUpload).Teammate1W2Specials = (*shift).OtherResults[1].SpecialCounts[1]

			if len((*shift).OtherResults[1].SpecialCounts) > 2 {
				shiftUpload.Teammate1W3Specials = (*shift).OtherResults[1].SpecialCounts[2]
			}
		}

		if len((*shift).OtherResults[1].WeaponList) > 0 {
			(*shiftUpload).Teammate1WeaponW1 = (*shift).OtherResults[1].WeaponList[0].ID

			if len((*shift).OtherResults[1].WeaponList) > 1 {
				(*shiftUpload).Teammate1WeaponW2 = (*shift).OtherResults[1].WeaponList[1].ID

				if len((*shift).OtherResults[1].WeaponList) > 2 {
					(*shiftUpload).Teammate1WeaponW3 = (*shift).OtherResults[1].WeaponList[2].ID
				}
			}
		}
	}
}

func shiftSetTeammate2(shift *types.Shift, shiftUpload *types.ShiftUpload) {
	if len((*shift).OtherResults) > 2 {
		(*shiftUpload).Teammate2DeathCount = (*shift).OtherResults[2].DeadCount
		(*shiftUpload).Teammate2DrizzlerKills = (*shift).OtherResults[2].BossKillCounts.Drizzler.Count
		(*shiftUpload).Teammate2FlyfishKills = (*shift).OtherResults[2].BossKillCounts.Flyfish.Count
		(*shiftUpload).Teammate2Gender = (*shift).OtherResults[2].PlayerType.Gender
		(*shiftUpload).Teammate2GoldenEggs = (*shift).OtherResults[2].GoldenEggs
		(*shiftUpload).Teammate2GoldieKills = (*shift).OtherResults[2].BossKillCounts.Goldie.Count
		(*shiftUpload).Teammate2GrillerKills = (*shift).OtherResults[2].BossKillCounts.Griller.Count
		(*shiftUpload).Teammate2ID = (*shift).OtherResults[2].Pid
		(*shiftUpload).Teammate2MawsKills = (*shift).OtherResults[2].BossKillCounts.Maws.Count
		(*shiftUpload).Teammate2Name = (*shift).OtherResults[2].Name
		(*shiftUpload).Teammate2PowerEggs = (*shift).OtherResults[2].PowerEggs
		(*shiftUpload).Teammate2ReviveCount = (*shift).OtherResults[2].HelpCount
		(*shiftUpload).Teammate2ScrapperKills = (*shift).OtherResults[2].BossKillCounts.Scrapper.Count
		(*shiftUpload).Teammate2Special = *(*shift).OtherResults[2].Special.ID
		(*shiftUpload).Teammate2Species = *(*shift).OtherResults[2].PlayerType.Species
		(*shiftUpload).Teammate2SteelEelKills = (*shift).OtherResults[2].BossKillCounts.SteelEel.Count
		(*shiftUpload).Teammate2SteelheadKills = (*shift).OtherResults[2].BossKillCounts.Steelhead.Count
		(*shiftUpload).Teammate2StingerKills = (*shift).OtherResults[2].BossKillCounts.Stinger.Count
		(*shiftUpload).Teammate2W1Specials = (*shift).OtherResults[2].SpecialCounts[0]

		if len((*shift).OtherResults[2].SpecialCounts) > 1 {
			(*shiftUpload).Teammate2W2Specials = (*shift).OtherResults[2].SpecialCounts[1]

			if len((*shift).OtherResults[2].SpecialCounts) > 2 {
				(*shiftUpload).Teammate2W3Specials = (*shift).OtherResults[2].SpecialCounts[2]
			}
		}

		if len((*shift).OtherResults[2].WeaponList) > 0 {
			(*shiftUpload).Teammate2WeaponW1 = (*shift).OtherResults[2].WeaponList[0].ID

			if len((*shift).OtherResults[2].WeaponList) > 1 {
				(*shiftUpload).Teammate2WeaponW2 = (*shift).OtherResults[2].WeaponList[1].ID

				if len((*shift).OtherResults[2].WeaponList) > 2 {
					(*shiftUpload).Teammate2WeaponW3 = (*shift).OtherResults[2].WeaponList[2].ID
				}
			}
		}
	}
}

func setBattlePayload(battle *types.Battle) types.BattleUpload {
	battleUpload := types.BattleUpload{}
	battleUpload.SplatnetJSON = battle
	battleUpload.SplatnetUpload = true
	battleUpload.StatInkUpload = false
	battleUpload.BattleNumber = (*battle).BattleNumber
	battleUpload.Rule = (*battle).Rule.Key
	battleUpload.MatchType = (*battle).GameMode.Key
	battleUpload.Stage = (*battle).Stage.ID
	battleUpload.Win = *(*battle).MyTeamResult.Key == "victory"
	battleSetHasDc(battle, &battleUpload)
	battleUpload.Time = (*battle).StartTime
	battleUpload.WinMeter = (*battle).WinMeter
	battleSetScoreTime(battle, &battleUpload)
	battleUpload.TagID = (*battle).TagID
	battleUpload.LeaguePoint = *(*battle).LeaguePoint
	battleUpload.SplatfestPoint = 0
	battleUpload.SplatfestTitleAfter = ""

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

func UploadBattle(battleUpload *types.BattleUpload, apiKey string, client *http.Client) {
	url := "https://splatstats.cass-dlcm.dev/two_battles/api/battles/"
	auth := map[string]string{
		"Authorization": "Token " + apiKey,
		"Content-Type":  "application/json",
	}

	if (*battleUpload).PlayerSplatnetID == "" || (*battleUpload).BattleNumber == "" {
		fmt.Println("Skipping battle due to missing data.")
		return
	}

	bodyMarshalled, err := json.Marshal(*battleUpload)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(bodyMarshalled))
	if err != nil {
		panic(err)
	}

	for key, element := range auth {
		req.Header.Add(key, element)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if bodyString := string(body); resp.StatusCode == 400 && bodyString == "{\"non_field_errors\":[\"The fields player_splatnet_id, battle_number must make a unique set.\"]}" {
		if _, err := fmt.Printf("Battle #%s already uploaded\n", (*battleUpload).BattleNumber); err != nil {
			panic(err)
		}
	} else if resp.StatusCode == 201 {
		if _, err := fmt.Printf("Battle #%s uploaded to %s\n", (*battleUpload).BattleNumber, resp.Header.Get("location")[0:45]+resp.Header.Get("location")[57:]); err != nil {
			panic(err)
		}
	} else {
		if _, err := fmt.Println(resp.Status); err != nil {
			panic(err)
		}
		if _, err := fmt.Println(bodyString); err != nil {
			panic(err)
		}
		panic(nil)
	}
}

func battleSetHasDc(battle *types.Battle, battleUpload *types.BattleUpload) {
	(*battleUpload).HasDisconnectedPlayer = false

	for i := range (*battle).MyTeamMembers {
		(*battleUpload).HasDisconnectedPlayer = (*battleUpload).HasDisconnectedPlayer || (*(*battle).MyTeamMembers[i].GamePaintPoint == 0 && *(*battle).MyTeamMembers[i].KillCount == 0 && *(*battle).MyTeamMembers[i].SpecialCount == 0 && *(*battle).MyTeamMembers[i].DeathCount == 0 && *(*battle).MyTeamMembers[i].AssistCount == 0)
	}

	for i := range (*battle).OtherTeamMembers {
		(*battleUpload).HasDisconnectedPlayer = (*battleUpload).HasDisconnectedPlayer || (*(*battle).OtherTeamMembers[i].GamePaintPoint == 0 && *(*battle).OtherTeamMembers[i].KillCount == 0 && *(*battle).OtherTeamMembers[i].SpecialCount == 0 && *(*battle).OtherTeamMembers[i].DeathCount == 0 && *(*battle).OtherTeamMembers[i].AssistCount == 0)
	}
}

func battleSetScoreTime(battle *types.Battle, battleUpload *types.BattleUpload) {
	if (*battle).MyTeamCount != nil {
		(*battleUpload).MyTeamCount = float64(*(*battle).MyTeamCount)
	} else {
		(*battleUpload).MyTeamCount = (*battle).MyTeamPercentage
	}

	if (*battle).OtherTeamCount != nil {
		(*battleUpload).OtherTeamCount = float64(*(*battle).OtherTeamCount)
	} else {
		(*battleUpload).OtherTeamCount = (*battle).OtherTeamPercentage
	}

	if battleUpload.Rule == "turf_war" {
		(*battleUpload).ElapsedTime = 180
	} else {
		(*battleUpload).ElapsedTime = (*battle).ElapsedTime
	}
}

func battleSetPlayer(battle *types.Battle, battleUpload *types.BattleUpload) {
	(*battleUpload).PlayerSplatnetID = *(*battle).PlayerResult.Player.PrincipalID
	(*battleUpload).PlayerName = *(*battle).PlayerResult.Player.Nickname
	(*battleUpload).PlayerWeapon = *(*battle).PlayerResult.Player.Weapon.ID
	(*battleUpload).PlayerRank = *(*battle).Udemae.Number
	(*battleUpload).PlayerSplatfestTitle = ""
	(*battleUpload).PlayerLevelStar = *(*battle).StarRank
	(*battleUpload).PlayerLevel = *(*battle).PlayerRank
	(*battleUpload).PlayerKills = *(*battle).PlayerResult.KillCount
	(*battleUpload).PlayerDeaths = *(*battle).PlayerResult.DeathCount
	(*battleUpload).PlayerAssists = *(*battle).PlayerResult.AssistCount
	(*battleUpload).PlayerSpecials = *(*battle).PlayerResult.SpecialCount
	(*battleUpload).PlayerGamePaintPoint = *(*battle).PlayerResult.GamePaintPoint
	(*battleUpload).PlayerGender = *(*battle).PlayerResult.Player.PlayerType.Gender
	(*battleUpload).PlayerSpecies = (*battle).PlayerResult.Player.PlayerType.Species
	(*battleUpload).PlayerXPower = (*battle).XPower
	(*battleUpload).PlayerHeadgear = *(*battle).PlayerResult.Player.Head.ID
	(*battleUpload).PlayerHeadgearMain = *(*battle).PlayerResult.Player.HeadSkills.Main.ID

	if len((*battle).PlayerResult.Player.HeadSkills.Subs) > 0 {
		(*battleUpload).PlayerHeadgearSub0 = *(*battle).PlayerResult.Player.HeadSkills.Subs[0].ID

		if len((*battle).PlayerResult.Player.HeadSkills.Subs) > 1 {
			(*battleUpload).PlayerHeadgearSub1 = *(*battle).PlayerResult.Player.HeadSkills.Subs[1].ID

			if len((*battle).PlayerResult.Player.HeadSkills.Subs) > 2 {
				(*battleUpload).PlayerHeadgearSub2 = *(*battle).PlayerResult.Player.HeadSkills.Subs[2].ID
			}
		}
	}

	(*battleUpload).PlayerClothes = *(*battle).PlayerResult.Player.Clothes.ID
	(*battleUpload).PlayerClothesMain = *(*battle).PlayerResult.Player.ClothesSkills.Main.ID

	if len((*battle).PlayerResult.Player.ClothesSkills.Subs) > 0 {
		(*battleUpload).PlayerClothesSub0 = *(*battle).PlayerResult.Player.ClothesSkills.Subs[0].ID

		if len((*battle).PlayerResult.Player.ClothesSkills.Subs) > 1 {
			(*battleUpload).PlayerClothesSub1 = *(*battle).PlayerResult.Player.ClothesSkills.Subs[1].ID

			if len((*battle).PlayerResult.Player.ClothesSkills.Subs) > 2 {
				(*battleUpload).PlayerClothesSub2 = *(*battle).PlayerResult.Player.ClothesSkills.Subs[2].ID
			}
		}
	}

	(*battleUpload).PlayerShoes = *(*battle).PlayerResult.Player.Shoes.ID
	(*battleUpload).PlayerShoesMain = *(*battle).PlayerResult.Player.ShoesSkills.Main.ID

	if len((*battle).PlayerResult.Player.ShoesSkills.Subs) > 0 {
		(*battleUpload).PlayerShoesSub0 = *(*battle).PlayerResult.Player.ShoesSkills.Subs[0].ID

		if len((*battle).PlayerResult.Player.ShoesSkills.Subs) > 1 {
			(*battleUpload).PlayerShoesSub1 = *(*battle).PlayerResult.Player.ShoesSkills.Subs[1].ID

			if len((*battle).PlayerResult.Player.ShoesSkills.Subs) > 2 {
				(*battleUpload).PlayerShoesSub2 = *(*battle).PlayerResult.Player.ShoesSkills.Subs[2].ID
			}
		}
	}
}

func battleSetTeammate0(battle *types.Battle, battleUpload *types.BattleUpload) {
	if len((*battle).MyTeamMembers) > 0 {
		teammate0 := (*battle).MyTeamMembers[0]
		(*battleUpload).Teammate0SplatnetID = *teammate0.Player.PrincipalID
		(*battleUpload).Teammate0Name = *teammate0.Player.Nickname
		(*battleUpload).Teammate0LevelStar = *teammate0.Player.StarRank
		(*battleUpload).Teammate0Level = *teammate0.Player.PlayerRank
		(*battleUpload).Teammate0Rank = *teammate0.Player.Udemae.Name
		(*battleUpload).Teammate0Weapon = *teammate0.Player.Weapon.ID
		(*battleUpload).Teammate0Gender = *teammate0.Player.PlayerType.Gender
		(*battleUpload).Teammate0Species = *teammate0.Player.PlayerType.Species
		(*battleUpload).Teammate0Kills = *teammate0.KillCount
		(*battleUpload).Teammate0Deaths = *teammate0.DeathCount
		(*battleUpload).Teammate0Assists = *teammate0.AssistCount
		(*battleUpload).Teammate0GamePaintPoint = *teammate0.GamePaintPoint
		(*battleUpload).Teammate0Specials = *teammate0.SpecialCount
		(*battleUpload).Teammate0Headgear = teammate0.Player.Head.ID
		(*battleUpload).Teammate0HeadgearMain = teammate0.Player.HeadSkills.Main.ID

		if len(teammate0.Player.HeadSkills.Subs) > 0 {
			(*battleUpload).Teammate0HeadgearSub0 = teammate0.Player.HeadSkills.Subs[0].ID

			if len(teammate0.Player.HeadSkills.Subs) > 1 {
				(*battleUpload).Teammate0HeadgearSub1 = teammate0.Player.HeadSkills.Subs[1].ID

				if len(teammate0.Player.HeadSkills.Subs) > 2 {
					(*battleUpload).Teammate0HeadgearSub2 = teammate0.Player.HeadSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Teammate0Clothes = teammate0.Player.Clothes.ID
		(*battleUpload).Teammate0ClothesMain = teammate0.Player.ClothesSkills.Main.ID

		if len(teammate0.Player.ClothesSkills.Subs) > 0 {
			(*battleUpload).Teammate0ClothesSub0 = teammate0.Player.ClothesSkills.Subs[0].ID

			if len(teammate0.Player.ClothesSkills.Subs) > 1 {
				(*battleUpload).Teammate0ClothesSub1 = teammate0.Player.ClothesSkills.Subs[1].ID

				if len(teammate0.Player.ClothesSkills.Subs) > 2 {
					(*battleUpload).Teammate0ClothesSub2 = teammate0.Player.ClothesSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Teammate0Shoes = teammate0.Player.Shoes.ID
		(*battleUpload).Teammate0ShoesMain = teammate0.Player.ShoesSkills.Main.ID

		if len(teammate0.Player.ShoesSkills.Subs) > 0 {
			(*battleUpload).Teammate0ShoesSub0 = teammate0.Player.ShoesSkills.Subs[0].ID

			if len(teammate0.Player.ShoesSkills.Subs) > 1 {
				(*battleUpload).Teammate0ShoesSub1 = teammate0.Player.ShoesSkills.Subs[1].ID

				if len(teammate0.Player.ShoesSkills.Subs) > 2 {
					(*battleUpload).Teammate0ShoesSub2 = teammate0.Player.ShoesSkills.Subs[2].ID
				}
			}
		}
	}
}

func battleSetTeammate1(battle *types.Battle, battleUpload *types.BattleUpload) {
	if len((*battle).MyTeamMembers) > 1 {
		teammate1 := (*battle).MyTeamMembers[1]
		(*battleUpload).Teammate1SplatnetID = *teammate1.Player.PrincipalID
		(*battleUpload).Teammate1Name = *teammate1.Player.Nickname
		(*battleUpload).Teammate1LevelStar = *teammate1.Player.StarRank
		(*battleUpload).Teammate1Level = *teammate1.Player.PlayerRank
		(*battleUpload).Teammate1Rank = *teammate1.Player.Udemae.Name
		(*battleUpload).Teammate1Weapon = *teammate1.Player.Weapon.ID
		(*battleUpload).Teammate1Gender = *teammate1.Player.PlayerType.Gender
		(*battleUpload).Teammate1Species = *teammate1.Player.PlayerType.Species
		(*battleUpload).Teammate1Kills = *teammate1.KillCount
		(*battleUpload).Teammate1Deaths = *teammate1.DeathCount
		(*battleUpload).Teammate1Assists = *teammate1.AssistCount
		(*battleUpload).Teammate1GamePaintPoint = *teammate1.GamePaintPoint
		(*battleUpload).Teammate1Specials = *teammate1.SpecialCount
		(*battleUpload).Teammate1Headgear = teammate1.Player.Head.ID
		(*battleUpload).Teammate1HeadgearMain = teammate1.Player.HeadSkills.Main.ID

		if len(teammate1.Player.HeadSkills.Subs) > 0 {
			(*battleUpload).Teammate1HeadgearSub0 = teammate1.Player.HeadSkills.Subs[0].ID

			if len(teammate1.Player.HeadSkills.Subs) > 1 {
				(*battleUpload).Teammate1HeadgearSub1 = teammate1.Player.HeadSkills.Subs[1].ID

				if len(teammate1.Player.HeadSkills.Subs) > 2 {
					(*battleUpload).Teammate1HeadgearSub2 = teammate1.Player.HeadSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Teammate1Clothes = teammate1.Player.Clothes.ID
		(*battleUpload).Teammate1ClothesMain = teammate1.Player.ClothesSkills.Main.ID

		if len(teammate1.Player.ClothesSkills.Subs) > 0 {
			(*battleUpload).Teammate1ClothesSub0 = teammate1.Player.ClothesSkills.Subs[0].ID

			if len(teammate1.Player.ClothesSkills.Subs) > 1 {
				(*battleUpload).Teammate1ClothesSub1 = teammate1.Player.ClothesSkills.Subs[1].ID

				if len(teammate1.Player.ClothesSkills.Subs) > 2 {
					(*battleUpload).Teammate1ClothesSub2 = teammate1.Player.ClothesSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Teammate1Shoes = teammate1.Player.Shoes.ID
		(*battleUpload).Teammate1ShoesMain = teammate1.Player.ShoesSkills.Main.ID

		if len(teammate1.Player.ShoesSkills.Subs) > 0 {
			(*battleUpload).Teammate1ShoesSub0 = teammate1.Player.ShoesSkills.Subs[0].ID

			if len(teammate1.Player.ShoesSkills.Subs) > 1 {
				(*battleUpload).Teammate1ShoesSub1 = teammate1.Player.ShoesSkills.Subs[1].ID

				if len(teammate1.Player.ShoesSkills.Subs) > 2 {
					(*battleUpload).Teammate1ShoesSub2 = teammate1.Player.ShoesSkills.Subs[2].ID
				}
			}
		}
	}
}

func battleSetTeammate2(battle *types.Battle, battleUpload *types.BattleUpload) {
	if len((*battle).MyTeamMembers) > 2 {
		teammate2 := (*battle).MyTeamMembers[2]
		(*battleUpload).Teammate2SplatnetID = *teammate2.Player.PrincipalID
		(*battleUpload).Teammate2Name = *teammate2.Player.Nickname
		(*battleUpload).Teammate2LevelStar = *teammate2.Player.StarRank
		(*battleUpload).Teammate2Level = *teammate2.Player.PlayerRank
		(*battleUpload).Teammate2Rank = *teammate2.Player.Udemae.Name
		(*battleUpload).Teammate2Weapon = *teammate2.Player.Weapon.ID
		(*battleUpload).Teammate2Gender = *teammate2.Player.PlayerType.Gender
		(*battleUpload).Teammate2Species = *teammate2.Player.PlayerType.Species
		(*battleUpload).Teammate2Kills = *teammate2.KillCount
		(*battleUpload).Teammate2Deaths = *teammate2.DeathCount
		(*battleUpload).Teammate2Assists = *teammate2.AssistCount
		(*battleUpload).Teammate2GamePaintPoint = *teammate2.GamePaintPoint
		(*battleUpload).Teammate2Specials = *teammate2.SpecialCount
		(*battleUpload).Teammate2Headgear = teammate2.Player.Head.ID
		(*battleUpload).Teammate2HeadgearMain = teammate2.Player.HeadSkills.Main.ID

		if len(teammate2.Player.HeadSkills.Subs) > 0 {
			(*battleUpload).Teammate2HeadgearSub0 = teammate2.Player.HeadSkills.Subs[0].ID

			if len(teammate2.Player.HeadSkills.Subs) > 1 {
				(*battleUpload).Teammate2HeadgearSub1 = teammate2.Player.HeadSkills.Subs[1].ID

				if len(teammate2.Player.HeadSkills.Subs) > 2 {
					(*battleUpload).Teammate2HeadgearSub2 = teammate2.Player.HeadSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Teammate2Clothes = teammate2.Player.Clothes.ID
		(*battleUpload).Teammate2ClothesMain = teammate2.Player.ClothesSkills.Main.ID

		if len(teammate2.Player.ClothesSkills.Subs) > 0 {
			(*battleUpload).Teammate2ClothesSub0 = teammate2.Player.ClothesSkills.Subs[0].ID

			if len(teammate2.Player.ClothesSkills.Subs) > 1 {
				(*battleUpload).Teammate2ClothesSub1 = teammate2.Player.ClothesSkills.Subs[1].ID

				if len(teammate2.Player.ClothesSkills.Subs) > 2 {
					(*battleUpload).Teammate2ClothesSub2 = teammate2.Player.ClothesSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Teammate2Shoes = teammate2.Player.Shoes.ID
		(*battleUpload).Teammate2ShoesMain = teammate2.Player.ShoesSkills.Main.ID

		if len(teammate2.Player.ShoesSkills.Subs) > 0 {
			(*battleUpload).Teammate2ShoesSub0 = teammate2.Player.ShoesSkills.Subs[0].ID

			if len(teammate2.Player.ShoesSkills.Subs) > 1 {
				(*battleUpload).Teammate2ShoesSub1 = teammate2.Player.ShoesSkills.Subs[1].ID

				if len(teammate2.Player.ShoesSkills.Subs) > 2 {
					(*battleUpload).Teammate2ShoesSub2 = teammate2.Player.ShoesSkills.Subs[2].ID
				}
			}
		}
	}
}

func battleSetOpponent0(battle *types.Battle, battleUpload *types.BattleUpload) {
	if len((*battle).OtherTeamMembers) > 0 {
		opponent0 := (*battle).OtherTeamMembers[0]
		(*battleUpload).Opponent0SplatnetID = *opponent0.Player.PrincipalID
		(*battleUpload).Opponent0Name = *opponent0.Player.Nickname
		(*battleUpload).Opponent0LevelStar = *opponent0.Player.StarRank
		(*battleUpload).Opponent0Level = *opponent0.Player.PlayerRank
		(*battleUpload).Opponent0Rank = *opponent0.Player.Udemae.Name
		(*battleUpload).Opponent0Weapon = *opponent0.Player.Weapon.ID
		(*battleUpload).Opponent0Gender = *opponent0.Player.PlayerType.Gender
		(*battleUpload).Opponent0Species = *opponent0.Player.PlayerType.Species
		(*battleUpload).Opponent0Kills = *opponent0.KillCount
		(*battleUpload).Opponent0Deaths = *opponent0.DeathCount
		(*battleUpload).Opponent0Assists = *opponent0.AssistCount
		(*battleUpload).Opponent0GamePaintPoint = *opponent0.GamePaintPoint
		(*battleUpload).Opponent0Specials = *opponent0.SpecialCount
		(*battleUpload).Opponent0Headgear = opponent0.Player.Head.ID
		(*battleUpload).Opponent0HeadgearMain = opponent0.Player.HeadSkills.Main.ID

		if len(opponent0.Player.HeadSkills.Subs) > 0 {
			(*battleUpload).Opponent0HeadgearSub0 = opponent0.Player.HeadSkills.Subs[0].ID

			if len(opponent0.Player.HeadSkills.Subs) > 1 {
				(*battleUpload).Opponent0HeadgearSub1 = opponent0.Player.HeadSkills.Subs[1].ID

				if len(opponent0.Player.HeadSkills.Subs) > 2 {
					(*battleUpload).Opponent0HeadgearSub2 = opponent0.Player.HeadSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Opponent0Clothes = opponent0.Player.Clothes.ID
		(*battleUpload).Opponent0ClothesMain = opponent0.Player.ClothesSkills.Main.ID

		if len(opponent0.Player.ClothesSkills.Subs) > 0 {
			(*battleUpload).Opponent0ClothesSub0 = opponent0.Player.ClothesSkills.Subs[0].ID

			if len(opponent0.Player.ClothesSkills.Subs) > 1 {
				(*battleUpload).Opponent0ClothesSub1 = opponent0.Player.ClothesSkills.Subs[1].ID

				if len(opponent0.Player.ClothesSkills.Subs) > 2 {
					(*battleUpload).Opponent0ClothesSub2 = opponent0.Player.ClothesSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Opponent0Shoes = opponent0.Player.Shoes.ID
		(*battleUpload).Opponent0ShoesMain = opponent0.Player.ShoesSkills.Main.ID

		if len(opponent0.Player.ShoesSkills.Subs) > 0 {
			(*battleUpload).Opponent0ShoesSub0 = opponent0.Player.ShoesSkills.Subs[0].ID

			if len(opponent0.Player.ShoesSkills.Subs) > 1 {
				(*battleUpload).Opponent0ShoesSub1 = opponent0.Player.ShoesSkills.Subs[1].ID

				if len(opponent0.Player.ShoesSkills.Subs) > 2 {
					(*battleUpload).Opponent0ShoesSub2 = opponent0.Player.ShoesSkills.Subs[2].ID
				}
			}
		}
	}
}

func battleSetOpponent1(battle *types.Battle, battleUpload *types.BattleUpload) {
	if len((*battle).OtherTeamMembers) > 1 {
		opponent1 := (*battle).OtherTeamMembers[1]
		(*battleUpload).Opponent1SplatnetID = *opponent1.Player.PrincipalID
		(*battleUpload).Opponent1Name = *opponent1.Player.Nickname
		(*battleUpload).Opponent1LevelStar = *opponent1.Player.StarRank
		(*battleUpload).Opponent1Level = *opponent1.Player.PlayerRank
		(*battleUpload).Opponent1Rank = *opponent1.Player.Udemae.Name
		(*battleUpload).Opponent1Weapon = *opponent1.Player.Weapon.ID
		(*battleUpload).Opponent1Gender = *opponent1.Player.PlayerType.Gender
		(*battleUpload).Opponent1Species = *opponent1.Player.PlayerType.Species
		(*battleUpload).Opponent1Kills = *opponent1.KillCount
		(*battleUpload).Opponent1Deaths = *opponent1.DeathCount
		(*battleUpload).Opponent1Assists = *opponent1.AssistCount
		(*battleUpload).Opponent1GamePaintPoint = *opponent1.GamePaintPoint
		(*battleUpload).Opponent1Specials = *opponent1.SpecialCount
		(*battleUpload).Opponent1Headgear = opponent1.Player.Head.ID
		(*battleUpload).Opponent1HeadgearMain = opponent1.Player.HeadSkills.Main.ID

		if len(opponent1.Player.HeadSkills.Subs) > 0 {
			(*battleUpload).Opponent1HeadgearSub0 = opponent1.Player.HeadSkills.Subs[0].ID

			if len(opponent1.Player.HeadSkills.Subs) > 1 {
				(*battleUpload).Opponent1HeadgearSub1 = opponent1.Player.HeadSkills.Subs[1].ID

				if len(opponent1.Player.HeadSkills.Subs) > 2 {
					(*battleUpload).Opponent1HeadgearSub2 = opponent1.Player.HeadSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Opponent1Clothes = opponent1.Player.Clothes.ID
		(*battleUpload).Opponent1ClothesMain = opponent1.Player.ClothesSkills.Main.ID

		if len(opponent1.Player.ClothesSkills.Subs) > 0 {
			(*battleUpload).Opponent1ClothesSub0 = opponent1.Player.ClothesSkills.Subs[0].ID

			if len(opponent1.Player.ClothesSkills.Subs) > 1 {
				(*battleUpload).Opponent1ClothesSub1 = opponent1.Player.ClothesSkills.Subs[1].ID

				if len(opponent1.Player.ClothesSkills.Subs) > 2 {
					(*battleUpload).Opponent1ClothesSub2 = opponent1.Player.ClothesSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Opponent1Shoes = opponent1.Player.Shoes.ID
		(*battleUpload).Opponent1ShoesMain = opponent1.Player.ShoesSkills.Main.ID

		if len(opponent1.Player.ShoesSkills.Subs) > 0 {
			(*battleUpload).Opponent1ShoesSub0 = opponent1.Player.ShoesSkills.Subs[0].ID

			if len(opponent1.Player.ShoesSkills.Subs) > 1 {
				(*battleUpload).Opponent1ShoesSub1 = opponent1.Player.ShoesSkills.Subs[1].ID

				if len(opponent1.Player.ShoesSkills.Subs) > 2 {
					(*battleUpload).Opponent1ShoesSub2 = opponent1.Player.ShoesSkills.Subs[2].ID
				}
			}
		}
	}
}

func battleSetOpponent2(battle *types.Battle, battleUpload *types.BattleUpload) {
	if len((*battle).OtherTeamMembers) > 2 {
		opponent2 := (*battle).OtherTeamMembers[2]
		(*battleUpload).Opponent2SplatnetID = *opponent2.Player.PrincipalID
		(*battleUpload).Opponent2Name = *opponent2.Player.Nickname
		(*battleUpload).Opponent2LevelStar = *opponent2.Player.StarRank
		(*battleUpload).Opponent2Level = *opponent2.Player.PlayerRank
		(*battleUpload).Opponent2Rank = *opponent2.Player.Udemae.Name
		(*battleUpload).Opponent2Weapon = *opponent2.Player.Weapon.ID
		(*battleUpload).Opponent2Gender = *opponent2.Player.PlayerType.Gender
		(*battleUpload).Opponent2Species = *opponent2.Player.PlayerType.Species
		(*battleUpload).Opponent2Kills = *opponent2.KillCount
		(*battleUpload).Opponent2Deaths = *opponent2.DeathCount
		(*battleUpload).Opponent2Assists = *opponent2.AssistCount
		(*battleUpload).Opponent2GamePaintPoint = *opponent2.GamePaintPoint
		(*battleUpload).Opponent2Specials = *opponent2.SpecialCount
		(*battleUpload).Opponent2Headgear = opponent2.Player.Head.ID
		(*battleUpload).Opponent2HeadgearMain = opponent2.Player.HeadSkills.Main.ID

		if len(opponent2.Player.HeadSkills.Subs) > 0 {
			(*battleUpload).Opponent2HeadgearSub0 = opponent2.Player.HeadSkills.Subs[0].ID

			if len(opponent2.Player.HeadSkills.Subs) > 1 {
				(*battleUpload).Opponent2HeadgearSub1 = opponent2.Player.HeadSkills.Subs[1].ID

				if len(opponent2.Player.HeadSkills.Subs) > 2 {
					(*battleUpload).Opponent2HeadgearSub2 = opponent2.Player.HeadSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Opponent2Clothes = opponent2.Player.Clothes.ID
		(*battleUpload).Opponent2ClothesMain = opponent2.Player.ClothesSkills.Main.ID

		if len(opponent2.Player.ClothesSkills.Subs) > 0 {
			(*battleUpload).Opponent2ClothesSub0 = opponent2.Player.ClothesSkills.Subs[0].ID

			if len(opponent2.Player.ClothesSkills.Subs) > 1 {
				(*battleUpload).Opponent2ClothesSub1 = opponent2.Player.ClothesSkills.Subs[1].ID

				if len(opponent2.Player.ClothesSkills.Subs) > 2 {
					(*battleUpload).Opponent2ClothesSub2 = opponent2.Player.ClothesSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Opponent2Shoes = opponent2.Player.Shoes.ID
		(*battleUpload).Opponent2ShoesMain = opponent2.Player.ShoesSkills.Main.ID

		if len(opponent2.Player.ShoesSkills.Subs) > 0 {
			(*battleUpload).Opponent2ShoesSub0 = opponent2.Player.ShoesSkills.Subs[0].ID

			if len(opponent2.Player.ShoesSkills.Subs) > 1 {
				(*battleUpload).Opponent2ShoesSub1 = opponent2.Player.ShoesSkills.Subs[1].ID

				if len(opponent2.Player.ShoesSkills.Subs) > 2 {
					(*battleUpload).Opponent2ShoesSub2 = opponent2.Player.ShoesSkills.Subs[2].ID
				}
			}
		}
	}
}

func battleSetOpponent3(battle *types.Battle, battleUpload *types.BattleUpload) {
	if len((*battle).OtherTeamMembers) > 3 {
		opponent3 := (*battle).OtherTeamMembers[3]
		(*battleUpload).Opponent3SplatnetID = *opponent3.Player.PrincipalID
		(*battleUpload).Opponent3Name = *opponent3.Player.Nickname
		(*battleUpload).Opponent3LevelStar = *opponent3.Player.StarRank
		(*battleUpload).Opponent3Level = *opponent3.Player.PlayerRank
		(*battleUpload).Opponent3Rank = *opponent3.Player.Udemae.Name
		(*battleUpload).Opponent3Weapon = *opponent3.Player.Weapon.ID
		(*battleUpload).Opponent3Gender = *opponent3.Player.PlayerType.Gender
		(*battleUpload).Opponent3Species = *opponent3.Player.PlayerType.Species
		(*battleUpload).Opponent3Kills = *opponent3.KillCount
		(*battleUpload).Opponent3Deaths = *opponent3.DeathCount
		(*battleUpload).Opponent3Assists = *opponent3.AssistCount
		(*battleUpload).Opponent3GamePaintPoint = *opponent3.GamePaintPoint
		(*battleUpload).Opponent3Specials = *opponent3.SpecialCount
		(*battleUpload).Opponent3Headgear = opponent3.Player.Head.ID
		(*battleUpload).Opponent3HeadgearMain = opponent3.Player.HeadSkills.Main.ID

		if len(opponent3.Player.HeadSkills.Subs) > 0 {
			(*battleUpload).Opponent3HeadgearSub0 = opponent3.Player.HeadSkills.Subs[0].ID

			if len(opponent3.Player.HeadSkills.Subs) > 1 {
				(*battleUpload).Opponent3HeadgearSub1 = opponent3.Player.HeadSkills.Subs[1].ID

				if len(opponent3.Player.HeadSkills.Subs) > 2 {
					(*battleUpload).Opponent3HeadgearSub2 = opponent3.Player.HeadSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Opponent3Clothes = opponent3.Player.Clothes.ID
		(*battleUpload).Opponent3ClothesMain = opponent3.Player.ClothesSkills.Main.ID

		if len(opponent3.Player.ClothesSkills.Subs) > 0 {
			(*battleUpload).Opponent3ClothesSub0 = opponent3.Player.ClothesSkills.Subs[0].ID

			if len(opponent3.Player.ClothesSkills.Subs) > 1 {
				(*battleUpload).Opponent3ClothesSub1 = opponent3.Player.ClothesSkills.Subs[1].ID

				if len(opponent3.Player.ClothesSkills.Subs) > 2 {
					(*battleUpload).Opponent3ClothesSub2 = opponent3.Player.ClothesSkills.Subs[2].ID
				}
			}
		}

		(*battleUpload).Opponent3Shoes = opponent3.Player.Shoes.ID
		(*battleUpload).Opponent3ShoesMain = opponent3.Player.ShoesSkills.Main.ID

		if len(opponent3.Player.ShoesSkills.Subs) > 0 {
			(*battleUpload).Opponent3ShoesSub0 = opponent3.Player.ShoesSkills.Subs[0].ID

			if len(opponent3.Player.ShoesSkills.Subs) > 1 {
				(*battleUpload).Opponent3ShoesSub1 = opponent3.Player.ShoesSkills.Subs[1].ID

				if len(opponent3.Player.ShoesSkills.Subs) > 2 {
					(*battleUpload).Opponent3ShoesSub2 = opponent3.Player.ShoesSkills.Subs[2].ID
				}
			}
		}
	}
}
