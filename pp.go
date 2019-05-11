package oppai

import (
	"errors"
	"math"
)

var (
	// ErrGamemodeUnsupported is returned if the gamemode is not osu! standard
	ErrGamemodeUnsupported = errors.New("this gamemode is not yet supported")

	// ErrInvalidScoreVersion is returned when the score to calculate isn't v1 nor v2
	ErrInvalidScoreVersion = errors.New("unsupported score")
)

/* ------------------------------------------------------------- */
/* pp calc                                                       */

/* base pp value for stars, used internally by ppv2 */
func ppBase(stars float64) float64 {
	return pow(5.0*math.Max(1.0, stars/0.0675)-4.0, 3.0) /
		100000.0
}

// PPv2 : structure to store ppv2 values
type PPv2 struct {
	Total float64
	Aim   float64
	Speed float64
	Acc   float64

	Mods int

	ComputedAccuracy Accuracy
}

// PPv2WithMods calculates the pp of the map with
// the mods passed and acc passed
func PPv2WithMods(
	aimStars, speedStars float64, b *Map, mods,
	n300, n100, n50, nmiss, combo int,
) *PPv2 {

	pp := &PPv2{
		Mods: mods,
	}

	pp.ppv2x(
		aimStars, speedStars, -1, b.NSliders, b.NCircles,
		len(b.Objects), b.Mode, mods, combo, n300, n100, n50, nmiss,
		b.AR, b.OD, 1, b, -1,
	)

	return pp
}

// PPv2ssWithMods calculates the pp of the map with the mods passed and 100% acc
func PPv2ssWithMods(aimStars, speedStars float64, b *Map, mods int) *PPv2 {
	pp := &PPv2{
		Mods: mods,
	}

	pp.ppv2x(
		aimStars, speedStars, -1, b.NSliders, b.NCircles,
		len(b.Objects), b.Mode, mods, -1, -1, 0, 0, 0,
		b.AR, b.OD, 1, b, -1,
	)

	return pp
}

// PPv2ss calculates the pp of the map with nomods and 100% acc
func PPv2ss(aimStars, speedStars float64, b *Map) *PPv2 {
	pp := &PPv2{}
	pp.ppv2x(
		aimStars, speedStars, -1, b.NSliders, b.NCircles,
		len(b.Objects), b.Mode, ModsNOMOD, -1, -1, 0, 0, 0,
		b.AR, b.OD, 1, b, -1,
	)

	return pp
}

func (pp *PPv2) ppv2x(aimStars, speedStars float64,
	maxCombo, nsliders, ncircles, nobjects, mode, mods,
	combo, n300, n100, n50, nmiss int, baseAR, baseOD float64,
	scoreVersion int, beatmap *Map, acc float64) (*PPv2, error) {

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
		return nil, ErrGamemodeUnsupported
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
			return nil, ErrInvalidScoreVersion
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
	ARbonus := 1.0

	if mapstats.AR > 10.33 {
		ARbonus += 0.3 * (float64(mapstats.AR) - 10.33)
	} else if mapstats.AR < 8.0 {
		ARbonus += 0.01 * (8.0 - float64(mapstats.AR))
	}

	/* aim pp ---------------------------------------------- */
	pp.Aim = ppBase(aimStars)
	pp.Aim *= lengthBonus
	pp.Aim *= missPenalty
	pp.Aim *= comboBreak
	pp.Aim *= ARbonus

	HDbonus := 1.0
	if (mods & ModsHD) != 0 {
		HDbonus += 0.04 * (12.0 - float64(mapstats.AR))
	}

	pp.Aim *= HDbonus

	if (mods & ModsFL) != 0 {
		var FLbonus = 1.0 + 0.35*math.Min(1.0, float64(nobjects)/200.0)

		if nobjects > 200 {
			FLbonus += 0.3 * math.Min(1, (float64(nobjects)-200.0)/300.0)
		}

		if nobjects > 500 {
			FLbonus += (float64(nobjects) - 500.0) / 1200.0
		}

		pp.Aim *= FLbonus
	}

	accBonus := 0.5 + accuracy/2.0
	ODsquared := mapstats.OD * mapstats.OD
	odBonus := float64(0.98 + ODsquared/2500.0)

	pp.Aim *= accBonus
	pp.Aim *= odBonus

	/* speed pp -------------------------------------------- */
	pp.Speed = ppBase(speedStars)
	pp.Speed *= lengthBonus
	pp.Speed *= missPenalty
	pp.Speed *= comboBreak
	if mapstats.AR > 10.33 {
		pp.Speed *= ARbonus
	}
	pp.Speed *= HDbonus

	pp.Speed *= 0.02 + accuracy
	pp.Speed *= 0.96 + float64(ODsquared)/1600.0

	/* acc pp ---------------------------------------------- */
	pp.Acc = pow(1.52163, float64(mapstats.OD)) *
		pow(realAcc, 24.0) * 2.83

	pp.Acc *= math.Min(1.15, pow(float64(ncircles)/1000.0, 0.3))

	if (mods & ModsHD) != 0 {
		pp.Acc *= 1.08
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
		pow(pp.Aim, 1.1)+pow(pp.Speed, 1.1)+pow(pp.Acc, 1.1),
		1.0/1.1,
	) * finalMultiplier

	return pp, nil
}
