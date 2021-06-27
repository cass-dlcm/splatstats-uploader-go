package enums

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FailureReasonEnum string

const (
	WIPE_OUT   FailureReasonEnum = "wipe_out"
	TIME_LIMIT FailureReasonEnum = "time_limit"
)

func (fre *FailureReasonEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type FRE FailureReasonEnum
	r := (*FRE)(fre)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *fre {
	case WIPE_OUT, TIME_LIMIT:
		return nil
	}
	return errors.New("Invalid FailureReasonEnum. Got: " + fmt.Sprint(fre))
}

type SalmonStageEnum string

const (
	smokeyard SalmonStageEnum = "Salmonid Smokeyard"
	polaris   SalmonStageEnum = "Ruins of Ark Polaris"
	grounds   SalmonStageEnum = "Spawning Grounds"
	bay       SalmonStageEnum = "Marooner's Bay"
	outpost   SalmonStageEnum = "Lost Outpost"
)

type SalmonSplatnetScheduleStageImageEnum string

const (
	smokeyardSplatnetImg SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/e9f7c7b35e6d46778cd3cbc0d89bd7e1bc3be493.png"
	polarisSplatnetImg   SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/50064ec6e97aac91e70df5fc2cfecf61ad8615fd.png"
	groundsSplatnetImg   SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/65c68c6f0641cc5654434b78a6f10b0ad32ccdee.png"
	baySplatnetImg       SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/e07d73b7d9f0c64e552b34a2e6c29b8564c63388.png"
	outpostSplatnetImg   SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/6d68f5baa75f3a94e5e9bfb89b82e7377e3ecd2c.png"
)

func (sse *SalmonStageEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SSE SalmonStageEnum
	r := (*SSE)(sse)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *sse {
	case smokeyard, polaris, grounds, bay, outpost:
		return nil
	}
	return errors.New("invalid type")
}

func (ssssie *SalmonSplatnetScheduleStageImageEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SSSSIE SalmonSplatnetScheduleStageImageEnum
	r := (*SSSSIE)(ssssie)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *ssssie {
	case smokeyardSplatnetImg, polarisSplatnetImg, groundsSplatnetImg, baySplatnetImg, outpostSplatnetImg:
		return nil
	}
	return errors.New("invalid type")
}

type SalmonWeaponEnum string

// List of SalmonWeaponEnum
const (
	SalmonSplooshOMatic      SalmonWeaponEnum = "Sploosh-o-matic"
	SalmonSplattershotJr     SalmonWeaponEnum = "Splattershot Jr."
	SalmonSplashOMatic       SalmonWeaponEnum = "Splash-o-matic"
	SalmonAerosprayMg        SalmonWeaponEnum = "Aerospray MG"
	SalmonSplattershot       SalmonWeaponEnum = "Splattershot"
	Salmon52gal              SalmonWeaponEnum = ".52 Gal"
	SalmonNZap85             SalmonWeaponEnum = "N-ZAP '85"
	SalmonSplattershotPro    SalmonWeaponEnum = "Splattershot Pro"
	Salmon96gal              SalmonWeaponEnum = ".96 Gal"
	SalmonJetSquelcher       SalmonWeaponEnum = "Jet Squelcher"
	SalmonLunaBlaster        SalmonWeaponEnum = "Luna Blaster"
	SalmonBlaster            SalmonWeaponEnum = "Blaster"
	SalmonRangeBlaster       SalmonWeaponEnum = "Range Blaster"
	SalmonClashBlaster       SalmonWeaponEnum = "Clash Blaster"
	SalmonRapidBlaster       SalmonWeaponEnum = "Rapid Blaster"
	SalmonRapidBlasterPro    SalmonWeaponEnum = "Rapid Blaster Pro"
	SalmonL3Nozzlenose       SalmonWeaponEnum = "L-3 Nozzlenose"
	SalmonH3Nozzlenose       SalmonWeaponEnum = "H-3 Nozzlenose"
	SalmonSqueezer           SalmonWeaponEnum = "Squeezer"
	SalmonCarbonRoller       SalmonWeaponEnum = "Carbon Roller"
	SalmonSplatRoller        SalmonWeaponEnum = "Splat Roller"
	SalmonDynamoRoller       SalmonWeaponEnum = "Dynamo Roller"
	SalmonFlingzaRoller      SalmonWeaponEnum = "Flingza Roller"
	SalmonInkbrush           SalmonWeaponEnum = "Inkbrush"
	SalmonOctobrush          SalmonWeaponEnum = "Octobrush"
	SalmonClassicSquiffer    SalmonWeaponEnum = "Classic Squiffer"
	SalmonSplatCharger       SalmonWeaponEnum = "Splat Charger"
	SalmonSplatterscope      SalmonWeaponEnum = "Splatterscope"
	SalmonELiter4K           SalmonWeaponEnum = "E-liter 4K"
	SalmonELiter4KScope      SalmonWeaponEnum = "E-liter 4K Scope"
	SalmonBamboozler14MkI    SalmonWeaponEnum = "Bamboozler 14 Mk I"
	SalmonGooTuber           SalmonWeaponEnum = "Goo Tuber"
	SalmonSlosher            SalmonWeaponEnum = "Slosher"
	SalmonSodaSlosher        SalmonWeaponEnum = "Soda Slosher"
	SalmonTriSlosher         SalmonWeaponEnum = "Tri-Slosher"
	SalmonSloshingMachine    SalmonWeaponEnum = "Sloshing Machine"
	SalmonBloblobber         SalmonWeaponEnum = "Bloblobber"
	SalmonExplosher          SalmonWeaponEnum = "Explosher"
	SalmonMiniSplatling      SalmonWeaponEnum = "Mini Splatling"
	SalmonHeavySplatling     SalmonWeaponEnum = "Heavy Splatling"
	SalmonHydraSplatling     SalmonWeaponEnum = "Hydra Splatling"
	SalmonBallpointSplatling SalmonWeaponEnum = "Ballpoint Splatling"
	SalmonNautilus47         SalmonWeaponEnum = "Nautilus 47"
	SalmonDappleDualies      SalmonWeaponEnum = "Dapple Dualies"
	SalmonSplatDualies       SalmonWeaponEnum = "Splat Dualies"
	SalmonGloogaDualies      SalmonWeaponEnum = "Glooga Dualies"
	SalmonDualieSquelchers   SalmonWeaponEnum = "Dualie Squelchers"
	SalmonDarkTetraDualies   SalmonWeaponEnum = "Dark Tetra Dualies"
	SalmonSplatBrella        SalmonWeaponEnum = "Splat Brella"
	SalmonTentaBrella        SalmonWeaponEnum = "Tenta Brella"
	SalmonUndercoverBrella   SalmonWeaponEnum = "Undercover Brella"
	SalmonGrizzcoBlaster     SalmonWeaponEnum = "Grizzco Blaster"
	SalmonGrizzcoBrella      SalmonWeaponEnum = "Grizzco Brella"
	SalmonGrizzcoCharger     SalmonWeaponEnum = "Grizzco Charger"
	SalmonGrizzcoSlosher     SalmonWeaponEnum = "Grizzco Slosher"
)

func (swe *SalmonWeaponEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SWE SalmonWeaponEnum
	r := (*SWE)(swe)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *swe {
	case SalmonSplooshOMatic, SalmonSplattershotJr, SalmonSplashOMatic, SalmonAerosprayMg, SalmonSplattershot,
		Salmon52gal, SalmonNZap85, SalmonSplattershotPro, Salmon96gal, SalmonJetSquelcher, SalmonLunaBlaster, SalmonBlaster,
		SalmonRangeBlaster, SalmonClashBlaster, SalmonRapidBlaster, SalmonRapidBlasterPro, SalmonL3Nozzlenose,
		SalmonH3Nozzlenose, SalmonSqueezer, SalmonCarbonRoller, SalmonSplatRoller, SalmonDynamoRoller, SalmonFlingzaRoller,
		SalmonInkbrush, SalmonOctobrush, SalmonClassicSquiffer, SalmonSplatCharger, SalmonSplatterscope, SalmonELiter4K,
		SalmonELiter4KScope, SalmonBamboozler14MkI, SalmonGooTuber, SalmonSlosher, SalmonSodaSlosher, SalmonTriSlosher,
		SalmonSloshingMachine, SalmonBloblobber, SalmonExplosher, SalmonMiniSplatling, SalmonHeavySplatling,
		SalmonHydraSplatling, SalmonBallpointSplatling, SalmonNautilus47, SalmonDappleDualies, SalmonSplatDualies,
		SalmonGloogaDualies, SalmonDualieSquelchers, SalmonDarkTetraDualies, SalmonSplatBrella, SalmonTentaBrella,
		SalmonUndercoverBrella, SalmonGrizzcoBlaster, SalmonGrizzcoBrella, SalmonGrizzcoCharger, SalmonGrizzcoSlosher:
		return nil
	}
	return errors.New("Invalid SalmonWeaponEnum. Got: " + fmt.Sprint(*swe))
}

type SalmonWeaponScheduleEnum string

// List of SalmonWeaponScheduleEnum
const (
	RandomGrizzcoSchedule      SalmonWeaponScheduleEnum = "Random Grizzco"
	RandomSchedule             SalmonWeaponScheduleEnum = "Random"
	SplooshOMaticSchedule      SalmonWeaponScheduleEnum = "Sploosh-o-matic"
	SplattershotJrSchedule     SalmonWeaponScheduleEnum = "Splattershot Jr."
	SplashOMaticSchedule       SalmonWeaponScheduleEnum = "Splash-o-matic"
	AerosprayMgSchedule        SalmonWeaponScheduleEnum = "Aerospray MG"
	SplattershotSchedule       SalmonWeaponScheduleEnum = "Splattershot"
	Point52GalSchedule         SalmonWeaponScheduleEnum = ".52 Gal"
	NZap85Schedule             SalmonWeaponScheduleEnum = "N-ZAP '85"
	SplattershotProSchedule    SalmonWeaponScheduleEnum = "Splattershot Pro"
	Point96GalSchedule         SalmonWeaponScheduleEnum = ".96 Gal"
	JetSquelcherSchedule       SalmonWeaponScheduleEnum = "Jet Squelcher"
	LunaBlasterSchedule        SalmonWeaponScheduleEnum = "Luna Blaster"
	BlasterSchedule            SalmonWeaponScheduleEnum = "Blaster"
	RangeBlasterSchedule       SalmonWeaponScheduleEnum = "Range Blaster"
	ClashBlasterSchedule       SalmonWeaponScheduleEnum = "Clash Blaster"
	RapidBlasterSchedule       SalmonWeaponScheduleEnum = "Rapid Blaster"
	RapidBlasterProSchedule    SalmonWeaponScheduleEnum = "Rapid Blaster Pro"
	L3NozzlenoseSchedule       SalmonWeaponScheduleEnum = "L-3 Nozzlenose"
	H3NozzlenoseSchedule       SalmonWeaponScheduleEnum = "H-3 Nozzlenose"
	SqueezerSchedule           SalmonWeaponScheduleEnum = "Squeezer"
	CarbonRollerSchedule       SalmonWeaponScheduleEnum = "Carbon Roller"
	SplatRollerSchedule        SalmonWeaponScheduleEnum = "Splat Roller"
	DynamoRollerSchedule       SalmonWeaponScheduleEnum = "Dynamo Roller"
	FlingzaRollerSchedule      SalmonWeaponScheduleEnum = "Flingza Roller"
	InkbrushSchedule           SalmonWeaponScheduleEnum = "Inkbrush"
	OctobrushSchedule          SalmonWeaponScheduleEnum = "Octobrush"
	ClassicSquifferSchedule    SalmonWeaponScheduleEnum = "Classic Squiffer"
	SplatChargerSchedule       SalmonWeaponScheduleEnum = "Splat Charger"
	SplatterscopeSchedule      SalmonWeaponScheduleEnum = "Splatterscope"
	ELiter4KSchedule           SalmonWeaponScheduleEnum = "E-liter 4K"
	ELiter4KScopeSchedule      SalmonWeaponScheduleEnum = "E-liter 4K Scope"
	Bamboozler14MkISchedule    SalmonWeaponScheduleEnum = "Bamboozler 14 Mk I"
	GooTuberSchedule           SalmonWeaponScheduleEnum = "Goo Tuber"
	SlosherSchedule            SalmonWeaponScheduleEnum = "Slosher"
	SodaSlosherSchedule        SalmonWeaponScheduleEnum = "Soda Slosher"
	TriSlosherSchedule         SalmonWeaponScheduleEnum = "Tri-Slosher"
	SloshingMachineSchedule    SalmonWeaponScheduleEnum = "Sloshing Machine"
	BloblobberSchedule         SalmonWeaponScheduleEnum = "Bloblobber"
	ExplosherSchedule          SalmonWeaponScheduleEnum = "Explosher"
	MiniSplatlingSchedule      SalmonWeaponScheduleEnum = "Mini Splatling"
	HeavySplatlingSchedule     SalmonWeaponScheduleEnum = "Heavy Splatling"
	HydraSplatlingSchedule     SalmonWeaponScheduleEnum = "Hydra Splatling"
	BallpointSplatlingSchedule SalmonWeaponScheduleEnum = "Ballpoint Splatling"
	Nautilus47Schedule         SalmonWeaponScheduleEnum = "Nautilus 47"
	DappleDualiesSchedule      SalmonWeaponScheduleEnum = "Dapple Dualies"
	SplatDualiesSchedule       SalmonWeaponScheduleEnum = "Splat Dualies"
	GloogaDualiesSchedule      SalmonWeaponScheduleEnum = "Glooga Dualies"
	DualieSquelchersSchedule   SalmonWeaponScheduleEnum = "Dualie Squelchers"
	DarkTetraDualiesSchedule   SalmonWeaponScheduleEnum = "Dark Tetra Dualies"
	SplatBrellaSchedule        SalmonWeaponScheduleEnum = "Splat Brella"
	TentaBrellaSchedule        SalmonWeaponScheduleEnum = "Tenta Brella"
	UndercoverBrellaSchedule   SalmonWeaponScheduleEnum = "Undercover Brella"
)

func (swse *SalmonWeaponScheduleEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SWSE SalmonWeaponScheduleEnum
	r := (*SWSE)(swse)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *swse {
	case RandomGrizzcoSchedule, RandomSchedule, SplooshOMaticSchedule, SplattershotJrSchedule, SplashOMaticSchedule,
		AerosprayMgSchedule, SplattershotSchedule, Point52GalSchedule, NZap85Schedule, SplattershotProSchedule,
		Point96GalSchedule, JetSquelcherSchedule, LunaBlasterSchedule, BlasterSchedule, RangeBlasterSchedule,
		ClashBlasterSchedule, RapidBlasterSchedule, RapidBlasterProSchedule, L3NozzlenoseSchedule, H3NozzlenoseSchedule,
		SqueezerSchedule, CarbonRollerSchedule, SplatRollerSchedule, DynamoRollerSchedule, FlingzaRollerSchedule,
		InkbrushSchedule, OctobrushSchedule, ClassicSquifferSchedule, SplatChargerSchedule, SplatterscopeSchedule,
		ELiter4KSchedule, ELiter4KScopeSchedule, Bamboozler14MkISchedule, GooTuberSchedule, SlosherSchedule,
		SodaSlosherSchedule, TriSlosherSchedule, SloshingMachineSchedule, BloblobberSchedule, ExplosherSchedule,
		MiniSplatlingSchedule, HeavySplatlingSchedule, HydraSplatlingSchedule, BallpointSplatlingSchedule,
		Nautilus47Schedule, DappleDualiesSchedule, SplatDualiesSchedule, GloogaDualiesSchedule,
		DualieSquelchersSchedule, DarkTetraDualiesSchedule, SplatBrellaSchedule, TentaBrellaSchedule,
		UndercoverBrellaSchedule:
		return nil
	}
	return errors.New("Invalid SalmonWeaponScheduleEnum. Got: " + fmt.Sprint(*swse))
}

type SalmonWeaponScheduleSpecialEnum string

const (
	RandomGrizzcoSpecialSchedule SalmonWeaponScheduleSpecialEnum = "Random Grizzco"
	RandomSpecialSchedule        SalmonWeaponScheduleSpecialEnum = "Random"
)

func (swsse *SalmonWeaponScheduleSpecialEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SWSSE SalmonWeaponScheduleSpecialEnum
	r := (*SWSSE)(swsse)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *swsse {
	case RandomGrizzcoSpecialSchedule, RandomSpecialSchedule:
		return nil
	}
	return errors.New("Invalid SalmonWeaponScheduleSpecialEnum. Got: " + fmt.Sprint(*swsse))
}
