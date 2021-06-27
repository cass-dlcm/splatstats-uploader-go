package types

import "github.com/cass-dlcm/splatstatsuploader/enums"

type Battle struct {
	UserId                  int64
	SplatnetJson            *BattleSplatnet         `json:"splatnet_json,omitempty"`
	SplatnetUpload          bool                    `json:"splatnet_upload"`
	StatInkJson             *BattleStatInk          `json:"stat_ink_json,omitempty"`
	StatInkUpload           bool                    `json:"stat_ink_upload"`
	BattleNumber            int                     `json:"battle_number"`
	PlayerSplatnetId        string                  `json:"player_splatnet_id"`
	ElapsedTime             int                     `json:"elapsed_time"`
	HasDisconnectedPlayer   bool                    `json:"has_disconnected_player"`
	LeaguePoint             *float64                `json:"league_point,omitempty"`
	MatchType               string                  `json:"match_type"`
	Rule                    string                  `json:"rule"`
	MyTeamCount             float64                 `json:"my_team_count"`
	OtherTeamCount          float64                 `json:"other_team_count"`
	SplatfestPoint          *float64                `json:"splatfest_point,omitempty"`
	SplatfestTitle          *string                 `json:"splatfest_title,omitempty"`
	SplatfestTitleAfter     *string                 `json:"splatfest_title_after"`
	Stage                   string                  `json:"stage"`
	TagId                   *string                 `json:"tag_id,omitempty"`
	Time                    int                     `json:"time,omitempty"`
	Win                     bool                    `json:"win,omitempty"`
	WinMeter                *float64                `json:"win_meter,omitempty"`
	Opponent0SplatnetId     *string                 `json:"opponent0_splatnet_id,omitempty"`
	Opponent0Name           *string                 `json:"opponent0_name,omitempty"`
	Opponent0Rank           *string                 `json:"opponent0_rank,omitempty"`
	Opponent0LevelStar      *int                    `json:"opponent0_level_star,omitempty"`
	Opponent0Level          *int                    `json:"opponent0_level,omitempty"`
	Opponent0Weapon         *enums.BattleWeaponEnum `json:"opponent0_weapon,omitempty"`
	Opponent0Gender         *enums.GenderEnum       `json:"opponent0_gender,omitempty"`
	Opponent0Species        *enums.SpeciesEnum      `json:"opponent0_species,omitempty"`
	Opponent0Assists        *int                    `json:"opponent0_assists,omitempty"`
	Opponent0Deaths         *int                    `json:"opponent0_deaths,omitempty"`
	Opponent0GamePaintPoint *int                    `json:"opponent0_game_paint_point,omitempty"`
	Opponent0Kills          *int                    `json:"opponent0_kills,omitempty"`
	Opponent0Specials       *int                    `json:"opponent0_specials,omitempty"`
	Opponent0Headgear       *string                 `json:"opponent0_headgear,omitempty"`
	Opponent0HeadgearMain   *string                 `json:"opponent0_headgear_main,omitempty"`
	Opponent0HeadgearSub0   *string                 `json:"opponent0_headgear_sub0,omitempty"`
	Opponent0HeadgearSub1   *string                 `json:"opponent0_headgear_sub1,omitempty"`
	Opponent0HeadgearSub2   *string                 `json:"opponent0_headgear_sub2,omitempty"`
	Opponent0Clothes        *string                 `json:"opponent0_clothes,omitempty"`
	Opponent0ClothesMain    *string                 `json:"opponent0_clothes_main,omitempty"`
	Opponent0ClothesSub0    *string                 `json:"opponent0_clothes_sub0,omitempty"`
	Opponent0ClothesSub1    *string                 `json:"opponent0_clothes_sub1,omitempty"`
	Opponent0ClothesSub2    *string                 `json:"opponent0_clothes_sub2,omitempty"`
	Opponent0Shoes          *string                 `json:"opponent0_shoes,omitempty"`
	Opponent0ShoesMain      *string                 `json:"opponent0_shoes_main,omitempty"`
	Opponent0ShoesSub0      *string                 `json:"opponent0_shoes_sub0,omitempty"`
	Opponent0ShoesSub1      *string                 `json:"opponent0_shoes_sub1,omitempty"`
	Opponent0ShoesSub2      *string                 `json:"opponent0_shoes_sub2,omitempty"`
	Opponent1SplatnetId     *string                 `json:"opponent1_splatnet_id,omitempty"`
	Opponent1Name           *string                 `json:"opponent1_name,omitempty"`
	Opponent1Rank           *string                 `json:"opponent1_rank,omitempty"`
	Opponent1LevelStar      *int                    `json:"opponent1_level_star,omitempty"`
	Opponent1Level          *int                    `json:"opponent1_level,omitempty"`
	Opponent1Weapon         *enums.BattleWeaponEnum `json:"opponent1_weapon,omitempty"`
	Opponent1Gender         *enums.GenderEnum       `json:"opponent1_gender,omitempty"`
	Opponent1Species        *enums.SpeciesEnum      `json:"opponent1_species,omitempty"`
	Opponent1Assists        *int                    `json:"opponent1_assists,omitempty"`
	Opponent1Deaths         *int                    `json:"opponent1_deaths,omitempty"`
	Opponent1GamePaintPoint *int                    `json:"opponent1_game_paint_point,omitempty"`
	Opponent1Kills          *int                    `json:"opponent1_kills,omitempty"`
	Opponent1Specials       *int                    `json:"opponent1_specials,omitempty"`
	Opponent1Headgear       *string                 `json:"opponent1_headgear,omitempty"`
	Opponent1HeadgearMain   *string                 `json:"opponent1_headgear_main,omitempty"`
	Opponent1HeadgearSub0   *string                 `json:"opponent1_headgear_sub0,omitempty"`
	Opponent1HeadgearSub1   *string                 `json:"opponent1_headgear_sub1,omitempty"`
	Opponent1HeadgearSub2   *string                 `json:"opponent1_headgear_sub2,omitempty"`
	Opponent1Clothes        *string                 `json:"opponent1_clothes,omitempty"`
	Opponent1ClothesMain    *string                 `json:"opponent1_clothes_main,omitempty"`
	Opponent1ClothesSub0    *string                 `json:"opponent1_clothes_sub0,omitempty"`
	Opponent1ClothesSub1    *string                 `json:"opponent1_clothes_sub1,omitempty"`
	Opponent1ClothesSub2    *string                 `json:"opponent1_clothes_sub2,omitempty"`
	Opponent1Shoes          *string                 `json:"opponent1_shoes,omitempty"`
	Opponent1ShoesMain      *string                 `json:"opponent1_shoes_main,omitempty"`
	Opponent1ShoesSub0      *string                 `json:"opponent1_shoes_sub0,omitempty"`
	Opponent1ShoesSub1      *string                 `json:"opponent1_shoes_sub1,omitempty"`
	Opponent1ShoesSub2      *string                 `json:"opponent1_shoes_sub2,omitempty"`
	Opponent2SplatnetId     *string                 `json:"opponent2_splatnet_id,omitempty"`
	Opponent2Name           *string                 `json:"opponent2_name,omitempty"`
	Opponent2Rank           *string                 `json:"opponent2_rank,omitempty"`
	Opponent2LevelStar      *int                    `json:"opponent2_level_star,omitempty"`
	Opponent2Level          *int                    `json:"opponent2_level,omitempty"`
	Opponent2Weapon         *enums.BattleWeaponEnum `json:"opponent2_weapon,omitempty"`
	Opponent2Gender         *enums.GenderEnum       `json:"opponent2_gender,omitempty"`
	Opponent2Species        *enums.SpeciesEnum      `json:"opponent2_species,omitempty"`
	Opponent2Assists        *int                    `json:"opponent2_assists,omitempty"`
	Opponent2Deaths         *int                    `json:"opponent2_deaths,omitempty"`
	Opponent2GamePaintPoint *int                    `json:"opponent2_game_paint_point,omitempty"`
	Opponent2Kills          *int                    `json:"opponent2_kills,omitempty"`
	Opponent2Specials       *int                    `json:"opponent2_specials,omitempty"`
	Opponent2Headgear       *string                 `json:"opponent2_headgear,omitempty"`
	Opponent2HeadgearMain   *string                 `json:"opponent2_headgear_main,omitempty"`
	Opponent2HeadgearSub0   *string                 `json:"opponent2_headgear_sub0,omitempty"`
	Opponent2HeadgearSub1   *string                 `json:"opponent2_headgear_sub1,omitempty"`
	Opponent2HeadgearSub2   *string                 `json:"opponent2_headgear_sub2,omitempty"`
	Opponent2Clothes        *string                 `json:"opponent2_clothes,omitempty"`
	Opponent2ClothesMain    *string                 `json:"opponent2_clothes_main,omitempty"`
	Opponent2ClothesSub0    *string                 `json:"opponent2_clothes_sub0,omitempty"`
	Opponent2ClothesSub1    *string                 `json:"opponent2_clothes_sub1,omitempty"`
	Opponent2ClothesSub2    *string                 `json:"opponent2_clothes_sub2,omitempty"`
	Opponent2Shoes          *string                 `json:"opponent2_shoes,omitempty"`
	Opponent2ShoesMain      *string                 `json:"opponent2_shoes_main,omitempty"`
	Opponent2ShoesSub0      *string                 `json:"opponent2_shoes_sub0,omitempty"`
	Opponent2ShoesSub1      *string                 `json:"opponent2_shoes_sub1,omitempty"`
	Opponent2ShoesSub2      *string                 `json:"opponent2_shoes_sub2,omitempty"`
	Opponent3SplatnetId     *string                 `json:"opponent3_splatnet_id,omitempty"`
	Opponent3Name           *string                 `json:"opponent3_name,omitempty"`
	Opponent3Rank           *string                 `json:"opponent3_rank,omitempty"`
	Opponent3LevelStar      *int                    `json:"opponent3_level_star,omitempty"`
	Opponent3Level          *int                    `json:"opponent3_level,omitempty"`
	Opponent3Weapon         *enums.BattleWeaponEnum `json:"opponent3_weapon,omitempty"`
	Opponent3Gender         *enums.GenderEnum       `json:"opponent3_gender,omitempty"`
	Opponent3Species        *enums.SpeciesEnum      `json:"opponent3_species,omitempty"`
	Opponent3Assists        *int                    `json:"opponent3_assists,omitempty"`
	Opponent3Deaths         *int                    `json:"opponent3_deaths,omitempty"`
	Opponent3GamePaintPoint *int                    `json:"opponent3_game_paint_point,omitempty"`
	Opponent3Kills          *int                    `json:"opponent3_kills,omitempty"`
	Opponent3Specials       *int                    `json:"opponent3_specials,omitempty"`
	Opponent3Headgear       *string                 `json:"opponent3_headgear,omitempty"`
	Opponent3HeadgearMain   *string                 `json:"opponent3_headgear_main,omitempty"`
	Opponent3HeadgearSub0   *string                 `json:"opponent3_headgear_sub0,omitempty"`
	Opponent3HeadgearSub1   *string                 `json:"opponent3_headgear_sub1,omitempty"`
	Opponent3HeadgearSub2   *string                 `json:"opponent3_headgear_sub2,omitempty"`
	Opponent3Clothes        *string                 `json:"opponent3_clothes,omitempty"`
	Opponent3ClothesMain    *string                 `json:"opponent3_clothes_main,omitempty"`
	Opponent3ClothesSub0    *string                 `json:"opponent3_clothes_sub0,omitempty"`
	Opponent3ClothesSub1    *string                 `json:"opponent3_clothes_sub1,omitempty"`
	Opponent3ClothesSub2    *string                 `json:"opponent3_clothes_sub2,omitempty"`
	Opponent3Shoes          *string                 `json:"opponent3_shoes,omitempty"`
	Opponent3ShoesMain      *string                 `json:"opponent3_shoes_main,omitempty"`
	Opponent3ShoesSub0      *string                 `json:"opponent3_shoes_sub0,omitempty"`
	Opponent3ShoesSub1      *string                 `json:"opponent3_shoes_sub1,omitempty"`
	Opponent3ShoesSub2      *string                 `json:"opponent3_shoes_sub2,omitempty"`
	Teammate0SplatnetId     *string                 `json:"teammate0_splatnet_id,omitempty"`
	Teammate0Name           *string                 `json:"teammate0_name,omitempty"`
	Teammate0Rank           *string                 `json:"teammate0_rank,omitempty"`
	Teammate0LevelStar      *int                    `json:"teammate0_level_star,omitempty"`
	Teammate0Level          *int                    `json:"teammate0_level,omitempty"`
	Teammate0Weapon         *enums.BattleWeaponEnum `json:"teammate0_weapon,omitempty"`
	Teammate0Gender         *enums.GenderEnum       `json:"teammate0_gender,omitempty"`
	Teammate0Species        *enums.SpeciesEnum      `json:"teammate0_species,omitempty"`
	Teammate0Assists        *int                    `json:"teammate0_assists,omitempty"`
	Teammate0Deaths         *int                    `json:"teammate0_deaths,omitempty"`
	Teammate0GamePaintPoint *int                    `json:"teammate0_game_paint_point,omitempty"`
	Teammate0Kills          *int                    `json:"teammate0_kills,omitempty"`
	Teammate0Specials       *int                    `json:"teammate0_specials,omitempty"`
	Teammate0Headgear       *string                 `json:"teammate0_headgear,omitempty"`
	Teammate0HeadgearMain   *string                 `json:"teammate0_headgear_main,omitempty"`
	Teammate0HeadgearSub0   *string                 `json:"teammate0_headgear_sub0,omitempty"`
	Teammate0HeadgearSub1   *string                 `json:"teammate0_headgear_sub1,omitempty"`
	Teammate0HeadgearSub2   *string                 `json:"teammate0_headgear_sub2,omitempty"`
	Teammate0Clothes        *string                 `json:"teammate0_clothes,omitempty"`
	Teammate0ClothesMain    *string                 `json:"teammate0_clothes_main,omitempty"`
	Teammate0ClothesSub0    *string                 `json:"teammate0_clothes_sub0,omitempty"`
	Teammate0ClothesSub1    *string                 `json:"teammate0_clothes_sub1,omitempty"`
	Teammate0ClothesSub2    *string                 `json:"teammate0_clothes_sub2,omitempty"`
	Teammate0Shoes          *string                 `json:"teammate0_shoes,omitempty"`
	Teammate0ShoesMain      *string                 `json:"teammate0_shoes_main,omitempty"`
	Teammate0ShoesSub0      *string                 `json:"teammate0_shoes_sub0,omitempty"`
	Teammate0ShoesSub1      *string                 `json:"teammate0_shoes_sub1,omitempty"`
	Teammate0ShoesSub2      *string                 `json:"teammate0_shoes_sub2,omitempty"`
	Teammate1SplatnetId     *string                 `json:"teammate1_splatnet_id,omitempty"`
	Teammate1Name           *string                 `json:"teammate1_name,omitempty"`
	Teammate1Rank           *string                 `json:"teammate1_rank,omitempty"`
	Teammate1LevelStar      *int                    `json:"teammate1_level_star,omitempty"`
	Teammate1Level          *int                    `json:"teammate1_level,omitempty"`
	Teammate1Weapon         *enums.BattleWeaponEnum `json:"teammate1_weapon,omitempty"`
	Teammate1Gender         *enums.GenderEnum       `json:"teammate1_gender,omitempty"`
	Teammate1Species        *enums.SpeciesEnum      `json:"teammate1_species,omitempty"`
	Teammate1Assists        *int                    `json:"teammate1_assists,omitempty"`
	Teammate1Deaths         *int                    `json:"teammate1_deaths,omitempty"`
	Teammate1GamePaintPoint *int                    `json:"teammate1_game_paint_point,omitempty"`
	Teammate1Kills          *int                    `json:"teammate1_kills,omitempty"`
	Teammate1Specials       *int                    `json:"teammate1_specials,omitempty"`
	Teammate1Headgear       *string                 `json:"teammate1_headgear,omitempty"`
	Teammate1HeadgearMain   *string                 `json:"teammate1_headgear_main,omitempty"`
	Teammate1HeadgearSub0   *string                 `json:"teammate1_headgear_sub0,omitempty"`
	Teammate1HeadgearSub1   *string                 `json:"teammate1_headgear_sub1,omitempty"`
	Teammate1HeadgearSub2   *string                 `json:"teammate1_headgear_sub2,omitempty"`
	Teammate1Clothes        *string                 `json:"teammate1_clothes,omitempty"`
	Teammate1ClothesMain    *string                 `json:"teammate1_clothes_main,omitempty"`
	Teammate1ClothesSub0    *string                 `json:"teammate1_clothes_sub0,omitempty"`
	Teammate1ClothesSub1    *string                 `json:"teammate1_clothes_sub1,omitempty"`
	Teammate1ClothesSub2    *string                 `json:"teammate1_clothes_sub2,omitempty"`
	Teammate1Shoes          *string                 `json:"teammate1_shoes,omitempty"`
	Teammate1ShoesMain      *string                 `json:"teammate1_shoes_main,omitempty"`
	Teammate1ShoesSub0      *string                 `json:"teammate1_shoes_sub0,omitempty"`
	Teammate1ShoesSub1      *string                 `json:"teammate1_shoes_sub1,omitempty"`
	Teammate1ShoesSub2      *string                 `json:"teammate1_shoes_sub2,omitempty"`
	Teammate2SplatnetId     *string                 `json:"teammate2_splatnet_id,omitempty"`
	Teammate2Name           *string                 `json:"teammate2_name,omitempty"`
	Teammate2Rank           *string                 `json:"teammate2_rank,omitempty"`
	Teammate2LevelStar      *int                    `json:"teammate2_level_star,omitempty"`
	Teammate2Level          *int                    `json:"teammate2_level,omitempty"`
	Teammate2Weapon         *enums.BattleWeaponEnum `json:"teammate2_weapon,omitempty"`
	Teammate2Gender         *enums.GenderEnum       `json:"teammate2_gender,omitempty"`
	Teammate2Species        *enums.SpeciesEnum      `json:"teammate2_species,omitempty"`
	Teammate2Assists        *int                    `json:"teammate2_assists,omitempty"`
	Teammate2Deaths         *int                    `json:"teammate2_deaths,omitempty"`
	Teammate2GamePaintPoint *int                    `json:"teammate2_game_paint_point,omitempty"`
	Teammate2Kills          *int                    `json:"teammate2_kills,omitempty"`
	Teammate2Specials       *int                    `json:"teammate2_specials,omitempty"`
	Teammate2Headgear       *string                 `json:"teammate2_headgear,omitempty"`
	Teammate2HeadgearMain   *string                 `json:"teammate2_headgear_main,omitempty"`
	Teammate2HeadgearSub0   *string                 `json:"teammate2_headgear_sub0,omitempty"`
	Teammate2HeadgearSub1   *string                 `json:"teammate2_headgear_sub1,omitempty"`
	Teammate2HeadgearSub2   *string                 `json:"teammate2_headgear_sub2,omitempty"`
	Teammate2Clothes        *string                 `json:"teammate2_clothes,omitempty"`
	Teammate2ClothesMain    *string                 `json:"teammate2_clothes_main,omitempty"`
	Teammate2ClothesSub0    *string                 `json:"teammate2_clothes_sub0,omitempty"`
	Teammate2ClothesSub1    *string                 `json:"teammate2_clothes_sub1,omitempty"`
	Teammate2ClothesSub2    *string                 `json:"teammate2_clothes_sub2,omitempty"`
	Teammate2Shoes          *string                 `json:"teammate2_shoes,omitempty"`
	Teammate2ShoesMain      *string                 `json:"teammate2_shoes_main,omitempty"`
	Teammate2ShoesSub0      *string                 `json:"teammate2_shoes_sub0,omitempty"`
	Teammate2ShoesSub1      *string                 `json:"teammate2_shoes_sub1,omitempty"`
	Teammate2ShoesSub2      *string                 `json:"teammate2_shoes_sub2,omitempty"`
	PlayerName              string                  `json:"player_name"`
	PlayerRank              *string                 `json:"player_rank,omitempty"`
	PlayerLevelStar         int                     `json:"player_level_star"`
	PlayerLevel             int                     `json:"player_level"`
	PlayerWeapon            enums.BattleWeaponEnum  `json:"player_weapon"`
	PlayerGender            enums.GenderEnum        `json:"player_gender"`
	PlayerSpecies           enums.SpeciesEnum       `json:"player_species"`
	PlayerAssists           int                     `json:"player_assists"`
	PlayerDeaths            int                     `json:"player_deaths"`
	PlayerGamePaintPoint    int                     `json:"player_game_paint_point"`
	PlayerKills             int                     `json:"player_kills"`
	PlayerSpecials          int                     `json:"player_specials"`
	PlayerHeadgear          string                  `json:"player_headgear"`
	PlayerHeadgearMain      string                  `json:"player_headgear_main"`
	PlayerHeadgearSub0      string                  `json:"player_headgear_sub0"`
	PlayerHeadgearSub1      string                  `json:"player_headgear_sub1"`
	PlayerHeadgearSub2      string                  `json:"player_headgear_sub2"`
	PlayerClothes           string                  `json:"player_clothes"`
	PlayerClothesMain       string                  `json:"player_clothes_main"`
	PlayerClothesSub0       string                  `json:"player_clothes_sub0"`
	PlayerClothesSub1       string                  `json:"player_clothes_sub1"`
	PlayerClothesSub2       string                  `json:"player_clothes_sub2"`
	PlayerShoes             string                  `json:"player_shoes"`
	PlayerShoesMain         string                  `json:"player_shoes_main"`
	PlayerShoesSub0         string                  `json:"player_shoes_sub0"`
	PlayerShoesSub1         string                  `json:"player_shoes_sub1"`
	PlayerShoesSub2         string                  `json:"player_shoes_sub2"`
}

type BattleSplatnet struct {
	Udemae              *BattleSplatnetUdemae        `json:"udemae,omitempty"`
	Stage               SplatnetTriple               `json:"stage"`
	OtherTeamCount      *int                         `json:"other_team_count,omitempty"`
	MyTeamCount         *int                         `json:"my_team_count,omitempty"`
	StarRank            int                          `json:"star_rank"`
	Rule                BattleSplatnetRule           `json:"rule"`
	PlayerResult        BattleSplatnetPlayerResult   `json:"player_result"`
	EstimateGachiPower  *int                         `json:"estimate_gachi_power"`
	ElapsedTime         *int                         `json:"elapsed_time"`
	StartTime           int                          `json:"start_time"`
	GameMode            SplatnetDouble               `json:"game_mode"`
	XPower              *interface{}                 `json:"x_power,omitempty"`
	BattleNumber        string                       `json:"battle_number"`
	Type                string                       `json:"type"`
	PlayerRank          int                          `json:"player_rank"`
	CrownPlayers        *interface{}                 `json:"crown_players,omitempty"`
	MyTeamMembers       []BattleSplatnetPlayerResult `json:"my_team_members"`
	OtherTeamMembers    []BattleSplatnetPlayerResult `json:"other_team_members"`
	WeaponPaintPoint    int                          `json:"weapon_paint_point"`
	Rank                *interface{}                 `json:"rank,omitempty"`
	MyTeamResult        SplatnetDouble               `json:"my_team_result"`
	EstimateXPower      *interface{}                 `json:"estimate_x_power,omitempty"`
	OtherTeamResult     SplatnetDouble               `json:"other_team_result"`
	LeaguePoint         *float64                     `json:"league_point,omitempty"`
	WinMeter            *float64                     `json:"win_meter,omitempty"`
	MyTeamPercentage    *float64                     `json:"my_team_percentage,omitempty"`
	OtherTeamPercentage *float64                     `json:"other_team_percentage,omitempty"`
	TagId               *string                      `json:"tag_id,omitempty"`
}

type BattleSplatnetUdemae struct {
	Name            string `json:"name,omitempty"`
	IsX             bool   `json:"is_x,omitempty"`
	IsNumberReached bool   `json:"is_number_reached,omitempty"`
	Number          int    `json:"number,omitempty"`
	SPlusNumber     int    `json:"s_plus_number,omitempty"`
}

type BattleSplatnetRule struct {
	Key           string `json:"key"`
	Name          string `json:"name"`
	MultilineName string `json:"multiline_name"`
}

type BattleSplatnetPlayerResult struct {
	DeathCount     int                              `json:"death_count"`
	GamePaintPoint int                              `json:"game_paint_point"`
	KillCount      int                              `json:"kill_count"`
	SpecialCount   int                              `json:"special_count"`
	AssistCount    int                              `json:"assist_count"`
	SortScore      int                              `json:"sort_score"`
	Player         BattleSplatnetPlayerResultPlayer `json:"player"`
}

type BattleSplatnetPlayerResultPlayer struct {
	HeadSkills    BattleSplatnetPlayerResultPlayerSkills   `json:"head_skills"`
	ShoesSkills   BattleSplatnetPlayerResultPlayerSkills   `json:"shoes_skills"`
	ClothesSkills BattleSplatnetPlayerResultPlayerSkills   `json:"clothes_skills"`
	PlayerRank    int                                      `json:"player_rank"`
	StarRank      int                                      `json:"star_rank"`
	Nickname      string                                   `json:"nickname"`
	PlayerType    SplatnetPlayerType                       `json:"player_type"`
	PrincipalId   string                                   `json:"principal_id"`
	Head          BattleSplatnetPlayerResultPlayerClothing `json:"head"`
	Clothes       BattleSplatnetPlayerResultPlayerClothing `json:"clothes"`
	Shoes         BattleSplatnetPlayerResultPlayerClothing `json:"shoes"`
	Udemae        *BattleSplatnetUdemae                    `json:"udemae,omitempty"`
	Weapon        BattleSplatnetPlayerResultPlayerWeapon   `json:"weapon"`
}

type BattleSplatnetPlayerResultPlayerSkills struct {
	Main SplatnetTriple   `json:"main,omitempty"`
	Subs []SplatnetTriple `json:"subs,omitempty"`
}

type BattleSplatnetPlayerResultPlayerClothing struct {
	Id        string                                        `json:"id"`
	Image     string                                        `json:"image"`
	Name      string                                        `json:"name"`
	Thumbnail string                                        `json:"thumbnail"`
	Kind      string                                        `json:"kind"`
	Rarity    int                                           `json:"rarity"`
	Brand     BattleSplatnetPlayerResultPlayerClothingBrand `json:"brand"`
}

type BattleSplatnetPlayerResultPlayerWeapon struct {
	Id        enums.BattleWeaponEnum `json:"id"`
	Image     string                 `json:"image"`
	Name      string                 `json:"name"`
	Thumbnail string                 `json:"thumbnail"`
	Sub       SplatnetQuad           `json:"sub"`
	Special   SplatnetQuad           `json:"special"`
}

type BattleSplatnetPlayerResultPlayerClothingBrand struct {
	Id            string         `json:"id"`
	Image         string         `json:"image"`
	Name          string         `json:"name"`
	FrequentSkill SplatnetTriple `json:"frequent_skill"`
}

type BattleStatInk struct {
	Id                             int                     `json:"id"`
	SplatnetNumber                 int                     `json:"splatnet_number"`
	Url                            string                  `json:"url"`
	User                           BattleStatInkUser       `json:"user"`
	Lobby                          StatInkKeyName          `json:"lobby"`
	Mode                           StatInkKeyName          `json:"mode"`
	Rule                           StatInkKeyName          `json:"rule"`
	Map                            BattleStatInkMap        `json:"map"`
	Weapon                         BattleStatInkWeapon     `json:"weapon,omitempty"`
	Freshness                      *BattleStatInkFreshness `json:"freshness,omitempty"`
	Rank                           *BattleStatInkRank      `json:"rank,omitempty"`
	RankExp                        *int                    `json:"rank_exp,omitempty"`
	RankAfter                      *BattleStatInkRank      `json:"rank_after,omitempty"`
	XPower                         *interface{}            `json:"x_power,omitempty"`
	XPowerAfter                    *interface{}            `json:"x_power_after,omitempty"`
	EstimateXPower                 *interface{}            `json:"estimate_x_power,omitempty"`
	Level                          int                     `json:"level"`
	LevelAfter                     int                     `json:"level_after"`
	StarRank                       int                     `json:"star_rank"`
	Result                         string                  `json:"result"`
	KnockOut                       bool                    `json:"knock_out"`
	RankInTeam                     int                     `json:"rank_in_team"`
	Kill                           int                     `json:"kill"`
	Death                          int                     `json:"death"`
	KillOrAssist                   int                     `json:"kill_or_assist"`
	Special                        int                     `json:"special"`
	KillRatio                      float64                 `json:"kill_ratio"`
	KillRate                       float64                 `json:"kill_rate"`
	MaxKillCombo                   *interface{}            `json:"max_kill_combo,omitempty"`
	MaxKillStreak                  *interface{}            `json:"max_kill_streak,omitempty"`
	DeathReasons                   *interface{}            `json:"death_reasons,omitempty"`
	MyPoint                        int                     `json:"my_point"`
	EstimateGachiPower             *int                    `json:"estimate_gachi_power,omitempty"`
	LeaguePoint                    *string                 `json:"league_point,omitempty"`
	MyTeamEstimateLeaguePoint      *int                    `json:"my_team_estimate_league_point,omitempty"`
	HisTeamEstimateLeaguePoint     *int                    `json:"his_team_estimate_league_point,omitempty"`
	MyTeamPoint                    *interface{}            `json:"my_team_point,omitempty"`
	HisTeamPoint                   *interface{}            `json:"his_team_point,omitempty"`
	MyTeamPercent                  *string                 `json:"my_team_percent,omitempty"`
	HisTeamPercent                 *string                 `json:"his_team_percent,omitempty"`
	MyTeamId                       *string                 `json:"my_team_id,omitempty"`
	HisTeamId                      *string                 `json:"his_team_id,omitempty"`
	Species                        StatInkKeyName          `json:"species"`
	Gender                         StatInkGender           `json:"gender"`
	FestTitle                      *StatInkKeyName         `json:"fest_title,omitempty"`
	FestExp                        *int                    `json:"fest_exp,omitempty"`
	FestTitleAfter                 *StatInkKeyName         `json:"fest_title_after,omitempty"`
	FestExpAfter                   *int                    `json:"fest_exp_after,omitempty"`
	FestPower                      *string                 `json:"fest_power,omitempty"`
	MyTeamEstimateFestPower        *int                    `json:"my_team_estimate_fest_power,omitempty"`
	HisTeamMyTeamEstimateFestPower *int                    `json:"his_team_my_team_estimate_fest_power,omitempty"`
	MyTeamFestTheme                *string                 `json:"my_team_fest_theme,omitempty"`
	MyTeamNickname                 *string                 `json:"my_team_nickname,omitempty"`
	HisTeamNickname                *string                 `json:"his_team_nickname,omitempty"`
	Clout                          *int                    `json:"clout,omitempty"`
	TotalClout                     *int                    `json:"total_clout,omitempty"`
	TotalCloutAfter                *int                    `json:"total_clout_after,omitempty"`
	MyTeamWinStreak                *int                    `json:"my_team_win_streak,omitempty"`
	HisTeamWinStreak               *int                    `json:"his_team_win_streak,omitempty"`
	SynergyBonus                   *float64                `json:"synergy_bonus,omitempty"`
	SpecialBattle                  *StatInkKeyName         `json:"special_battle,omitempty"`
	ImageResult                    *string                 `json:"image_result"`
	ImageGear                      *string                 `json:"image_gear"`
	Gears                          BattleStatInkGears      `json:"gears"`
	Period                         int                     `json:"period"`
	PeriodRange                    string                  `json:"period_range"`
	Players                        []BattleStatInkPlayer   `json:"players"`
	Events                         *interface{}            `json:"events,omitempty"`
	SplatnetJson                   *interface{}            `json:"splatnet_json,omitempty"`
	Agent                          BattleStatInkAgent      `json:"agent"`
	Automated                      bool                    `json:"automated"`
	Environment                    *interface{}            `json:"environment,omitempty"`
	LinkUrl                        string                  `json:"link_url"`
	Note                           *interface{}            `json:"note,omitempty"`
	GameVersion                    string                  `json:"game_version"`
	NawabariBonus                  *int                    `json:"nawabari_bonus,omitempty"`
	StartAt                        StatInkTime             `json:"start_at"`
	EndAt                          StatInkTime             `json:"end_at"`
	RegisterAt                     StatInkTime             `json:"register_at"`
	MyTeamCount                    *int
	HisTeamCount                   *int
}

type BattleStatInkUser struct {
	Id         int                    `json:"id"`
	Name       string                 `json:"name"`
	ScreenName string                 `json:"screen_name"`
	Url        string                 `json:"url"`
	JoinAt     StatInkTime            `json:"join_at"`
	Profile    StatInkProfile         `json:"profile"`
	Stat       *interface{}           `json:"stat,omitempty"`
	Stats      BattleStatInkUserStats `json:"stats"`
}

type BattleStatInkMap struct {
	Key       string      `json:"key"`
	Name      StatInkName `json:"name"`
	Splatnet  int         `json:"splatnet"`
	Area      int         `json:"area"`
	ReleaseAt StatInkTime `json:"release_at"`
	ShortName StatInkName `json:"short_name"`
}

type BattleStatInkWeapon struct {
	Key         enums.BattleStatinkWeaponEnum `json:"key"`
	Name        StatInkName                   `json:"name"`
	Splatnet    int                           `json:"splatnet"`
	Type        BattleStatInkWeaponType       `json:"type"`
	ReskinOf    *string                       `json:"reskin_of"`
	MainRef     string                        `json:"main_ref"`
	Sub         StatInkKeyName                `json:"sub"`
	Special     StatInkKeyName                `json:"special"`
	MainPowerUp StatInkKeyName                `json:"main_power_up"`
}

type BattleStatInkFreshness struct {
	Freshness float64     `json:"freshness"`
	Title     StatInkName `json:"title"`
}

type BattleStatInkRank struct {
	Key  string         `json:"key"`
	Name StatInkName    `json:"name"`
	Zone StatInkKeyName `json:"zone"`
}

type BattleStatInkGears struct {
	Headgear BattleStatInkGearsClothes `json:"headgear"`
	Clothing BattleStatInkGearsClothes `json:"clothing"`
	Shoes    BattleStatInkGearsClothes `json:"shoes"`
}

type BattleStatInkPlayer struct {
	Team         string              `json:"team"`
	IsMe         bool                `json:"is_me"`
	Weapon       BattleStatInkWeapon `json:"weapon"`
	Level        int                 `json:"level"`
	Rank         *BattleStatInkRank  `json:"rank,omitempty"`
	StarRank     int                 `json:"star_rank"`
	RankInTeam   int                 `json:"rank_in_team"`
	Kill         int                 `json:"kill"`
	Death        int                 `json:"death"`
	KillOrAssist int                 `json:"kill_or_assist"`
	Special      int                 `json:"special"`
	MyKill       *interface{}        `json:"my_kill,omitempty"`
	Point        int                 `json:"point"`
	Name         string              `json:"name"`
	Species      StatInkKeyName      `json:"species"`
	Gender       StatInkGender       `json:"gender"`
	FestTitle    *StatInkKeyName     `json:"fest_title,omitempty"`
	SplatnetId   string              `json:"splatnet_id"`
	Top500       bool                `json:"top_500"`
	Icon         string              `json:"icon"`
}

type BattleStatInkAgent struct {
	Name            string                       `json:"name"`
	Version         string                       `json:"version"`
	GameVersion     *interface{}                 `json:"game_version,omitempty"`
	GameVersionDate *interface{}                 `json:"game_version_date,omitempty"`
	Custom          *interface{}                 `json:"custom,omitempty"`
	Variables       *BattleStatInkAgentVariables `json:"variables,omitempty"`
}

type BattleStatInkUserStats struct {
	V1 *interface{}              `json:"v1,omitempty"`
	V2 *BattleStatInkUserStatsV2 `json:"v2,omitempty"`
}

type BattleStatInkWeaponType struct {
	Key      string         `json:"key"`
	Name     StatInkName    `json:"name"`
	Category StatInkKeyName `json:"category"`
}

type BattleStatInkGearsClothes struct {
	Gear               BattleStatInkGearsClothesGear `json:"gear"`
	PrimaryAbility     StatInkKeyName                `json:"primary_ability"`
	SecondaryAbilities []StatInkKeyName              `json:"secondary_abilities"`
}

type BattleStatInkAgentVariables struct {
	UploadMode string `json:"upload_mode,omitempty"`
}

type BattleStatInkUserStatsV2 struct {
	UpdatedAt StatInkTime                      `json:"updated_at"`
	Entire    BattleStatInkUserStatsV2Entire   `json:"entire"`
	Nawabari  BattleStatInkUserStatsV2Nawabari `json:"nawabari"`
	Gachi     BattleStatInkUserStatsV2Gachi    `json:"gachi"`
}

type BattleStatInkGearsClothesGear struct {
	Key            string         `json:"key"`
	Name           StatInkName    `json:"name"`
	Splatnet       int            `json:"splatnet"`
	Type           StatInkKeyName `json:"type"`
	Brand          StatInkKeyName `json:"brand"`
	PrimaryAbility StatInkKeyName `json:"primary_ability"`
}

type BattleStatInkUserStatsV2Entire struct {
	Battles     int     `json:"battles"`
	WinPct      float64 `json:"win_pct"`
	KillRatio   float64 `json:"kill_ratio"`
	KillTotal   int     `json:"kill_total"`
	KillAvg     float64 `json:"kill_avg"`
	KillPerMin  float64 `json:"kill_per_min"`
	DeathTotal  int     `json:"death_total"`
	DeathAvg    float64 `json:"death_avg"`
	DeathPerMin float64 `json:"death_per_min"`
}

type BattleStatInkUserStatsV2Nawabari struct {
	Battles     int     `json:"battles"`
	WinPct      float64 `json:"win_pct"`
	KillRatio   float64 `json:"kill_ratio"`
	KillTotal   int     `json:"kill_total"`
	KillAvg     float64 `json:"kill_avg"`
	KillPerMin  float64 `json:"kill_per_min"`
	DeathTotal  int     `json:"death_total"`
	DeathAvg    float64 `json:"death_avg"`
	DeathPerMin float64 `json:"death_per_min"`
	TotalInked  int     `json:"total_inked"`
	MaxInked    int     `json:"max_inked"`
	AvgInked    float64 `json:"avg_inked"`
}

type BattleStatInkUserStatsV2Gachi struct {
	Battles     int                                `json:"battles"`
	WinPct      float64                            `json:"win_pct"`
	KillRatio   float64                            `json:"kill_ratio"`
	KillTotal   int                                `json:"kill_total"`
	KillAvg     float64                            `json:"kill_avg"`
	KillPerMin  float64                            `json:"kill_per_min"`
	DeathTotal  int                                `json:"death_total"`
	DeathAvg    float64                            `json:"death_avg"`
	DeathPerMin float64                            `json:"death_per_min"`
	Rules       BattleStatInkUserStatsV2GachiRules `json:"rules"`
}

type BattleStatInkUserStatsV2GachiRules struct {
	Area   BattleStatInkUserStatsV2GachiRulesSub `json:"area"`
	Yagura BattleStatInkUserStatsV2GachiRulesSub `json:"yagura"`
	Hoko   BattleStatInkUserStatsV2GachiRulesSub `json:"hoko"`
	Asari  BattleStatInkUserStatsV2GachiRulesSub `json:"asari"`
}

type BattleStatInkUserStatsV2GachiRulesSub struct {
	RankPeak    string `json:"rank_peak"`
	RankCurrent string `json:"rank_current"`
	//XPowerPeak    *interface{} `json:"x_power_peak,omitempty"`
	//XPowerCurrent *interface{} `json:"x_power_current,omitempty"`
}
