package types

import (
	"github.com/shopspring/decimal"
)

type Shift struct {
	DangerRate float64 `json:"danger_rate"`
	JobResult  struct {
		IsClear       *bool   `json:"is_clear"`
		FailureReason *string `json:"failure_reason"`
		FailureWave   *int    `json:"failure_wave"`
	} `json:"job_result"`
	JobScore        *int `json:"job_score"`
	JobID           *int `json:"job_id"`
	JobRate         *int `json:"job_rate"`
	GradePoint      *int `json:"grade_point"`
	GradePointDelta *int `json:"grade_point_delta"`
	OtherResults    []struct {
		SpecialCounts []*int `json:"special_counts"`
		Special       struct {
			ImageB *string `json:"image_b"`
			Name   *string `json:"name"`
			ID     *string `json:"id"`
			ImageA *string `json:"image_a"`
		} `json:"special"`
		Pid        *string `json:"pid"`
		PlayerType struct {
			Gender  *string `json:"style"`
			Species *string `json:"species"`
		} `json:"player_type"`
		WeaponList []struct {
			ID     *string `json:"id"`
			Weapon struct {
				ID        *string `json:"id"`
				Image     *string `json:"image"`
				Name      *string `json:"name"`
				Thumbnail *string `json:"thumbnail"`
			} `json:"weapon"`
		} `json:"weapon_list"`
		Name           *string `json:"name"`
		DeadCount      *int    `json:"dead_count"`
		GoldenEggs     *int    `json:"golden_ikura_num"`
		BossKillCounts struct {
			Goldie struct {
				Boss struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"3"`
			Steelhead struct {
				Boss struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"6"`
			Flyfish struct {
				Boss struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"9"`
			Scrapper struct {
				Count *int `json:"count"`
				Boss  struct {
					Key  *string `json:"key"`
					Name *string `json:"name"`
				} `json:"boss"`
			} `json:"12"`
			SteelEel struct {
				Count *int `json:"count"`
				Boss  struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
			} `json:"13"`
			Stinger struct {
				Count *int `json:"count"`
				Boss  struct {
					Key  *string `json:"key"`
					Name *string `json:"name"`
				} `json:"boss"`
			} `json:"14"`
			Maws struct {
				Boss struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"15"`
			Griller struct {
				Boss struct {
					Key  *string `json:"key"`
					Name *string `json:"name"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"16"`
			Drizzler struct {
				Count *int `json:"count"`
				Boss  struct {
					Key  *string `json:"key"`
					Name *string `json:"name"`
				} `json:"boss"`
			} `json:"21"`
		} `json:"boss_kill_counts"`
		PowerEggs *int `json:"ikura_num"`
		HelpCount *int `json:"help_count"`
	} `json:"other_results"`
	KumaPoint  *int `json:"kuma_point"`
	StartTime  *int `json:"start_time"`
	PlayerType struct {
		Species *string `json:"species"`
		Gender  *string `json:"style"`
	} `json:"player_type"`
	PlayTime   *int `json:"play_time"`
	BossCounts struct {
		Goldie struct {
			Count *int `json:"count"`
			Boss  struct {
				Name *string `json:"name"`
				Key  *string `json:"key"`
			} `json:"boss"`
		} `json:"3"`
		Steelhead struct {
			Boss struct {
				Name *string `json:"name"`
				Key  *string `json:"key"`
			} `json:"boss"`
			Count *int `json:"count"`
		} `json:"6"`
		Flyfish struct {
			Boss struct {
				Name *string `json:"name"`
				Key  *string `json:"key"`
			} `json:"boss"`
			Count *int `json:"count"`
		} `json:"9"`
		Scrapper struct {
			Boss struct {
				Name *string `json:"name"`
				Key  *string `json:"key"`
			} `json:"boss"`
			Count *int `json:"count"`
		} `json:"12"`
		SteelEel struct {
			Count *int `json:"count"`
			Boss  struct {
				Name *string `json:"name"`
				Key  *string `json:"key"`
			} `json:"boss"`
		} `json:"13"`
		Stinger struct {
			Count *int `json:"count"`
			Boss  struct {
				Key  *string `json:"key"`
				Name *string `json:"name"`
			} `json:"boss"`
		} `json:"14"`
		Maws struct {
			Boss struct {
				Name *string `json:"name"`
				Key  *string `json:"key"`
			} `json:"boss"`
			Count *int `json:"count"`
		} `json:"15"`
		Griller struct {
			Boss struct {
				Key  *string `json:"key"`
				Name *string `json:"name"`
			} `json:"boss"`
			Count *int `json:"count"`
		} `json:"16"`
		Drizzler struct {
			Boss struct {
				Name *string `json:"name"`
				Key  *string `json:"key"`
			} `json:"boss"`
			Count *int `json:"count"`
		} `json:"21"`
	} `json:"boss_counts"`
	EndTime  *int `json:"end_time"`
	MyResult struct {
		GoldenEggs     *int `json:"golden_ikura_num"`
		DeadCount      *int `json:"dead_count"`
		HelpCount      *int `json:"help_count"`
		BossKillCounts struct {
			Goldie struct {
				Boss struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"3"`
			Steelhead struct {
				Count *int `json:"count"`
				Boss  struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
			} `json:"6"`
			Flyfish struct {
				Boss struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"9"`
			Scrapper struct {
				Boss struct {
					Key  *string `json:"key"`
					Name *string `json:"name"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"12"`
			SteelEel struct {
				Boss struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"13"`
			Stinger struct {
				Count *int `json:"count"`
				Boss  struct {
					Key  *string `json:"key"`
					Name *string `json:"name"`
				} `json:"boss"`
			} `json:"14"`
			Maws struct {
				Count *int `json:"count"`
				Boss  struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
			} `json:"15"`
			Griller struct {
				Count *int `json:"count"`
				Boss  struct {
					Name *string `json:"name"`
					Key  *string `json:"key"`
				} `json:"boss"`
			} `json:"16"`
			Drizzler struct {
				Boss struct {
					Key  *string `json:"key"`
					Name *string `json:"name"`
				} `json:"boss"`
				Count *int `json:"count"`
			} `json:"21"`
		} `json:"boss_kill_counts"`
		PowerEggs *int `json:"ikura_num"`
		Special   struct {
			Name   *string `json:"name"`
			ImageB *string `json:"image_b"`
			ImageA *string `json:"image_a"`
			ID     *string `json:"id"`
		} `json:"special"`
		SpecialCounts []*int  `json:"special_counts"`
		Name          *string `json:"name"`
		WeaponList    []struct {
			ID     *string `json:"id"`
			Weapon struct {
				ID        *string `json:"id"`
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
				Image     *string `json:"image"`
			} `json:"weapon"`
		} `json:"weapon_list"`
		PlayerType struct {
			Species *string `json:"species"`
			Gender  *string `json:"style"`
		} `json:"player_type"`
		Pid *string `json:"pid"`
	} `json:"my_result"`
	WaveDetails []struct {
		WaterLevel struct {
			Key  *string `json:"key"`
			Name *string `json:"name"`
		} `json:"water_level"`
		EventType struct {
			Key  *string `json:"key"`
			Name *string `json:"name"`
		} `json:"event_type"`
		GoldenEggs   *int `json:"golden_ikura_num"`
		GoldenAppear *int `json:"golden_ikura_pop_num"`
		PowerEggs    *int `json:"ikura_num"`
		QuotaNum     *int `json:"quota_num"`
	} `json:"wave_details"`
	Grade struct {
		ID        *string `json:"id"`
		ShortName *string `json:"short_name"`
		LongName  *string `json:"long_name"`
		Name      *string `json:"name"`
	} `json:"grade"`
	Schedule struct {
		StartTime *int `json:"start_time"`
		Weapons   []struct {
			Weapon struct {
				ID        *string `json:"id"`
				Image     *string `json:"image"`
				Name      *string `json:"name"`
				Thumbnail *string `json:"thumbnail"`
			} `json:"weapon"`
			ID *string `json:"id"`
		} `json:"weapons"`
		EndTime *int `json:"end_time"`
		Stage   struct {
			Image *string `json:"image"`
			Name  *string `json:"name"`
		} `json:"stage"`
	} `json:"schedule"`
}

type Battle struct {
	Udemae struct {
		Name            *string      `json:"name"`
		IsX             *bool        `json:"is_x"`
		IsNumberReached *bool        `json:"is_number_reached"`
		Number          *int         `json:"number"`
		SPlusNumber     *interface{} `json:"s_plus_number"`
	} `json:"udemae"`
	Stage struct {
		ID    *string `json:"id"`
		Image *string `json:"image"`
		Name  *string `json:"name"`
	} `json:"stage"`
	OtherTeamCount *int `json:"other_team_count"`
	MyTeamCount    *int `json:"my_team_count"`
	StarRank       *int `json:"star_rank"`
	Rule           struct {
		Name          *string `json:"name"`
		Key           *string `json:"key"`
		MultilineName *string `json:"multiline_name"`
	} `json:"rule"`
	PlayerResult struct {
		DeathCount     *int `json:"death_count"`
		GamePaintPoint *int `json:"game_paint_point"`
		KillCount      *int `json:"kill_count"`
		SpecialCount   *int `json:"special_count"`
		AssistCount    *int `json:"assist_count"`
		SortScore      *int `json:"sort_score"`
		Player         struct {
			HeadSkills struct {
				Main struct {
					Name  *string `json:"name"`
					ID    *string `json:"id"`
					Image *string `json:"image"`
				} `json:"main"`
				Subs []struct {
					Name  *string `json:"name"`
					ID    *string `json:"id"`
					Image *string `json:"image"`
				} `json:"subs"`
			} `json:"head_skills"`
			ShoesSkills struct {
				Subs []struct {
					Name  *string `json:"name"`
					ID    *string `json:"id"`
					Image *string `json:"image"`
				} `json:"subs"`
				Main struct {
					Name  *string `json:"name"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
				} `json:"main"`
			} `json:"shoes_skills"`
			PlayerRank *int `json:"player_rank"`
			PlayerType struct {
				Gender  *string `json:"style"`
				Species *string `json:"species"`
			} `json:"player_type"`
			PrincipalID *string `json:"principal_id"`
			Head        struct {
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
				Image     *string `json:"image"`
				Kind      *string `json:"kind"`
				Rarity    *int    `json:"rarity"`
				Brand     struct {
					FrequentSkill struct {
						Image *string `json:"image"`
						ID    *string `json:"id"`
						Name  *string `json:"name"`
					} `json:"frequent_skill"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
					Name  *string `json:"name"`
				} `json:"brand"`
				ID *string `json:"id"`
			} `json:"head"`
			Nickname *string `json:"nickname"`
			Clothes  struct {
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
				Image     *string `json:"image"`
				Kind      *string `json:"kind"`
				Rarity    *int    `json:"rarity"`
				ID        *string `json:"id"`
				Brand     struct {
					Name          *string `json:"name"`
					FrequentSkill struct {
						Name  *string `json:"name"`
						ID    *string `json:"id"`
						Image *string `json:"image"`
					} `json:"frequent_skill"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
				} `json:"brand"`
			} `json:"clothes"`
			Udemae struct {
				SPlusNumber     *interface{} `json:"s_plus_number"`
				Number          *int         `json:"number"`
				IsNumberReached *bool        `json:"is_number_reached"`
				IsX             *bool        `json:"is_x"`
				Name            *string      `json:"name"`
			} `json:"udemae"`
			Shoes struct {
				Image     *string `json:"image"`
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
				ID        *string `json:"id"`
				Brand     struct {
					Image         *string `json:"image"`
					FrequentSkill struct {
						ID    *string `json:"id"`
						Image *string `json:"image"`
						Name  *string `json:"name"`
					} `json:"frequent_skill"`
					ID   *string `json:"id"`
					Name *string `json:"name"`
				} `json:"brand"`
				Rarity *int    `json:"rarity"`
				Kind   *string `json:"kind"`
			} `json:"shoes"`
			StarRank *int `json:"star_rank"`
			Weapon   struct {
				Sub struct {
					Name   *string `json:"name"`
					ImageA *string `json:"image_a"`
					ID     *string `json:"id"`
					ImageB *string `json:"image_b"`
				} `json:"sub"`
				Special struct {
					ImageA *string `json:"image_a"`
					Name   *string `json:"name"`
					ID     *string `json:"id"`
					ImageB *string `json:"image_b"`
				} `json:"special"`
				Image     *string `json:"image"`
				ID        *string `json:"id"`
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
			} `json:"weapon"`
			ClothesSkills struct {
				Subs []struct {
					Name  *string `json:"name"`
					ID    *string `json:"id"`
					Image *string `json:"image"`
				} `json:"subs"`
				Main struct {
					Name  *string `json:"name"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
				} `json:"main"`
			} `json:"clothes_skills"`
		} `json:"player"`
	} `json:"player_result"`
	EstimateGachiPower *int `json:"estimate_gachi_power"`
	ElapsedTime        *int `json:"elapsed_time"`
	StartTime          *int `json:"start_time"`
	GameMode           struct {
		Key  *string `json:"key"`
		Name *string `json:"name"`
	} `json:"game_mode"`
	XPower        decimal.NullDecimal `json:"x_power"`
	BattleNumber  *string             `json:"battle_number"`
	Type          *string             `json:"type"`
	PlayerRank    *int                `json:"player_rank"`
	CrownPlayers  *interface{}        `json:"crown_players"`
	MyTeamMembers []struct {
		GamePaintPoint *int `json:"game_paint_point"`
		DeathCount     *int `json:"death_count"`
		SortScore      *int `json:"sort_score"`
		AssistCount    *int `json:"assist_count"`
		KillCount      *int `json:"kill_count"`
		SpecialCount   *int `json:"special_count"`
		Player         struct {
			Shoes struct {
				Brand struct {
					FrequentSkill struct {
						ID    *string `json:"id"`
						Image *string `json:"image"`
						Name  *string `json:"name"`
					} `json:"frequent_skill"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
					Name  *string `json:"name"`
				} `json:"brand"`
				ID        *string `json:"id"`
				Kind      *string `json:"kind"`
				Rarity    *int    `json:"rarity"`
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
				Image     *string `json:"image"`
			} `json:"shoes"`
			Weapon struct {
				Name      *string `json:"name"`
				Thumbnail *string `json:"thumbnail"`
				Image     *string `json:"image"`
				ID        *string `json:"id"`
				Special   struct {
					ImageB *string `json:"image_b"`
					ID     *string `json:"id"`
					ImageA *string `json:"image_a"`
					Name   *string `json:"name"`
				} `json:"special"`
				Sub struct {
					ID     *string `json:"id"`
					Name   *string `json:"name"`
					ImageA *string `json:"image_a"`
					ImageB *string `json:"image_b"`
				} `json:"sub"`
			} `json:"weapon"`
			StarRank      *int `json:"star_rank"`
			ClothesSkills struct {
				Subs []struct {
					Name  *string `json:"name"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
				} `json:"subs"`
				Main struct {
					Name  *string `json:"name"`
					ID    *string `json:"id"`
					Image *string `json:"image"`
				} `json:"main"`
			} `json:"clothes_skills"`
			PrincipalID *string `json:"principal_id"`
			Head        struct {
				ID    *string `json:"id"`
				Brand struct {
					Name          *string `json:"name"`
					FrequentSkill struct {
						Name  *string `json:"name"`
						ID    *string `json:"id"`
						Image *string `json:"image"`
					} `json:"frequent_skill"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
				} `json:"brand"`
				Kind      *string `json:"kind"`
				Rarity    *int    `json:"rarity"`
				Name      *string `json:"name"`
				Thumbnail *string `json:"thumbnail"`
				Image     *string `json:"image"`
			} `json:"head"`
			Clothes struct {
				Image     *string `json:"image"`
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
				Rarity    *int    `json:"rarity"`
				Kind      *string `json:"kind"`
				Brand     struct {
					Name          *string `json:"name"`
					ID            *string `json:"id"`
					FrequentSkill struct {
						Name  *string `json:"name"`
						ID    *string `json:"id"`
						Image *string `json:"image"`
					} `json:"frequent_skill"`
					Image *string `json:"image"`
				} `json:"brand"`
				ID *string `json:"id"`
			} `json:"clothes"`
			Nickname *string `json:"nickname"`
			Udemae   struct {
				SPlusNumber *interface{} `json:"s_plus_number"`
				IsX         *bool        `json:"is_x"`
				Name        *string      `json:"name"`
			} `json:"udemae"`
			PlayerType struct {
				Species *string `json:"species"`
				Gender  *string `json:"style"`
			} `json:"player_type"`
			HeadSkills struct {
				Main struct {
					Name  *string `json:"name"`
					ID    *string `json:"id"`
					Image *string `json:"image"`
				} `json:"main"`
				Subs []struct {
					ID    *string `json:"id"`
					Image *string `json:"image"`
					Name  *string `json:"name"`
				} `json:"subs"`
			} `json:"head_skills"`
			PlayerRank  *int `json:"player_rank"`
			ShoesSkills struct {
				Subs []struct {
					ID    *string `json:"id"`
					Image *string `json:"image"`
					Name  *string `json:"name"`
				} `json:"subs"`
				Main struct {
					Name  *string `json:"name"`
					ID    *string `json:"id"`
					Image *string `json:"image"`
				} `json:"main"`
			} `json:"shoes_skills"`
		} `json:"player"`
	} `json:"my_team_members"`
	OtherTeamMembers []struct {
		GamePaintPoint *int `json:"game_paint_point"`
		DeathCount     *int `json:"death_count"`
		SpecialCount   *int `json:"special_count"`
		KillCount      *int `json:"kill_count"`
		SortScore      *int `json:"sort_score"`
		AssistCount    *int `json:"assist_count"`
		Player         struct {
			PlayerRank  *int `json:"player_rank"`
			ShoesSkills struct {
				Subs []struct {
					ID    *string `json:"id"`
					Image *string `json:"image"`
					Name  *string `json:"name"`
				} `json:"subs"`
				Main struct {
					Name  *string `json:"name"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
				} `json:"main"`
			} `json:"shoes_skills"`
			HeadSkills struct {
				Subs []struct {
					ID    *string `json:"id"`
					Image *string `json:"image"`
					Name  *string `json:"name"`
				} `json:"subs"`
				Main struct {
					Image *string `json:"image"`
					ID    *string `json:"id"`
					Name  *string `json:"name"`
				} `json:"main"`
			} `json:"head_skills"`
			PlayerType struct {
				Gender  *string `json:"style"`
				Species *string `json:"species"`
			} `json:"player_type"`
			Udemae struct {
				SPlusNumber *interface{} `json:"s_plus_number"`
				IsX         *bool        `json:"is_x"`
				Name        *string      `json:"name"`
			} `json:"udemae"`
			Clothes struct {
				Name      *string `json:"name"`
				Thumbnail *string `json:"thumbnail"`
				Image     *string `json:"image"`
				Kind      *string `json:"kind"`
				Rarity    *int    `json:"rarity"`
				Brand     struct {
					Name          *string `json:"name"`
					FrequentSkill struct {
						Name  *string `json:"name"`
						ID    *string `json:"id"`
						Image *string `json:"image"`
					} `json:"frequent_skill"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
				} `json:"brand"`
				ID *string `json:"id"`
			} `json:"clothes"`
			Nickname *string `json:"nickname"`
			Head     struct {
				Kind   *string `json:"kind"`
				Rarity *int    `json:"rarity"`
				Brand  struct {
					Name  *string `json:"name"`
					Image *string `json:"image"`
					ID    *string `json:"id"`
				} `json:"brand"`
				ID        *string `json:"id"`
				Name      *string `json:"name"`
				Thumbnail *string `json:"thumbnail"`
				Image     *string `json:"image"`
			} `json:"head"`
			PrincipalID   *string `json:"principal_id"`
			ClothesSkills struct {
				Main struct {
					Image *string `json:"image"`
					ID    *string `json:"id"`
					Name  *string `json:"name"`
				} `json:"main"`
				Subs []struct {
					ID    *string `json:"id"`
					Image *string `json:"image"`
					Name  *string `json:"name"`
				} `json:"subs"`
			} `json:"clothes_skills"`
			StarRank *int `json:"star_rank"`
			Weapon   struct {
				ID        *string `json:"id"`
				Image     *string `json:"image"`
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
				Sub       struct {
					ImageB *string `json:"image_b"`
					Name   *string `json:"name"`
					ImageA *string `json:"image_a"`
					ID     *string `json:"id"`
				} `json:"sub"`
				Special struct {
					ImageA *string `json:"image_a"`
					Name   *string `json:"name"`
					ID     *string `json:"id"`
					ImageB *string `json:"image_b"`
				} `json:"special"`
			} `json:"weapon"`
			Shoes struct {
				Kind   *string `json:"kind"`
				Rarity *int    `json:"rarity"`
				Brand  struct {
					Name          *string `json:"name"`
					Image         *string `json:"image"`
					FrequentSkill struct {
						Name  *string `json:"name"`
						Image *string `json:"image"`
						ID    *string `json:"id"`
					} `json:"frequent_skill"`
					ID *string `json:"id"`
				} `json:"brand"`
				ID        *string `json:"id"`
				Thumbnail *string `json:"thumbnail"`
				Name      *string `json:"name"`
				Image     *string `json:"image"`
			} `json:"shoes"`
		} `json:"player"`
	} `json:"other_team_members"`
	WeaponPaintPoint *int         `json:"weapon_paint_point"`
	Rank             *interface{} `json:"rank"`
	MyTeamResult     struct {
		Name *string `json:"name"`
		Key  *string `json:"key"`
	} `json:"my_team_result"`
	EstimateXPower  *interface{} `json:"estimate_x_power"`
	OtherTeamResult struct {
		Key  *string `json:"key"`
		Name *string `json:"name"`
	} `json:"other_team_result"`
	LeaguePoint         decimal.NullDecimal `json:"league_point"`
	WinMeter            decimal.NullDecimal `json:"win_meter"`
	MyTeamPercentage    decimal.NullDecimal `json:"my_team_percentage"`
	OtherTeamPercentage decimal.NullDecimal `json:"other_team_percentage"`
	TagID               *string             `json:"tag_id"`
}

type ShiftUpload struct {
	DangerRate              float64 `json:"danger_rate"`
	DrizzlerCount           *int    `json:"drizzler_count,omitempty"`
	Endtime                 *string `json:"endtime,omitempty"`
	FailureWave             *int    `json:"failure_wave,omitempty"`
	FlyfishCount            *int    `json:"flyfish_count,omitempty"`
	GoldieCount             *int    `json:"goldie_count,omitempty"`
	GradePoint              *int    `json:"grade_point,omitempty"`
	GradePointDelta         *int    `json:"grade_point_delta,omitempty"`
	GrillerCount            *int    `json:"griller_count,omitempty"`
	IsClear                 *bool   `json:"is_clear,omitempty"`
	JobFailureReason        *string `json:"job_failure_reason,omitempty"`
	JobID                   *int    `json:"job_id,omitempty"`
	MawsCount               *int    `json:"maws_count,omitempty"`
	PlayerDeathCount        *int    `json:"player_death_count,omitempty"`
	PlayerDrizzlerKills     *int    `json:"player_drizzler_kills,omitempty"`
	PlayerFlyfishKills      *int    `json:"player_flyfish_kills,omitempty"`
	PlayerGender            *string `json:"player_gender,omitempty"`
	PlayerGoldenEggs        *int    `json:"player_golden_eggs,omitempty"`
	PlayerGoldieKills       *int    `json:"player_goldie_kills,omitempty"`
	PlayerGrillerKills      *int    `json:"player_griller_kills,omitempty"`
	PlayerID                *string `json:"player_id,omitempty"`
	PlayerMawsKills         *int    `json:"player_maws_kills,omitempty"`
	PlayerName              *string `json:"player_name,omitempty"`
	PlayerPowerEggs         *int    `json:"player_power_eggs,omitempty"`
	PlayerReviveCount       *int    `json:"player_revive_count,omitempty"`
	PlayerScrapperKills     *int    `json:"player_scrapper_kills,omitempty"`
	PlayerSpecial           *string `json:"player_special,omitempty"`
	PlayerSpecies           *string `json:"player_species,omitempty"`
	PlayerSteelEelKills     *int    `json:"player_steel_eel_kills,omitempty"`
	PlayerSteelheadKills    *int    `json:"player_steelhead_kills,omitempty"`
	PlayerStingerKills      *int    `json:"player_stinger_kills,omitempty"`
	PlayerTitle             *string `json:"player_title,omitempty"`
	PlayerW1Specials        *int    `json:"player_w1_specials,omitempty"`
	PlayerW2Specials        *int    `json:"player_w2_specials,omitempty"`
	PlayerW3Specials        *int    `json:"player_w3_specials,omitempty"`
	PlayerWeaponW1          *string `json:"player_weapon_w1,omitempty"`
	PlayerWeaponW2          *string `json:"player_weapon_w2,omitempty"`
	PlayerWeaponW3          *string `json:"player_weapon_w3,omitempty"`
	Playtime                *string `json:"playtime,omitempty"`
	ScheduleEndtime         *string `json:"schedule_endtime,omitempty"`
	ScheduleStarttime       *string `json:"schedule_starttime,omitempty"`
	ScheduleWeapon0         *string `json:"schedule_weapon_0,omitempty"`
	ScheduleWeapon1         *string `json:"schedule_weapon_1,omitempty"`
	ScheduleWeapon2         *string `json:"schedule_weapon_2,omitempty"`
	ScheduleWeapon3         *string `json:"schedule_weapon_3,omitempty"`
	ScrapperCount           *int    `json:"scrapper_count,omitempty"`
	SplatnetJSON            Shift   `json:"splatnet_json"`
	SplatnetUpload          *bool   `json:"splatnet_upload,omitempty"`
	Stage                   *string `json:"stage,omitempty"`
	StatInkUpload           *bool   `json:"stat_ink_upload,omitempty"`
	Starttime               *string `json:"starttime,omitempty"`
	SteelEelCount           *int    `json:"steel_eel_count,omitempty"`
	SteelheadCount          *int    `json:"steelhead_count,omitempty"`
	StingerCount            *int    `json:"stinger_count,omitempty"`
	Teammate0DeathCount     *int    `json:"teammate0_death_count,omitempty"`
	Teammate0DrizzlerKills  *int    `json:"teammate0_drizzler_kills,omitempty"`
	Teammate0FlyfishKills   *int    `json:"teammate0_flyfish_kills,omitempty"`
	Teammate0Gender         *string `json:"teammate0_gender,omitempty"`
	Teammate0GoldenEggs     *int    `json:"teammate0_golden_eggs,omitempty"`
	Teammate0GoldieKills    *int    `json:"teammate0_goldie_kills,omitempty"`
	Teammate0GrillerKills   *int    `json:"teammate0_griller_kills,omitempty"`
	Teammate0ID             *string `json:"teammate0_id,omitempty"`
	Teammate0MawsKills      *int    `json:"teammate0_maws_kills,omitempty"`
	Teammate0Name           *string `json:"teammate0_name,omitempty"`
	Teammate0PowerEggs      *int    `json:"teammate0_power_eggs,omitempty"`
	Teammate0ReviveCount    *int    `json:"teammate0_revive_count,omitempty"`
	Teammate0ScrapperKills  *int    `json:"teammate0_scrapper_kills,omitempty"`
	Teammate0Special        *string `json:"teammate0_special,omitempty"`
	Teammate0Species        *string `json:"teammate0_species,omitempty"`
	Teammate0SteelEelKills  *int    `json:"teammate0_steel_eel_kills,omitempty"`
	Teammate0SteelheadKills *int    `json:"teammate0_steelhead_kills,omitempty"`
	Teammate0StingerKills   *int    `json:"teammate0_stinger_kills,omitempty"`
	Teammate0W1Specials     *int    `json:"teammate0_w1_specials,omitempty"`
	Teammate0W2Specials     *int    `json:"teammate0_w2_specials,omitempty"`
	Teammate0W3Specials     *int    `json:"teammate0_w3_specials,omitempty"`
	Teammate0WeaponW1       *string `json:"teammate0_weapon_w1,omitempty"`
	Teammate0WeaponW2       *string `json:"teammate0_weapon_w2,omitempty"`
	Teammate0WeaponW3       *string `json:"teammate0_weapon_w3,omitempty"`
	Teammate1DeathCount     *int    `json:"teammate1_death_count,omitempty"`
	Teammate1DrizzlerKills  *int    `json:"teammate1_drizzler_kills,omitempty"`
	Teammate1FlyfishKills   *int    `json:"teammate1_flyfish_kills,omitempty"`
	Teammate1Gender         *string `json:"teammate1_gender,omitempty"`
	Teammate1GoldenEggs     *int    `json:"teammate1_golden_eggs,omitempty"`
	Teammate1GoldieKills    *int    `json:"teammate1_goldie_kills,omitempty"`
	Teammate1GrillerKills   *int    `json:"teammate1_griller_kills,omitempty"`
	Teammate1ID             *string `json:"teammate1_id,omitempty"`
	Teammate1MawsKills      *int    `json:"teammate1_maws_kills,omitempty"`
	Teammate1Name           *string `json:"teammate1_name,omitempty"`
	Teammate1PowerEggs      *int    `json:"teammate1_power_eggs,omitempty"`
	Teammate1ReviveCount    *int    `json:"teammate1_revive_count,omitempty"`
	Teammate1ScrapperKills  *int    `json:"teammate1_scrapper_kills,omitempty"`
	Teammate1Special        *string `json:"teammate1_special,omitempty"`
	Teammate1Species        *string `json:"teammate1_species,omitempty"`
	Teammate1SteelEelKills  *int    `json:"teammate1_steel_eel_kills,omitempty"`
	Teammate1SteelheadKills *int    `json:"teammate1_steelhead_kills,omitempty"`
	Teammate1StingerKills   *int    `json:"teammate1_stinger_kills,omitempty"`
	Teammate1W1Specials     *int    `json:"teammate1_w1_specials,omitempty"`
	Teammate1W2Specials     *int    `json:"teammate1_w2_specials,omitempty"`
	Teammate1W3Specials     *int    `json:"teammate1_w3_specials,omitempty"`
	Teammate1WeaponW1       *string `json:"teammate1_weapon_w1,omitempty"`
	Teammate1WeaponW2       *string `json:"teammate1_weapon_w2,omitempty"`
	Teammate1WeaponW3       *string `json:"teammate1_weapon_w3,omitempty"`
	Teammate2DeathCount     *int    `json:"teammate2_death_count,omitempty"`
	Teammate2DrizzlerKills  *int    `json:"teammate2_drizzler_kills,omitempty"`
	Teammate2FlyfishKills   *int    `json:"teammate2_flyfish_kills,omitempty"`
	Teammate2Gender         *string `json:"teammate2_gender,omitempty"`
	Teammate2GoldenEggs     *int    `json:"teammate2_golden_eggs,omitempty"`
	Teammate2GoldieKills    *int    `json:"teammate2_goldie_kills,omitempty"`
	Teammate2GrillerKills   *int    `json:"teammate2_griller_kills,omitempty"`
	Teammate2ID             *string `json:"teammate2_id,omitempty"`
	Teammate2MawsKills      *int    `json:"teammate2_maws_kills,omitempty"`
	Teammate2Name           *string `json:"teammate2_name,omitempty"`
	Teammate2PowerEggs      *int    `json:"teammate2_power_eggs,omitempty"`
	Teammate2ReviveCount    *int    `json:"teammate2_revive_count,omitempty"`
	Teammate2ScrapperKills  *int    `json:"teammate2_scrapper_kills,omitempty"`
	Teammate2Special        *string `json:"teammate2_special,omitempty"`
	Teammate2Species        *string `json:"teammate2_species,omitempty"`
	Teammate2SteelEelKills  *int    `json:"teammate2_steel_eel_kills,omitempty"`
	Teammate2SteelheadKills *int    `json:"teammate2_steelhead_kills,omitempty"`
	Teammate2StingerKills   *int    `json:"teammate2_stinger_kills,omitempty"`
	Teammate2W1Specials     *int    `json:"teammate2_w1_specials,omitempty"`
	Teammate2W2Specials     *int    `json:"teammate2_w2_specials,omitempty"`
	Teammate2W3Specials     *int    `json:"teammate2_w3_specials,omitempty"`
	Teammate2WeaponW1       *string `json:"teammate2_weapon_w1,omitempty"`
	Teammate2WeaponW2       *string `json:"teammate2_weapon_w2,omitempty"`
	Teammate2WeaponW3       *string `json:"teammate2_weapon_w3,omitempty"`
	Wave1EventType          *string `json:"wave_1_event_type,omitempty"`
	Wave1GoldenAppear       *int    `json:"wave_1_golden_appear,omitempty"`
	Wave1GoldenDelivered    *int    `json:"wave_1_golden_delivered,omitempty"`
	Wave1PowerEggs          *int    `json:"wave_1_power_eggs,omitempty"`
	Wave1Quota              *int    `json:"wave_1_quota,omitempty"`
	Wave1WaterLevel         *string `json:"wave_1_water_level,omitempty"`
	Wave2EventType          *string `json:"wave_2_event_type,omitempty"`
	Wave2GoldenAppear       *int    `json:"wave_2_golden_appear,omitempty"`
	Wave2GoldenDelivered    *int    `json:"wave_2_golden_delivered,omitempty"`
	Wave2PowerEggs          *int    `json:"wave_2_power_eggs,omitempty"`
	Wave2Quota              *int    `json:"wave_2_quota,omitempty"`
	Wave2WaterLevel         *string `json:"wave_2_water_level,omitempty"`
	Wave3EventType          *string `json:"wave_3_event_type,omitempty"`
	Wave3GoldenAppear       *int    `json:"wave_3_golden_appear,omitempty"`
	Wave3GoldenDelivered    *int    `json:"wave_3_golden_delivered,omitempty"`
	Wave3PowerEggs          *int    `json:"wave_3_power_eggs,omitempty"`
	Wave3Quota              *int    `json:"wave_3_quota,omitempty"`
	Wave3WaterLevel         *string `json:"wave_3_water_level,omitempty"`
}

type BattleUpload struct {
	BattleNumber            *string             `json:"battle_number,omitempty"`
	ElapsedTime             *int                `json:"elapsed_time,omitempty"`
	HasDisconnectedPlayer   *bool               `json:"has_disconnected_player,omitempty"`
	LeaguePoint             decimal.NullDecimal `json:"league_point"`
	MatchType               *string             `json:"match_type,omitempty"`
	MyTeamCount             decimal.NullDecimal `json:"my_team_count"`
	Opponent0Assists        *int                `json:"opponent0_assists,omitempty"`
	Opponent0Clothes        *string             `json:"opponent0_clothes,omitempty"`
	Opponent0ClothesMain    *string             `json:"opponent0_clothes_main,omitempty"`
	Opponent0ClothesSub0    *string             `json:"opponent0_clothes_sub0,omitempty"`
	Opponent0ClothesSub1    *string             `json:"opponent0_clothes_sub1,omitempty"`
	Opponent0ClothesSub2    *string             `json:"opponent0_clothes_sub2,omitempty"`
	Opponent0Deaths         *int                `json:"opponent0_deaths,omitempty"`
	Opponent0GamePaintPoint *int                `json:"opponent0_game_paint_point,omitempty"`
	Opponent0Gender         *string             `json:"opponent0_gender,omitempty"`
	Opponent0Headgear       *string             `json:"opponent0_headgear,omitempty"`
	Opponent0HeadgearMain   *string             `json:"opponent0_headgear_main,omitempty"`
	Opponent0HeadgearSub0   *string             `json:"opponent0_headgear_sub0,omitempty"`
	Opponent0HeadgearSub1   *string             `json:"opponent0_headgear_sub1,omitempty"`
	Opponent0HeadgearSub2   *string             `json:"opponent0_headgear_sub2,omitempty"`
	Opponent0Kills          *int                `json:"opponent0_kills,omitempty"`
	Opponent0Level          *int                `json:"opponent0_level,omitempty"`
	Opponent0LevelStar      *int                `json:"opponent0_level_star,omitempty"`
	Opponent0Name           *string             `json:"opponent0_name,omitempty"`
	Opponent0Rank           *string             `json:"opponent0_rank,omitempty"`
	Opponent0Shoes          *string             `json:"opponent0_shoes,omitempty"`
	Opponent0ShoesMain      *string             `json:"opponent0_shoes_main,omitempty"`
	Opponent0ShoesSub0      *string             `json:"opponent0_shoes_sub0,omitempty"`
	Opponent0ShoesSub1      *string             `json:"opponent0_shoes_sub1,omitempty"`
	Opponent0ShoesSub2      *string             `json:"opponent0_shoes_sub2,omitempty"`
	Opponent0Specials       *int                `json:"opponent0_specials,omitempty"`
	Opponent0Species        *string             `json:"opponent0_species,omitempty"`
	Opponent0SplatnetID     *string             `json:"opponent0_splatnet_id,omitempty"`
	Opponent0Weapon         *string             `json:"opponent0_weapon,omitempty"`
	Opponent1Assists        *int                `json:"opponent1_assists,omitempty"`
	Opponent1Clothes        *string             `json:"opponent1_clothes,omitempty"`
	Opponent1ClothesMain    *string             `json:"opponent1_clothes_main,omitempty"`
	Opponent1ClothesSub0    *string             `json:"opponent1_clothes_sub0,omitempty"`
	Opponent1ClothesSub1    *string             `json:"opponent1_clothes_sub1,omitempty"`
	Opponent1ClothesSub2    *string             `json:"opponent1_clothes_sub2,omitempty"`
	Opponent1Deaths         *int                `json:"opponent1_deaths,omitempty"`
	Opponent1GamePaintPoint *int                `json:"opponent1_game_paint_point,omitempty"`
	Opponent1Gender         *string             `json:"opponent1_gender,omitempty"`
	Opponent1Headgear       *string             `json:"opponent1_headgear,omitempty"`
	Opponent1HeadgearMain   *string             `json:"opponent1_headgear_main,omitempty"`
	Opponent1HeadgearSub0   *string             `json:"opponent1_headgear_sub0,omitempty"`
	Opponent1HeadgearSub1   *string             `json:"opponent1_headgear_sub1,omitempty"`
	Opponent1HeadgearSub2   *string             `json:"opponent1_headgear_sub2,omitempty"`
	Opponent1Kills          *int                `json:"opponent1_kills,omitempty"`
	Opponent1Level          *int                `json:"opponent1_level,omitempty"`
	Opponent1LevelStar      *int                `json:"opponent1_level_star,omitempty"`
	Opponent1Name           *string             `json:"opponent1_name,omitempty"`
	Opponent1Rank           *string             `json:"opponent1_rank,omitempty"`
	Opponent1Shoes          *string             `json:"opponent1_shoes,omitempty"`
	Opponent1ShoesMain      *string             `json:"opponent1_shoes_main,omitempty"`
	Opponent1ShoesSub0      *string             `json:"opponent1_shoes_sub0,omitempty"`
	Opponent1ShoesSub1      *string             `json:"opponent1_shoes_sub1,omitempty"`
	Opponent1ShoesSub2      *string             `json:"opponent1_shoes_sub2,omitempty"`
	Opponent1Specials       *int                `json:"opponent1_specials,omitempty"`
	Opponent1Species        *string             `json:"opponent1_species,omitempty"`
	Opponent1SplatnetID     *string             `json:"opponent1_splatnet_id,omitempty"`
	Opponent1Weapon         *string             `json:"opponent1_weapon,omitempty"`
	Opponent2Assists        *int                `json:"opponent2_assists,omitempty"`
	Opponent2Clothes        *string             `json:"opponent2_clothes,omitempty"`
	Opponent2ClothesMain    *string             `json:"opponent2_clothes_main,omitempty"`
	Opponent2ClothesSub0    *string             `json:"opponent2_clothes_sub0,omitempty"`
	Opponent2ClothesSub1    *string             `json:"opponent2_clothes_sub1,omitempty"`
	Opponent2ClothesSub2    *string             `json:"opponent2_clothes_sub2,omitempty"`
	Opponent2Deaths         *int                `json:"opponent2_deaths,omitempty"`
	Opponent2GamePaintPoint *int                `json:"opponent2_game_paint_point,omitempty"`
	Opponent2Gender         *string             `json:"opponent2_gender,omitempty"`
	Opponent2Headgear       *string             `json:"opponent2_headgear,omitempty"`
	Opponent2HeadgearMain   *string             `json:"opponent2_headgear_main,omitempty"`
	Opponent2HeadgearSub0   *string             `json:"opponent2_headgear_sub0,omitempty"`
	Opponent2HeadgearSub1   *string             `json:"opponent2_headgear_sub1,omitempty"`
	Opponent2HeadgearSub2   *string             `json:"opponent2_headgear_sub2,omitempty"`
	Opponent2Kills          *int                `json:"opponent2_kills,omitempty"`
	Opponent2Level          *int                `json:"opponent2_level,omitempty"`
	Opponent2LevelStar      *int                `json:"opponent2_level_star,omitempty"`
	Opponent2Name           *string             `json:"opponent2_name,omitempty"`
	Opponent2Rank           *string             `json:"opponent2_rank,omitempty"`
	Opponent2Shoes          *string             `json:"opponent2_shoes,omitempty"`
	Opponent2ShoesMain      *string             `json:"opponent2_shoes_main,omitempty"`
	Opponent2ShoesSub0      *string             `json:"opponent2_shoes_sub0,omitempty"`
	Opponent2ShoesSub1      *string             `json:"opponent2_shoes_sub1,omitempty"`
	Opponent2ShoesSub2      *string             `json:"opponent2_shoes_sub2,omitempty"`
	Opponent2Specials       *int                `json:"opponent2_specials,omitempty"`
	Opponent2Species        *string             `json:"opponent2_species,omitempty"`
	Opponent2SplatnetID     *string             `json:"opponent2_splatnet_id,omitempty"`
	Opponent2Weapon         *string             `json:"opponent2_weapon,omitempty"`
	Opponent3Assists        *int                `json:"opponent3_assists,omitempty"`
	Opponent3Clothes        *string             `json:"opponent3_clothes,omitempty"`
	Opponent3ClothesMain    *string             `json:"opponent3_clothes_main,omitempty"`
	Opponent3ClothesSub0    *string             `json:"opponent3_clothes_sub0,omitempty"`
	Opponent3ClothesSub1    *string             `json:"opponent3_clothes_sub1,omitempty"`
	Opponent3ClothesSub2    *string             `json:"opponent3_clothes_sub2,omitempty"`
	Opponent3Deaths         *int                `json:"opponent3_deaths,omitempty"`
	Opponent3GamePaintPoint *int                `json:"opponent3_game_paint_point,omitempty"`
	Opponent3Gender         *string             `json:"opponent3_gender,omitempty"`
	Opponent3Headgear       *string             `json:"opponent3_headgear,omitempty"`
	Opponent3HeadgearMain   *string             `json:"opponent3_headgear_main,omitempty"`
	Opponent3HeadgearSub0   *string             `json:"opponent3_headgear_sub0,omitempty"`
	Opponent3HeadgearSub1   *string             `json:"opponent3_headgear_sub1,omitempty"`
	Opponent3HeadgearSub2   *string             `json:"opponent3_headgear_sub2,omitempty"`
	Opponent3Kills          *int                `json:"opponent3_kills,omitempty"`
	Opponent3Level          *int                `json:"opponent3_level,omitempty"`
	Opponent3LevelStar      *int                `json:"opponent3_level_star,omitempty"`
	Opponent3Name           *string             `json:"opponent3_name,omitempty"`
	Opponent3Rank           *string             `json:"opponent3_rank,omitempty"`
	Opponent3Shoes          *string             `json:"opponent3_shoes,omitempty"`
	Opponent3ShoesMain      *string             `json:"opponent3_shoes_main,omitempty"`
	Opponent3ShoesSub0      *string             `json:"opponent3_shoes_sub0,omitempty"`
	Opponent3ShoesSub1      *string             `json:"opponent3_shoes_sub1,omitempty"`
	Opponent3ShoesSub2      *string             `json:"opponent3_shoes_sub2,omitempty"`
	Opponent3Specials       *int                `json:"opponent3_specials,omitempty"`
	Opponent3Species        *string             `json:"opponent3_species,omitempty"`
	Opponent3SplatnetID     *string             `json:"opponent3_splatnet_id,omitempty"`
	Opponent3Weapon         *string             `json:"opponent3_weapon,omitempty"`
	OtherTeamCount          decimal.NullDecimal `json:"other_team_count"`
	PlayerAssists           *int                `json:"player_assists,omitempty"`
	PlayerClothes           *string             `json:"player_clothes,omitempty"`
	PlayerClothesMain       *string             `json:"player_clothes_main,omitempty"`
	PlayerClothesSub0       *string             `json:"player_clothes_sub0,omitempty"`
	PlayerClothesSub1       *string             `json:"player_clothes_sub1,omitempty"`
	PlayerClothesSub2       *string             `json:"player_clothes_sub2,omitempty"`
	PlayerDeaths            *int                `json:"player_deaths,omitempty"`
	PlayerGamePaintPoint    *int                `json:"player_game_paint_point,omitempty"`
	PlayerGender            *string             `json:"player_gender,omitempty"`
	PlayerHeadgear          *string             `json:"player_headgear,omitempty"`
	PlayerHeadgearMain      *string             `json:"player_headgear_main,omitempty"`
	PlayerHeadgearSub0      *string             `json:"player_headgear_sub0,omitempty"`
	PlayerHeadgearSub1      *string             `json:"player_headgear_sub1,omitempty"`
	PlayerHeadgearSub2      *string             `json:"player_headgear_sub2,omitempty"`
	PlayerKills             *int                `json:"player_kills,omitempty"`
	PlayerLevel             *int                `json:"player_level,omitempty"`
	PlayerLevelStar         *int                `json:"player_level_star,omitempty"`
	PlayerName              *string             `json:"player_name,omitempty"`
	PlayerRank              *int                `json:"player_rank,omitempty"`
	PlayerShoes             *string             `json:"player_shoes,omitempty"`
	PlayerShoesMain         *string             `json:"player_shoes_main,omitempty"`
	PlayerShoesSub0         *string             `json:"player_shoes_sub0,omitempty"`
	PlayerShoesSub1         *string             `json:"player_shoes_sub1,omitempty"`
	PlayerShoesSub2         *string             `json:"player_shoes_sub2,omitempty"`
	PlayerSpecials          *int                `json:"player_specials,omitempty"`
	PlayerSpecies           *string             `json:"player_species,omitempty"`
	PlayerSplatfestTitle    *string             `json:"player_splatfest_title,omitempty"`
	PlayerSplatnetID        *string             `json:"player_splatnet_id,omitempty"`
	PlayerWeapon            *string             `json:"player_weapon,omitempty"`
	PlayerXPower            decimal.NullDecimal `json:"player_x_power"`
	Rule                    *string             `json:"rule,omitempty"`
	SplatfestPoint          decimal.NullDecimal `json:"splatfest_point"`
	SplatfestTitleAfter     *string             `json:"splatfest_title_after,omitempty"`
	SplatnetJSON            Battle              `json:"splatnet_json"`
	SplatnetUpload          *bool               `json:"splatnet_upload,omitempty"`
	Stage                   *string             `json:"stage,omitempty"`
	StatInkUpload           *bool               `json:"stat_ink_upload,omitempty"`
	TagID                   *string             `json:"tag_id,omitempty"`
	Teammate0Assists        *int                `json:"teammate0_assists,omitempty"`
	Teammate0Clothes        *string             `json:"teammate0_clothes,omitempty"`
	Teammate0ClothesMain    *string             `json:"teammate0_clothes_main,omitempty"`
	Teammate0ClothesSub0    *string             `json:"teammate0_clothes_sub0,omitempty"`
	Teammate0ClothesSub1    *string             `json:"teammate0_clothes_sub1,omitempty"`
	Teammate0ClothesSub2    *string             `json:"teammate0_clothes_sub2,omitempty"`
	Teammate0Deaths         *int                `json:"teammate0_deaths,omitempty"`
	Teammate0GamePaintPoint *int                `json:"teammate0_game_paint_point,omitempty"`
	Teammate0Gender         *string             `json:"teammate0_gender,omitempty"`
	Teammate0Headgear       *string             `json:"teammate0_headgear,omitempty"`
	Teammate0HeadgearMain   *string             `json:"teammate0_headgear_main,omitempty"`
	Teammate0HeadgearSub0   *string             `json:"teammate0_headgear_sub0,omitempty"`
	Teammate0HeadgearSub1   *string             `json:"teammate0_headgear_sub1,omitempty"`
	Teammate0HeadgearSub2   *string             `json:"teammate0_headgear_sub2,omitempty"`
	Teammate0Kills          *int                `json:"teammate0_kills,omitempty"`
	Teammate0Level          *int                `json:"teammate0_level,omitempty"`
	Teammate0LevelStar      *int                `json:"teammate0_level_star,omitempty"`
	Teammate0Name           *string             `json:"teammate0_name,omitempty"`
	Teammate0Rank           *string             `json:"teammate0_rank,omitempty"`
	Teammate0Shoes          *string             `json:"teammate0_shoes,omitempty"`
	Teammate0ShoesMain      *string             `json:"teammate0_shoes_main,omitempty"`
	Teammate0ShoesSub0      *string             `json:"teammate0_shoes_sub0,omitempty"`
	Teammate0ShoesSub1      *string             `json:"teammate0_shoes_sub1,omitempty"`
	Teammate0ShoesSub2      *string             `json:"teammate0_shoes_sub2,omitempty"`
	Teammate0Specials       *int                `json:"teammate0_specials,omitempty"`
	Teammate0Species        *string             `json:"teammate0_species,omitempty"`
	Teammate0SplatnetID     *string             `json:"teammate0_splatnet_id,omitempty"`
	Teammate0Weapon         *string             `json:"teammate0_weapon,omitempty"`
	Teammate1Assists        *int                `json:"teammate1_assists,omitempty"`
	Teammate1Clothes        *string             `json:"teammate1_clothes,omitempty"`
	Teammate1ClothesMain    *string             `json:"teammate1_clothes_main,omitempty"`
	Teammate1ClothesSub0    *string             `json:"teammate1_clothes_sub0,omitempty"`
	Teammate1ClothesSub1    *string             `json:"teammate1_clothes_sub1,omitempty"`
	Teammate1ClothesSub2    *string             `json:"teammate1_clothes_sub2,omitempty"`
	Teammate1Deaths         *int                `json:"teammate1_deaths,omitempty"`
	Teammate1GamePaintPoint *int                `json:"teammate1_game_paint_point,omitempty"`
	Teammate1Gender         *string             `json:"teammate1_gender,omitempty"`
	Teammate1Headgear       *string             `json:"teammate1_headgear,omitempty"`
	Teammate1HeadgearMain   *string             `json:"teammate1_headgear_main,omitempty"`
	Teammate1HeadgearSub0   *string             `json:"teammate1_headgear_sub0,omitempty"`
	Teammate1HeadgearSub1   *string             `json:"teammate1_headgear_sub1,omitempty"`
	Teammate1HeadgearSub2   *string             `json:"teammate1_headgear_sub2,omitempty"`
	Teammate1Kills          *int                `json:"teammate1_kills,omitempty"`
	Teammate1Level          *int                `json:"teammate1_level,omitempty"`
	Teammate1LevelStar      *int                `json:"teammate1_level_star,omitempty"`
	Teammate1Name           *string             `json:"teammate1_name,omitempty"`
	Teammate1Rank           *string             `json:"teammate1_rank,omitempty"`
	Teammate1Shoes          *string             `json:"teammate1_shoes,omitempty"`
	Teammate1ShoesMain      *string             `json:"teammate1_shoes_main,omitempty"`
	Teammate1ShoesSub0      *string             `json:"teammate1_shoes_sub0,omitempty"`
	Teammate1ShoesSub1      *string             `json:"teammate1_shoes_sub1,omitempty"`
	Teammate1ShoesSub2      *string             `json:"teammate1_shoes_sub2,omitempty"`
	Teammate1Specials       *int                `json:"teammate1_specials,omitempty"`
	Teammate1Species        *string             `json:"teammate1_species,omitempty"`
	Teammate1SplatnetID     *string             `json:"teammate1_splatnet_id,omitempty"`
	Teammate1Weapon         *string             `json:"teammate1_weapon,omitempty"`
	Teammate2Assists        *int                `json:"teammate2_assists,omitempty"`
	Teammate2Clothes        *string             `json:"teammate2_clothes,omitempty"`
	Teammate2ClothesMain    *string             `json:"teammate2_clothes_main,omitempty"`
	Teammate2ClothesSub0    *string             `json:"teammate2_clothes_sub0,omitempty"`
	Teammate2ClothesSub1    *string             `json:"teammate2_clothes_sub1,omitempty"`
	Teammate2ClothesSub2    *string             `json:"teammate2_clothes_sub2,omitempty"`
	Teammate2Deaths         *int                `json:"teammate2_deaths,omitempty"`
	Teammate2GamePaintPoint *int                `json:"teammate2_game_paint_point,omitempty"`
	Teammate2Gender         *string             `json:"teammate2_gender,omitempty"`
	Teammate2Headgear       *string             `json:"teammate2_headgear,omitempty"`
	Teammate2HeadgearMain   *string             `json:"teammate2_headgear_main,omitempty"`
	Teammate2HeadgearSub0   *string             `json:"teammate2_headgear_sub0,omitempty"`
	Teammate2HeadgearSub1   *string             `json:"teammate2_headgear_sub1,omitempty"`
	Teammate2HeadgearSub2   *string             `json:"teammate2_headgear_sub2,omitempty"`
	Teammate2Kills          *int                `json:"teammate2_kills,omitempty"`
	Teammate2Level          *int                `json:"teammate2_level,omitempty"`
	Teammate2LevelStar      *int                `json:"teammate2_level_star,omitempty"`
	Teammate2Name           *string             `json:"teammate2_name,omitempty"`
	Teammate2Rank           *string             `json:"teammate2_rank,omitempty"`
	Teammate2Shoes          *string             `json:"teammate2_shoes,omitempty"`
	Teammate2ShoesMain      *string             `json:"teammate2_shoes_main,omitempty"`
	Teammate2ShoesSub0      *string             `json:"teammate2_shoes_sub0,omitempty"`
	Teammate2ShoesSub1      *string             `json:"teammate2_shoes_sub1,omitempty"`
	Teammate2ShoesSub2      *string             `json:"teammate2_shoes_sub2,omitempty"`
	Teammate2Specials       *int                `json:"teammate2_specials,omitempty"`
	Teammate2Species        *string             `json:"teammate2_species,omitempty"`
	Teammate2SplatnetID     *string             `json:"teammate2_splatnet_id,omitempty"`
	Teammate2Weapon         *string             `json:"teammate2_weapon,omitempty"`
	Time                    *int                `json:"time,omitempty"`
	Win                     *bool               `json:"win,omitempty"`
	WinMeter                decimal.NullDecimal `json:"win_meter"`
}

type BattleList struct {
	UniqueID string `json:"unique_id"`
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
	Results []Battle `json:"results"`
}
