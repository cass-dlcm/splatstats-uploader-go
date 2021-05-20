package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"cass-dlcm.dev/splatstatsuploader/types"
)

func Monitor(m int, s bool, salmon bool, api_key string, version string, client *http.Client) {

}

func File(salmon bool, api_key string, version string, client *http.Client) {
	if salmon {
		files, err := ioutil.ReadDir("./two_salmon/")
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			jsonFile, err := os.Open("./two_salmon/" + file.Name())
			if err != nil {
				panic(err)
			}
			defer jsonFile.Close()
			var shift types.Shift
			byteValue, err := ioutil.ReadAll(jsonFile)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(byteValue, &shift)
			uploadSalmon(shift, api_key, version, client)
		}
	} else {

	}
}

func GetSplatnet(s bool, salmon bool, api_key string, version string, client *http.Client) {

}

func uploadSalmon(shift types.Shift, api_key string, version string, client *http.Client) {
	shiftUpload := types.ShiftUpload{}
	shiftUpload.SplatnetJSON = shift
	shiftUpload.SplatnetUpload = true
	shiftUpload.StatInkUpload = false
	shiftUpload.DangerRate = shift.DangerRate
	shiftUpload.DrizzlerCount = shift.BossCounts.Drizzler.Count
	shiftUpload.Endtime = time.Unix(int64(shift.EndTime), 0).Format("2006-01-02 15:04:05")
	shiftUpload.FailureWave = shift.JobResult.FailureWave
	shiftUpload.FlyfishCount = shift.BossCounts.Flyfish.Count
	shiftUpload.GoldieCount = shift.BossCounts.Goldie.Count
	shiftUpload.GradePoint = shift.GradePoint
	shiftUpload.GradePointDelta = shift.GradePointDelta
	shiftUpload.GrillerCount = shift.BossCounts.Griller.Count
	shiftUpload.IsClear = shift.JobResult.IsClear
	shiftUpload.JobFailureReason = shift.JobResult.FailureReason
	shiftUpload.JobID = shift.JobID
	shiftUpload.MawsCount = shift.BossCounts.Maws.Count
	shiftUpload.PlayerDeathCount = shift.MyResult.DeadCount
	shiftUpload.PlayerDrizzlerKills = shift.MyResult.BossKillCounts.Drizzler.Count
	shiftUpload.PlayerFlyfishKills = shift.MyResult.BossKillCounts.Flyfish.Count
	shiftUpload.PlayerGender = shift.MyResult.PlayerType.Gender
	shiftUpload.PlayerGoldenEggs = shift.MyResult.GoldenEggs
	shiftUpload.PlayerGoldieKills = shift.MyResult.BossKillCounts.Goldie.Count
	shiftUpload.PlayerGrillerKills = shift.MyResult.BossKillCounts.Griller.Count
	shiftUpload.PlayerID = shift.MyResult.Pid
	shiftUpload.PlayerMawsKills = shift.MyResult.BossKillCounts.Maws.Count
	shiftUpload.PlayerName = shift.MyResult.Name
	shiftUpload.PlayerPowerEggs = shift.MyResult.PowerEggs
	shiftUpload.PlayerReviveCount = shift.MyResult.HelpCount
	shiftUpload.PlayerScrapperKills = shift.MyResult.BossKillCounts.Scrapper.Count
	shiftUpload.PlayerSpecial = shift.MyResult.Special.ID
	shiftUpload.PlayerSpecies = shift.MyResult.PlayerType.Species
	shiftUpload.PlayerSteelEelKills = shift.MyResult.BossKillCounts.SteelEel.Count
	shiftUpload.PlayerSteelheadKills = shift.MyResult.BossKillCounts.Steelhead.Count
	shiftUpload.PlayerStingerKills = shift.MyResult.BossKillCounts.Stinger.Count
	shiftUpload.PlayerTitle = shift.Grade.ID
	shiftUpload.PlayerW1Specials = shift.MyResult.SpecialCounts[0]
	shiftUpload.PlayerWeaponW1 = shift.MyResult.WeaponList[0].ID
	if shiftUpload.FailureWave > 1 {
		shiftUpload.PlayerW2Specials = shift.MyResult.SpecialCounts[1]
		shiftUpload.PlayerWeaponW2 = shift.MyResult.WeaponList[1].ID
		if shiftUpload.FailureWave > 2 {
			shiftUpload.PlayerW3Specials = shift.MyResult.SpecialCounts[2]
			shiftUpload.PlayerWeaponW3 = shift.MyResult.WeaponList[2].ID
		} else {
			shiftUpload.PlayerW3Specials = 0
			shiftUpload.PlayerWeaponW3 = ""
		}
	} else {
		shiftUpload.PlayerW2Specials = 0
		shiftUpload.PlayerW3Specials = 0
		shiftUpload.PlayerWeaponW2 = ""
		shiftUpload.PlayerWeaponW3 = ""
	}
	shiftUpload.Playtime = time.Unix(int64(shift.PlayTime), 0).Format("2006-01-02 15:04:05")
	shiftUpload.ScheduleEndtime = time.Unix(int64(shift.Schedule.EndTime), 0).Format("2006-01-02 15:04:05")
	shiftUpload.ScheduleStarttime = time.Unix(int64(shift.Schedule.StartTime), 0).Format("2006-01-02 15:04:05")
	shiftUpload.ScheduleWeapon0 = shift.Schedule.Weapons[0].ID
	shiftUpload.ScheduleWeapon1 = shift.Schedule.Weapons[1].ID
	shiftUpload.ScheduleWeapon2 = shift.Schedule.Weapons[2].ID
	shiftUpload.ScheduleWeapon3 = shift.Schedule.Weapons[3].ID
	shiftUpload.ScrapperCount = shift.BossCounts.Scrapper.Count
	shiftUpload.Stage = shift.Schedule.Stage.Name
	shiftUpload.Starttime = time.Unix(int64(shift.StartTime), 0).Format("2006-01-02 15:04:05")
	shiftUpload.SteelEelCount = shift.BossCounts.SteelEel.Count
	shiftUpload.SteelheadCount = shift.BossCounts.Steelhead.Count
	shiftUpload.StingerCount = shift.BossCounts.Stinger.Count
	if len(shift.OtherResults) > 0 {
		shiftUpload.Teammate0DeathCount = shift.OtherResults[0].DeadCount
		shiftUpload.Teammate0DrizzlerKills = shift.OtherResults[0].BossKillCounts.Drizzler.Count
		shiftUpload.Teammate0FlyfishKills = shift.OtherResults[0].BossKillCounts.Flyfish.Count
		shiftUpload.Teammate0Gender = shift.OtherResults[0].PlayerType.Gender
		shiftUpload.Teammate0GoldenEggs = shift.OtherResults[0].GoldenEggs
		shiftUpload.Teammate0GoldieKills = shift.OtherResults[0].BossKillCounts.Goldie.Count
		shiftUpload.Teammate0GrillerKills = shift.OtherResults[0].BossKillCounts.Griller.Count
		shiftUpload.Teammate0ID = shift.OtherResults[0].Pid
		shiftUpload.Teammate0MawsKills = shift.OtherResults[0].BossKillCounts.Maws.Count
		shiftUpload.Teammate0Name = shift.OtherResults[0].Name
		shiftUpload.Teammate0PowerEggs = shift.OtherResults[0].PowerEggs
		shiftUpload.Teammate0ReviveCount = shift.OtherResults[0].HelpCount
		shiftUpload.Teammate0ScrapperKills = shift.OtherResults[0].BossKillCounts.Scrapper.Count
		shiftUpload.Teammate0Special = shift.OtherResults[0].Special.ID
		shiftUpload.Teammate0Species = shift.OtherResults[0].PlayerType.Species
		shiftUpload.Teammate0SteelEelKills = shift.OtherResults[0].BossKillCounts.SteelEel.Count
		shiftUpload.Teammate0SteelheadKills = shift.OtherResults[0].BossKillCounts.Steelhead.Count
		shiftUpload.Teammate0StingerKills = shift.OtherResults[0].BossKillCounts.Stinger.Count
		shiftUpload.Teammate0W1Specials = shift.OtherResults[0].SpecialCounts[0]
		shiftUpload.Teammate0WeaponW1 = shift.OtherResults[0].WeaponList[0].ID
		if len(shift.OtherResults[0].WeaponList) > 1 {
			shiftUpload.Teammate0WeaponW2 = shift.OtherResults[0].WeaponList[1].ID
			if len(shift.OtherResults[0].WeaponList) > 2 {
				shiftUpload.Teammate0WeaponW3 = shift.OtherResults[0].WeaponList[2].ID
			} else {
				shiftUpload.Teammate0WeaponW3 = ""
			}
		} else {
			shiftUpload.Teammate0WeaponW2 = ""
			shiftUpload.Teammate0WeaponW3 = ""
		}
		if shiftUpload.FailureWave > 1 {
			shiftUpload.Teammate0W2Specials = shift.OtherResults[0].SpecialCounts[1]
			if shiftUpload.FailureWave > 2 {
				shiftUpload.Teammate0W3Specials = shift.OtherResults[0].SpecialCounts[2]
			} else {
				shiftUpload.Teammate0W3Specials = 0
			}
		} else {
			shiftUpload.Teammate0W2Specials = 0
			shiftUpload.Teammate0W3Specials = 0
		}
		if len(shift.OtherResults) > 1 {
			shiftUpload.Teammate1DeathCount = shift.OtherResults[1].DeadCount
			shiftUpload.Teammate1DrizzlerKills = shift.OtherResults[1].BossKillCounts.Drizzler.Count
			shiftUpload.Teammate1FlyfishKills = shift.OtherResults[1].BossKillCounts.Flyfish.Count
			shiftUpload.Teammate1Gender = shift.OtherResults[1].PlayerType.Gender
			shiftUpload.Teammate1GoldenEggs = shift.OtherResults[1].GoldenEggs
			shiftUpload.Teammate1GoldieKills = shift.OtherResults[1].BossKillCounts.Goldie.Count
			shiftUpload.Teammate1GrillerKills = shift.OtherResults[1].BossKillCounts.Griller.Count
			shiftUpload.Teammate1ID = shift.OtherResults[1].Pid
			shiftUpload.Teammate1MawsKills = shift.OtherResults[1].BossKillCounts.Maws.Count
			shiftUpload.Teammate1Name = shift.OtherResults[1].Name
			shiftUpload.Teammate1PowerEggs = shift.OtherResults[1].PowerEggs
			shiftUpload.Teammate1ReviveCount = shift.OtherResults[1].HelpCount
			shiftUpload.Teammate1ScrapperKills = shift.OtherResults[1].BossKillCounts.Scrapper.Count
			shiftUpload.Teammate1Special = shift.OtherResults[1].Special.ID
			shiftUpload.Teammate1Species = shift.OtherResults[1].PlayerType.Species
			shiftUpload.Teammate1SteelEelKills = shift.OtherResults[1].BossKillCounts.SteelEel.Count
			shiftUpload.Teammate1SteelheadKills = shift.OtherResults[1].BossKillCounts.Steelhead.Count
			shiftUpload.Teammate1StingerKills = shift.OtherResults[1].BossKillCounts.Stinger.Count
			shiftUpload.Teammate1W1Specials = shift.OtherResults[1].SpecialCounts[0]
			if len(shift.OtherResults[1].WeaponList) > 0 {
				shiftUpload.Teammate1WeaponW1 = shift.OtherResults[1].WeaponList[0].ID
			} else {
				shiftUpload.Teammate1WeaponW1 = ""
			}
			if len(shift.OtherResults[1].WeaponList) > 2 {
				shiftUpload.Teammate1WeaponW3 = shift.OtherResults[1].WeaponList[2].ID
			} else {
				shiftUpload.Teammate1WeaponW3 = ""
			}
			if len(shift.OtherResults[1].WeaponList) > 1 {
				shiftUpload.Teammate1W2Specials = shift.OtherResults[1].SpecialCounts[1]
				shiftUpload.Teammate1WeaponW2 = shift.OtherResults[1].WeaponList[1].ID
				if shiftUpload.FailureWave > 2 {
					shiftUpload.Teammate1W3Specials = shift.OtherResults[1].SpecialCounts[2]
				} else {
					shiftUpload.Teammate1W3Specials = 0
				}
			} else {
				shiftUpload.Teammate1W2Specials = 0
				shiftUpload.Teammate1W3Specials = 0
				shiftUpload.Teammate1WeaponW2 = ""
				shiftUpload.Teammate1WeaponW3 = ""
			}
			if len(shift.OtherResults) > 2 {
				shiftUpload.Teammate2DeathCount = shift.OtherResults[2].DeadCount
				shiftUpload.Teammate2DrizzlerKills = shift.OtherResults[2].BossKillCounts.Drizzler.Count
				shiftUpload.Teammate2FlyfishKills = shift.OtherResults[2].BossKillCounts.Flyfish.Count
				shiftUpload.Teammate2Gender = shift.OtherResults[2].PlayerType.Gender
				shiftUpload.Teammate2GoldenEggs = shift.OtherResults[2].GoldenEggs
				shiftUpload.Teammate2GoldieKills = shift.OtherResults[2].BossKillCounts.Goldie.Count
				shiftUpload.Teammate2GrillerKills = shift.OtherResults[2].BossKillCounts.Griller.Count
				shiftUpload.Teammate2ID = shift.OtherResults[2].Pid
				shiftUpload.Teammate2MawsKills = shift.OtherResults[2].BossKillCounts.Maws.Count
				shiftUpload.Teammate2Name = shift.OtherResults[2].Name
				shiftUpload.Teammate2PowerEggs = shift.OtherResults[2].PowerEggs
				shiftUpload.Teammate2ReviveCount = shift.OtherResults[2].HelpCount
				shiftUpload.Teammate2ScrapperKills = shift.OtherResults[2].BossKillCounts.Scrapper.Count
				shiftUpload.Teammate2Special = shift.OtherResults[2].Special.ID
				shiftUpload.Teammate2Species = shift.OtherResults[2].PlayerType.Species
				shiftUpload.Teammate2SteelEelKills = shift.OtherResults[2].BossKillCounts.SteelEel.Count
				shiftUpload.Teammate2SteelheadKills = shift.OtherResults[2].BossKillCounts.Steelhead.Count
				shiftUpload.Teammate2StingerKills = shift.OtherResults[2].BossKillCounts.Stinger.Count
				shiftUpload.Teammate2W1Specials = shift.OtherResults[2].SpecialCounts[0]
				shiftUpload.Teammate2WeaponW1 = shift.OtherResults[2].WeaponList[0].ID
				if len(shift.OtherResults[2].WeaponList) > 1 {
					shiftUpload.Teammate2WeaponW2 = shift.OtherResults[2].WeaponList[1].ID
					if len(shift.OtherResults[2].WeaponList) > 2 {
						shiftUpload.Teammate2WeaponW3 = shift.OtherResults[2].WeaponList[2].ID
					} else {
						shiftUpload.Teammate2WeaponW3 = ""
					}
				} else {
					shiftUpload.Teammate2WeaponW2 = ""
					shiftUpload.Teammate2WeaponW3 = ""
				}
				if shiftUpload.FailureWave > 1 {
					shiftUpload.Teammate2W2Specials = shift.OtherResults[2].SpecialCounts[1]
					if len(shiftUpload.SplatnetJSON.OtherResults[2].SpecialCounts) > 2 {
						shiftUpload.Teammate2W3Specials = shift.OtherResults[2].SpecialCounts[2]
					} else {
						shiftUpload.Teammate2W3Specials = 0
					}
				} else {
					shiftUpload.Teammate2W2Specials = 0
					shiftUpload.Teammate2W3Specials = 0
				}
			}
		}
	}
	shiftUpload.Wave1EventType = shift.WaveDetails[0].EventType.Key
	shiftUpload.Wave1GoldenAppear = shift.WaveDetails[0].GoldenAppear
	shiftUpload.Wave1GoldenDelivered = shift.WaveDetails[0].GoldenEggs
	shiftUpload.Wave1PowerEggs = shift.WaveDetails[0].PowerEggs
	shiftUpload.Wave1Quota = shift.WaveDetails[0].QuotaNum
	shiftUpload.Wave1WaterLevel = shift.WaveDetails[0].WaterLevel.Key
	if len(shift.WaveDetails) > 1 {
		shiftUpload.Wave2EventType = shift.WaveDetails[1].EventType.Key
		shiftUpload.Wave2GoldenAppear = shift.WaveDetails[1].GoldenAppear
		shiftUpload.Wave2GoldenDelivered = shift.WaveDetails[1].GoldenEggs
		shiftUpload.Wave2PowerEggs = shift.WaveDetails[1].PowerEggs
		shiftUpload.Wave2Quota = shift.WaveDetails[1].QuotaNum
		shiftUpload.Wave2WaterLevel = shift.WaveDetails[1].WaterLevel.Key
		if len(shift.WaveDetails) > 2 {
			shiftUpload.Wave3EventType = shift.WaveDetails[2].EventType.Key
			shiftUpload.Wave3GoldenAppear = shift.WaveDetails[2].GoldenAppear
			shiftUpload.Wave3GoldenDelivered = shift.WaveDetails[2].GoldenEggs
			shiftUpload.Wave3PowerEggs = shift.WaveDetails[2].PowerEggs
			shiftUpload.Wave3Quota = shift.WaveDetails[2].QuotaNum
			shiftUpload.Wave3WaterLevel = shift.WaveDetails[2].WaterLevel.Key
		} else {
			shiftUpload.Wave3EventType = ""
			shiftUpload.Wave3GoldenAppear = 0
			shiftUpload.Wave3GoldenDelivered = 0
			shiftUpload.Wave3PowerEggs = 0
			shiftUpload.Wave3Quota = 0
			shiftUpload.Wave3WaterLevel = ""
		}
	} else {
		shiftUpload.Wave3EventType = ""
		shiftUpload.Wave3GoldenAppear = 0
		shiftUpload.Wave3GoldenDelivered = 0
		shiftUpload.Wave3PowerEggs = 0
		shiftUpload.Wave3Quota = 0
		shiftUpload.Wave3WaterLevel = ""
		shiftUpload.Wave2EventType = ""
		shiftUpload.Wave2GoldenAppear = 0
		shiftUpload.Wave2GoldenDelivered = 0
		shiftUpload.Wave2PowerEggs = 0
		shiftUpload.Wave2Quota = 0
		shiftUpload.Wave2WaterLevel = ""
	}
	url := "https://splatstats.cass-dlcm.dev/two_salmon/api/shifts/"
	auth := map[string]string{
		"Authorization": "Token " + api_key,
		"Content-Type":  "application/json",
	}
	body_marshalled, err := json.Marshal(shiftUpload)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body_marshalled))
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
		fmt.Println(err)
	}
	bodyString := string(body)
	if resp.StatusCode == 400 && bodyString == "{\"non_field_errors\":[\"The fields player_id, job_id must make a unique set.\"]}" {
		fmt.Printf("Shift #%d already uploaded\n", shiftUpload.JobID)
	} else if resp.StatusCode == 201 {
		fmt.Printf("Shift #%d uploaded to %s\n", shiftUpload.JobID, resp.Header.Get("location"))
	} else {
		fmt.Println(resp.Status)
		fmt.Println(bodyString)
		panic(nil)
	}
}
