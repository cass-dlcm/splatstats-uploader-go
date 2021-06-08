package statink2splatstats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cass-dlcm/splatstatsuploader/data"
	"github.com/cass-dlcm/splatstatsuploader/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
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

func Migrate(s bool, salmon bool, apiKey string, client *http.Client) {
	if salmon {
		var allData []types.ShiftStatInk

		downloadShifts(s, &allData, client)

		for i := range allData {
			shift := transformShift(&allData[i])
			data.UploadSalmon(&shift, apiKey, client)
		}
	} else {
		var allData []types.BattleStatInk

		downloadBattles(s, &allData, client)

		for i := range allData {
			battle := transformBattle(&allData[i])
			data.UploadBattle(&battle, apiKey, client)
		}
	}
}

func downloadShifts(s bool, allData *[]types.ShiftStatInk, client *http.Client) {
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
		q.Add("count", "50")
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

		if len(tempData) > 0 {
			currentID = tempData[len(tempData)-1].ID
		}

		for i := range tempData {
			*allData = append(*allData, tempData[i])

			if s {
				shiftMarshal, err := json.Marshal(tempData[i])
				if err != nil {
					panic(err)
				}

				if err := ioutil.WriteFile("two_salmon_statink/"+fmt.Sprint(*(tempData[i].SplatnetNumber))+".json", shiftMarshal, 0600); err != nil {
					panic(err)
				}
			}
		}
	}
}

func transformShift(statInkShift *types.ShiftStatInk) types.ShiftUpload {
	var shift types.ShiftUpload

	var err error
	shift.DangerRate, err = strconv.ParseFloat((*statInkShift).DangerRate, 64)

	if err != nil {
		panic(err)
	}

	shift.StatInkUpload = true
	shift.SplatnetUpload = false
	shift.Playtime = (*statInkShift).StartAt.Iso8601.Format("2006-01-02 15:04:05")
	shift.ScheduleStarttime = (*statInkShift).ShiftStartAt.Iso8601.Format("2006-01-02 15:04:05")
	shift.GradePointDelta = *((*statInkShift).TitleExpAfter) - (*statInkShift).TitleExp
	shift.StatInkJson = statInkShift
	shift.Stage = (*statInkShift).Stage.Name.EnUs
	shift.GradePoint = (*statInkShift).TitleExpAfter
	shift.JobFailureReason = (*statInkShift).FailReason.Key
	shift.IsClear = (*statInkShift).IsCleared
	shift.FailureWave = (*statInkShift).ClearWaves + 1
	shift.JobID = (*statInkShift).SplatnetNumber

	transformWaves(statInkShift, &shift)
	transformPlayerDataSalmon(statInkShift, &shift)
	transformTeammate0Salmon(statInkShift, &shift)
	transformTeammate1Salmon(statInkShift, &shift)
	transformTeammate2Salmon(statInkShift, &shift)
	transformBossAppearances(statInkShift, &shift)

	return shift
}

func transformWaves(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	if len((*statInkShift).Waves) > 0 {
		(*shiftUpload).Wave1GoldenDelivered = (*statInkShift).Waves[0].GoldenEggDelivered
		(*shiftUpload).Wave1EventType = (*statInkShift).Waves[0].KnownOccurrence.Splatnet
		(*shiftUpload).Wave1GoldenAppear = (*statInkShift).Waves[0].GoldenEggAppearances
		(*shiftUpload).Wave1Quota = (*statInkShift).Waves[0].GoldenEggQuota
		(*shiftUpload).Wave1PowerEggs = (*statInkShift).Waves[0].PowerEggCollected
		(*shiftUpload).Wave1WaterLevel = (*statInkShift).Waves[0].WaterLevel.Splatnet

		if len((*statInkShift).Waves) > 1 {
			(*shiftUpload).Wave2GoldenDelivered = (*statInkShift).Waves[1].GoldenEggDelivered
			(*shiftUpload).Wave2EventType = (*statInkShift).Waves[1].KnownOccurrence.Splatnet
			(*shiftUpload).Wave2GoldenAppear = (*statInkShift).Waves[1].GoldenEggAppearances
			(*shiftUpload).Wave2Quota = (*statInkShift).Waves[1].GoldenEggQuota
			(*shiftUpload).Wave2PowerEggs = (*statInkShift).Waves[1].PowerEggCollected
			(*shiftUpload).Wave2WaterLevel = (*statInkShift).Waves[1].WaterLevel.Splatnet

			if len((*statInkShift).Waves) > 2 {
				(*shiftUpload).Wave3GoldenDelivered = (*statInkShift).Waves[2].GoldenEggDelivered
				(*shiftUpload).Wave3EventType = (*statInkShift).Waves[2].KnownOccurrence.Splatnet
				(*shiftUpload).Wave3GoldenAppear = (*statInkShift).Waves[2].GoldenEggAppearances
				(*shiftUpload).Wave3Quota = (*statInkShift).Waves[2].GoldenEggQuota
				(*shiftUpload).Wave3PowerEggs = (*statInkShift).Waves[2].PowerEggCollected
				(*shiftUpload).Wave3WaterLevel = (*statInkShift).Waves[2].WaterLevel.Splatnet
			}
		}
	}
}

func transformPlayerDataSalmon(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	(*shiftUpload).PlayerTitle = fmt.Sprint((*statInkShift).TitleAfter.Splatnet)
	(*shiftUpload).PlayerSpecial = fmt.Sprint((*statInkShift).MyData.Special.Splatnet)
	(*shiftUpload).PlayerSpecies = (*statInkShift).MyData.Species.Key + "s"
	(*shiftUpload).PlayerGender = (*statInkShift).MyData.Gender.Key
	(*shiftUpload).PlayerGoldenEggs = (*statInkShift).MyData.GoldenEggDelivered
	(*shiftUpload).PlayerPowerEggs = (*statInkShift).MyData.PowerEggCollected
	(*shiftUpload).PlayerName = (*statInkShift).MyData.Name
	(*shiftUpload).PlayerReviveCount = (*statInkShift).MyData.Rescue
	(*shiftUpload).PlayerDeathCount = (*statInkShift).MyData.Death
	(*shiftUpload).PlayerID = (*statInkShift).MyData.SplatnetID

	if len((*statInkShift).MyData.SpecialUses) > 0 {
		(*shiftUpload).PlayerW1Specials = (*statInkShift).MyData.SpecialUses[0]

		if len((*statInkShift).MyData.SpecialUses) > 1 {
			(*shiftUpload).PlayerW2Specials = (*statInkShift).MyData.SpecialUses[1]

			if len((*statInkShift).MyData.SpecialUses) > 2 {
				(*shiftUpload).PlayerW3Specials = (*statInkShift).MyData.SpecialUses[2]
			}
		}
	}

	if len((*statInkShift).MyData.Weapons) > 0 {
		(*shiftUpload).PlayerWeaponW1 = fmt.Sprint((*statInkShift).MyData.Weapons[0].Splatnet)

		if len((*statInkShift).MyData.Weapons) > 1 {
			(*shiftUpload).PlayerWeaponW2 = fmt.Sprint((*statInkShift).MyData.Weapons[1].Splatnet)

			if len((*statInkShift).MyData.Weapons) > 2 {
				(*shiftUpload).PlayerWeaponW3 = fmt.Sprint((*statInkShift).MyData.Weapons[2].Splatnet)
			}
		}
	}

	transformPlayerBossKills(statInkShift, shiftUpload)
}

func transformPlayerBossKills(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	for _, boss := range (*statInkShift).MyData.BossKills {
		switch boss.Boss.Splatnet {
		case 3:
			(*shiftUpload).PlayerGoldieKills = boss.Count
		case 6:
			(*shiftUpload).PlayerSteelheadKills = boss.Count
		case 9:
			(*shiftUpload).PlayerFlyfishKills = boss.Count
		case 12:
			(*shiftUpload).PlayerScrapperKills = boss.Count
		case 13:
			(*shiftUpload).PlayerSteelEelKills = boss.Count
		case 14:
			(*shiftUpload).PlayerStingerKills = boss.Count
		case 15:
			(*shiftUpload).PlayerMawsKills = boss.Count
		case 16:
			(*shiftUpload).PlayerGrillerKills = boss.Count
		case 21:
			(*shiftUpload).PlayerDrizzlerKills = boss.Count
		}
	}
}

func transformTeammate0Salmon(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	(*shiftUpload).Teammate0Special = fmt.Sprint((*statInkShift).Teammates[0].Special.Splatnet)
	(*shiftUpload).Teammate0Species = (*statInkShift).Teammates[0].Species.Key + "s"
	(*shiftUpload).Teammate0Gender = (*statInkShift).Teammates[0].Gender.Key
	(*shiftUpload).Teammate0GoldenEggs = (*statInkShift).Teammates[0].GoldenEggDelivered
	(*shiftUpload).Teammate0PowerEggs = (*statInkShift).Teammates[0].PowerEggCollected
	(*shiftUpload).Teammate0Name = (*statInkShift).Teammates[0].Name
	(*shiftUpload).Teammate0ReviveCount = (*statInkShift).Teammates[0].Rescue
	(*shiftUpload).Teammate0DeathCount = (*statInkShift).Teammates[0].Death
	(*shiftUpload).Teammate0ID = (*statInkShift).Teammates[0].SplatnetID
	transformTeammate0BossKills(statInkShift, shiftUpload)

	if len((*statInkShift).Teammates[0].SpecialUses) > 0 {
		(*shiftUpload).Teammate0W1Specials = (*statInkShift).Teammates[0].SpecialUses[0]

		if len((*statInkShift).Teammates[0].SpecialUses) > 1 {
			(*shiftUpload).Teammate0W2Specials = (*statInkShift).Teammates[0].SpecialUses[1]

			if len((*statInkShift).Teammates[0].SpecialUses) > 2 {
				(*shiftUpload).Teammate0W3Specials = (*statInkShift).Teammates[0].SpecialUses[2]
			}
		}
	}

	if len((*statInkShift).Teammates[0].Weapons) > 0 {
		(*shiftUpload).Teammate0WeaponW1 = fmt.Sprint((*statInkShift).Teammates[0].Weapons[0].Splatnet)

		if len((*statInkShift).Teammates[0].Weapons) > 1 {
			(*shiftUpload).Teammate0WeaponW2 = fmt.Sprint((*statInkShift).Teammates[0].Weapons[1].Splatnet)

			if len((*statInkShift).Teammates[0].Weapons) > 2 {
				(*shiftUpload).Teammate0WeaponW3 = fmt.Sprint((*statInkShift).Teammates[0].Weapons[2].Splatnet)
			}
		}
	}
}

func transformTeammate0BossKills(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	for _, boss := range (*statInkShift).Teammates[0].BossKills {
		switch boss.Boss.Splatnet {
		case 3:
			(*shiftUpload).Teammate0GoldieKills = boss.Count
		case 6:
			(*shiftUpload).Teammate0SteelheadKills = boss.Count
		case 9:
			(*shiftUpload).Teammate0FlyfishKills = boss.Count
		case 12:
			(*shiftUpload).Teammate0ScrapperKills = boss.Count
		case 13:
			(*shiftUpload).Teammate0SteelEelKills = boss.Count
		case 14:
			(*shiftUpload).Teammate0StingerKills = boss.Count
		case 15:
			(*shiftUpload).Teammate0MawsKills = boss.Count
		case 16:
			(*shiftUpload).Teammate0GrillerKills = boss.Count
		case 21:
			(*shiftUpload).Teammate0DrizzlerKills = boss.Count
		}
	}
}

func transformTeammate1Salmon(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	(*shiftUpload).Teammate1Special = fmt.Sprint((*statInkShift).Teammates[1].Special.Splatnet)
	(*shiftUpload).Teammate1Species = (*statInkShift).Teammates[1].Species.Key + "s"
	(*shiftUpload).Teammate1Gender = (*statInkShift).Teammates[1].Gender.Key
	(*shiftUpload).Teammate1GoldenEggs = (*statInkShift).Teammates[1].GoldenEggDelivered
	(*shiftUpload).Teammate1PowerEggs = (*statInkShift).Teammates[1].PowerEggCollected
	(*shiftUpload).Teammate1Name = (*statInkShift).Teammates[1].Name
	(*shiftUpload).Teammate1ReviveCount = (*statInkShift).Teammates[1].Rescue
	(*shiftUpload).Teammate1DeathCount = (*statInkShift).Teammates[1].Death
	(*shiftUpload).Teammate1ID = (*statInkShift).Teammates[1].SplatnetID

	if len((*statInkShift).Teammates[1].SpecialUses) > 0 {
		(*shiftUpload).Teammate1W1Specials = (*statInkShift).Teammates[1].SpecialUses[0]

		if len((*statInkShift).Teammates[1].SpecialUses) > 1 {
			(*shiftUpload).Teammate1W2Specials = (*statInkShift).Teammates[1].SpecialUses[1]

			if len((*statInkShift).Teammates[1].SpecialUses) > 2 {
				(*shiftUpload).Teammate1W3Specials = (*statInkShift).Teammates[1].SpecialUses[2]
			}
		}
	}

	if len((*statInkShift).Teammates[1].Weapons) > 0 {
		(*shiftUpload).Teammate1WeaponW1 = fmt.Sprint((*statInkShift).Teammates[1].Weapons[0].Splatnet)

		if len((*statInkShift).Teammates[1].Weapons) > 1 {
			(*shiftUpload).Teammate1WeaponW2 = fmt.Sprint((*statInkShift).Teammates[1].Weapons[1].Splatnet)

			if len((*statInkShift).Teammates[1].Weapons) > 2 {
				(*shiftUpload).Teammate1WeaponW3 = fmt.Sprint((*statInkShift).Teammates[1].Weapons[2].Splatnet)
			}
		}
	}

	transformTeammate1BossKills(statInkShift, shiftUpload)
}

func transformTeammate1BossKills(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	for _, boss := range (*statInkShift).Teammates[1].BossKills {
		switch boss.Boss.Splatnet {
		case 3:
			(*shiftUpload).Teammate1GoldieKills = boss.Count
		case 6:
			(*shiftUpload).Teammate1SteelheadKills = boss.Count
		case 9:
			(*shiftUpload).Teammate1FlyfishKills = boss.Count
		case 12:
			(*shiftUpload).Teammate1ScrapperKills = boss.Count
		case 13:
			(*shiftUpload).Teammate1SteelEelKills = boss.Count
		case 14:
			(*shiftUpload).Teammate1StingerKills = boss.Count
		case 15:
			(*shiftUpload).Teammate1MawsKills = boss.Count
		case 16:
			(*shiftUpload).Teammate1GrillerKills = boss.Count
		case 21:
			(*shiftUpload).Teammate1DrizzlerKills = boss.Count
		}
	}
}

func transformTeammate2Salmon(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	(*shiftUpload).Teammate2Special = fmt.Sprint((*statInkShift).Teammates[2].Special.Splatnet)
	(*shiftUpload).Teammate2Species = (*statInkShift).Teammates[2].Species.Key + "s"
	(*shiftUpload).Teammate2Gender = (*statInkShift).Teammates[2].Gender.Key
	(*shiftUpload).Teammate2GoldenEggs = (*statInkShift).Teammates[2].GoldenEggDelivered
	(*shiftUpload).Teammate2PowerEggs = (*statInkShift).Teammates[2].PowerEggCollected
	(*shiftUpload).Teammate2Name = (*statInkShift).Teammates[2].Name
	(*shiftUpload).Teammate2ReviveCount = (*statInkShift).Teammates[2].Rescue
	(*shiftUpload).Teammate2DeathCount = (*statInkShift).Teammates[2].Death
	(*shiftUpload).Teammate2ID = (*statInkShift).Teammates[2].SplatnetID

	if len((*statInkShift).Teammates[2].Weapons) > 0 {
		(*shiftUpload).Teammate2WeaponW1 = fmt.Sprint((*statInkShift).Teammates[2].Weapons[0].Splatnet)

		if len((*statInkShift).Teammates[2].Weapons) > 1 {
			(*shiftUpload).Teammate2WeaponW2 = fmt.Sprint((*statInkShift).Teammates[2].Weapons[1].Splatnet)

			if len((*statInkShift).Teammates[2].Weapons) > 2 {
				(*shiftUpload).Teammate2WeaponW3 = fmt.Sprint((*statInkShift).Teammates[2].Weapons[2].Splatnet)
			}
		}
	}

	if len((*statInkShift).Teammates[2].SpecialUses) > 0 {
		(*shiftUpload).Teammate2W1Specials = (*statInkShift).Teammates[2].SpecialUses[0]

		if len((*statInkShift).Teammates[2].SpecialUses) > 1 {
			(*shiftUpload).Teammate2W2Specials = (*statInkShift).Teammates[2].SpecialUses[1]

			if len((*statInkShift).Teammates[2].SpecialUses) > 2 {
				(*shiftUpload).Teammate2W3Specials = (*statInkShift).Teammates[2].SpecialUses[2]
			}
		}
	}

	transformTeammate2BossKills(statInkShift, shiftUpload)
}

func transformTeammate2BossKills(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	for _, boss := range (*statInkShift).Teammates[2].BossKills {
		switch boss.Boss.Splatnet {
		case 3:
			(*shiftUpload).Teammate2GoldieKills = boss.Count
		case 6:
			(*shiftUpload).Teammate2SteelheadKills = boss.Count
		case 9:
			(*shiftUpload).Teammate2FlyfishKills = boss.Count
		case 12:
			(*shiftUpload).Teammate2ScrapperKills = boss.Count
		case 13:
			(*shiftUpload).Teammate2SteelEelKills = boss.Count
		case 14:
			(*shiftUpload).Teammate2StingerKills = boss.Count
		case 15:
			(*shiftUpload).Teammate2MawsKills = boss.Count
		case 16:
			(*shiftUpload).Teammate2GrillerKills = boss.Count
		case 21:
			(*shiftUpload).Teammate2DrizzlerKills = boss.Count
		}
	}
}

func transformBossAppearances(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	for _, boss := range (*statInkShift).BossAppearances {
		switch boss.Boss.Splatnet {
		case 3:
			(*shiftUpload).GoldieCount = boss.Count
		case 6:
			(*shiftUpload).SteelheadCount = boss.Count
		case 9:
			(*shiftUpload).FlyfishCount = boss.Count
		case 12:
			(*shiftUpload).ScrapperCount = boss.Count
		case 13:
			(*shiftUpload).SteelEelCount = boss.Count
		case 14:
			(*shiftUpload).StingerCount = boss.Count
		case 15:
			(*shiftUpload).MawsCount = boss.Count
		case 16:
			(*shiftUpload).GrillerCount = boss.Count
		case 21:
			(*shiftUpload).DrizzlerCount = boss.Count
		}
	}
}

func downloadBattles(s bool, allData *[]types.BattleStatInk, client *http.Client) {
	if viper.GetString("statink_api_key") == "" {
		enterStatinkApiKey()
	}

	statinkApiKey := viper.GetString("statink_api_key")
	reqUrl := "https://stat.ink/api/v2/user-battle"
	previousID := -1

	currentID := 1
	for currentID > previousID {
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

		req, err := http.NewRequestWithContext(ctx, "GET", reqUrl, nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Authorization", "Bearer "+statinkApiKey)

		q := req.URL.Query()
		q.Add("newer_than", fmt.Sprint(currentID))
		q.Add("order", "asc")
		q.Add("count", "50")
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		fmt.Println(resp.Status)

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(errors.Wrap(err, resp.Body.Close().Error()))
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}

		var tempData types.BattleStatInkArray
		if err := json.Unmarshal(bodyBytes, &tempData); err != nil {
			panic(err)
		}

		if err := resp.Body.Close(); err != nil {
			panic(err)
		}

		previousID = currentID

		if len(tempData) > 0 {
			currentID = tempData[len(tempData)-1].ID
		}

		for i := range tempData {
			*allData = append(*allData, tempData[i])

			if s {
				battleMarshal, err := json.Marshal(tempData[i])
				if err != nil {
					panic(err)
				}

				if err := ioutil.WriteFile("two_battle_statink/"+fmt.Sprint(tempData[i].SplatnetNumber)+".json", battleMarshal, 0600); err != nil {
					panic(err)
				}
			}
		}
	}
}

func transformBattle(statInkBattle *types.BattleStatInk) types.BattleUpload {
	var battle types.BattleUpload

	var err error

	battle.SplatnetUpload = false
	battle.StatInkUpload = true
	battle.StatInkJson = statInkBattle

	switch (*statInkBattle).Rule.Key {
	case "asari":
		battle.Rule = "clam_blitz"
	case "nawabari":
		battle.Rule = "turf_war"
	case "area":
		battle.Rule = "splat_zones"
	case "hoko":
		battle.Rule = "rainmaker"
	case "yagura":
		battle.Rule = "tower_control"
	}

	switch (*statInkBattle).Lobby.Key {
	case "standard":
		switch (*statInkBattle).Mode.Key {
		case "gachi":
			battle.MatchType = "gachi"
		case "fest":
			battle.MatchType = "fes_solo"
		case "regular":
			battle.MatchType = "turf_war"
		}
	case "private":
		battle.MatchType = "private"
	case "squad_2":
		battle.MatchType = "league_pair"
	case "squad_4":
		switch (*statInkBattle).Mode.Key {
		case "fest":
			battle.MatchType = "fes_team"
		case "gachi":
			battle.MatchType = "league_team"
		}
	case "fest_normal":
		battle.MatchType = "fes_team"
	}

	battle.Stage = fmt.Sprint((*statInkBattle).Map.Splatnet)
	battle.Win = (*statInkBattle).Result == "win"
	battle.HasDisconnectedPlayer = false

	for i := range (*statInkBattle).Players {
		battle.HasDisconnectedPlayer = battle.HasDisconnectedPlayer || ((*statInkBattle).Players[i].Point == 0 && (*statInkBattle).Players[i].KillOrAssist == 0 && (*statInkBattle).Players[i].Death == 0)
	}

	battle.Time = (*statInkBattle).StartAt.Time
	battle.BattleNumber = fmt.Sprint((*statInkBattle).SplatnetNumber)
	battle.WinMeter = (*statInkBattle).Freshness.Freshness

	if (*statInkBattle).MyTeamCount != nil {
		battle.MyTeamCount = float64(*(*statInkBattle).MyTeamCount)
	} else if (*statInkBattle).MyTeamPercent != nil {
		f, err := strconv.ParseFloat(*(*statInkBattle).MyTeamPercent, 64)
		if err != nil {
			panic(err)
		}
		battle.MyTeamCount = f
	}

	if (*statInkBattle).HisTeamCount != nil {
		battle.OtherTeamCount = float64(*(*statInkBattle).HisTeamCount)
	} else if (*statInkBattle).HisTeamPercent != nil {
		battle.OtherTeamCount, err = strconv.ParseFloat(*(*statInkBattle).HisTeamPercent, 64)
		if err != nil {
			panic(err)
		}
	}

	battle.ElapsedTime = (*statInkBattle).EndAt.Time - (*statInkBattle).StartAt.Time
	battle.TagID = (*statInkBattle).MyTeamID

	if (*statInkBattle).LeaguePoint != "" {
		battle.LeaguePoint, err = strconv.ParseFloat((*statInkBattle).LeaguePoint, 64)
		if err != nil {
			panic(err)
		}
	}

	if (*statInkBattle).FestPower != "" {
		battle.SplatfestPoint, err = strconv.ParseFloat((*statInkBattle).FestPower, 64)
		if err != nil {
			panic(err)
		}
	}

	battle.SplatfestTitleAfter = (*statInkBattle).FestTitleAfter.Key

	battle.PlayerGender = (*statInkBattle).Gender.Key
	if (*statInkBattle).Species.Key != "" {
		battle.PlayerSpecies = (*statInkBattle).Species.Key + "s"
	}
	if (*statInkBattle).Gears.Headgear.Gear.Splatnet != nil {
		battle.PlayerHeadgear = fmt.Sprint(*(*statInkBattle).Gears.Headgear.Gear.Splatnet)
	}
	if (*statInkBattle).Gears.Clothing.Gear.Splatnet != nil {
		battle.PlayerClothes = fmt.Sprint(*(*statInkBattle).Gears.Clothing.Gear.Splatnet)
	}
	if (*statInkBattle).Gears.Shoes.Gear.Splatnet != nil {
		battle.PlayerShoes = fmt.Sprint(*(*statInkBattle).Gears.Shoes.Gear.Splatnet)
	}

	abilities := map[string]string{
		"Ink Saver (Main)":   "0",
		"Ink Saver (Sub)":    "1",
		"Ink Recovery Up":    "2",
		"Run Speed Up":       "3",
		"Swim Speed Up":      "4",
		"Special Charge Up":  "5",
		"Special Saver":      "6",
		"Special Power Up":   "7",
		"Quick Respawn":      "8",
		"Quick Super Jump":   "9",
		"Sub Power Up":       "10",
		"Ink Resistance Up":  "11",
		"Opening Gambit":     "100",
		"Last Ditch Effort":  "101",
		"Tenacity":           "102",
		"Comeback":           "103",
		"Ninja Squid":        "104",
		"Haunt":              "105",
		"Thermal Ink":        "106",
		"Respawn Punisher":   "107",
		"Ability Doubler":    "108",
		"Stealth Jump":       "109",
		"Object Shredder":    "110",
		"Drop Roller":        "111",
		"Bomb Defense Up DX": "200",
		"Main Power Up":      "201",
	}

	battle.PlayerHeadgearMain = abilities[(*statInkBattle).Gears.Headgear.PrimaryAbility.Name.EnUs]
	battle.PlayerClothesMain = abilities[(*statInkBattle).Gears.Clothing.PrimaryAbility.Name.EnUs]
	battle.PlayerShoesMain = abilities[(*statInkBattle).Gears.Shoes.PrimaryAbility.Name.EnUs]

	if len((*statInkBattle).Gears.Headgear.SecondaryAbilities) > 0 {
		battle.PlayerHeadgearSub0 = abilities[(*statInkBattle).Gears.Headgear.SecondaryAbilities[0].Name.EnUs]

		if len((*statInkBattle).Gears.Headgear.SecondaryAbilities) > 1 {
			battle.PlayerHeadgearSub1 = abilities[(*statInkBattle).Gears.Headgear.SecondaryAbilities[1].Name.EnUs]

			if len((*statInkBattle).Gears.Headgear.SecondaryAbilities) > 2 {
				battle.PlayerHeadgearSub2 = abilities[(*statInkBattle).Gears.Headgear.SecondaryAbilities[2].Name.EnUs]
			}
		}
	}

	if len((*statInkBattle).Gears.Clothing.SecondaryAbilities) > 0 {
		battle.PlayerClothesSub0 = abilities[(*statInkBattle).Gears.Clothing.SecondaryAbilities[0].Name.EnUs]

		if len((*statInkBattle).Gears.Clothing.SecondaryAbilities) > 1 {
			battle.PlayerClothesSub1 = abilities[(*statInkBattle).Gears.Clothing.SecondaryAbilities[1].Name.EnUs]

			if len((*statInkBattle).Gears.Clothing.SecondaryAbilities) > 2 {
				battle.PlayerClothesSub2 = abilities[(*statInkBattle).Gears.Clothing.SecondaryAbilities[2].Name.EnUs]
			}
		}
	}

	if len((*statInkBattle).Gears.Shoes.SecondaryAbilities) > 0 {
		battle.PlayerShoesSub0 = abilities[(*statInkBattle).Gears.Shoes.SecondaryAbilities[0].Name.EnUs]

		if len((*statInkBattle).Gears.Shoes.SecondaryAbilities) > 1 {
			battle.PlayerShoesSub1 = abilities[(*statInkBattle).Gears.Shoes.SecondaryAbilities[1].Name.EnUs]

			if len((*statInkBattle).Gears.Shoes.SecondaryAbilities) > 2 {
				battle.PlayerShoesSub2 = abilities[(*statInkBattle).Gears.Shoes.SecondaryAbilities[2].Name.EnUs]
			}
		}
	}

	var rank string

	switch battle.Rule {
	case "clam_blitz":
		rank = (*statInkBattle).User.Stats.V2.Gachi.Rules.Asari.RankCurrent
	case "turf_war":
		rank = ""
	case "splat_zones":
		rank = (*statInkBattle).User.Stats.V2.Gachi.Rules.Area.RankCurrent
	case "rainmaker":
		rank = (*statInkBattle).User.Stats.V2.Gachi.Rules.Hoko.RankCurrent
	case "tower_control":
		rank = (*statInkBattle).User.Stats.V2.Gachi.Rules.Yagura.RankCurrent
	}

	switch rank {
	case "C-":
		battle.PlayerRank = 0
	case "C":
		battle.PlayerRank = 1
	case "C+":
		battle.PlayerRank = 2
	case "B-":
		battle.PlayerRank = 3
	case "B":
		battle.PlayerRank = 4
	case "B+":
		battle.PlayerRank = 5
	case "A-":
		battle.PlayerRank = 6
	case "A":
		battle.PlayerRank = 7
	case "A+":
		battle.PlayerRank = 8
	case "S":
		battle.PlayerRank = 9
	case "S+ 0":
		battle.PlayerRank = 10
	case "S+ 1":
		battle.PlayerRank = 11
	case "S+ 2":
		battle.PlayerRank = 12
	case "S+ 3":
		battle.PlayerRank = 13
	case "S+ 4":
		battle.PlayerRank = 14
	case "S+ 5":
		battle.PlayerRank = 15
	case "S+ 6":
		battle.PlayerRank = 16
	case "S+ 7":
		battle.PlayerRank = 17
	case "S+ 8":
		battle.PlayerRank = 18
	case "S+ 9":
		battle.PlayerRank = 19
	case "X":
		battle.PlayerRank = 20
	}

	battle.PlayerLevel = (*statInkBattle).Level
	battle.PlayerLevelStar = (*statInkBattle).StarRank
	battle.PlayerKills = (*statInkBattle).Kill
	battle.PlayerDeaths = (*statInkBattle).Death
	battle.PlayerAssists = (*statInkBattle).KillOrAssist - (*statInkBattle).Kill
	battle.PlayerSpecials = (*statInkBattle).Special
	battle.PlayerGamePaintPoint = (*statInkBattle).MyPoint
	battle.PlayerSplatfestTitle = (*statInkBattle).FestTitle.Key

	otherCount := 0
	myCount := 0

	for i := range (*statInkBattle).Players {
		player := &(*statInkBattle).Players[i]
		switch (*player).Team {
		case "my":
			if (*player).IsMe {
				battle.PlayerWeapon = fmt.Sprint((*player).Weapon.Splatnet)
				battle.PlayerSplatnetID = (*player).SplatnetID
				battle.PlayerName = (*player).Name
			} else {
				switch myCount {
				case 0:
					battle.Teammate0SplatnetID = (*player).SplatnetID
					battle.Teammate0Name = (*player).Name
					battle.Teammate0LevelStar = (*player).StarRank
					battle.Teammate0Level = (*player).Level
					battle.Teammate0Rank = (*player).Rank.Name.EnUs
					battle.Teammate0Weapon = fmt.Sprint((*player).Weapon.Splatnet)
					battle.Teammate0Gender = (*player).Gender.Key
					if (*player).Species.Key != "" {
						battle.Teammate0Species = (*player).Species.Key + "s"
					}
					battle.Teammate0Kills = (*player).Kill
					battle.Teammate0Deaths = (*player).Death
					battle.Teammate0Assists = (*player).KillOrAssist - (*player).Kill
					battle.Teammate0GamePaintPoint = (*player).Point
					battle.Teammate0Specials = (*player).Special
				case 1:
					battle.Teammate1SplatnetID = (*player).SplatnetID
					battle.Teammate1Name = (*player).Name
					battle.Teammate1LevelStar = (*player).StarRank
					battle.Teammate1Level = (*player).Level
					battle.Teammate1Rank = (*player).Rank.Name.EnUs
					battle.Teammate1Weapon = fmt.Sprint((*player).Weapon.Splatnet)
					battle.Teammate1Gender = (*player).Gender.Key
					if (*player).Species.Key != "" {
						battle.Teammate1Species = (*player).Species.Key + "s"
					}
					battle.Teammate1Kills = (*player).Kill
					battle.Teammate1Deaths = (*player).Death
					battle.Teammate1Assists = (*player).KillOrAssist - (*player).Kill
					battle.Teammate1GamePaintPoint = (*player).Point
					battle.Teammate1Specials = (*player).Special
				case 2:
					battle.Teammate2SplatnetID = (*player).SplatnetID
					battle.Teammate2Name = (*player).Name
					battle.Teammate2LevelStar = (*player).StarRank
					battle.Teammate2Level = (*player).Level
					battle.Teammate2Rank = (*player).Rank.Name.EnUs
					battle.Teammate2Weapon = fmt.Sprint((*player).Weapon.Splatnet)
					battle.Teammate2Gender = (*player).Gender.Key
					if (*player).Species.Key != "" {
						battle.Teammate2Species = (*player).Species.Key + "s"
					}
					battle.Teammate2Kills = (*player).Kill
					battle.Teammate2Deaths = (*player).Death
					battle.Teammate2Assists = (*player).KillOrAssist - (*player).Kill
					battle.Teammate2GamePaintPoint = (*player).Point
					battle.Teammate2Specials = (*player).Special
				}
				myCount++
			}
		case "his":
			switch otherCount {
			case 0:
				battle.Opponent0SplatnetID = (*player).SplatnetID
				battle.Opponent0Name = (*player).Name
				battle.Opponent0LevelStar = (*player).StarRank
				battle.Opponent0Level = (*player).Level
				battle.Opponent0Rank = (*player).Rank.Name.EnUs
				battle.Opponent0Weapon = fmt.Sprint((*player).Weapon.Splatnet)
				battle.Opponent0Gender = (*player).Gender.Key
				if (*player).Species.Key != "" {
					battle.Opponent0Species = (*player).Species.Key + "s"
				}
				battle.Opponent0Kills = (*player).Kill
				battle.Opponent0Deaths = (*player).Death
				battle.Opponent0Assists = (*player).KillOrAssist - (*player).Kill
				battle.Opponent0GamePaintPoint = (*player).Point
				battle.Opponent0Specials = (*player).Special
			case 1:
				battle.Opponent1SplatnetID = (*player).SplatnetID
				battle.Opponent1Name = (*player).Name
				battle.Opponent1LevelStar = (*player).StarRank
				battle.Opponent1Level = (*player).Level
				battle.Opponent1Rank = (*player).Rank.Name.EnUs
				battle.Opponent1Weapon = fmt.Sprint((*player).Weapon.Splatnet)
				battle.Opponent1Gender = (*player).Gender.Key
				if (*player).Species.Key != "" {
					battle.Opponent1Species = (*player).Species.Key + "s"
				}
				battle.Opponent1Kills = (*player).Kill
				battle.Opponent1Deaths = (*player).Death
				battle.Opponent1Assists = (*player).KillOrAssist - (*player).Kill
				battle.Opponent1GamePaintPoint = (*player).Point
				battle.Opponent1Specials = (*player).Special
			case 2:
				battle.Opponent2SplatnetID = (*player).SplatnetID
				battle.Opponent2Name = (*player).Name
				battle.Opponent2LevelStar = (*player).StarRank
				battle.Opponent2Level = (*player).Level
				battle.Opponent2Rank = (*player).Rank.Name.EnUs
				battle.Opponent2Weapon = fmt.Sprint((*player).Weapon.Splatnet)
				battle.Opponent2Gender = (*player).Gender.Key
				if (*player).Species.Key != "" {
					battle.Opponent2Species = (*player).Species.Key + "s"
				}
				battle.Opponent2Kills = (*player).Kill
				battle.Opponent2Deaths = (*player).Death
				battle.Opponent2Assists = (*player).KillOrAssist - (*player).Kill
				battle.Opponent2GamePaintPoint = (*player).Point
				battle.Opponent2Specials = (*player).Special
			case 3:
				battle.Opponent3SplatnetID = (*player).SplatnetID
				battle.Opponent3Name = (*player).Name
				battle.Opponent3LevelStar = (*player).StarRank
				battle.Opponent3Level = (*player).Level
				battle.Opponent3Rank = (*player).Rank.Name.EnUs
				battle.Opponent3Weapon = fmt.Sprint((*player).Weapon.Splatnet)
				battle.Opponent3Gender = (*player).Gender.Key
				if (*player).Species.Key != "" {
					battle.Opponent3Species = (*player).Species.Key + "s"
				}
				battle.Opponent3Kills = (*player).Kill
				battle.Opponent3Deaths = (*player).Death
				battle.Opponent3Assists = (*player).KillOrAssist - (*player).Kill
				battle.Opponent3GamePaintPoint = (*player).Point
				battle.Opponent3Specials = (*player).Special
			}
			otherCount++
		}
	}

	return battle
}
