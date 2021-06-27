package types

// BattleList is a list of up to 50 partial Battles, with additional metadata.
type BattleList struct {
	Code     *string `json:"code"`
	UniqueID string  `json:"unique_id"`
	Summary  struct {
		AssistCountAverage  float64 `json:"assist_count_average"`
		VictoryRate         float64 `json:"victory_rate"`
		VictoryCount        int     `json:"victory_count"`
		DeathCountAverage   float64 `json:"death_count_average"`
		Count               int     `json:"count"`
		SpecialCountAverage float64 `json:"special_count_average"`
		DefeatCount         int     `json:"defeat_count"`
		KillCountAverage    float64 `json:"kill_count_average"`
	} `json:"summary"`
	Results []BattleSplatnet `json:"results"`
}

// ShiftList is a list of up to 50 partial Shifts, with additional metadata.
type ShiftList struct {
	Code    *string `json:"code"`
	Summary struct {
		Card struct {
			GoldenIkuraTotal int `json:"golden_ikura_total"`
			HelpTotal        int `json:"help_total"`
			KumaPointTotal   int `json:"kuma_point_total"`
			IkuraTotal       int `json:"ikura_total"`
			KumaPoint        int `json:"kuma_point"`
			JobNum           int `json:"job_num"`
		} `json:"card"`
		Stats []struct {
			DeadTotal            int   `json:"dead_total"`
			MyGoldenIkuraTotal   int   `json:"my_golden_ikura_total"`
			GradePoint           int   `json:"grade_point"`
			TeamGoldenIkuraTotal int   `json:"team_golden_ikura_total"`
			HelpTotal            int   `json:"help_total"`
			TeamIkuraTotal       int   `json:"team_ikura_total"`
			StartTime            int   `json:"start_time"`
			MyIkuraTotal         int   `json:"my_ikura_total"`
			FailureCounts        []int `json:"failure_counts"`
			Schedule             struct {
				Stage struct {
					Image string `json:"image"`
					Name  string `json:"name"`
				} `json:"stage"`
				EndTime   int `json:"end_time"`
				StartTime int `json:"start_time"`
				Weapons   []struct {
					Weapon struct {
						ID        string `json:"id"`
						Image     string `json:"image"`
						Name      string `json:"name"`
						Thumbnail string `json:"thumbnail"`
					} `json:"weapon"`
					ID string `json:"id"`
				} `json:"weapons"`
			} `json:"schedule"`
			JobNum         int `json:"job_num"`
			KumaPointTotal int `json:"kuma_point_total"`
			EndTime        int `json:"end_time"`
			ClearNum       int `json:"clear_num"`
			Grade          struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"grade"`
		} `json:"stats"`
	} `json:"summary"`
	RewardGear struct {
		Thumbnail string `json:"thumbnail"`
		Kind      string `json:"kind"`
		ID        string `json:"id"`
		Name      string `json:"name"`
		Brand     struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"brand"`
		Rarity int    `json:"rarity"`
		Image  string `json:"image"`
	} `json:"reward_gear"`
	Results []ShiftSplatnet `json:"results"`
}

type ShiftStatInkArray []ShiftStatInk

type BattleStatInkArray []BattleStatInk
