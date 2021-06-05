package statink2splatstats

import (
	"context"
	"encoding/json"
	"fmt"
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

	DownloadShifts(&allData, client)

	for i := range allData {
		var shift types.ShiftUpload

		TransformShift(&allData[i], &shift)
	}
}

func DownloadShifts(allData *[]types.ShiftStatInk, client *http.Client) {
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

func TransformShift(statInkShift *types.ShiftStatInk, shiftUpload *types.ShiftUpload) {
	f, err := strconv.ParseFloat((*statInkShift).DangerRate, 64)
	if err != nil {
		panic(err)
	}

	*(*shiftUpload).DangerRate = f
	*(*shiftUpload).StatInkUpload = true
	*(*shiftUpload).SplatnetUpload = false
	*(*shiftUpload).StatInkJson = *statInkShift

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

	*(*shiftUpload).Stage = (*statInkShift).Stage.Splatnet
	*(*shiftUpload).Playtime = (*statInkShift).StartAt.Iso8601.Format("2006-01-02 15:04:05")
	*(*shiftUpload).ScheduleStarttime = (*statInkShift).ShiftStartAt.Iso8601.Format("2006-01-02 15:04:05")
	*(*shiftUpload).GradePointDelta = (*statInkShift).TitleExpAfter - (*statInkShift).TitleExp
	*(*shiftUpload).GradePoint = (*statInkShift).TitleExpAfter
	*(*shiftUpload).JobFailureReason = (*statInkShift).FailReason
	*(*shiftUpload).IsClear = (*statInkShift).IsCleared
	*(*shiftUpload).FailureWave = (*statInkShift).ClearWaves
	*(*shiftUpload).JobID = (*statInkShift).SplatnetNumber

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

func MigrateBattles(api_key string, client *http.Client) {
	if viper.GetString("statink_api_key") == "" {
		enterStatinkApiKey()
	}
	// statinkApiKey := viper.GetString("statink_api_key")
}
