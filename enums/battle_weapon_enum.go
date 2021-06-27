package enums

import (
	"encoding/json"
	"errors"
	"fmt"
)

type BattleWeaponEnum string

const (
	SPLOOSH           BattleWeaponEnum = "0"
	NeoSploosh        BattleWeaponEnum = "1"
	Sploosh7          BattleWeaponEnum = "2"
	JR                BattleWeaponEnum = "10"
	CustomJr          BattleWeaponEnum = "11"
	KensaJr           BattleWeaponEnum = "12"
	SPLASH            BattleWeaponEnum = "20"
	NeoSplash         BattleWeaponEnum = "21"
	AeroMg            BattleWeaponEnum = "30"
	AeroRg            BattleWeaponEnum = "31"
	AeroPg            BattleWeaponEnum = "32"
	SPLATTERSHOT      BattleWeaponEnum = "40"
	TtekSplattershot  BattleWeaponEnum = "41"
	KensaSplattershot BattleWeaponEnum = "42"
	HeroShot          BattleWeaponEnum = "45"
	OctoShot          BattleWeaponEnum = "46"
	Point52Gal        BattleWeaponEnum = "50"
	Point52GalDeco    BattleWeaponEnum = "51"
	KensaPoint52Gal   BattleWeaponEnum = "52"
	Nzap85            BattleWeaponEnum = "60"
	Nzap89            BattleWeaponEnum = "61"
	Nzap83            BattleWeaponEnum = "62"
	PRO               BattleWeaponEnum = "70"
	ForgePro          BattleWeaponEnum = "71"
	KensaPro          BattleWeaponEnum = "72"
	Point96Gal        BattleWeaponEnum = "80"
	Point96GalDeco    BattleWeaponEnum = "81"
	JET               BattleWeaponEnum = "90"
	CustomJet         BattleWeaponEnum = "91"
	LUNA              BattleWeaponEnum = "200"
	LunaNeo           BattleWeaponEnum = "201"
	KensaLuna         BattleWeaponEnum = "202"
	BLASTER           BattleWeaponEnum = "210"
	CustomBlaster     BattleWeaponEnum = "211"
	HeroBlaster       BattleWeaponEnum = "215"
	RANGE             BattleWeaponEnum = "220"
	CustomRange       BattleWeaponEnum = "221"
	GrimRange         BattleWeaponEnum = "222"
	CLASH             BattleWeaponEnum = "230"
	ClashNeo          BattleWeaponEnum = "231"
	RAPID             BattleWeaponEnum = "240"
	RapidDeco         BattleWeaponEnum = "241"
	KensaRapid        BattleWeaponEnum = "242"
	RapidPro          BattleWeaponEnum = "250"
	RapidProDeco      BattleWeaponEnum = "251"
	L3                BattleWeaponEnum = "300"
	L3D               BattleWeaponEnum = "301"
	KensaL3           BattleWeaponEnum = "302"
	H3                BattleWeaponEnum = "310"
	H3D               BattleWeaponEnum = "311"
	CherryH3          BattleWeaponEnum = "312"
	SQUEEZER          BattleWeaponEnum = "400"
	FoilSqueezer      BattleWeaponEnum = "401"
	CARBON            BattleWeaponEnum = "1000"
	CarbonDeco        BattleWeaponEnum = "1001"
	ROLLER            BattleWeaponEnum = "1010"
	KrakOnRoller      BattleWeaponEnum = "1011"
	KensaRoller       BattleWeaponEnum = "1012"
	HeroRoller        BattleWeaponEnum = "1015"
	DYNAMO            BattleWeaponEnum = "1020"
	GoldDynamo        BattleWeaponEnum = "1021"
	KensaDynamo       BattleWeaponEnum = "1022"
	FLINGZA           BattleWeaponEnum = "1030"
	FoilFlingza       BattleWeaponEnum = "1031"
	INKBRUSH          BattleWeaponEnum = "1100"
	InkbrushNouveau   BattleWeaponEnum = "1101"
	PermanentInkbrush BattleWeaponEnum = "1102"
	OCTOBRUSH         BattleWeaponEnum = "1110"
	OctobrushNoveau   BattleWeaponEnum = "1111"
	KensaOctobrush    BattleWeaponEnum = "1112"
	HEROBRUSH         BattleWeaponEnum = "1115"
	SQUIFFER          BattleWeaponEnum = "2000"
	NewSquiffer       BattleWeaponEnum = "2001"
	FreshSquiffer     BattleWeaponEnum = "2002"
	CHARGER           BattleWeaponEnum = "2010"
	FirefinCharger    BattleWeaponEnum = "2011"
	KensaCharger      BattleWeaponEnum = "2012"
	HeroCharger       BattleWeaponEnum = "2015"
	SCOPE             BattleWeaponEnum = "2020"
	FirefinScope      BattleWeaponEnum = "2021"
	KensaScope        BattleWeaponEnum = "2022"
	ELITER            BattleWeaponEnum = "2030"
	CustomEliter      BattleWeaponEnum = "2031"
	EliterScope       BattleWeaponEnum = "2040"
	CustomEliterScope BattleWeaponEnum = "2041"
	BAMBOOZLER        BattleWeaponEnum = "2050"
	Bamboozler2       BattleWeaponEnum = "2051"
	Bamboozler3       BattleWeaponEnum = "2052"
	GOO               BattleWeaponEnum = "2060"
	CustomGoo         BattleWeaponEnum = "2061"
	SLOSHER           BattleWeaponEnum = "3000"
	SlosherDeco       BattleWeaponEnum = "3001"
	SodaSlosher       BattleWeaponEnum = "3002"
	HeroSlosher       BattleWeaponEnum = "3005"
	TRI               BattleWeaponEnum = "3010"
	TriNouveau        BattleWeaponEnum = "3011"
	MACHINE           BattleWeaponEnum = "3020"
	MachineNeo        BattleWeaponEnum = "3021"
	KensaMachine      BattleWeaponEnum = "3022"
	BLOB              BattleWeaponEnum = "3030"
	BlobDeco          BattleWeaponEnum = "3031"
	EXPLOSHER         BattleWeaponEnum = "3040"
	CustomExplosher   BattleWeaponEnum = "3041"
	MINI              BattleWeaponEnum = "4000"
	ZinkMini          BattleWeaponEnum = "4001"
	KensaMini         BattleWeaponEnum = "4002"
	HEAVY             BattleWeaponEnum = "4010"
	HeavyDeco         BattleWeaponEnum = "4011"
	HeavyRemix        BattleWeaponEnum = "4012"
	HeroSplatling     BattleWeaponEnum = "4015"
	HYDRA             BattleWeaponEnum = "4020"
	CustomHydra       BattleWeaponEnum = "4021"
	BALLPOINT         BattleWeaponEnum = "4030"
	BallpointNouveau  BattleWeaponEnum = "4031"
	Naut47            BattleWeaponEnum = "4040"
	Naut79            BattleWeaponEnum = "4041"
	DAPPLE            BattleWeaponEnum = "5000"
	DappleNouveau     BattleWeaponEnum = "5001"
	ClearDapple       BattleWeaponEnum = "5002"
	DUALIES           BattleWeaponEnum = "5010"
	EnperryDualies    BattleWeaponEnum = "5011"
	KensaDualies      BattleWeaponEnum = "5012"
	HeroDualies       BattleWeaponEnum = "5015"
	GLOOGA            BattleWeaponEnum = "5020"
	GloogaDeco        BattleWeaponEnum = "5021"
	KensaGlooga       BattleWeaponEnum = "5022"
	SQUELCHERS        BattleWeaponEnum = "5030"
	CustomSquelchers  BattleWeaponEnum = "5031"
	TETRA             BattleWeaponEnum = "5040"
	LightTetra        BattleWeaponEnum = "5041"
	BRELLA            BattleWeaponEnum = "6000"
	SorrellaBrella    BattleWeaponEnum = "6001"
	HeroBrella        BattleWeaponEnum = "6005"
	TENTA             BattleWeaponEnum = "6010"
	TentaSorella      BattleWeaponEnum = "6011"
	TentaCamo         BattleWeaponEnum = "6012"
	Undercover        BattleWeaponEnum = "6020"
	UndercoverSorella BattleWeaponEnum = "6021"
	KensaUndercover   BattleWeaponEnum = "6022"
)

func (bwe *BattleWeaponEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type BWE BattleWeaponEnum
	r := (*BWE)(bwe)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *bwe {
	case SPLOOSH, NeoSploosh, Sploosh7, JR, CustomJr, KensaJr, SPLASH, NeoSplash, AeroMg, AeroRg, AeroPg,
		SPLATTERSHOT, TtekSplattershot, KensaSplattershot, HeroShot, OctoShot, Point52Gal, Point52GalDeco,
		KensaPoint52Gal, Nzap85, Nzap89, Nzap83, PRO, ForgePro, KensaPro, Point96Gal, Point96GalDeco,
		JET, CustomJet, LUNA, LunaNeo, KensaLuna, BLASTER, CustomBlaster, HeroBlaster, RANGE, CustomRange,
		GrimRange, CLASH, ClashNeo, RAPID, RapidDeco, KensaRapid, RapidPro, RapidProDeco, L3, L3D, KensaL3,
		H3, H3D, CherryH3, SQUEEZER, FoilSqueezer, CARBON, CarbonDeco, ROLLER, KrakOnRoller, KensaRoller,
		HeroRoller, DYNAMO, GoldDynamo, KensaDynamo, FLINGZA, FoilFlingza, INKBRUSH, InkbrushNouveau,
		PermanentInkbrush, OCTOBRUSH, OctobrushNoveau, KensaOctobrush, HEROBRUSH, SQUIFFER, NewSquiffer,
		FreshSquiffer, CHARGER, FirefinCharger, KensaCharger, HeroCharger, SCOPE, FirefinScope, KensaScope,
		ELITER, CustomEliter, EliterScope, CustomEliterScope, BAMBOOZLER, Bamboozler2, Bamboozler3, GOO,
		CustomGoo, SLOSHER, SlosherDeco, SodaSlosher, HeroSlosher, TRI, TriNouveau, MACHINE, MachineNeo,
		KensaMachine, BLOB, BlobDeco, EXPLOSHER, CustomExplosher, MINI, ZinkMini, KensaMini, HEAVY, HeavyDeco,
		HeavyRemix, HeroSplatling, HYDRA, CustomHydra, BALLPOINT, BallpointNouveau, Naut47, Naut79, DAPPLE,
		DappleNouveau, ClearDapple, DUALIES, EnperryDualies, KensaDualies, HeroDualies, GLOOGA, GloogaDeco,
		KensaGlooga, SQUELCHERS, CustomSquelchers, TETRA, LightTetra, BRELLA, SorrellaBrella, HeroBrella, TENTA,
		TentaSorella, TentaCamo, Undercover, UndercoverSorella, KensaUndercover:
		return nil
	}
	return errors.New("Invalid BattleWeaponEnum. Got: " + fmt.Sprint(*bwe))
}

type BattleStatinkWeaponEnum string

const (
	statInkSploosh           BattleStatinkWeaponEnum = "bold"
	statInkNeoSploosh        BattleStatinkWeaponEnum = "bold_neo"
	statInkSploosh7          BattleStatinkWeaponEnum = "bold_7"
	statInkJr                BattleStatinkWeaponEnum = "wakaba"
	statInkCustomJr          BattleStatinkWeaponEnum = "momiji"
	statInkKensaJr           BattleStatinkWeaponEnum = "ochiba"
	statInkSplash            BattleStatinkWeaponEnum = "sharp"
	statInkNeoSplash         BattleStatinkWeaponEnum = "sharp_neo"
	statInkAeroMG            BattleStatinkWeaponEnum = "promodeler_mg"
	statInkAeroRG            BattleStatinkWeaponEnum = "promodeler_rg"
	statInkAeroPG            BattleStatinkWeaponEnum = "promodeler_pg"
	statInkSplattershot      BattleStatinkWeaponEnum = "sshooter"
	statInkTtekSplattershot  BattleStatinkWeaponEnum = "sshooter_collabo"
	statInkKensaSplattershot BattleStatinkWeaponEnum = "sshooter_becchu"
	statInkHeroShot          BattleStatinkWeaponEnum = "heroshooter_replica"
	statInkOctoShot          BattleStatinkWeaponEnum = "octoshooter_replica"
	statInkPoint52Gal        BattleStatinkWeaponEnum = "52gal"
	statInkPoint52GalDeco    BattleStatinkWeaponEnum = "52gal_deco"
	statInkKensaPoint52Gal   BattleStatinkWeaponEnum = "52gal_becchu"
	statInkNZap85            BattleStatinkWeaponEnum = "nzap85"
	statInkNZap89            BattleStatinkWeaponEnum = "nzap89"
	statInkNZap83            BattleStatinkWeaponEnum = "nzap83"
	statInkPro               BattleStatinkWeaponEnum = "prime"
	statInkForgePro          BattleStatinkWeaponEnum = "prime_collabo"
	statInkKensaPro          BattleStatinkWeaponEnum = "prime_becchu"
	statInkPoint96Gal        BattleStatinkWeaponEnum = "96gal"
	statInkPoint96GalDeco    BattleStatinkWeaponEnum = "96gal_deco"
	statInkJet               BattleStatinkWeaponEnum = "jetsweeper"
	statInkCustomJet         BattleStatinkWeaponEnum = "jetsweeper_custom"
	statInkLuna              BattleStatinkWeaponEnum = "nova"
	statInkLunaNeo           BattleStatinkWeaponEnum = "nova_neo"
	statInkKensaLuna         BattleStatinkWeaponEnum = "nova_becchu"
	statInkBlaster           BattleStatinkWeaponEnum = "hotblaster"
	statInkCustomBlaster     BattleStatinkWeaponEnum = "hotblaster_custom"
	statInkHeroBlaster       BattleStatinkWeaponEnum = "heroblaster_replica"
	statInkRange             BattleStatinkWeaponEnum = "longblaster"
	statInkCustomRange       BattleStatinkWeaponEnum = "longblaster_custom"
	statInkGrimRange         BattleStatinkWeaponEnum = "longblaster_necro"
	statInkClash             BattleStatinkWeaponEnum = "clashblaster"
	statInkClashNeo          BattleStatinkWeaponEnum = "clashblaster_neo"
	statInkRapid             BattleStatinkWeaponEnum = "rapid"
	statInkRapidDeco         BattleStatinkWeaponEnum = "rapid_deco"
	statInkKensaRapid        BattleStatinkWeaponEnum = "rapid_becchu"
	statInkRapidPro          BattleStatinkWeaponEnum = "rapid_elite"
	statInkRapidProDeco      BattleStatinkWeaponEnum = "rapid_elite_deco"
	statInkL3                BattleStatinkWeaponEnum = "l3reelgun"
	statInkL3D               BattleStatinkWeaponEnum = "l3reelgun_d"
	statInkKensaL3           BattleStatinkWeaponEnum = "l3reelgun_becchu"
	statInkH3                BattleStatinkWeaponEnum = "h3reelgun"
	statInkH3D               BattleStatinkWeaponEnum = "h3reelgun_d"
	statInkCherryH3          BattleStatinkWeaponEnum = "h3reelgun_cherry"
	statInkSqueezer          BattleStatinkWeaponEnum = "bottlegeyser"
	statInkFoilSqueezer      BattleStatinkWeaponEnum = "bottlegeyser_foil"
	statInkCarbon            BattleStatinkWeaponEnum = "carbon"
	statInkCarbonDeco        BattleStatinkWeaponEnum = "carbon_deco"
	statInkRoller            BattleStatinkWeaponEnum = "splatroller"
	statInkKrakOnRoller      BattleStatinkWeaponEnum = "splatroller_collabo"
	statInkKensaRoller       BattleStatinkWeaponEnum = "splatroller_becchu"
	statInkHeroRoller        BattleStatinkWeaponEnum = "heroroller_replica"
	statInkDynamo            BattleStatinkWeaponEnum = "dynamo"
	statInkGoldDynamo        BattleStatinkWeaponEnum = "dynamo_tesla"
	statInkKensaDynamo       BattleStatinkWeaponEnum = "dynamo_becchu"
	statInkFlingza           BattleStatinkWeaponEnum = "variableroller"
	statInkFoilFlingza       BattleStatinkWeaponEnum = "variableroller_foil"
	statInkInkbrush          BattleStatinkWeaponEnum = "pablo"
	statInkInkbrushNouveau   BattleStatinkWeaponEnum = "pablo_hue"
	statInkPermanentInkbrush BattleStatinkWeaponEnum = "pablo_permanent"
	statInkOctobrush         BattleStatinkWeaponEnum = "hokusai"
	statInkOctobrushNoveau   BattleStatinkWeaponEnum = "hokusai_hue"
	statInkKensaOctobrush    BattleStatinkWeaponEnum = "hokusai_becchu"
	statInkHeroBrush         BattleStatinkWeaponEnum = "herobrush_replica"
	statInkSquiffer          BattleStatinkWeaponEnum = "squiclean_a"
	statInkNewSquiffer       BattleStatinkWeaponEnum = "squiclean_b"
	statInkFreshSquiffer     BattleStatinkWeaponEnum = "squiclean_g"
	statInkCharger           BattleStatinkWeaponEnum = "splatcharger"
	statInkFirefinCharger    BattleStatinkWeaponEnum = "splatcharger_collabo"
	statInkKensaCharger      BattleStatinkWeaponEnum = "splatcharger_becchu"
	statInkHeroCharger       BattleStatinkWeaponEnum = "herocharger_replica"
	statInkScope             BattleStatinkWeaponEnum = "splatscope"
	statInkFirefinScope      BattleStatinkWeaponEnum = "splatscope_collabo"
	statInkKensaScope        BattleStatinkWeaponEnum = "splatscope_becchu"
	statInkEliter            BattleStatinkWeaponEnum = "liter4k"
	statInkCustomEliter      BattleStatinkWeaponEnum = "liter4k_custom"
	statInkEliterScope       BattleStatinkWeaponEnum = "liter4k_scope"
	statInkCustomEliterScope BattleStatinkWeaponEnum = "liter4k_scope_custom"
	statInkBamboozler        BattleStatinkWeaponEnum = "bamboo14mk1"
	statInkBamboozler2       BattleStatinkWeaponEnum = "bamboo14mk2"
	statInkBamboozler3       BattleStatinkWeaponEnum = "bamboo14mk3"
	statInkGoo               BattleStatinkWeaponEnum = "soytuber"
	statInkCustomGoo         BattleStatinkWeaponEnum = "soytuber_custom"
	statInkSlosher           BattleStatinkWeaponEnum = "bucketslosher"
	statInkSlosherDeco       BattleStatinkWeaponEnum = "bucketslosher_deco"
	statInkSodaSlosher       BattleStatinkWeaponEnum = "bucketslosher_soda"
	statInkHeroSlosher       BattleStatinkWeaponEnum = "heroslosher_replica"
	statInkTri               BattleStatinkWeaponEnum = "hissen"
	statInkTriNouveau        BattleStatinkWeaponEnum = "hissen_hue"
	statInkMachine           BattleStatinkWeaponEnum = "screwslosher"
	statInkMachineNeo        BattleStatinkWeaponEnum = "screwslosher_neo"
	statInkKensaMachine      BattleStatinkWeaponEnum = "screwslosher_becchu"
	statInkBlob              BattleStatinkWeaponEnum = "furo"
	statInkBlobDeco          BattleStatinkWeaponEnum = "furo_deco"
	statInkExplosher         BattleStatinkWeaponEnum = "explosher"
	statInkCustomExplosher   BattleStatinkWeaponEnum = "explosher_custom"
	statInkMini              BattleStatinkWeaponEnum = "splatspinner"
	statInkZinkMini          BattleStatinkWeaponEnum = "splatspinner_collabo"
	statInkKensaMini         BattleStatinkWeaponEnum = "splatspinner_becchu"
	statInkHeavy             BattleStatinkWeaponEnum = "barrelspinner"
	statInkHeavyDeco         BattleStatinkWeaponEnum = "barrelspinner_deco"
	statInkHeavyRemix        BattleStatinkWeaponEnum = "barrelspinner_remix"
	statInkHeroSplatling     BattleStatinkWeaponEnum = "herospinner_replica"
	statInkHydra             BattleStatinkWeaponEnum = "hydra"
	statInkCustomHydra       BattleStatinkWeaponEnum = "hydra_custom"
	statInkBallpoint         BattleStatinkWeaponEnum = "kugelschreiber"
	statInkBallpointNouveau  BattleStatinkWeaponEnum = "kugelschreiber_hue"
	statInkNaut47            BattleStatinkWeaponEnum = "nautilus47"
	statInkNaut79            BattleStatinkWeaponEnum = "nautilus79"
	statInkDapple            BattleStatinkWeaponEnum = "sputtery"
	statInkDappleNouveau     BattleStatinkWeaponEnum = "sputtery_hue"
	statInkClearDapple       BattleStatinkWeaponEnum = "sputtery_clear"
	statInkDualies           BattleStatinkWeaponEnum = "maneuver"
	statInkEnperryDualies    BattleStatinkWeaponEnum = "maneuver_collabo"
	statInkKensaDualies      BattleStatinkWeaponEnum = "maneuver_becchu"
	statInkHeroDualies       BattleStatinkWeaponEnum = "heromaneuver_replica"
	statInkGlooga            BattleStatinkWeaponEnum = "kelvin525"
	statInkGloogaDeco        BattleStatinkWeaponEnum = "kelvin525_deco"
	statInkKensaGlooga       BattleStatinkWeaponEnum = "kelvin525_becchu"
	statInkSquelchers        BattleStatinkWeaponEnum = "dualsweeper"
	statInkCustomSquelchers  BattleStatinkWeaponEnum = "dualsweeper_custom"
	statInkTetra             BattleStatinkWeaponEnum = "quadhopper_black"
	statInkLightTetra        BattleStatinkWeaponEnum = "quadhopper_white"
	statInkBrella            BattleStatinkWeaponEnum = "parashelter"
	statInkSorrellaBrella    BattleStatinkWeaponEnum = "parashelter_sorella"
	statInkHeroBrella        BattleStatinkWeaponEnum = "heroshelter_replica"
	statInkTenta             BattleStatinkWeaponEnum = "campingshelter"
	statInkTentaSorella      BattleStatinkWeaponEnum = "campingshelter_sorella"
	statInkTentaCamo         BattleStatinkWeaponEnum = "campingshelter_camo"
	statInkUndercover        BattleStatinkWeaponEnum = "spygadget"
	statInkUndercoverSorella BattleStatinkWeaponEnum = "spygadget_sorella"
	statInkKensaUndercover   BattleStatinkWeaponEnum = "spygadget_becchu"
	statInkEmptyWeapon       BattleStatinkWeaponEnum = ""
)

func (bswe *BattleStatinkWeaponEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type BSWE BattleStatinkWeaponEnum
	r := (*BSWE)(bswe)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *bswe {
	case statInkSploosh, statInkNeoSploosh, statInkSploosh7, statInkJr, statInkCustomJr, statInkKensaJr, statInkSplash,
		statInkNeoSplash, statInkAeroMG, statInkAeroRG, statInkAeroPG, statInkSplattershot, statInkTtekSplattershot,
		statInkKensaSplattershot, statInkHeroShot, statInkOctoShot, statInkPoint52Gal, statInkPoint52GalDeco,
		statInkKensaPoint52Gal, statInkNZap85, statInkNZap89, statInkNZap83, statInkPro, statInkForgePro,
		statInkKensaPro, statInkPoint96Gal, statInkPoint96GalDeco, statInkJet, statInkCustomJet, statInkLuna,
		statInkLunaNeo, statInkKensaLuna, statInkBlaster, statInkCustomBlaster, statInkHeroBlaster, statInkRange,
		statInkCustomRange, statInkGrimRange, statInkClash, statInkClashNeo, statInkRapid, statInkRapidDeco,
		statInkKensaRapid, statInkRapidPro, statInkRapidProDeco, statInkL3, statInkL3D, statInkKensaL3, statInkH3,
		statInkH3D, statInkCherryH3, statInkSqueezer, statInkFoilSqueezer, statInkCarbon, statInkCarbonDeco,
		statInkRoller, statInkKrakOnRoller, statInkKensaRoller, statInkHeroRoller, statInkDynamo, statInkGoldDynamo,
		statInkKensaDynamo, statInkFlingza, statInkFoilFlingza, statInkInkbrush, statInkInkbrushNouveau,
		statInkPermanentInkbrush, statInkOctobrush, statInkOctobrushNoveau, statInkKensaOctobrush, statInkHeroBrush,
		statInkSquiffer, statInkNewSquiffer, statInkFreshSquiffer, statInkCharger, statInkFirefinCharger,
		statInkKensaCharger, statInkHeroCharger, statInkScope, statInkFirefinScope, statInkKensaScope, statInkEliter,
		statInkCustomEliter, statInkEliterScope, statInkCustomEliterScope, statInkBamboozler, statInkBamboozler2,
		statInkBamboozler3, statInkGoo, statInkCustomGoo, statInkSlosher, statInkSlosherDeco, statInkSodaSlosher,
		statInkHeroSlosher, statInkTri, statInkTriNouveau, statInkMachine, statInkMachineNeo, statInkKensaMachine,
		statInkBlob, statInkBlobDeco, statInkExplosher, statInkCustomExplosher, statInkMini, statInkZinkMini,
		statInkKensaMini, statInkHeavy, statInkHeavyDeco, statInkHeavyRemix, statInkHeroSplatling, statInkHydra,
		statInkCustomHydra, statInkBallpoint, statInkBallpointNouveau, statInkNaut47, statInkNaut79, statInkDapple,
		statInkDappleNouveau, statInkClearDapple, statInkDualies, statInkEnperryDualies, statInkKensaDualies,
		statInkHeroDualies, statInkGlooga, statInkGloogaDeco, statInkKensaGlooga, statInkSquelchers,
		statInkCustomSquelchers, statInkTetra, statInkLightTetra, statInkBrella, statInkSorrellaBrella,
		statInkHeroBrella, statInkTenta, statInkTentaSorella, statInkTentaCamo, statInkUndercover,
		statInkUndercoverSorella, statInkKensaUndercover, statInkEmptyWeapon:
		return nil
	}
	return errors.New("Invalid BattleStatinkWeaponEnum. Got: " + fmt.Sprint(*bswe))
}
