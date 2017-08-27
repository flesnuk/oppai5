package main

import "math"

/* beatmap stats should be populated with the base beatmap stats and passed to
* mods_apply which will modify the stats for the given mods
 */

// MapStats ...
type MapStats struct {
	AR, OD, CS, HP float32
	speed          float32
}

// stats constants
const (
	OD0MS  float64 = 79.5
	OD10MS float64 = 19.5
	AR0MS  float64 = 1800.0
	AR5MS  float64 = 1200.0
	AR10MS float64 = 450.0

	ODMsStep  float64 = (OD0MS - OD10MS) / 10.0
	ARMsStep1 float64 = (AR0MS - AR5MS) / 5.0
	ARMsStep2 float64 = (AR5MS - AR10MS) / 5.0

	ApplyAR int = 1 << 0
	ApplyOD int = 1 << 1
	ApplyCS int = 1 << 2
	ApplyHP int = 1 << 3
)

// ModsApply applies mods to MapStats
func ModsApply(mods int, mapstats *MapStats, flags int) *MapStats {
	mapstats.speed = 1.0

	if (mods & ModsMapChanging) == 0 {
		return mapstats
	}

	if (mods & (ModsDT | ModsNC)) != 0 {
		mapstats.speed = 1.5
	}

	if (mods & ModsHT) != 0 {
		mapstats.speed *= 0.75
	}

	var OdArHpMultiplier float32 = 1.0

	if (mods & ModsHR) != 0 {
		OdArHpMultiplier = 1.4
	}

	if (mods & ModsEZ) != 0 {
		OdArHpMultiplier *= 0.5
	}

	if (flags & ApplyAR) != 0 {
		mapstats.AR *= OdArHpMultiplier

		// convert AR into milliseconds window
		var arms float64
		if mapstats.AR < 5.0 {
			arms = AR0MS - ARMsStep1*float64(mapstats.AR)
		} else {
			arms = AR5MS - ARMsStep2*(float64(mapstats.AR)-5.0)
		}

		/* stats must be capped to 0-10 before HT/DT which brings
		them to a range of -4.42->11.08 for OD and -5->11 for AR */
		arms = math.Min(AR0MS, math.Max(AR10MS, arms))
		arms /= float64(mapstats.speed)

		if arms > AR5MS {
			mapstats.AR = float32((AR0MS - arms) / ARMsStep1)
		} else {
			mapstats.AR = float32(5.0 + (AR5MS-arms)/ARMsStep2)
		}

	}

	if (flags & ApplyOD) != 0 {
		mapstats.OD *= OdArHpMultiplier
		odms := OD0MS - math.Ceil(ODMsStep*float64(mapstats.OD))
		odms = math.Min(OD0MS, math.Max(OD10MS, odms))
		odms /= float64(mapstats.speed)
		mapstats.OD = float32((OD0MS - odms) / ODMsStep)
	}

	if (flags & ApplyCS) != 0 {
		if (mods & ModsHR) != 0 {
			mapstats.CS *= 1.3
		}

		if (mods & ModsEZ) != 0 {
			mapstats.CS *= 0.5
		}

		mapstats.CS = float32(math.Min(10.0, float64(mapstats.CS)))
	}

	if (flags & ApplyHP) != 0 {
		mapstats.HP = float32(
			math.Min(10.0, float64(mapstats.HP*OdArHpMultiplier)))
	}

	return mapstats
}
