package types

import "github.com/cass-dlcm/splatstatsuploader/enums"

type Shift struct {
	UserId                  int64                           `json:"user_id"`
	PlayerSplatnetId        string                          `json:"player_splatnet_id"`
	JobId                   int64                           `json:"job_id"`
	SplatnetUpload          bool                            `json:"splatnet_upload"`
	StatInkUpload           bool                            `json:"stat_ink_upload"`
	SplatnetJson            *ShiftSplatnet                  `json:"splatnet_json,omitempty"`
	StatInkJson             *ShiftStatInk                   `json:"stat_ink_json,omitempty"`
	StartTime               int64                           `json:"start_time"`
	PlayTime                int64                           `json:"play_time"`
	EndTime                 int64                           `json:"end_time"`
	DangerRate              float64                         `json:"danger_rate"`
	IsClear                 bool                            `json:"is_clear"`
	JobFailureReason        *enums.FailureReasonEnum        `json:"job_failure_reason,omitempty"`
	FailureWave             *int                            `json:"failure_wave,omitempty"`
	GradePoint              int                             `json:"grade_point"`
	GradePointDelta         int                             `json:"grade_point_delta"`
	JobScore                int                             `json:"job_score"`
	DrizzlerCount           int                             `json:"drizzler_count"`
	FlyfishCount            int                             `json:"flyfish_count"`
	GoldieCount             int                             `json:"goldie_count"`
	GrillerCount            int                             `json:"griller_count"`
	MawsCount               int                             `json:"maws_count"`
	ScrapperCount           int                             `json:"scrapper_count"`
	SteelEelCount           int                             `json:"steel_eel_count"`
	SteelheadCount          int                             `json:"steelhead_count"`
	StingerCount            int                             `json:"stinger_count"`
	Stage                   enums.SalmonStageEnum           `json:"stage"`
	PlayerName              string                          `json:"player_name"`
	PlayerDeathCount        int                             `json:"player_death_count"`
	PlayerReviveCount       int                             `json:"player_revive_count"`
	PlayerGoldenEggs        int                             `json:"player_golden_eggs"`
	PlayerPowerEggs         int                             `json:"player_power_eggs"`
	PlayerSpecial           string                          `json:"player_special"`
	PlayerTitle             string                          `json:"player_title"`
	PlayerSpecies           enums.SpeciesEnum               `json:"player_species"`
	PlayerGender            enums.GenderEnum                `json:"player_gender"`
	PlayerW1Specials        int                             `json:"player_w1_specials"`
	PlayerW2Specials        *int                            `json:"player_w2_specials,omitempty"`
	PlayerW3Specials        *int                            `json:"player_w3_specials,omitempty"`
	PlayerW1Weapon          enums.SalmonWeaponEnum          `json:"player_w1_weapon"`
	PlayerW2Weapon          *enums.SalmonWeaponEnum         `json:"player_w2_weapon,omitempty"`
	PlayerW3Weapon          *enums.SalmonWeaponEnum         `json:"player_w3_weapon,omitempty"`
	PlayerDrizzlerKills     int                             `json:"player_drizzler_kills"`
	PlayerFlyfishKills      int                             `json:"player_flyfish_kills"`
	PlayerGoldieKills       int                             `json:"player_goldie_kills"`
	PlayerGrillerKills      int                             `json:"player_griller_kills"`
	PlayerMawsKills         int                             `json:"player_maws_kills"`
	PlayerScrapperKills     int                             `json:"player_scrapper_kills"`
	PlayerSteelEelKills     int                             `json:"player_steel_eel_kills"`
	PlayerSteelheadKills    int                             `json:"player_steelhead_kills"`
	PlayerStingerKills      int                             `json:"player_stinger_kills"`
	Teammate0SplatnetId     *string                         `json:"teammate_0_splatnet_id,omitempty"`
	Teammate0Name           *string                         `json:"teammate_0_name,omitempty"`
	Teammate0DeathCount     *int                            `json:"teammate_0_death_count,omitempty"`
	Teammate0ReviveCount    *int                            `json:"teammate_0_revive_count,omitempty"`
	Teammate0GoldenEggs     *int                            `json:"teammate_0_golden_eggs,omitempty"`
	Teammate0PowerEggs      *int                            `json:"teammate_0_power_eggs,omitempty"`
	Teammate0Special        *string                         `json:"teammate_0_special,omitempty"`
	Teammate0Species        *enums.SpeciesEnum              `json:"teammate_0_species,omitempty"`
	Teammate0Gender         *enums.GenderEnum               `json:"teammate_0_gender,omitempty"`
	Teammate0W1Specials     *int                            `json:"teammate_0_w1_specials,omitempty"`
	Teammate0W2Specials     *int                            `json:"teammate_0_w2_specials,omitempty"`
	Teammate0W3Specials     *int                            `json:"teammate_0_w3_specials,omitempty"`
	Teammate0W1Weapon       *enums.SalmonWeaponEnum         `json:"teammate_0_w1_weapon,omitempty"`
	Teammate0W2Weapon       *enums.SalmonWeaponEnum         `json:"teammate_0_w2_weapon,omitempty"`
	Teammate0W3Weapon       *enums.SalmonWeaponEnum         `json:"teammate_0_w3_weapon,omitempty"`
	Teammate0DrizzlerKills  *int                            `json:"teammate_0_drizzler_kills,omitempty"`
	Teammate0FlyfishKills   *int                            `json:"teammate_0_flyfish_kills,omitempty"`
	Teammate0GoldieKills    *int                            `json:"teammate_0_goldie_kills,omitempty"`
	Teammate0GrillerKills   *int                            `json:"teammate_0_griller_kills,omitempty"`
	Teammate0MawsKills      *int                            `json:"teammate_0_maws_kills,omitempty"`
	Teammate0ScrapperKills  *int                            `json:"teammate_0_scrapper_kills,omitempty"`
	Teammate0SteelEelKills  *int                            `json:"teammate_0_steel_eel_kills,omitempty"`
	Teammate0SteelheadKills *int                            `json:"teammate_0_steelhead_kills,omitempty"`
	Teammate0StingerKills   *int                            `json:"teammate_0_stinger_kills,omitempty"`
	Teammate1SplatnetId     *string                         `json:"teammate_1_splatnet_id,omitempty"`
	Teammate1Name           *string                         `json:"teammate_1_name,omitempty"`
	Teammate1DeathCount     *int                            `json:"teammate_1_death_count,omitempty"`
	Teammate1ReviveCount    *int                            `json:"teammate_1_revive_count,omitempty"`
	Teammate1GoldenEggs     *int                            `json:"teammate_1_golden_eggs,omitempty"`
	Teammate1PowerEggs      *int                            `json:"teammate_1_power_eggs,omitempty"`
	Teammate1Special        *string                         `json:"teammate_1_special,omitempty"`
	Teammate1Species        *enums.SpeciesEnum              `json:"teammate_1_species,omitempty"`
	Teammate1Gender         *enums.GenderEnum               `json:"teammate_1_gender,omitempty"`
	Teammate1W1Specials     *int                            `json:"teammate_1_w1_specials,omitempty"`
	Teammate1W2Specials     *int                            `json:"teammate_1_w2_specials,omitempty"`
	Teammate1W3Specials     *int                            `json:"teammate_1_w3_specials,omitempty"`
	Teammate1W1Weapon       *enums.SalmonWeaponEnum         `json:"teammate_1_w1_weapon,omitempty"`
	Teammate1W2Weapon       *enums.SalmonWeaponEnum         `json:"teammate_1_w2_weapon,omitempty"`
	Teammate1W3Weapon       *enums.SalmonWeaponEnum         `json:"teammate_1_w3_weapon,omitempty"`
	Teammate1DrizzlerKills  *int                            `json:"teammate_1_drizzler_kills,omitempty"`
	Teammate1FlyfishKills   *int                            `json:"teammate_1_flyfish_kills,omitempty"`
	Teammate1GoldieKills    *int                            `json:"teammate_1_goldie_kills,omitempty"`
	Teammate1GrillerKills   *int                            `json:"teammate_1_griller_kills,omitempty"`
	Teammate1MawsKills      *int                            `json:"teammate_1_maws_kills,omitempty"`
	Teammate1ScrapperKills  *int                            `json:"teammate_1_scrapper_kills,omitempty"`
	Teammate1SteelEelKills  *int                            `json:"teammate_1_steel_eel_kills,omitempty"`
	Teammate1SteelheadKills *int                            `json:"teammate_1_steelhead_kills,omitempty"`
	Teammate1StingerKills   *int                            `json:"teammate_1_stinger_kills,omitempty"`
	Teammate2SplatnetId     *string                         `json:"teammate_2_splatnet_id,omitempty"`
	Teammate2Name           *string                         `json:"teammate_2_name,omitempty"`
	Teammate2DeathCount     *int                            `json:"teammate_2_death_count,omitempty"`
	Teammate2ReviveCount    *int                            `json:"teammate_2_revive_count,omitempty"`
	Teammate2GoldenEggs     *int                            `json:"teammate_2_golden_eggs,omitempty"`
	Teammate2PowerEggs      *int                            `json:"teammate_2_power_eggs,omitempty"`
	Teammate2Special        *string                         `json:"teammate_2_special,omitempty"`
	Teammate2Species        *enums.SpeciesEnum              `json:"teammate_2_species,omitempty"`
	Teammate2Gender         *enums.GenderEnum               `json:"teammate_2_gender,omitempty"`
	Teammate2W1Specials     *int                            `json:"teammate_2_w1_specials,omitempty"`
	Teammate2W2Specials     *int                            `json:"teammate_2_w2_specials,omitempty"`
	Teammate2W3Specials     *int                            `json:"teammate_2_w3_specials,omitempty"`
	Teammate2W1Weapon       *enums.SalmonWeaponEnum         `json:"teammate_2_w1_weapon,omitempty"`
	Teammate2W2Weapon       *enums.SalmonWeaponEnum         `json:"teammate_2_w2_weapon,omitempty"`
	Teammate2W3Weapon       *enums.SalmonWeaponEnum         `json:"teammate_2_w3_weapon,omitempty"`
	Teammate2DrizzlerKills  *int                            `json:"teammate_2_drizzler_kills,omitempty"`
	Teammate2FlyfishKills   *int                            `json:"teammate_2_flyfish_kills,omitempty"`
	Teammate2GoldieKills    *int                            `json:"teammate_2_goldie_kills,omitempty"`
	Teammate2GrillerKills   *int                            `json:"teammate_2_griller_kills,omitempty"`
	Teammate2MawsKills      *int                            `json:"teammate_2_maws_kills,omitempty"`
	Teammate2ScrapperKills  *int                            `json:"teammate_2_scrapper_kills,omitempty"`
	Teammate2SteelEelKills  *int                            `json:"teammate_2_steel_eel_kills,omitempty"`
	Teammate2SteelheadKills *int                            `json:"teammate_2_steelhead_kills,omitempty"`
	Teammate2StingerKills   *int                            `json:"teammate_2_stinger_kills,omitempty"`
	ScheduleEndTime         *int64                          `json:"schedule_end_time,omitempty"`
	ScheduleStartTime       int64                           `json:"schedule_start_time,omitempty"`
	ScheduleWeapon0         *enums.SalmonWeaponScheduleEnum `json:"schedule_weapon_0,omitempty"`
	ScheduleWeapon1         *enums.SalmonWeaponScheduleEnum `json:"schedule_weapon_1,omitempty"`
	ScheduleWeapon2         *enums.SalmonWeaponScheduleEnum `json:"schedule_weapon_2,omitempty"`
	ScheduleWeapon3         *enums.SalmonWeaponScheduleEnum `json:"schedule_weapon_3,omitempty"`
	Wave1WaterLevel         string                          `json:"wave_1_water_level"`
	Wave1EventType          string                          `json:"wave_1_event_type"`
	Wave1GoldenDelivered    int                             `json:"wave_1_golden_ikura_num"`
	Wave1GoldenAppear       int                             `json:"wave_1_golden_ikura_pop_num"`
	Wave1PowerEggs          int                             `json:"wave_1_ikura_num"`
	Wave1Quota              int                             `json:"wave_1_quota_num"`
	Wave2WaterLevel         *string                         `json:"wave_2_water_level,omitempty"`
	Wave2EventType          *string                         `json:"wave_2_event_type,omitempty"`
	Wave2GoldenDelivered    *int                            `json:"wave_2_golden_ikura_num,omitempty"`
	Wave2GoldenAppear       *int                            `json:"wave_2_golden_ikura_pop_num,omitempty"`
	Wave2PowerEggs          *int                            `json:"wave_2_ikura_num,omitempty"`
	Wave2Quota              *int                            `json:"wave_2_quota_num,omitempty"`
	Wave3WaterLevel         *string                         `json:"wave_3_water_level,omitempty"`
	Wave3EventType          *string                         `json:"wave_3_event_type,omitempty"`
	Wave3GoldenDelivered    *int                            `json:"wave_3_golden_ikura_num,omitempty"`
	Wave3GoldenAppear       *int                            `json:"wave_3_golden_ikura_pop_num,omitempty"`
	Wave3PowerEggs          *int                            `json:"wave_3_ikura_num,omitempty"`
	Wave3Quota              *int                            `json:"wave_3_quota_num,omitempty"`
}

type ShiftSplatnet struct {
	JobId           int64                   `json:"job_id"`
	DangerRate      float64                 `json:"danger_rate"`
	JobResult       ShiftSplatnetJobResult  `json:"job_result"`
	JobScore        int                     `json:"job_score"`
	JobRate         int                     `json:"job_rate"`
	GradePoint      int                     `json:"grade_point"`
	GradePointDelta int                     `json:"grade_point_delta"`
	OtherResults    []ShiftSplatnetPlayer   `json:"other_results"`
	KumaPoint       int                     `json:"kuma_point"`
	StartTime       int64                   `json:"start_time"`
	PlayerType      SplatnetPlayerType      `json:"player_type"`
	PlayTime        int64                   `json:"play_time"`
	BossCounts      ShiftSplatnetBossCounts `json:"boss_counts"`
	EndTime         int64                   `json:"end_time"`
	MyResult        ShiftSplatnetPlayer     `json:"my_result"`
	WaveDetails     []ShiftSplatnetWave     `json:"wave_details"`
	Grade           ShiftSplatnetGrade      `json:"grade"`
	Schedule        ShiftSplatnetSchedule   `json:"schedule"`
}

type ShiftSplatnetJobResult struct {
	IsClear       bool                     `json:"is_clear,omitempty"`
	FailureReason *enums.FailureReasonEnum `json:"failure_reason,omitempty"`
	FailureWave   *int                     `json:"failure_wave,omitempty"`
}

type ShiftSplatnetPlayer struct {
	SpecialCounts  []int                           `json:"special_counts"`
	Special        SplatnetQuad                    `json:"special"`
	Pid            string                          `json:"pid"`
	PlayerType     SplatnetPlayerType              `json:"player_type"`
	WeaponList     []ShiftSplatnetPlayerWeaponList `json:"weapon_list"`
	Name           string                          `json:"name"`
	DeadCount      int                             `json:"dead_count"`
	GoldenEggs     int                             `json:"golden_ikura_num"`
	BossKillCounts ShiftSplatnetBossCounts         `json:"boss_kill_counts"`
	PowerEggs      int                             `json:"ikura_num"`
	HelpCount      int                             `json:"help_count"`
}

type ShiftSplatnetPlayerWeaponList struct {
	Id     string                              `json:"id"`
	Weapon ShiftSplatnetPlayerWeaponListWeapon `json:"weapon"`
}

type ShiftSplatnetPlayerWeaponListWeapon struct {
	Id        enums.SalmonWeaponEnum `json:"id"`
	Image     string                 `json:"image"`
	Name      string                 `json:"name"`
	Thumbnail string                 `json:"thumbnail"`
}

type ShiftSplatnetBossCounts struct {
	Goldie    ShiftSplatnetBossCountsBoss `json:"3"`
	Steelhead ShiftSplatnetBossCountsBoss `json:"6"`
	Flyfish   ShiftSplatnetBossCountsBoss `json:"9"`
	Scrapper  ShiftSplatnetBossCountsBoss `json:"12"`
	SteelEel  ShiftSplatnetBossCountsBoss `json:"13"`
	Stinger   ShiftSplatnetBossCountsBoss `json:"14"`
	Maws      ShiftSplatnetBossCountsBoss `json:"15"`
	Griller   ShiftSplatnetBossCountsBoss `json:"16"`
	Drizzler  ShiftSplatnetBossCountsBoss `json:"21"`
}

type ShiftSplatnetBossCountsBoss struct {
	Boss  SplatnetDouble `json:"boss"`
	Count int            `json:"count"`
}

type ShiftSplatnetWave struct {
	WaterLevel   SplatnetDouble `json:"water_level"`
	EventType    SplatnetDouble `json:"event_type"`
	GoldenEggs   int            `json:"golden_ikura_num"`
	GoldenAppear int            `json:"golden_ikura_pop_num"`
	PowerEggs    int            `json:"ikura_num"`
	QuotaNum     int            `json:"quota_num"`
}

type ShiftSplatnetGrade struct {
	Id        string `json:"id,omitempty"`
	ShortName string `json:"short_name,omitempty"`
	LongName  string `json:"long_name,omitempty"`
	Name      string `json:"name,omitempty"`
}

type ShiftSplatnetSchedule struct {
	StartTime int64                         `json:"start_time"`
	Weapons   []ShiftSplatnetScheduleWeapon `json:"weapons"`
	EndTime   int64                         `json:"end_time"`
	Stage     ShiftSplatnetScheduleStage    `json:"stage"`
}

type ShiftSplatnetScheduleWeapon struct {
	Id                string                                    `json:"id"`
	Weapon            *ShiftSplatnetScheduleWeaponWeapon        `json:"weapon"`
	CoopSpecialWeapon *ShiftSplatnetScheduleWeaponSpecialWeapon `json:"coop_special_weapon"`
}

type ShiftSplatnetScheduleWeaponWeapon struct {
	Id        string                         `json:"id"`
	Image     string                         `json:"image"`
	Name      enums.SalmonWeaponScheduleEnum `json:"name"`
	Thumbnail string                         `json:"thumbnail"`
}

type ShiftSplatnetScheduleWeaponSpecialWeapon struct {
	Image string                                `json:"image"`
	Name  enums.SalmonWeaponScheduleSpecialEnum `json:"name"`
}

type ShiftSplatnetScheduleStage struct {
	Image enums.SalmonSplatnetScheduleStageImageEnum `json:"image"`
	Name  enums.SalmonStageEnum                      `json:"name"`
}

type ShiftStatInk struct {
	Id              int                     `json:"id"`
	Uuid            string                  `json:"uuid"`
	SplatnetNumber  int64                   `json:"splatnet_number"`
	Url             string                  `json:"url"`
	ApiEndpoint     string                  `json:"api_endpoint"`
	User            ShiftStatInkUser        `json:"user"`
	Stage           ShiftStatInkStage       `json:"stage"`
	IsCleared       bool                    `json:"is_cleared"`
	FailReason      *ShiftStatInkFailReason `json:"fail_reason,omitempty"`
	ClearWaves      int                     `json:"clear_waves"`
	DangerRate      string                  `json:"danger_rate"`
	Quota           []int                   `json:"quota"`
	Title           ShiftStatInkTitle       `json:"title"`
	TitleExp        int                     `json:"title_exp"`
	TitleAfter      ShiftStatInkTitle       `json:"title_after"`
	TitleExpAfter   int                     `json:"title_exp_after"`
	BossAppearances []ShiftStatInkBossData  `json:"boss_appearances"`
	Waves           []ShiftStatInkWave      `json:"waves"`
	MyData          ShiftStatInkPlayer      `json:"my_data"`
	Teammates       []ShiftStatInkPlayer    `json:"teammates"`
	Agent           ShiftStatInkAgent       `json:"agent"`
	Automated       bool                    `json:"automated"`
	Note            *interface{}            `json:"note,omitempty"`
	LinkUrl         *interface{}            `json:"link_url,omitempty"`
	ShiftStartAt    StatInkTime             `json:"shift_start_at"`
	StartAt         StatInkTime             `json:"start_at"`
	EndAt           *interface{}            `json:"end_at,omitempty"`
	RegisterAt      StatInkTime             `json:"register_at"`
}

type ShiftStatInkFailReason struct {
	Key  enums.FailureReasonEnum `json:"key"`
	Name StatInkName             `json:"name"`
}

type ShiftStatInkUser struct {
	Id         int                   `json:"id"`
	Name       string                `json:"name"`
	ScreenName string                `json:"screen_name"`
	Url        string                `json:"url"`
	SalmonUrl  string                `json:"salmon_url"`
	BattleUrl  string                `json:"battle_url"`
	JoinAt     StatInkTime           `json:"join_at"`
	Profile    StatInkProfile        `json:"profile"`
	Stats      ShiftStatInkUserStats `json:"stats"`
}

type ShiftStatInkUserStats struct {
	WorkCount       int         `json:"work_count"`
	TotalGoldenEggs int         `json:"total_golden_eggs"`
	TotalEggs       int         `json:"total_eggs"`
	TotalRescued    int         `json:"total_rescued"`
	TotalPoint      int         `json:"total_point"`
	AsOf            StatInkTime `json:"as_of"`
	RegisteredAt    StatInkTime `json:"registered_at"`
}

type ShiftStatInkTripleInt struct {
	Key      string      `json:"key"`
	Name     StatInkName `json:"name"`
	Splatnet int         `json:"splatnet"`
}

type ShiftStatInkTripleString struct {
	Key      string      `json:"key"`
	Name     StatInkName `json:"name"`
	Splatnet string      `json:"splatnet"`
}

type ShiftStatInkStage struct {
	Key      string                `json:"key"`
	Name     ShiftStatInkStageName `json:"name"`
	Splatnet string                `json:"splatnet"`
}

type ShiftStatInkStageName struct {
	DeDE string                `json:"de_DE,omitempty"`
	EnGB string                `json:"en_GB,omitempty"`
	EnUS enums.SalmonStageEnum `json:"en_US,omitempty"`
	EsES string                `json:"es_ES,omitempty"`
	EsMX string                `json:"es_MX,omitempty"`
	FrCA string                `json:"fr_CA,omitempty"`
	FrFR string                `json:"fr_FR,omitempty"`
	ItIT string                `json:"it_IT,omitempty"`
	JaJP string                `json:"ja_JP,omitempty"`
	NlNL string                `json:"nl_NL,omitempty"`
	RuRU string                `json:"ru_RU,omitempty"`
	ZhCN string                `json:"zh_CN,omitempty"`
	ZhTW string                `json:"zh_TW,omitempty"`
}

type ShiftStatInkTitle struct {
	Splatnet    int         `json:"splatnet"`
	GenericName StatInkName `json:"generic_name"`
}

type ShiftStatInkBossData struct {
	Boss  ShiftStatInkBossDataBoss `json:"boss"`
	Count int                      `json:"count"`
}

type ShiftStatInkBossDataBoss struct {
	Splatnet    int    `json:"splatnet"`
	SplatnetStr string `json:"splatnetStr"`
}

type ShiftStatInkWave struct {
	KnownOccurrence      *ShiftStatInkTripleString `json:"known_occurrence,omitempty"`
	WaterLevel           ShiftStatInkTripleString `json:"water_level,omitempty"`
	GoldenEggQuota       int                       `json:"golden_egg_quota"`
	GoldenEggAppearances int                       `json:"golden_egg_appearances"`
	GoldenEggDelivered   int                       `json:"golden_egg_delivered"`
	PowerEggCollected    int                       `json:"power_egg_collected"`
}

type ShiftStatInkPlayer struct {
	SplatnetId         string                  `json:"splatnet_id"`
	Name               string                  `json:"name"`
	Special            ShiftStatInkTripleInt   `json:"special"`
	Rescue             int                     `json:"rescue"`
	Death              int                     `json:"death"`
	GoldenEggDelivered int                     `json:"golden_egg_delivered"`
	PowerEggCollected  int                     `json:"power_egg_collected"`
	Species            StatInkKeyName          `json:"species"`
	Gender             StatInkGender           `json:"gender"`
	SpecialUses        []int                   `json:"special_uses"`
	Weapons            []ShiftStatInkTripleInt `json:"weapons"`
	BossKills          []ShiftStatInkBossData  `json:"boss_kills"`
}

type ShiftStatInkAgent struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
