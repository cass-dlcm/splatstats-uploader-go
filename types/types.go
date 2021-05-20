package types

import "github.com/shopspring/decimal"

type Shift struct {
	DangerRate float64 `json:"danger_rate"`
	JobResult  struct {
		IsClear       bool   `json:"is_clear"`
		FailureReason string `json:"failure_reason"`
		FailureWave   int    `json:"failure_wave"`
	} `json:"job_result"`
	JobScore        int `json:"job_score"`
	JobID           int `json:"job_id"`
	JobRate         int `json:"job_rate"`
	GradePoint      int `json:"grade_point"`
	GradePointDelta int `json:"grade_point_delta"`
	OtherResults    []struct {
		SpecialCounts []int `json:"special_counts"`
		Special       struct {
			ImageB string `json:"image_b"`
			Name   string `json:"name"`
			ID     string `json:"id"`
			ImageA string `json:"image_a"`
		} `json:"special"`
		Pid        string `json:"pid"`
		PlayerType struct {
			Gender  string `json:"style"`
			Species string `json:"species"`
		} `json:"player_type"`
		WeaponList []struct {
			ID     string `json:"id"`
			Weapon struct {
				ID        string `json:"id"`
				Image     string `json:"image"`
				Name      string `json:"name"`
				Thumbnail string `json:"thumbnail"`
			} `json:"weapon"`
		} `json:"weapon_list"`
		Name           string `json:"name"`
		DeadCount      int    `json:"dead_count"`
		GoldenEggs     int    `json:"golden_ikura_num"`
		BossKillCounts struct {
			Goldie struct {
				Boss struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"3"`
			Steelhead struct {
				Boss struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"6"`
			Flyfish struct {
				Boss struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"9"`
			Scrapper struct {
				Count int `json:"count"`
				Boss  struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"boss"`
			} `json:"12"`
			SteelEel struct {
				Count int `json:"count"`
				Boss  struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
			} `json:"13"`
			Stinger struct {
				Count int `json:"count"`
				Boss  struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"boss"`
			} `json:"14"`
			Maws struct {
				Boss struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"15"`
			Griller struct {
				Boss struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"16"`
			Drizzler struct {
				Count int `json:"count"`
				Boss  struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"boss"`
			} `json:"21"`
		} `json:"boss_kill_counts"`
		PowerEggs int `json:"ikura_num"`
		HelpCount int `json:"help_count"`
	} `json:"other_results"`
	KumaPoint  int `json:"kuma_point"`
	StartTime  int `json:"start_time"`
	PlayerType struct {
		Species string `json:"species"`
		Gender  string `json:"style"`
	} `json:"player_type"`
	PlayTime   int `json:"play_time"`
	BossCounts struct {
		Goldie struct {
			Count int `json:"count"`
			Boss  struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"boss"`
		} `json:"3"`
		Steelhead struct {
			Boss struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"boss"`
			Count int `json:"count"`
		} `json:"6"`
		Flyfish struct {
			Boss struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"boss"`
			Count int `json:"count"`
		} `json:"9"`
		Scrapper struct {
			Boss struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"boss"`
			Count int `json:"count"`
		} `json:"12"`
		SteelEel struct {
			Count int `json:"count"`
			Boss  struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"boss"`
		} `json:"13"`
		Stinger struct {
			Count int `json:"count"`
			Boss  struct {
				Key  string `json:"key"`
				Name string `json:"name"`
			} `json:"boss"`
		} `json:"14"`
		Maws struct {
			Boss struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"boss"`
			Count int `json:"count"`
		} `json:"15"`
		Griller struct {
			Boss struct {
				Key  string `json:"key"`
				Name string `json:"name"`
			} `json:"boss"`
			Count int `json:"count"`
		} `json:"16"`
		Drizzler struct {
			Boss struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"boss"`
			Count int `json:"count"`
		} `json:"21"`
	} `json:"boss_counts"`
	EndTime  int `json:"end_time"`
	MyResult struct {
		GoldenEggs     int `json:"golden_ikura_num"`
		DeadCount      int `json:"dead_count"`
		HelpCount      int `json:"help_count"`
		BossKillCounts struct {
			Goldie struct {
				Boss struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"3"`
			Steelhead struct {
				Count int `json:"count"`
				Boss  struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
			} `json:"6"`
			Flyfish struct {
				Boss struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"9"`
			Scrapper struct {
				Boss struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"12"`
			SteelEel struct {
				Boss struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"13"`
			Stinger struct {
				Count int `json:"count"`
				Boss  struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"boss"`
			} `json:"14"`
			Maws struct {
				Count int `json:"count"`
				Boss  struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
			} `json:"15"`
			Griller struct {
				Count int `json:"count"`
				Boss  struct {
					Name string `json:"name"`
					Key  string `json:"key"`
				} `json:"boss"`
			} `json:"16"`
			Drizzler struct {
				Boss struct {
					Key  string `json:"key"`
					Name string `json:"name"`
				} `json:"boss"`
				Count int `json:"count"`
			} `json:"21"`
		} `json:"boss_kill_counts"`
		PowerEggs int `json:"ikura_num"`
		Special   struct {
			Name   string `json:"name"`
			ImageB string `json:"image_b"`
			ImageA string `json:"image_a"`
			ID     string `json:"id"`
		} `json:"special"`
		SpecialCounts []int  `json:"special_counts"`
		Name          string `json:"name"`
		WeaponList    []struct {
			ID     string `json:"id"`
			Weapon struct {
				ID        string `json:"id"`
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
				Image     string `json:"image"`
			} `json:"weapon"`
		} `json:"weapon_list"`
		PlayerType struct {
			Species string `json:"species"`
			Gender  string `json:"style"`
		} `json:"player_type"`
		Pid string `json:"pid"`
	} `json:"my_result"`
	WaveDetails []struct {
		WaterLevel struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"water_level"`
		EventType struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"event_type"`
		GoldenEggs   int `json:"golden_ikura_num"`
		GoldenAppear int `json:"golden_ikura_pop_num"`
		PowerEggs    int `json:"ikura_num"`
		QuotaNum     int `json:"quota_num"`
	} `json:"wave_details"`
	Grade struct {
		ID        string `json:"id"`
		ShortName string `json:"short_name"`
		LongName  string `json:"long_name"`
		Name      string `json:"name"`
	} `json:"grade"`
	Schedule struct {
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
		EndTime int `json:"end_time"`
		Stage   struct {
			Image string `json:"image"`
			Name  string `json:"name"`
		} `json:"stage"`
	} `json:"schedule"`
}

type Battle struct {
	Udemae struct {
		Name            string      `json:"name"`
		IsX             bool        `json:"is_x"`
		IsNumberReached bool        `json:"is_number_reached"`
		Number          int         `json:"number"`
		SPlusNumber     interface{} `json:"s_plus_number"`
	} `json:"udemae"`
	Stage struct {
		ID    string `json:"id"`
		Image string `json:"image"`
		Name  string `json:"name"`
	} `json:"stage"`
	OtherTeamCount int `json:"other_team_count"`
	MyTeamCount    int `json:"my_team_count"`
	StarRank       int `json:"star_rank"`
	Rule           struct {
		Name          string `json:"name"`
		Key           string `json:"key"`
		MultilineName string `json:"multiline_name"`
	} `json:"rule"`
	PlayerResult struct {
		DeathCount     int `json:"death_count"`
		GamePaintPoint int `json:"game_paint_point"`
		KillCount      int `json:"kill_count"`
		SpecialCount   int `json:"special_count"`
		AssistCount    int `json:"assist_count"`
		SortScore      int `json:"sort_score"`
		Player         struct {
			HeadSkills struct {
				Main struct {
					Name  string `json:"name"`
					ID    string `json:"id"`
					Image string `json:"image"`
				} `json:"main"`
				Subs []struct {
					Name  string `json:"name"`
					ID    string `json:"id"`
					Image string `json:"image"`
				} `json:"subs"`
			} `json:"head_skills"`
			ShoesSkills struct {
				Subs []struct {
					Name  string `json:"name"`
					ID    string `json:"id"`
					Image string `json:"image"`
				} `json:"subs"`
				Main struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					ID    string `json:"id"`
				} `json:"main"`
			} `json:"shoes_skills"`
			PlayerRank int `json:"player_rank"`
			PlayerType struct {
				Gender  string `json:"style"`
				Species string `json:"species"`
			} `json:"player_type"`
			PrincipalID string `json:"principal_id"`
			Head        struct {
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
				Image     string `json:"image"`
				Kind      string `json:"kind"`
				Rarity    int    `json:"rarity"`
				Brand     struct {
					FrequentSkill struct {
						Image string `json:"image"`
						ID    string `json:"id"`
						Name  string `json:"name"`
					} `json:"frequent_skill"`
					Image string `json:"image"`
					ID    string `json:"id"`
					Name  string `json:"name"`
				} `json:"brand"`
				ID string `json:"id"`
			} `json:"head"`
			Nickname string `json:"nickname"`
			Clothes  struct {
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
				Image     string `json:"image"`
				Kind      string `json:"kind"`
				Rarity    int    `json:"rarity"`
				ID        string `json:"id"`
				Brand     struct {
					Name          string `json:"name"`
					FrequentSkill struct {
						Name  string `json:"name"`
						ID    string `json:"id"`
						Image string `json:"image"`
					} `json:"frequent_skill"`
					Image string `json:"image"`
					ID    string `json:"id"`
				} `json:"brand"`
			} `json:"clothes"`
			Udemae struct {
				SPlusNumber     interface{} `json:"s_plus_number"`
				Number          int         `json:"number"`
				IsNumberReached bool        `json:"is_number_reached"`
				IsX             bool        `json:"is_x"`
				Name            string      `json:"name"`
			} `json:"udemae"`
			Shoes struct {
				Image     string `json:"image"`
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
				ID        string `json:"id"`
				Brand     struct {
					Image         string `json:"image"`
					FrequentSkill struct {
						ID    string `json:"id"`
						Image string `json:"image"`
						Name  string `json:"name"`
					} `json:"frequent_skill"`
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"brand"`
				Rarity int    `json:"rarity"`
				Kind   string `json:"kind"`
			} `json:"shoes"`
			StarRank int `json:"star_rank"`
			Weapon   struct {
				Sub struct {
					Name   string `json:"name"`
					ImageA string `json:"image_a"`
					ID     string `json:"id"`
					ImageB string `json:"image_b"`
				} `json:"sub"`
				Special struct {
					ImageA string `json:"image_a"`
					Name   string `json:"name"`
					ID     string `json:"id"`
					ImageB string `json:"image_b"`
				} `json:"special"`
				Image     string `json:"image"`
				ID        string `json:"id"`
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
			} `json:"weapon"`
			ClothesSkills struct {
				Subs []struct {
					Name  string `json:"name"`
					ID    string `json:"id"`
					Image string `json:"image"`
				} `json:"subs"`
				Main struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					ID    string `json:"id"`
				} `json:"main"`
			} `json:"clothes_skills"`
		} `json:"player"`
	} `json:"player_result"`
	EstimateGachiPower int `json:"estimate_gachi_power"`
	ElapsedTime        int `json:"elapsed_time"`
	StartTime          int `json:"start_time"`
	GameMode           struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"game_mode"`
	XPower        interface{} `json:"x_power"`
	BattleNumber  string      `json:"battle_number"`
	Type          string      `json:"type"`
	PlayerRank    int         `json:"player_rank"`
	CrownPlayers  interface{} `json:"crown_players"`
	MyTeamMembers []struct {
		GamePaintPoint int `json:"game_paint_point"`
		DeathCount     int `json:"death_count"`
		SortScore      int `json:"sort_score"`
		AssistCount    int `json:"assist_count"`
		KillCount      int `json:"kill_count"`
		SpecialCount   int `json:"special_count"`
		Player         struct {
			Shoes struct {
				Brand struct {
					FrequentSkill struct {
						ID    string `json:"id"`
						Image string `json:"image"`
						Name  string `json:"name"`
					} `json:"frequent_skill"`
					Image string `json:"image"`
					ID    string `json:"id"`
					Name  string `json:"name"`
				} `json:"brand"`
				ID        string `json:"id"`
				Kind      string `json:"kind"`
				Rarity    int    `json:"rarity"`
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
				Image     string `json:"image"`
			} `json:"shoes"`
			Weapon struct {
				Name      string `json:"name"`
				Thumbnail string `json:"thumbnail"`
				Image     string `json:"image"`
				ID        string `json:"id"`
				Special   struct {
					ImageB string `json:"image_b"`
					ID     string `json:"id"`
					ImageA string `json:"image_a"`
					Name   string `json:"name"`
				} `json:"special"`
				Sub struct {
					ID     string `json:"id"`
					Name   string `json:"name"`
					ImageA string `json:"image_a"`
					ImageB string `json:"image_b"`
				} `json:"sub"`
			} `json:"weapon"`
			StarRank      int `json:"star_rank"`
			ClothesSkills struct {
				Subs []struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					ID    string `json:"id"`
				} `json:"subs"`
				Main struct {
					Name  string `json:"name"`
					ID    string `json:"id"`
					Image string `json:"image"`
				} `json:"main"`
			} `json:"clothes_skills"`
			PrincipalID string `json:"principal_id"`
			Head        struct {
				ID    string `json:"id"`
				Brand struct {
					Name          string `json:"name"`
					FrequentSkill struct {
						Name  string `json:"name"`
						ID    string `json:"id"`
						Image string `json:"image"`
					} `json:"frequent_skill"`
					Image string `json:"image"`
					ID    string `json:"id"`
				} `json:"brand"`
				Kind      string `json:"kind"`
				Rarity    int    `json:"rarity"`
				Name      string `json:"name"`
				Thumbnail string `json:"thumbnail"`
				Image     string `json:"image"`
			} `json:"head"`
			Clothes struct {
				Image     string `json:"image"`
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
				Rarity    int    `json:"rarity"`
				Kind      string `json:"kind"`
				Brand     struct {
					Name          string `json:"name"`
					ID            string `json:"id"`
					FrequentSkill struct {
						Name  string `json:"name"`
						ID    string `json:"id"`
						Image string `json:"image"`
					} `json:"frequent_skill"`
					Image string `json:"image"`
				} `json:"brand"`
				ID string `json:"id"`
			} `json:"clothes"`
			Nickname string `json:"nickname"`
			Udemae   struct {
				SPlusNumber interface{} `json:"s_plus_number"`
				IsX         bool        `json:"is_x"`
				Name        string      `json:"name"`
			} `json:"udemae"`
			PlayerType struct {
				Species string `json:"species"`
				Gender  string `json:"style"`
			} `json:"player_type"`
			HeadSkills struct {
				Main struct {
					Name  string `json:"name"`
					ID    string `json:"id"`
					Image string `json:"image"`
				} `json:"main"`
				Subs []struct {
					ID    string `json:"id"`
					Image string `json:"image"`
					Name  string `json:"name"`
				} `json:"subs"`
			} `json:"head_skills"`
			PlayerRank  int `json:"player_rank"`
			ShoesSkills struct {
				Subs []struct {
					ID    string `json:"id"`
					Image string `json:"image"`
					Name  string `json:"name"`
				} `json:"subs"`
				Main struct {
					Name  string `json:"name"`
					ID    string `json:"id"`
					Image string `json:"image"`
				} `json:"main"`
			} `json:"shoes_skills"`
		} `json:"player"`
	} `json:"my_team_members"`
	OtherTeamMembers []struct {
		GamePaintPoint int `json:"game_paint_point"`
		DeathCount     int `json:"death_count"`
		SpecialCount   int `json:"special_count"`
		KillCount      int `json:"kill_count"`
		SortScore      int `json:"sort_score"`
		AssistCount    int `json:"assist_count"`
		Player         struct {
			PlayerRank  int `json:"player_rank"`
			ShoesSkills struct {
				Subs []struct {
					ID    string `json:"id"`
					Image string `json:"image"`
					Name  string `json:"name"`
				} `json:"subs"`
				Main struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					ID    string `json:"id"`
				} `json:"main"`
			} `json:"shoes_skills"`
			HeadSkills struct {
				Subs []interface{} `json:"subs"`
				Main struct {
					Image string `json:"image"`
					ID    string `json:"id"`
					Name  string `json:"name"`
				} `json:"main"`
			} `json:"head_skills"`
			PlayerType struct {
				Gender  string `json:"style"`
				Species string `json:"species"`
			} `json:"player_type"`
			Udemae struct {
				SPlusNumber interface{} `json:"s_plus_number"`
				IsX         bool        `json:"is_x"`
				Name        string      `json:"name"`
			} `json:"udemae"`
			Clothes struct {
				Name      string `json:"name"`
				Thumbnail string `json:"thumbnail"`
				Image     string `json:"image"`
				Kind      string `json:"kind"`
				Rarity    int    `json:"rarity"`
				Brand     struct {
					Name          string `json:"name"`
					FrequentSkill struct {
						Name  string `json:"name"`
						ID    string `json:"id"`
						Image string `json:"image"`
					} `json:"frequent_skill"`
					Image string `json:"image"`
					ID    string `json:"id"`
				} `json:"brand"`
				ID string `json:"id"`
			} `json:"clothes"`
			Nickname string `json:"nickname"`
			Head     struct {
				Kind   string `json:"kind"`
				Rarity int    `json:"rarity"`
				Brand  struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					ID    string `json:"id"`
				} `json:"brand"`
				ID        string `json:"id"`
				Name      string `json:"name"`
				Thumbnail string `json:"thumbnail"`
				Image     string `json:"image"`
			} `json:"head"`
			PrincipalID   string `json:"principal_id"`
			ClothesSkills struct {
				Main struct {
					Image string `json:"image"`
					ID    string `json:"id"`
					Name  string `json:"name"`
				} `json:"main"`
				Subs []struct {
					ID    string `json:"id"`
					Image string `json:"image"`
					Name  string `json:"name"`
				} `json:"subs"`
			} `json:"clothes_skills"`
			StarRank int `json:"star_rank"`
			Weapon   struct {
				ID        string `json:"id"`
				Image     string `json:"image"`
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
				Sub       struct {
					ImageB string `json:"image_b"`
					Name   string `json:"name"`
					ImageA string `json:"image_a"`
					ID     string `json:"id"`
				} `json:"sub"`
				Special struct {
					ImageA string `json:"image_a"`
					Name   string `json:"name"`
					ID     string `json:"id"`
					ImageB string `json:"image_b"`
				} `json:"special"`
			} `json:"weapon"`
			Shoes struct {
				Kind   string `json:"kind"`
				Rarity int    `json:"rarity"`
				Brand  struct {
					Name          string `json:"name"`
					Image         string `json:"image"`
					FrequentSkill struct {
						Name  string `json:"name"`
						Image string `json:"image"`
						ID    string `json:"id"`
					} `json:"frequent_skill"`
					ID string `json:"id"`
				} `json:"brand"`
				ID        string `json:"id"`
				Thumbnail string `json:"thumbnail"`
				Name      string `json:"name"`
				Image     string `json:"image"`
			} `json:"shoes"`
		} `json:"player"`
	} `json:"other_team_members"`
	WeaponPaintPoint int         `json:"weapon_paint_point"`
	Rank             interface{} `json:"rank"`
	MyTeamResult     struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	} `json:"my_team_result"`
	EstimateXPower  interface{} `json:"estimate_x_power"`
	OtherTeamResult struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"other_team_result"`
}

type ShiftUpload struct {
	DangerRate              float64 `json:"danger_rate"`
	DrizzlerCount           int     `json:"drizzler_count"`
	Endtime                 string  `json:"endtime"`
	FailureWave             int     `json:"failure_wave"`
	FlyfishCount            int     `json:"flyfish_count"`
	GoldieCount             int     `json:"goldie_count"`
	GradePoint              int     `json:"grade_point"`
	GradePointDelta         int     `json:"grade_point_delta"`
	GrillerCount            int     `json:"griller_count"`
	IsClear                 bool    `json:"is_clear"`
	JobFailureReason        string  `json:"job_failure_reason"`
	JobID                   int     `json:"job_id"`
	MawsCount               int     `json:"maws_count"`
	PlayerDeathCount        int     `json:"player_death_count"`
	PlayerDrizzlerKills     int     `json:"player_drizzler_kills"`
	PlayerFlyfishKills      int     `json:"player_flyfish_kills"`
	PlayerGender            string  `json:"player_gender"`
	PlayerGoldenEggs        int     `json:"player_golden_eggs"`
	PlayerGoldieKills       int     `json:"player_goldie_kills"`
	PlayerGrillerKills      int     `json:"player_griller_kills"`
	PlayerID                string  `json:"player_id"`
	PlayerMawsKills         int     `json:"player_maws_kills"`
	PlayerName              string  `json:"player_name"`
	PlayerPowerEggs         int     `json:"player_power_eggs"`
	PlayerReviveCount       int     `json:"player_revive_count"`
	PlayerScrapperKills     int     `json:"player_scrapper_kills"`
	PlayerSpecial           string  `json:"player_special"`
	PlayerSpecies           string  `json:"player_species"`
	PlayerSteelEelKills     int     `json:"player_steel_eel_kills"`
	PlayerSteelheadKills    int     `json:"player_steelhead_kills"`
	PlayerStingerKills      int     `json:"player_stinger_kills"`
	PlayerTitle             string  `json:"player_title"`
	PlayerW1Specials        int     `json:"player_w1_specials"`
	PlayerW2Specials        int     `json:"player_w2_specials"`
	PlayerW3Specials        int     `json:"player_w3_specials"`
	PlayerWeaponW1          string  `json:"player_weapon_w1"`
	PlayerWeaponW2          string  `json:"player_weapon_w2"`
	PlayerWeaponW3          string  `json:"player_weapon_w3"`
	Playtime                string  `json:"playtime"`
	ScheduleEndtime         string  `json:"schedule_endtime"`
	ScheduleStarttime       string  `json:"schedule_starttime"`
	ScheduleWeapon0         string  `json:"schedule_weapon_0"`
	ScheduleWeapon1         string  `json:"schedule_weapon_1"`
	ScheduleWeapon2         string  `json:"schedule_weapon_2"`
	ScheduleWeapon3         string  `json:"schedule_weapon_3"`
	ScrapperCount           int     `json:"scrapper_count"`
	SplatnetJSON            Shift   `json:"splatnet_json"`
	SplatnetUpload          bool    `json:"splatnet_upload"`
	Stage                   string  `json:"stage"`
	StatInkUpload           bool    `json:"stat_ink_upload"`
	Starttime               string  `json:"starttime"`
	SteelEelCount           int     `json:"steel_eel_count"`
	SteelheadCount          int     `json:"steelhead_count"`
	StingerCount            int     `json:"stinger_count"`
	Teammate0DeathCount     int     `json:"teammate0_death_count"`
	Teammate0DrizzlerKills  int     `json:"teammate0_drizzler_kills"`
	Teammate0FlyfishKills   int     `json:"teammate0_flyfish_kills"`
	Teammate0Gender         string  `json:"teammate0_gender"`
	Teammate0GoldenEggs     int     `json:"teammate0_golden_eggs"`
	Teammate0GoldieKills    int     `json:"teammate0_goldie_kills"`
	Teammate0GrillerKills   int     `json:"teammate0_griller_kills"`
	Teammate0ID             string  `json:"teammate0_id"`
	Teammate0MawsKills      int     `json:"teammate0_maws_kills"`
	Teammate0Name           string  `json:"teammate0_name"`
	Teammate0PowerEggs      int     `json:"teammate0_power_eggs"`
	Teammate0ReviveCount    int     `json:"teammate0_revive_count"`
	Teammate0ScrapperKills  int     `json:"teammate0_scrapper_kills"`
	Teammate0Special        string  `json:"teammate0_special"`
	Teammate0Species        string  `json:"teammate0_species"`
	Teammate0SteelEelKills  int     `json:"teammate0_steel_eel_kills"`
	Teammate0SteelheadKills int     `json:"teammate0_steelhead_kills"`
	Teammate0StingerKills   int     `json:"teammate0_stinger_kills"`
	Teammate0W1Specials     int     `json:"teammate0_w1_specials"`
	Teammate0W2Specials     int     `json:"teammate0_w2_specials"`
	Teammate0W3Specials     int     `json:"teammate0_w3_specials"`
	Teammate0WeaponW1       string  `json:"teammate0_weapon_w1"`
	Teammate0WeaponW2       string  `json:"teammate0_weapon_w2"`
	Teammate0WeaponW3       string  `json:"teammate0_weapon_w3"`
	Teammate1DeathCount     int     `json:"teammate1_death_count"`
	Teammate1DrizzlerKills  int     `json:"teammate1_drizzler_kills"`
	Teammate1FlyfishKills   int     `json:"teammate1_flyfish_kills"`
	Teammate1Gender         string  `json:"teammate1_gender"`
	Teammate1GoldenEggs     int     `json:"teammate1_golden_eggs"`
	Teammate1GoldieKills    int     `json:"teammate1_goldie_kills"`
	Teammate1GrillerKills   int     `json:"teammate1_griller_kills"`
	Teammate1ID             string  `json:"teammate1_id"`
	Teammate1MawsKills      int     `json:"teammate1_maws_kills"`
	Teammate1Name           string  `json:"teammate1_name"`
	Teammate1PowerEggs      int     `json:"teammate1_power_eggs"`
	Teammate1ReviveCount    int     `json:"teammate1_revive_count"`
	Teammate1ScrapperKills  int     `json:"teammate1_scrapper_kills"`
	Teammate1Special        string  `json:"teammate1_special"`
	Teammate1Species        string  `json:"teammate1_species"`
	Teammate1SteelEelKills  int     `json:"teammate1_steel_eel_kills"`
	Teammate1SteelheadKills int     `json:"teammate1_steelhead_kills"`
	Teammate1StingerKills   int     `json:"teammate1_stinger_kills"`
	Teammate1W1Specials     int     `json:"teammate1_w1_specials"`
	Teammate1W2Specials     int     `json:"teammate1_w2_specials"`
	Teammate1W3Specials     int     `json:"teammate1_w3_specials"`
	Teammate1WeaponW1       string  `json:"teammate1_weapon_w1"`
	Teammate1WeaponW2       string  `json:"teammate1_weapon_w2"`
	Teammate1WeaponW3       string  `json:"teammate1_weapon_w3"`
	Teammate2DeathCount     int     `json:"teammate2_death_count"`
	Teammate2DrizzlerKills  int     `json:"teammate2_drizzler_kills"`
	Teammate2FlyfishKills   int     `json:"teammate2_flyfish_kills"`
	Teammate2Gender         string  `json:"teammate2_gender"`
	Teammate2GoldenEggs     int     `json:"teammate2_golden_eggs"`
	Teammate2GoldieKills    int     `json:"teammate2_goldie_kills"`
	Teammate2GrillerKills   int     `json:"teammate2_griller_kills"`
	Teammate2ID             string  `json:"teammate2_id"`
	Teammate2MawsKills      int     `json:"teammate2_maws_kills"`
	Teammate2Name           string  `json:"teammate2_name"`
	Teammate2PowerEggs      int     `json:"teammate2_power_eggs"`
	Teammate2ReviveCount    int     `json:"teammate2_revive_count"`
	Teammate2ScrapperKills  int     `json:"teammate2_scrapper_kills"`
	Teammate2Special        string  `json:"teammate2_special"`
	Teammate2Species        string  `json:"teammate2_species"`
	Teammate2SteelEelKills  int     `json:"teammate2_steel_eel_kills"`
	Teammate2SteelheadKills int     `json:"teammate2_steelhead_kills"`
	Teammate2StingerKills   int     `json:"teammate2_stinger_kills"`
	Teammate2W1Specials     int     `json:"teammate2_w1_specials"`
	Teammate2W2Specials     int     `json:"teammate2_w2_specials"`
	Teammate2W3Specials     int     `json:"teammate2_w3_specials"`
	Teammate2WeaponW1       string  `json:"teammate2_weapon_w1"`
	Teammate2WeaponW2       string  `json:"teammate2_weapon_w2"`
	Teammate2WeaponW3       string  `json:"teammate2_weapon_w3"`
	Wave1EventType          string  `json:"wave_1_event_type"`
	Wave1GoldenAppear       int     `json:"wave_1_golden_appear"`
	Wave1GoldenDelivered    int     `json:"wave_1_golden_delivered"`
	Wave1PowerEggs          int     `json:"wave_1_power_eggs"`
	Wave1Quota              int     `json:"wave_1_quota"`
	Wave1WaterLevel         string  `json:"wave_1_water_level"`
	Wave2EventType          string  `json:"wave_2_event_type"`
	Wave2GoldenAppear       int     `json:"wave_2_golden_appear"`
	Wave2GoldenDelivered    int     `json:"wave_2_golden_delivered"`
	Wave2PowerEggs          int     `json:"wave_2_power_eggs"`
	Wave2Quota              int     `json:"wave_2_quota"`
	Wave2WaterLevel         string  `json:"wave_2_water_level"`
	Wave3EventType          string  `json:"wave_3_event_type"`
	Wave3GoldenAppear       int     `json:"wave_3_golden_appear"`
	Wave3GoldenDelivered    int     `json:"wave_3_golden_delivered"`
	Wave3PowerEggs          int     `json:"wave_3_power_eggs"`
	Wave3Quota              int     `json:"wave_3_quota"`
	Wave3WaterLevel         string  `json:"wave_3_water_level"`
}

type BattleUpload struct {
	BattleNumber            string              `json:"battle_number"`
	ElapsedTime             int                 `json:"elapsed_time"`
	HasDisconnectedPlayer   bool                `json:"has_disconnected_player"`
	LeaguePoint             decimal.NullDecimal `json:"league_point"`
	MatchType               string              `json:"match_type"`
	MyTeamCount             int                 `json:"my_team_count"`
	Opponent0Assists        int                 `json:"opponent0_assists"`
	Opponent0Clothes        string              `json:"opponent0_clothes"`
	Opponent0ClothesMain    string              `json:"opponent0_clothes_main"`
	Opponent0ClothesSub0    string              `json:"opponent0_clothes_sub0"`
	Opponent0ClothesSub1    string              `json:"opponent0_clothes_sub1"`
	Opponent0ClothesSub2    string              `json:"opponent0_clothes_sub2"`
	Opponent0Deaths         int                 `json:"opponent0_deaths"`
	Opponent0GamePaintPoint int                 `json:"opponent0_game_paint_point"`
	Opponent0Gender         string              `json:"opponent0_gender"`
	Opponent0Headgear       string              `json:"opponent0_headgear"`
	Opponent0HeadgearMain   string              `json:"opponent0_headgear_main"`
	Opponent0HeadgearSub0   string              `json:"opponent0_headgear_sub0"`
	Opponent0HeadgearSub1   string              `json:"opponent0_headgear_sub1"`
	Opponent0HeadgearSub2   string              `json:"opponent0_headgear_sub2"`
	Opponent0Kills          int                 `json:"opponent0_kills"`
	Opponent0Level          int                 `json:"opponent0_level"`
	Opponent0LevelStar      int                 `json:"opponent0_level_star"`
	Opponent0Name           string              `json:"opponent0_name"`
	Opponent0Rank           int                 `json:"opponent0_rank"`
	Opponent0Shoes          string              `json:"opponent0_shoes"`
	Opponent0ShoesMain      string              `json:"opponent0_shoes_main"`
	Opponent0ShoesSub0      string              `json:"opponent0_shoes_sub0"`
	Opponent0ShoesSub1      string              `json:"opponent0_shoes_sub1"`
	Opponent0ShoesSub2      string              `json:"opponent0_shoes_sub2"`
	Opponent0Specials       int                 `json:"opponent0_specials"`
	Opponent0Species        string              `json:"opponent0_species"`
	Opponent0SplatnetID     string              `json:"opponent0_splatnet_id"`
	Opponent0Weapon         string              `json:"opponent0_weapon"`
	Opponent1Assists        int                 `json:"opponent1_assists"`
	Opponent1Clothes        string              `json:"opponent1_clothes"`
	Opponent1ClothesMain    string              `json:"opponent1_clothes_main"`
	Opponent1ClothesSub0    string              `json:"opponent1_clothes_sub0"`
	Opponent1ClothesSub1    string              `json:"opponent1_clothes_sub1"`
	Opponent1ClothesSub2    string              `json:"opponent1_clothes_sub2"`
	Opponent1Deaths         int                 `json:"opponent1_deaths"`
	Opponent1GamePaintPoint int                 `json:"opponent1_game_paint_point"`
	Opponent1Gender         string              `json:"opponent1_gender"`
	Opponent1Headgear       string              `json:"opponent1_headgear"`
	Opponent1HeadgearMain   string              `json:"opponent1_headgear_main"`
	Opponent1HeadgearSub0   string              `json:"opponent1_headgear_sub0"`
	Opponent1HeadgearSub1   string              `json:"opponent1_headgear_sub1"`
	Opponent1HeadgearSub2   string              `json:"opponent1_headgear_sub2"`
	Opponent1Kills          int                 `json:"opponent1_kills"`
	Opponent1Level          int                 `json:"opponent1_level"`
	Opponent1LevelStar      int                 `json:"opponent1_level_star"`
	Opponent1Name           string              `json:"opponent1_name"`
	Opponent1Rank           int                 `json:"opponent1_rank"`
	Opponent1Shoes          string              `json:"opponent1_shoes"`
	Opponent1ShoesMain      string              `json:"opponent1_shoes_main"`
	Opponent1ShoesSub0      string              `json:"opponent1_shoes_sub0"`
	Opponent1ShoesSub1      string              `json:"opponent1_shoes_sub1"`
	Opponent1ShoesSub2      string              `json:"opponent1_shoes_sub2"`
	Opponent1Specials       int                 `json:"opponent1_specials"`
	Opponent1Species        string              `json:"opponent1_species"`
	Opponent1SplatnetID     string              `json:"opponent1_splatnet_id"`
	Opponent1Weapon         string              `json:"opponent1_weapon"`
	Opponent2Assists        int                 `json:"opponent2_assists"`
	Opponent2Clothes        string              `json:"opponent2_clothes"`
	Opponent2ClothesMain    string              `json:"opponent2_clothes_main"`
	Opponent2ClothesSub0    string              `json:"opponent2_clothes_sub0"`
	Opponent2ClothesSub1    string              `json:"opponent2_clothes_sub1"`
	Opponent2ClothesSub2    string              `json:"opponent2_clothes_sub2"`
	Opponent2Deaths         int                 `json:"opponent2_deaths"`
	Opponent2GamePaintPoint int                 `json:"opponent2_game_paint_point"`
	Opponent2Gender         string              `json:"opponent2_gender"`
	Opponent2Headgear       string              `json:"opponent2_headgear"`
	Opponent2HeadgearMain   string              `json:"opponent2_headgear_main"`
	Opponent2HeadgearSub0   string              `json:"opponent2_headgear_sub0"`
	Opponent2HeadgearSub1   string              `json:"opponent2_headgear_sub1"`
	Opponent2HeadgearSub2   string              `json:"opponent2_headgear_sub2"`
	Opponent2Kills          int                 `json:"opponent2_kills"`
	Opponent2Level          int                 `json:"opponent2_level"`
	Opponent2LevelStar      int                 `json:"opponent2_level_star"`
	Opponent2Name           string              `json:"opponent2_name"`
	Opponent2Rank           int                 `json:"opponent2_rank"`
	Opponent2Shoes          string              `json:"opponent2_shoes"`
	Opponent2ShoesMain      string              `json:"opponent2_shoes_main"`
	Opponent2ShoesSub0      string              `json:"opponent2_shoes_sub0"`
	Opponent2ShoesSub1      string              `json:"opponent2_shoes_sub1"`
	Opponent2ShoesSub2      string              `json:"opponent2_shoes_sub2"`
	Opponent2Specials       int                 `json:"opponent2_specials"`
	Opponent2Species        string              `json:"opponent2_species"`
	Opponent2SplatnetID     string              `json:"opponent2_splatnet_id"`
	Opponent2Weapon         string              `json:"opponent2_weapon"`
	Opponent3Assists        int                 `json:"opponent3_assists"`
	Opponent3Clothes        string              `json:"opponent3_clothes"`
	Opponent3ClothesMain    string              `json:"opponent3_clothes_main"`
	Opponent3ClothesSub0    string              `json:"opponent3_clothes_sub0"`
	Opponent3ClothesSub1    string              `json:"opponent3_clothes_sub1"`
	Opponent3ClothesSub2    string              `json:"opponent3_clothes_sub2"`
	Opponent3Deaths         int                 `json:"opponent3_deaths"`
	Opponent3GamePaintPoint int                 `json:"opponent3_game_paint_point"`
	Opponent3Gender         string              `json:"opponent3_gender"`
	Opponent3Headgear       string              `json:"opponent3_headgear"`
	Opponent3HeadgearMain   string              `json:"opponent3_headgear_main"`
	Opponent3HeadgearSub0   string              `json:"opponent3_headgear_sub0"`
	Opponent3HeadgearSub1   string              `json:"opponent3_headgear_sub1"`
	Opponent3HeadgearSub2   string              `json:"opponent3_headgear_sub2"`
	Opponent3Kills          int                 `json:"opponent3_kills"`
	Opponent3Level          int                 `json:"opponent3_level"`
	Opponent3LevelStar      int                 `json:"opponent3_level_star"`
	Opponent3Name           string              `json:"opponent3_name"`
	Opponent3Rank           int                 `json:"opponent3_rank"`
	Opponent3Shoes          string              `json:"opponent3_shoes"`
	Opponent3ShoesMain      string              `json:"opponent3_shoes_main"`
	Opponent3ShoesSub0      string              `json:"opponent3_shoes_sub0"`
	Opponent3ShoesSub1      string              `json:"opponent3_shoes_sub1"`
	Opponent3ShoesSub2      string              `json:"opponent3_shoes_sub2"`
	Opponent3Specials       int                 `json:"opponent3_specials"`
	Opponent3Species        string              `json:"opponent3_species"`
	Opponent3SplatnetID     string              `json:"opponent3_splatnet_id"`
	Opponent3Weapon         string              `json:"opponent3_weapon"`
	OtherTeamCount          int                 `json:"other_team_count"`
	PlayerAssists           int                 `json:"player_assists"`
	PlayerClothes           string              `json:"player_clothes"`
	PlayerClothesMain       string              `json:"player_clothes_main"`
	PlayerClothesSub0       string              `json:"player_clothes_sub0"`
	PlayerClothesSub1       string              `json:"player_clothes_sub1"`
	PlayerClothesSub2       string              `json:"player_clothes_sub2"`
	PlayerDeaths            int                 `json:"player_deaths"`
	PlayerGamePaintPoint    int                 `json:"player_game_paint_point"`
	PlayerGender            string              `json:"player_gender"`
	PlayerHeadgear          string              `json:"player_headgear"`
	PlayerHeadgearMain      string              `json:"player_headgear_main"`
	PlayerHeadgearSub0      string              `json:"player_headgear_sub0"`
	PlayerHeadgearSub1      string              `json:"player_headgear_sub1"`
	PlayerHeadgearSub2      string              `json:"player_headgear_sub2"`
	PlayerKills             int                 `json:"player_kills"`
	PlayerLevel             int                 `json:"player_level"`
	PlayerLevelStar         int                 `json:"player_level_star"`
	PlayerName              string              `json:"player_name"`
	PlayerRank              int                 `json:"player_rank"`
	PlayerShoes             string              `json:"player_shoes"`
	PlayerShoesMain         string              `json:"player_shoes_main"`
	PlayerShoesSub0         string              `json:"player_shoes_sub0"`
	PlayerShoesSub1         string              `json:"player_shoes_sub1"`
	PlayerShoesSub2         string              `json:"player_shoes_sub2"`
	PlayerSpecials          int                 `json:"player_specials"`
	PlayerSpecies           string              `json:"player_species"`
	PlayerSplatfestTitle    string              `json:"player_splatfest_title"`
	PlayerSplatnetID        string              `json:"player_splatnet_id"`
	PlayerWeapon            string              `json:"player_weapon"`
	PlayerXPower            decimal.NullDecimal `json:"player_x_power"`
	Rule                    string              `json:"rule"`
	SplatfestPoint          decimal.NullDecimal `json:"splatfest_point"`
	SplatfestTitleAfter     string              `json:"splatfest_title_after"`
	SplatnetJSON            Battle              `json:"splatnet_json"`
	SplatnetUpload          bool                `json:"splatnet_upload"`
	Stage                   string              `json:"stage"`
	StatInkUpload           bool                `json:"stat_ink_upload"`
	TagID                   string              `json:"tag_id"`
	Teammate0Assists        int                 `json:"teammate0_assists"`
	Teammate0Clothes        string              `json:"teammate0_clothes"`
	Teammate0ClothesMain    string              `json:"teammate0_clothes_main"`
	Teammate0ClothesSub0    string              `json:"teammate0_clothes_sub0"`
	Teammate0ClothesSub1    string              `json:"teammate0_clothes_sub1"`
	Teammate0ClothesSub2    string              `json:"teammate0_clothes_sub2"`
	Teammate0Deaths         int                 `json:"teammate0_deaths"`
	Teammate0GamePaintPoint int                 `json:"teammate0_game_paint_point"`
	Teammate0Gender         string              `json:"teammate0_gender"`
	Teammate0Headgear       string              `json:"teammate0_headgear"`
	Teammate0HeadgearMain   string              `json:"teammate0_headgear_main"`
	Teammate0HeadgearSub0   string              `json:"teammate0_headgear_sub0"`
	Teammate0HeadgearSub1   string              `json:"teammate0_headgear_sub1"`
	Teammate0HeadgearSub2   string              `json:"teammate0_headgear_sub2"`
	Teammate0Kills          int                 `json:"teammate0_kills"`
	Teammate0Level          int                 `json:"teammate0_level"`
	Teammate0LevelStar      int                 `json:"teammate0_level_star"`
	Teammate0Name           string              `json:"teammate0_name"`
	Teammate0Rank           int                 `json:"teammate0_rank"`
	Teammate0Shoes          string              `json:"teammate0_shoes"`
	Teammate0ShoesMain      string              `json:"teammate0_shoes_main"`
	Teammate0ShoesSub0      string              `json:"teammate0_shoes_sub0"`
	Teammate0ShoesSub1      string              `json:"teammate0_shoes_sub1"`
	Teammate0ShoesSub2      string              `json:"teammate0_shoes_sub2"`
	Teammate0Specials       int                 `json:"teammate0_specials"`
	Teammate0Species        string              `json:"teammate0_species"`
	Teammate0SplatnetID     string              `json:"teammate0_splatnet_id"`
	Teammate0Weapon         string              `json:"teammate0_weapon"`
	Teammate1Assists        int                 `json:"teammate1_assists"`
	Teammate1Clothes        string              `json:"teammate1_clothes"`
	Teammate1ClothesMain    string              `json:"teammate1_clothes_main"`
	Teammate1ClothesSub0    string              `json:"teammate1_clothes_sub0"`
	Teammate1ClothesSub1    string              `json:"teammate1_clothes_sub1"`
	Teammate1ClothesSub2    string              `json:"teammate1_clothes_sub2"`
	Teammate1Deaths         int                 `json:"teammate1_deaths"`
	Teammate1GamePaintPoint int                 `json:"teammate1_game_paint_point"`
	Teammate1Gender         string              `json:"teammate1_gender"`
	Teammate1Headgear       string              `json:"teammate1_headgear"`
	Teammate1HeadgearMain   string              `json:"teammate1_headgear_main"`
	Teammate1HeadgearSub0   string              `json:"teammate1_headgear_sub0"`
	Teammate1HeadgearSub1   string              `json:"teammate1_headgear_sub1"`
	Teammate1HeadgearSub2   string              `json:"teammate1_headgear_sub2"`
	Teammate1Kills          int                 `json:"teammate1_kills"`
	Teammate1Level          int                 `json:"teammate1_level"`
	Teammate1LevelStar      int                 `json:"teammate1_level_star"`
	Teammate1Name           string              `json:"teammate1_name"`
	Teammate1Rank           int                 `json:"teammate1_rank"`
	Teammate1Shoes          string              `json:"teammate1_shoes"`
	Teammate1ShoesMain      string              `json:"teammate1_shoes_main"`
	Teammate1ShoesSub0      string              `json:"teammate1_shoes_sub0"`
	Teammate1ShoesSub1      string              `json:"teammate1_shoes_sub1"`
	Teammate1ShoesSub2      string              `json:"teammate1_shoes_sub2"`
	Teammate1Specials       int                 `json:"teammate1_specials"`
	Teammate1Species        string              `json:"teammate1_species"`
	Teammate1SplatnetID     string              `json:"teammate1_splatnet_id"`
	Teammate1Weapon         string              `json:"teammate1_weapon"`
	Teammate2Assists        int                 `json:"teammate2_assists"`
	Teammate2Clothes        string              `json:"teammate2_clothes"`
	Teammate2ClothesMain    string              `json:"teammate2_clothes_main"`
	Teammate2ClothesSub0    string              `json:"teammate2_clothes_sub0"`
	Teammate2ClothesSub1    string              `json:"teammate2_clothes_sub1"`
	Teammate2ClothesSub2    string              `json:"teammate2_clothes_sub2"`
	Teammate2Deaths         int                 `json:"teammate2_deaths"`
	Teammate2GamePaintPoint int                 `json:"teammate2_game_paint_point"`
	Teammate2Gender         string              `json:"teammate2_gender"`
	Teammate2Headgear       string              `json:"teammate2_headgear"`
	Teammate2HeadgearMain   string              `json:"teammate2_headgear_main"`
	Teammate2HeadgearSub0   string              `json:"teammate2_headgear_sub0"`
	Teammate2HeadgearSub1   string              `json:"teammate2_headgear_sub1"`
	Teammate2HeadgearSub2   string              `json:"teammate2_headgear_sub2"`
	Teammate2Kills          int                 `json:"teammate2_kills"`
	Teammate2Level          int                 `json:"teammate2_level"`
	Teammate2LevelStar      int                 `json:"teammate2_level_star"`
	Teammate2Name           string              `json:"teammate2_name"`
	Teammate2Rank           int                 `json:"teammate2_rank"`
	Teammate2Shoes          string              `json:"teammate2_shoes"`
	Teammate2ShoesMain      string              `json:"teammate2_shoes_main"`
	Teammate2ShoesSub0      string              `json:"teammate2_shoes_sub0"`
	Teammate2ShoesSub1      string              `json:"teammate2_shoes_sub1"`
	Teammate2ShoesSub2      string              `json:"teammate2_shoes_sub2"`
	Teammate2Specials       int                 `json:"teammate2_specials"`
	Teammate2Species        string              `json:"teammate2_species"`
	Teammate2SplatnetID     string              `json:"teammate2_splatnet_id"`
	Teammate2Weapon         string              `json:"teammate2_weapon"`
	Time                    int                 `json:"time"`
	Win                     bool                `json:"win"`
	WinMeter                decimal.NullDecimal `json:"win_meter"`
}
