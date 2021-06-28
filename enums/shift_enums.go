package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
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
	SalmonSplooshOMatic      SalmonWeaponEnum = "0"
	SalmonSplattershotJr     SalmonWeaponEnum = "10"
	SalmonSplashOMatic       SalmonWeaponEnum = "20"
	SalmonAerosprayMg        SalmonWeaponEnum = "30"
	SalmonSplattershot       SalmonWeaponEnum = "40"
	Salmon52gal              SalmonWeaponEnum = "50"
	SalmonNZap85             SalmonWeaponEnum = "60"
	SalmonSplattershotPro    SalmonWeaponEnum = "70"
	Salmon96gal              SalmonWeaponEnum = "80"
	SalmonJetSquelcher       SalmonWeaponEnum = "90"
	SalmonLunaBlaster        SalmonWeaponEnum = "200"
	SalmonBlaster            SalmonWeaponEnum = "210"
	SalmonRangeBlaster       SalmonWeaponEnum = "220"
	SalmonClashBlaster       SalmonWeaponEnum = "230"
	SalmonRapidBlaster       SalmonWeaponEnum = "240"
	SalmonRapidBlasterPro    SalmonWeaponEnum = "250"
	SalmonL3Nozzlenose       SalmonWeaponEnum = "300"
	SalmonH3Nozzlenose       SalmonWeaponEnum = "310"
	SalmonSqueezer           SalmonWeaponEnum = "400"
	SalmonCarbonRoller       SalmonWeaponEnum = "1000"
	SalmonSplatRoller        SalmonWeaponEnum = "1010"
	SalmonDynamoRoller       SalmonWeaponEnum = "1020"
	SalmonFlingzaRoller      SalmonWeaponEnum = "1030"
	SalmonInkbrush           SalmonWeaponEnum = "1100"
	SalmonOctobrush          SalmonWeaponEnum = "1110"
	SalmonClassicSquiffer    SalmonWeaponEnum = "2000"
	SalmonSplatCharger       SalmonWeaponEnum = "2010"
	SalmonSplatterscope      SalmonWeaponEnum = "2020"
	SalmonELiter4K           SalmonWeaponEnum = "2030"
	SalmonELiter4KScope      SalmonWeaponEnum = "2040"
	SalmonBamboozler14MkI    SalmonWeaponEnum = "2050"
	SalmonGooTuber           SalmonWeaponEnum = "2060"
	SalmonSlosher            SalmonWeaponEnum = "3000"
	SalmonTriSlosher         SalmonWeaponEnum = "3010"
	SalmonSloshingMachine    SalmonWeaponEnum = "3020"
	SalmonBloblobber         SalmonWeaponEnum = "3030"
	SalmonExplosher          SalmonWeaponEnum = "3040"
	SalmonMiniSplatling      SalmonWeaponEnum = "4000"
	SalmonHeavySplatling     SalmonWeaponEnum = "4010"
	SalmonHydraSplatling     SalmonWeaponEnum = "4020"
	SalmonBallpointSplatling SalmonWeaponEnum = "4030"
	SalmonNautilus47         SalmonWeaponEnum = "4040"
	SalmonDappleDualies      SalmonWeaponEnum = "5000"
	SalmonSplatDualies       SalmonWeaponEnum = "5010"
	SalmonGloogaDualies      SalmonWeaponEnum = "5020"
	SalmonDualieSquelchers   SalmonWeaponEnum = "5030"
	SalmonDarkTetraDualies   SalmonWeaponEnum = "5040"
	SalmonSplatBrella        SalmonWeaponEnum = "6000"
	SalmonTentaBrella        SalmonWeaponEnum = "6010"
	SalmonUndercoverBrella   SalmonWeaponEnum = "6020"
	SalmonGrizzcoBlaster     SalmonWeaponEnum = "20000"
	SalmonGrizzcoBrella      SalmonWeaponEnum = "20010"
	SalmonGrizzcoCharger     SalmonWeaponEnum = "20020"
	SalmonGrizzcoSlosher     SalmonWeaponEnum = "20030"
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
		SalmonELiter4KScope, SalmonBamboozler14MkI, SalmonGooTuber, SalmonSlosher, SalmonTriSlosher,
		SalmonSloshingMachine, SalmonBloblobber, SalmonExplosher, SalmonMiniSplatling, SalmonHeavySplatling,
		SalmonHydraSplatling, SalmonBallpointSplatling, SalmonNautilus47, SalmonDappleDualies, SalmonSplatDualies,
		SalmonGloogaDualies, SalmonDualieSquelchers, SalmonDarkTetraDualies, SalmonSplatBrella, SalmonTentaBrella,
		SalmonUndercoverBrella, SalmonGrizzcoBlaster, SalmonGrizzcoBrella, SalmonGrizzcoCharger, SalmonGrizzcoSlosher:
		return nil
	}
	return errors.New("Invalid SalmonWeaponEnum. Got: " + fmt.Sprint(*swe))
}

func (swe SalmonWeaponEnum) GetDisplay(printer *message.Printer) string {
	switch swe {
	case SalmonSplooshOMatic:
		return printer.Sprintf("Sploosh-o-matic")
	case SalmonSplattershotJr:
		return printer.Sprintf("Splattershot Jr.")
	case SalmonSplashOMatic:
		return printer.Sprintf("Splash-o-matic")
	case SalmonAerosprayMg:
		return printer.Sprintf("Aerospray MG")
	case SalmonSplattershot:
		return printer.Sprintf("Splattershot")
	case Salmon52gal:
		return printer.Sprintf(".52 Gal")
	case SalmonNZap85:
		return printer.Sprintf("N-ZAP '85")
	case SalmonSplattershotPro:
		return printer.Sprintf("Splattershot Pro")
	case Salmon96gal:
		return printer.Sprintf(".96 Gal")
	case SalmonJetSquelcher:
		return printer.Sprintf("Jet Squelcher")
	case SalmonLunaBlaster:
		return printer.Sprintf("Luna Blaster")
	case SalmonBlaster:
		return printer.Sprintf("Blaster")
	case SalmonRangeBlaster:
		return printer.Sprintf("Range Blaster")
	case SalmonClashBlaster:
		return printer.Sprintf("Clash Blaster")
	case SalmonRapidBlaster:
		return printer.Sprintf("Rapid Blaster")
	case SalmonRapidBlasterPro:
		return printer.Sprintf("Rapid Blaster Pro")
	case SalmonL3Nozzlenose:
		return printer.Sprintf("L-3 Nozzlenose")
	case SalmonH3Nozzlenose:
		return printer.Sprintf("H-3 Nozzlenose")
	case SalmonSqueezer:
		return printer.Sprintf("Squeezer")
	case SalmonCarbonRoller:
		return printer.Sprintf("Carbon Roller")
	case SalmonSplatRoller:
		return printer.Sprintf("Splat Roller")
	case SalmonDynamoRoller:
		return printer.Sprintf("Dynamo Roller")
	case SalmonFlingzaRoller:
		return printer.Sprintf("Flingza Roller")
	case SalmonInkbrush:
		return printer.Sprintf("Inkbrush")
	case SalmonOctobrush:
		return printer.Sprintf("Octobrush")
	case SalmonClassicSquiffer:
		return printer.Sprintf("Classic Squiffer")
	case SalmonSplatCharger:
		return printer.Sprintf("Splat Charger")
	case SalmonSplatterscope:
		return printer.Sprintf("Splatterscope")
	case SalmonELiter4K:
		return printer.Sprintf("E-liter 4K")
	case SalmonELiter4KScope:
		return printer.Sprintf("E-liter 4K Scope")
	case SalmonBamboozler14MkI:
		return printer.Sprintf("Bamboozler 14 Mk I")
	case SalmonGooTuber:
		return printer.Sprintf("Goo Tuber")
	case SalmonSlosher:
		return printer.Sprintf("Slosher")
	case SalmonTriSlosher:
		return printer.Sprintf("Tri-Slosher")
	case SalmonSloshingMachine:
		return printer.Sprintf("Sloshing Machine")
	case SalmonBloblobber:
		return printer.Sprintf("Bloblobber")
	case SalmonExplosher:
		return printer.Sprintf("Explosher")
	case SalmonMiniSplatling:
		return printer.Sprintf("Mini Splatling")
	case SalmonHeavySplatling:
		return printer.Sprintf("Heavy Splatling")
	case SalmonHydraSplatling:
		return printer.Sprintf("Hydra Splatling")
	case SalmonBallpointSplatling:
		return printer.Sprintf("Ballpoint Splatling")
	case SalmonNautilus47:
		return printer.Sprintf("Nautilus 47")
	case SalmonDappleDualies:
		return printer.Sprintf("Dapple Dualies")
	case SalmonSplatDualies:
		return printer.Sprintf("Splat Dualies")
	case SalmonGloogaDualies:
		return printer.Sprintf("Glooga Dualies")
	case SalmonDualieSquelchers:
		return printer.Sprintf("Dualie Squelchers")
	case SalmonDarkTetraDualies:
		return printer.Sprintf("Dark Tetra Dualies")
	case SalmonSplatBrella:
		return printer.Sprintf("Splat Brella")
	case SalmonTentaBrella:
		return printer.Sprintf("Tenta Brella")
	case SalmonUndercoverBrella:
		return printer.Sprintf("Undercover Brella")
	case SalmonGrizzcoBlaster:
		return printer.Sprintf("Grizzco Blaster")
	case SalmonGrizzcoBrella:
		return printer.Sprintf("Grizzco Brella")
	case SalmonGrizzcoCharger:
		return printer.Sprintf("Grizzco Charger")
	case SalmonGrizzcoSlosher:
		return printer.Sprintf("Grizzco Slosher")
	}
	return ""
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
