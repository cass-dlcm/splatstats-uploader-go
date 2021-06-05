package statink2splatstats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cass-dlcm/splatstatsuploader/data"
	"github.com/cass-dlcm/splatstatsuploader/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func enterStatinkApiKey() {
	var apiKey string

	if _, err := fmt.Println("Enter your stat.ink API key here: "); err != nil {
		panic(err)
	}

	if _, err := fmt.Scanln(&apiKey); err != nil {
		panic(err)
	}

	viper.Set("statink_api_key", apiKey)

	if err := viper.WriteConfig(); err != nil {
		panic(err)
	}
}

func MigrateSalmon(apiKey string, client *http.Client) {
	var allData []types.ShiftStatInk

	downloadShifts(&allData, client)

	for i := range allData {
		var shift types.ShiftUpload

		transformShift(&allData[i], &shift)
		data.UploadSalmon(&shift, apiKey, client)
	}
}

func downloadShifts(allData *[]types.ShiftStatInk, client *http.Client) {
	if viper.GetString("statink_api_key") == "" {
		enterStatinkApiKey()
	}

	statinkApiKey := viper.GetString("statink_api_key")
	reqUrl := "https://stat.ink/api/v2/user-salmon"
	previousID := -1

	currentID := 1
	for currentID > previousID {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		req, err := http.NewRequestWithContext(ctx, "GET", reqUrl, nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Authorization", "Bearer "+statinkApiKey)

		q := req.URL.Query()
		q.Add("newer_than", fmt.Sprint(currentID))
		q.Add("order", "asc")
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(errors.Wrap(err, resp.Body.Close().Error()))
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}

		var tempData types.ShiftStatInkArray
		if err := json.Unmarshal(bodyBytes, &tempData); err != nil {
			panic(err)
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}

		previousID = currentID
		currentID = tempData[len(tempData)-1].ID

		for i := range tempData {
			*allData = append(*allData, tempData[i])
		}
	}
}

func transformShift(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	f, err := strconv.ParseFloat((*statInkShift).DangerRate, 64)
	if err != nil {
		panic(err)
	}

	*(*shiftUpload).DangerRate = f
	*(*shiftUpload).StatInkUpload = true
	*(*shiftUpload).SplatnetUpload = false
	*(*shiftUpload).StatInkJson = *statInkShift
	*(*shiftUpload).Stage = (*statInkShift).Stage.Splatnet
	*(*shiftUpload).Playtime = (*statInkShift).StartAt.Iso8601.Format("2006-01-02 15:04:05")
	*(*shiftUpload).ScheduleStarttime = (*statInkShift).ShiftStartAt.Iso8601.Format("2006-01-02 15:04:05")
	*(*shiftUpload).GradePointDelta = (*statInkShift).TitleExpAfter - (*statInkShift).TitleExp
	*(*shiftUpload).GradePoint = (*statInkShift).TitleExpAfter
	*(*shiftUpload).JobFailureReason = (*statInkShift).FailReason
	*(*shiftUpload).IsClear = (*statInkShift).IsCleared
	*(*shiftUpload).FailureWave = (*statInkShift).ClearWaves
	*(*shiftUpload).JobID = (*statInkShift).SplatnetNumber

	transformWaves(statInkShift, shiftUpload)
	transformPlayerDataSalmon(statInkShift, shiftUpload)
	transformBossAppearances(statInkShift, shiftUpload)
	transformTeammate0Salmon(statInkShift, shiftUpload)
	transformTeammate1Salmon(statInkShift, shiftUpload)
	transformTeammate2Salmon(statInkShift, shiftUpload)
}

func transformWaves(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	if len((*statInkShift).Waves) > 0 {
		*(*shiftUpload).Wave1GoldenDelivered = (*statInkShift).Waves[0].GoldenEggDelivered
		*(*shiftUpload).Wave1EventType = (*statInkShift).Waves[0].KnownOccurrence.Splatnet
		*(*shiftUpload).Wave1GoldenAppear = (*statInkShift).Waves[0].GoldenEggAppearances
		*(*shiftUpload).Wave1Quota = (*statInkShift).Waves[0].GoldenEggQuota
		*(*shiftUpload).Wave1PowerEggs = (*statInkShift).Waves[0].PowerEggCollected
		*(*shiftUpload).Wave1WaterLevel = (*statInkShift).Waves[0].WaterLevel.Splatnet

		if len((*statInkShift).Waves) > 1 {
			*(*shiftUpload).Wave2GoldenDelivered = (*statInkShift).Waves[1].GoldenEggDelivered
			*(*shiftUpload).Wave2EventType = (*statInkShift).Waves[1].KnownOccurrence.Splatnet
			*(*shiftUpload).Wave2GoldenAppear = (*statInkShift).Waves[1].GoldenEggAppearances
			*(*shiftUpload).Wave2Quota = (*statInkShift).Waves[1].GoldenEggQuota
			*(*shiftUpload).Wave2PowerEggs = (*statInkShift).Waves[1].PowerEggCollected
			*(*shiftUpload).Wave2WaterLevel = (*statInkShift).Waves[1].WaterLevel.Splatnet

			if len((*statInkShift).Waves) > 2 {
				*(*shiftUpload).Wave3GoldenDelivered = (*statInkShift).Waves[2].GoldenEggDelivered
				*(*shiftUpload).Wave3EventType = (*statInkShift).Waves[2].KnownOccurrence.Splatnet
				*(*shiftUpload).Wave3GoldenAppear = (*statInkShift).Waves[2].GoldenEggAppearances
				*(*shiftUpload).Wave3Quota = (*statInkShift).Waves[2].GoldenEggQuota
				*(*shiftUpload).Wave3PowerEggs = (*statInkShift).Waves[2].PowerEggCollected
				*(*shiftUpload).Wave3WaterLevel = (*statInkShift).Waves[2].WaterLevel.Splatnet
			}
		}
	}
}

func transformPlayerDataSalmon(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	*(*shiftUpload).PlayerSpecies = (*statInkShift).MyData.Species.Key
	*(*shiftUpload).PlayerGender = (*statInkShift).MyData.Gender.Key
	*(*shiftUpload).PlayerTitle = fmt.Sprint((*statInkShift).TitleAfter.Splatnet)
	*(*shiftUpload).PlayerGoldenEggs = (*statInkShift).MyData.GoldenEggDelivered
	*(*shiftUpload).PlayerPowerEggs = (*statInkShift).MyData.PowerEggCollected
	*(*shiftUpload).PlayerName = (*statInkShift).MyData.Name
	*(*shiftUpload).PlayerSpecial = fmt.Sprint((*statInkShift).MyData.Special.Splatnet)

	if len((*statInkShift).MyData.Weapons) > 0 {
		*(*shiftUpload).PlayerWeaponW1 = fmt.Sprint((*statInkShift).MyData.Weapons[0].Splatnet)

		if len((*statInkShift).MyData.Weapons) > 1 {
			*(*shiftUpload).PlayerWeaponW2 = fmt.Sprint((*statInkShift).MyData.Weapons[1].Splatnet)

			if len((*statInkShift).MyData.Weapons) > 2 {
				*(*shiftUpload).PlayerWeaponW3 = fmt.Sprint((*statInkShift).MyData.Weapons[2].Splatnet)
			}
		}
	}

	*(*shiftUpload).PlayerReviveCount = (*statInkShift).MyData.Rescue
	*(*shiftUpload).PlayerDeathCount = (*statInkShift).MyData.Death
	*(*shiftUpload).PlayerID = (*statInkShift).MyData.SplatnetID

	for _, boss := range (*statInkShift).MyData.BossKills {
		switch boss.Boss.Splatnet {
		case 3:
			*(*shiftUpload).PlayerGoldieKills = boss.Count
		case 6:
			*(*shiftUpload).PlayerSteelheadKills = boss.Count
		case 9:
			*(*shiftUpload).PlayerFlyfishKills = boss.Count
		case 12:
			*(*shiftUpload).PlayerScrapperKills = boss.Count
		case 13:
			*(*shiftUpload).PlayerSteelEelKills = boss.Count
		case 14:
			*(*shiftUpload).PlayerStingerKills = boss.Count
		case 15:
			*(*shiftUpload).PlayerMawsKills = boss.Count
		case 16:
			*(*shiftUpload).PlayerGrillerKills = boss.Count
		case 21:
			*(*shiftUpload).PlayerDrizzlerKills = boss.Count
		}
	}

	if len((*statInkShift).MyData.SpecialUses) > 0 {
		*(*shiftUpload).PlayerW1Specials = (*statInkShift).MyData.SpecialUses[0]

		if len((*statInkShift).MyData.SpecialUses) > 1 {
			*(*shiftUpload).PlayerW2Specials = (*statInkShift).MyData.SpecialUses[1]

			if len((*statInkShift).MyData.SpecialUses) > 2 {
				*(*shiftUpload).PlayerW3Specials = (*statInkShift).MyData.SpecialUses[2]
			}
		}
	}
}

func transformTeammate0Salmon(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	if len((*statInkShift).Teammates) < 1 {
		return
	}

	*(*shiftUpload).Teammate0Species = (*statInkShift).Teammates[0].Species.Key
	*(*shiftUpload).Teammate0Gender = (*statInkShift).Teammates[0].Gender.Key
	*(*shiftUpload).Teammate0GoldenEggs = (*statInkShift).Teammates[0].GoldenEggDelivered
	*(*shiftUpload).Teammate0PowerEggs = (*statInkShift).Teammates[0].PowerEggCollected
	*(*shiftUpload).Teammate0Name = (*statInkShift).Teammates[0].Name
	*(*shiftUpload).Teammate0Special = fmt.Sprint((*statInkShift).Teammates[0].Special.Splatnet)

	if len((*statInkShift).Teammates[0].Weapons) > 0 {
		*(*shiftUpload).Teammate0WeaponW1 = fmt.Sprint((*statInkShift).Teammates[0].Weapons[0].Splatnet)

		if len((*statInkShift).Teammates[0].Weapons) > 1 {
			*(*shiftUpload).Teammate0WeaponW2 = fmt.Sprint((*statInkShift).Teammates[0].Weapons[1].Splatnet)

			if len((*statInkShift).Teammates[0].Weapons) > 2 {
				*(*shiftUpload).Teammate0WeaponW3 = fmt.Sprint((*statInkShift).Teammates[0].Weapons[2].Splatnet)
			}
		}
	}

	*(*shiftUpload).Teammate0ReviveCount = (*statInkShift).Teammates[0].Rescue
	*(*shiftUpload).Teammate0DeathCount = (*statInkShift).Teammates[0].Death
	*(*shiftUpload).Teammate0ID = (*statInkShift).Teammates[0].SplatnetID

	for _, boss := range (*statInkShift).Teammates[0].BossKills {
		switch boss.Boss.Splatnet {
		case 3:
			*(*shiftUpload).Teammate0GoldieKills = boss.Count
		case 6:
			*(*shiftUpload).Teammate0SteelheadKills = boss.Count
		case 9:
			*(*shiftUpload).Teammate0FlyfishKills = boss.Count
		case 12:
			*(*shiftUpload).Teammate0ScrapperKills = boss.Count
		case 13:
			*(*shiftUpload).Teammate0SteelEelKills = boss.Count
		case 14:
			*(*shiftUpload).Teammate0StingerKills = boss.Count
		case 15:
			*(*shiftUpload).Teammate0MawsKills = boss.Count
		case 16:
			*(*shiftUpload).Teammate0GrillerKills = boss.Count
		case 21:
			*(*shiftUpload).Teammate0DrizzlerKills = boss.Count
		}
	}

	if len((*statInkShift).Teammates[0].SpecialUses) > 0 {
		*(*shiftUpload).Teammate0W1Specials = (*statInkShift).Teammates[0].SpecialUses[0]

		if len((*statInkShift).Teammates[0].SpecialUses) > 1 {
			*(*shiftUpload).Teammate0W2Specials = (*statInkShift).Teammates[0].SpecialUses[1]

			if len((*statInkShift).Teammates[0].SpecialUses) > 2 {
				*(*shiftUpload).Teammate0W3Specials = (*statInkShift).Teammates[0].SpecialUses[2]
			}
		}
	}
}

func transformTeammate1Salmon(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	if len((*statInkShift).Teammates) < 2 {
		return
	}

	*(*shiftUpload).Teammate1Species = (*statInkShift).Teammates[1].Species.Key
	*(*shiftUpload).Teammate1Gender = (*statInkShift).Teammates[1].Gender.Key
	*(*shiftUpload).Teammate1GoldenEggs = (*statInkShift).Teammates[1].GoldenEggDelivered
	*(*shiftUpload).Teammate1PowerEggs = (*statInkShift).Teammates[1].PowerEggCollected
	*(*shiftUpload).Teammate1Name = (*statInkShift).Teammates[1].Name
	*(*shiftUpload).Teammate1Special = fmt.Sprint((*statInkShift).Teammates[1].Special.Splatnet)

	if len((*statInkShift).Teammates[1].Weapons) > 0 {
		*(*shiftUpload).Teammate1WeaponW1 = fmt.Sprint((*statInkShift).Teammates[1].Weapons[0].Splatnet)

		if len((*statInkShift).Teammates[1].Weapons) > 1 {
			*(*shiftUpload).Teammate1WeaponW2 = fmt.Sprint((*statInkShift).Teammates[1].Weapons[1].Splatnet)

			if len((*statInkShift).Teammates[1].Weapons) > 2 {
				*(*shiftUpload).Teammate1WeaponW3 = fmt.Sprint((*statInkShift).Teammates[1].Weapons[2].Splatnet)
			}
		}
	}

	*(*shiftUpload).Teammate1ReviveCount = (*statInkShift).Teammates[1].Rescue
	*(*shiftUpload).Teammate1DeathCount = (*statInkShift).Teammates[1].Death
	*(*shiftUpload).Teammate1ID = (*statInkShift).Teammates[1].SplatnetID

	for _, boss := range (*statInkShift).Teammates[1].BossKills {
		switch boss.Boss.Splatnet {
		case 3:
			*(*shiftUpload).Teammate1GoldieKills = boss.Count
		case 6:
			*(*shiftUpload).Teammate1SteelheadKills = boss.Count
		case 9:
			*(*shiftUpload).Teammate1FlyfishKills = boss.Count
		case 12:
			*(*shiftUpload).Teammate1ScrapperKills = boss.Count
		case 13:
			*(*shiftUpload).Teammate1SteelEelKills = boss.Count
		case 14:
			*(*shiftUpload).Teammate1StingerKills = boss.Count
		case 15:
			*(*shiftUpload).Teammate1MawsKills = boss.Count
		case 16:
			*(*shiftUpload).Teammate1GrillerKills = boss.Count
		case 21:
			*(*shiftUpload).Teammate1DrizzlerKills = boss.Count
		}
	}

	if len((*statInkShift).Teammates[1].SpecialUses) > 0 {
		*(*shiftUpload).Teammate1W1Specials = (*statInkShift).Teammates[1].SpecialUses[0]

		if len((*statInkShift).Teammates[1].SpecialUses) > 1 {
			*(*shiftUpload).Teammate1W2Specials = (*statInkShift).Teammates[1].SpecialUses[1]

			if len((*statInkShift).Teammates[1].SpecialUses) > 2 {
				*(*shiftUpload).Teammate1W3Specials = (*statInkShift).Teammates[1].SpecialUses[2]
			}
		}
	}
}

func transformTeammate2Salmon(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	if len((*statInkShift).Teammates) < 3 {
		return
	}

	*(*shiftUpload).Teammate2Species = (*statInkShift).Teammates[2].Species.Key
	*(*shiftUpload).Teammate2Gender = (*statInkShift).Teammates[2].Gender.Key
	*(*shiftUpload).Teammate2GoldenEggs = (*statInkShift).Teammates[2].GoldenEggDelivered
	*(*shiftUpload).Teammate2PowerEggs = (*statInkShift).Teammates[2].PowerEggCollected
	*(*shiftUpload).Teammate2Name = (*statInkShift).Teammates[2].Name
	*(*shiftUpload).Teammate2Special = fmt.Sprint((*statInkShift).Teammates[2].Special.Splatnet)

	if len((*statInkShift).Teammates[2].Weapons) > 0 {
		*(*shiftUpload).Teammate2WeaponW1 = fmt.Sprint((*statInkShift).Teammates[2].Weapons[0].Splatnet)

		if len((*statInkShift).Teammates[2].Weapons) > 1 {
			*(*shiftUpload).Teammate2WeaponW2 = fmt.Sprint((*statInkShift).Teammates[2].Weapons[1].Splatnet)

			if len((*statInkShift).Teammates[2].Weapons) > 2 {
				*(*shiftUpload).Teammate2WeaponW3 = fmt.Sprint((*statInkShift).Teammates[2].Weapons[2].Splatnet)
			}
		}
	}

	*(*shiftUpload).Teammate2ReviveCount = (*statInkShift).Teammates[2].Rescue
	*(*shiftUpload).Teammate2DeathCount = (*statInkShift).Teammates[2].Death
	*(*shiftUpload).Teammate2ID = (*statInkShift).Teammates[2].SplatnetID

	for _, boss := range (*statInkShift).Teammates[2].BossKills {
		switch boss.Boss.Splatnet {
		case 3:
			*(*shiftUpload).Teammate2GoldieKills = boss.Count
		case 6:
			*(*shiftUpload).Teammate2SteelheadKills = boss.Count
		case 9:
			*(*shiftUpload).Teammate2FlyfishKills = boss.Count
		case 12:
			*(*shiftUpload).Teammate2ScrapperKills = boss.Count
		case 13:
			*(*shiftUpload).Teammate2SteelEelKills = boss.Count
		case 14:
			*(*shiftUpload).Teammate2StingerKills = boss.Count
		case 15:
			*(*shiftUpload).Teammate2MawsKills = boss.Count
		case 16:
			*(*shiftUpload).Teammate2GrillerKills = boss.Count
		case 21:
			*(*shiftUpload).Teammate2DrizzlerKills = boss.Count
		}
	}

	if len((*statInkShift).Teammates[2].SpecialUses) > 0 {
		*(*shiftUpload).Teammate2W1Specials = (*statInkShift).Teammates[2].SpecialUses[0]

		if len((*statInkShift).Teammates[2].SpecialUses) > 1 {
			*(*shiftUpload).Teammate2W2Specials = (*statInkShift).Teammates[2].SpecialUses[1]

			if len((*statInkShift).Teammates[2].SpecialUses) > 2 {
				*(*shiftUpload).Teammate2W3Specials = (*statInkShift).Teammates[2].SpecialUses[2]
			}
		}
	}
}

func transformBossAppearances(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	for _, boss := range (*statInkShift).BossAppearances {
		switch boss.Boss.Splatnet {
		case 3:
			*(*shiftUpload).GoldieCount = boss.Count
		case 6:
			*(*shiftUpload).SteelheadCount = boss.Count
		case 9:
			*(*shiftUpload).FlyfishCount = boss.Count
		case 12:
			*(*shiftUpload).ScrapperCount = boss.Count
		case 13:
			*(*shiftUpload).SteelEelCount = boss.Count
		case 14:
			*(*shiftUpload).StingerCount = boss.Count
		case 15:
			*(*shiftUpload).MawsCount = boss.Count
		case 16:
			*(*shiftUpload).GrillerCount = boss.Count
		case 21:
			*(*shiftUpload).DrizzlerCount = boss.Count
		}
	}
}

func MigrateBattles(apiKey string, client *http.Client) {
	if viper.GetString("statink_api_key") == "" {
		enterStatinkApiKey()
	}
	// statinkApiKey := viper.GetString("statink_api_key")
}
