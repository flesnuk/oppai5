package oppai

import (
	"math"
)

/* ------------------------------------------------------------- */
/* pp calc                                                       */

/* base pp value for stars, used internally by ppv2 */
func ppBase(stars float64) float64 {
	return pow(5.0*math.Max(1.0, stars/0.0675)-4.0, 3.0) /
		100000.0
}

// PPv2Parameters : parameters to be passed to PPv2.
/* aim_stars, speed_stars, max_combo, nsliders, ncircles, nobjects,
* base_ar, base_od are required.
 */
type PPv2Parameters struct {
	Beatmap                      *Map
	AimStars, SpeedStars         float64
	MaxCombo                     int
	NSliders, NCircles, NObjects int
	BaseAR                       float32 // the base AR (before applying mods).
	BaseOD                       float32 // the base OD (before applying mods).
	Mode                         int     // gamemode
	Mods                         int     // the mods bitmask, same as osu! api
	Combo                        int     // Max combo achieved, if -1 it will default to maxCombo-nmiss
	// if N300 = -1 it will default to nobjects - n100 - n50 - nmiss
	N300, N100, N50, NMiss int
	ScoreVer               int // scorev1 (1) or scorev2 (2).
}

// PPv2 : structure to store ppv2 values
type PPv2 struct {
	Total, Aim, Speed, Acc float64
	ComputedAccuracy       Accuracy
}

func (pp *PPv2) ppv2x(aimStars, speedStars float64,
	maxCombo, nsliders, ncircles, nobjects, mode, mods,
	combo, n300, n100, n50, nmiss int, baseAR, baseOD float32,
	scoreVersion int, beatmap *Map, acc float64) PPv2 {

	if beatmap != nil {
		mode = beatmap.Mode
		baseAR = beatmap.AR
		baseOD = beatmap.OD
		maxCombo = beatmap.maxCombo()
		nsliders = beatmap.NSliders
		ncircles = beatmap.NCircles
		nobjects = len(beatmap.Objects)
	}

	if mode != ModeStd {
		panic("this gamemode is not yet supported")
	}

	if maxCombo <= 0 {
		info("W: max_combo <= 0, changing to 1")
		maxCombo = 1
	}

	if combo < 0 {
		combo = maxCombo - nmiss
	}

	if n300 < 0 {
		n300 = nobjects - n100 - n50 - nmiss
	}

	/* accuracy -------------------------------------------- */
	var realAcc float64
	var accuracy float64

	if acc >= 0.0 {
		realAcc = acc
		accuracy = acc
	} else {
		pp.ComputedAccuracy = Accuracy{
			N300:    n300,
			N100:    n100,
			N50:     n50,
			NMisses: nmiss,
		}

		accuracy = pp.ComputedAccuracy.Value()
		realAcc = accuracy

		switch scoreVersion {
		case 1:
			/* scorev1 ignores sliders since they are free 300s
			and for some reason also ignores spinners */
			nspinners := nobjects - nsliders - ncircles

			realAcc = (&Accuracy{
				N300:    n300 - nsliders - nspinners,
				N100:    n100,
				N50:     n50,
				NMisses: nmiss,
			}).Value()

			realAcc = math.Max(0.0, realAcc)
		case 2:
			ncircles = nobjects
		default:
			panic("unsupported score")
		}
	}

	/* global values --------------------------------------- */
	nobjectsOver2k := float64(nobjects) / 2000.0
	lengthBonus := 0.95 + 0.4*math.Min(1.0, nobjectsOver2k)

	if nobjects > 2000 {
		lengthBonus += math.Log10(nobjectsOver2k) * 0.5
	}

	missPenalty := pow(0.97, float64(nmiss))
	comboBreak := pow(float64(combo), 0.8) /
		pow(float64(maxCombo), 0.8)

	/* calculate stats with mods */
	mapstats := &MapStats{
		AR: baseAR,
		OD: baseOD,
	}
	ModsApply(mods, mapstats, ApplyAR|ApplyOD)

	/* ar bonus -------------------------------------------- */
	arBonus := 1.0

	if mapstats.AR > 10.33 {
		arBonus += 0.45 * (float64(mapstats.AR) - 10.33)
	} else if mapstats.AR < 8.0 {
		lowArBonus := 0.01 * (8.0 - float64(mapstats.AR))

		if (mods & ModsHD) != 0 {
			lowArBonus *= 2.0
		}

		arBonus += lowArBonus
	}

	/* aim pp ---------------------------------------------- */
	pp.Aim = ppBase(aimStars)
	pp.Aim *= lengthBonus
	pp.Aim *= missPenalty
	pp.Aim *= comboBreak
	pp.Aim *= arBonus

	if (mods & ModsHD) != 0 {
		/* 1.04f bonus for AR10, 1.06f for AR9, 1.02f for AR11 */
		pp.Aim *= 1.02 + (11.0 - mapstats.AR) / 50.0
	}

	if (mods & ModsFL) != 0 {
		pp.Aim *= 1.45 * lengthBonus
	}
	
	/* acc bonus (bad aim can lead to bad acc, reused in speed) */
	accBonus := 0.5 + accuracy/2.0
	
	/* od bonus (high od requires better aim timing to acc, reuse in spd) */
	odBonus := float64(0.98 + (mapstats.OD*mapstats.OD)/2500.0)

	pp.Aim *= accBonus
	pp.Aim *= odBonus

	/* speed pp -------------------------------------------- */
	pp.Speed = ppBase(speedStars)
	pp.Speed *= lengthBonus
	pp.Speed *= missPenalty
	pp.Speed *= comboBreak
	pp.Speed *= accBonus
	pp.Speed *= odBonus
	
	if (mods & ModsHD) != 0 {
		pp.Speed *= 1.18
	}

	/* acc pp ---------------------------------------------- */
	pp.Acc = pow(1.52163, float64(mapstats.OD)) *
		pow(realAcc, 24.0) * 2.83

	pp.Acc *= math.Min(1.15, pow(float64(ncircles)/1000.0, 0.3))

	if (mods & ModsHD) != 0 {
		pp.Acc *= 1.02
	}

	if (mods & ModsFL) != 0 {
		pp.Acc *= 1.02
	}
	/* total pp -------------------------------------------- */
	finalMultiplier := 1.12

	if (mods & ModsNF) != 0 {
		finalMultiplier *= 0.90
	}

	if (mods & ModsSO) != 0 {
		finalMultiplier *= 0.95
	}

	pp.Total = pow(
		pow(pp.Aim, 1.1)+pow(pp.Speed, 1.1)+
			pow(pp.Acc, 1.1),
		1.0/1.1) * finalMultiplier

	return *pp
}

// PPv2WithMods calculates the pp of the map with the mods passed and acc passed
func (pp *PPv2) PPv2WithMods(aimStars, speedStars float64, b *Map, mods, n300, n100, n50, nmiss, combo int) {
	pp.ppv2x(aimStars, speedStars, -1, b.NSliders, b.NCircles,
		len(b.Objects), b.Mode, mods, combo, n300, n100, n50, nmiss,
		b.AR, b.OD, 1, b, -1)
}

// PPv2ssWithMods calculates the pp of the map with the mods passed and 100% acc
func (pp *PPv2) PPv2ssWithMods(aimStars, speedStars float64, b *Map, mods int) {
	pp.ppv2x(aimStars, speedStars, -1, b.NSliders, b.NCircles,
		len(b.Objects), b.Mode, mods, -1, -1, 0, 0, 0,
		b.AR, b.OD, 1, b, -1)
}

// PPv2ss calculates the pp of the map with nomods and 100% acc
func (pp *PPv2) PPv2ss(aimStars, speedStars float64, b *Map) {
	pp.ppv2x(aimStars, speedStars, -1, b.NSliders, b.NCircles,
		len(b.Objects), b.Mode, ModsNOMOD, -1, -1, 0, 0, 0,
		b.AR, b.OD, 1, b, -1)
}

// MultiAccPP returns the PP for every acc step
type MultiAccPP struct {
	P95   float64
	P98   float64
	P99   float64
	P99p5 float64
	P100  float64
}

// PPv2StepWithMods calculates the pp of the map with the mods passed and step acc
func (pp *PPv2) PPv2StepWithMods(aimStars, speedStars float64, b *Map, mods int) MultiAccPP {
	return MultiAccPP{
		P95: pp.ppv2x(aimStars, speedStars, -1, b.NSliders, b.NCircles,
			len(b.Objects), b.Mode, mods, -1, -1, 0, 0, 0,
			b.AR, b.OD, 1, b, 0.95).Total,
		P98: pp.ppv2x(aimStars, speedStars, -1, b.NSliders, b.NCircles,
			len(b.Objects), b.Mode, mods, -1, -1, 0, 0, 0,
			b.AR, b.OD, 1, b, 0.98).Total,
		P99: pp.ppv2x(aimStars, speedStars, -1, b.NSliders, b.NCircles,
			len(b.Objects), b.Mode, mods, -1, -1, 0, 0, 0,
			b.AR, b.OD, 1, b, 0.99).Total,
		P99p5: pp.ppv2x(aimStars, speedStars, -1, b.NSliders, b.NCircles,
			len(b.Objects), b.Mode, mods, -1, -1, 0, 0, 0,
			b.AR, b.OD, 1, b, 0.995).Total,
		P100: pp.ppv2x(aimStars, speedStars, -1, b.NSliders, b.NCircles,
			len(b.Objects), b.Mode, mods, -1, -1, 0, 0, 0,
			b.AR, b.OD, 1, b, 1.0).Total,
	}
}
