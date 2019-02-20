package oppai

import "math"

/* ------------------------------------------------------------- */
/* difficulty calculator                                         */

// constants for difficulty calculator
const (
	/** almost the normalized circle diameter. */
	AlmostDiameter float64 = 90.0

	/**
	 * arbitrary thresholds to determine when a stream is spaced
	 * enough that it becomes hard to alternate.
	 */
	SingleSpacing float64 = 125.0

	/**
	 * max strains are weighted from highest to lowest, this is how
	 * much the weight decays.
	 */
	DecayWeight float64 = 0.9

	/**
	 * strains are calculated by analyzing the map in chunks and taking
	 * the peak strains in each chunk. this is the length of a strain
	 * interval in milliseconds
	 */
	StrainStep float64 = 400.0

	/** non-normalized diameter where the small circle buff starts. */
	CirclesizeBuffThreshold float64 = 30.0

	/** global stars multiplier. */
	StarScalingFactor float64 = 0.0675

	/** in osu! pixels */
	PlayfieldWidth  float64 = 512.0
	PlayfieldHeight float64 = 384.0

	/**
	 * 50% of the difference between aim and speed is added to total
	 * star rating to compensate for aim/speed only maps
	 */
	ExtremeScalingFactor float64 = 0.5

	MinSpeedBonus float64 = 75.0
	MaxSpeedBonus float64 = 45.0
	AngleBonusScale float64 = 90.0
	AimTimingThreshold float64 = 107.0
	SpeedAngleBonusBegin float64 = 5 * math.Pi / 6
	AimAngleBonusBegin float64 = math.Pi / 3
)

// DecayBase : strain decay per interval.
var DecayBase = []float64{0.3, 0.15}

// WeightScaling : balances speed and aim.
var WeightScaling = []float64{1400.0, 26.25}

// PlayfieldCenter ...
var PlayfieldCenter = Vector2{
	X: PlayfieldWidth / 2.0,
	Y: PlayfieldHeight / 2.0,
}

func dSpacingWeight(Type int, distance float64, delta_time float64,
	prev_distance float64, prev_delta_time float64, angle float64) float64 {
	strain_time := math.Max(delta_time, 50.0)
	prev_strain_time := math.Max(prev_delta_time, 50.0)
	var angle_bonus float64
	switch Type {
	case DiffAim:
		result := 0.0
		if !math.IsNaN(angle) && angle > AimAngleBonusBegin {
			angle_bonus = math.Sqrt(
				math.Max(prev_distance - AngleBonusScale, 0.0) *
				math.Pow(math.Sin(angle - AimAngleBonusBegin), 2.0) *
				math.Max(distance - AngleBonusScale, 0.0))
			result = 1.5 * math.Pow(math.Max(0.0, angle_bonus), 0.99) /
				math.Max(AimTimingThreshold, prev_strain_time)
		}
		weighted_distance := math.Pow(distance, 0.99)
		return math.Max(result +
			weighted_distance /
			math.Max(AimTimingThreshold, strain_time),
			weighted_distance / strain_time)
	case DiffSpeed:
		distance := math.Min(distance, SingleSpacing)
		delta_time := math.Max(delta_time, MaxSpeedBonus)
		speed_bonus := 1.0
		if delta_time < MinSpeedBonus {
			speed_bonus += math.Pow((MinSpeedBonus - delta_time) / 40.0, 2.0)
		}
		angle_bonus := 1.0
		if !math.IsNaN(angle) && angle < SpeedAngleBonusBegin  {
			s := math.Sin(1.5 * (SpeedAngleBonusBegin - angle))
			angle_bonus += math.Pow(s, 2) / 3.57
			if angle < math.Pi / 2.0 {
				angle_bonus = 1.28
				if distance < AngleBonusScale && angle < math.Pi / 4.0 {
					angle_bonus += (1.0 - angle_bonus) *
						math.Min((AngleBonusScale - distance) / 10.0, 1.0)
				}else if distance < AngleBonusScale {
					angle_bonus += (1.0 - angle_bonus) *
						math.Min((AngleBonusScale - distance) / 10.0, 1.0) *
						math.Sin((math.Pi / 2.0 - angle) * 4.0 / math.Pi)
				}
			}
		}
		return (
			(1.0 + (speed_bonus - 1.0) * 0.75) * angle_bonus *
			(0.95 + speed_bonus * math.Pow(distance / SingleSpacing, 3.5))) /
			strain_time
	}
	panic("this diff type does not exist")
}

/**
* calculates the strain for one difficulty type and stores it in
* obj. this assumes that normpos is already computed.
* this also sets is_single if type is DIFF_SPEED
 */
func dStrain(Type int, obj *HitObject, prev HitObject, speedMul float64) {
	var value float64
	timeElapsed := (obj.Time - prev.Time) / speedMul
	var decay = pow(DecayBase[Type], timeElapsed / 1000.0)

	obj.DeltaTime = timeElapsed

	if (obj.Type & (ObjSlider | ObjCircle)) != 0 {
		var distance = obj.Normpos.sub(prev.Normpos).len()
		obj.DDistance = distance

		if Type == DiffSpeed {
			obj.IsSingle = distance > SingleSpacing
		}

		value = dSpacingWeight(Type, obj.DDistance, timeElapsed,
			prev.DDistance, prev.DeltaTime, obj.Angle)
		value *= WeightScaling[Type]
	}
	obj.Strains[Type] = prev.Strains[Type]*decay + value
}

// DiffCalc difficulty calculator
type DiffCalc struct {
	Total float64 // star rating
	Aim   float64 // aim stars

	/** aim difficulty (used to calc length bonus) */
	AimDifficulty float64

	/** aim length bonus (unused at the moment) */
	AimLengthBonus float64

	Speed float64 // speed stars

	/** speed difficulty (used to calc length bonus) */
	SpeedDifficulty float64

	/** speed length bonus (unused at the moment) */
	SpeedLengthBonus float64

	/**
	 * number of notes that are considered singletaps by the
	 * difficulty calculator.
	 */
	NSingles int
	/**
	 * number of taps slower or equal to the singletap threshold
	 * value.
	 */
	NSinglesThreshold int
	Beatmap           Map // the beatmap we want to calculate the difficulty for
	speedMul          float64
	strains           []float64
	mapStats          *MapStats
}

func lengthBonus(stars float64, difficulty float64) float64 {
	return 0.32 + 0.5 *
		(math.Log10(difficulty + stars) - math.Log10(stars))
}

type diffValues struct {
	Difficulty	float64
	Total 		float64
}

func (d *DiffCalc) calcIndividual(Type int) diffValues {
	d.strains = make([]float64, 0, 256)

	var strainStep = StrainStep * d.speedMul
	var intervalEnd = math.Ceil(d.Beatmap.Objects[0].Time / strainStep) * strainStep
	var maxStrain float64

	// calculate all strains
	for i := 0; i < len(d.Beatmap.Objects); i++ {
		var obj = d.Beatmap.Objects[i]

		var prev *HitObject
		if i > 0 {
			prev = d.Beatmap.Objects[i-1]
		}

		if prev != nil {
			dStrain(Type, obj, *prev, d.speedMul)
		}

		for obj.Time > intervalEnd {
			/* add max strain for this interval */
			d.strains = append(d.strains, maxStrain)

			if prev != nil {
				/* decay last object's strains until the next
				   interval and use that as the initial max
				   strain */

				var decay = math.Pow(DecayBase[Type],
					(intervalEnd-prev.Time)/1000.0)
				maxStrain = prev.Strains[Type] * decay
			} else {
				maxStrain = 0.0
			}

			intervalEnd += strainStep
		}

		maxStrain = math.Max(maxStrain, obj.Strains[Type])
	}

	d.strains = append(d.strains, maxStrain)

	/* weight the top strains sorted from highest to lowest */
	weight := 1.0
	var total float64
	var difficulty float64

	reverseSortFloat64s(d.strains)

	for _, strain := range d.strains {
		total += math.Pow(strain, 1.2)
		difficulty += strain * weight
		weight *= DecayWeight
	}

	return diffValues{difficulty, total}
}

// DefaultSingletapThreshold default value for singletap_threshold.
const DefaultSingletapThreshold float64 = 125.0

// Calc calculates beatmap difficulty and stores it in total, aim,
/* speed, nsingles, nsingles_speed fields.
 * @param singletap_threshold the smallest milliseconds interval
 *        that will be considered singletappable. for example,
 *        125ms is 240 1/2 singletaps ((60000 / 240) / 2)
 * @return self
 */
func (d *DiffCalc) Calc(mods int, singletapThreshold float64) DiffCalc {

	mapstats := &MapStats{
		CS: d.Beatmap.CS,
		AR: d.Beatmap.AR,
		HP: d.Beatmap.HP,
		OD: d.Beatmap.OD,
	}
	ModsApply(mods, mapstats, 15)
	d.mapStats = mapstats
	d.speedMul = float64(mapstats.speed)

	radius := (PlayfieldWidth / 16.0) *
		(1.0 - 0.7*(float64(mapstats.CS)-5.0)/5.0)

	/* positions are normalized on circle radius so that we can
	calc as if everything was the same circlesize */
	scalingFactor := 52.0 / radius

	if radius < CirclesizeBuffThreshold {
		scalingFactor *= 1.0 +
			math.Min(CirclesizeBuffThreshold - radius, 5.0) / 50.0
	}

	normalizedCenter := PlayfieldCenter.mul(scalingFactor)

	var prev1 *HitObject
	var prev2 *HitObject

	/* calculate normalized positions */
	for i := 0; i < len(d.Beatmap.Objects); i++ {
		obj := d.Beatmap.Objects[i]
		if (obj.Type & ObjSpinner) != 0 {
			obj.Normpos = Vector2(normalizedCenter)

		} else {
			var pos Vector2

			if (obj.Type & ObjSlider) != 0 {
				pos = obj.Data.(Slider).pos
			} else if (obj.Type & ObjCircle) != 0 {
				pos = obj.Data.(Circle).pos
			} else {
				info("W: unknown object type")
				pos = Vector2{}
			}

			obj.Normpos = Vector2(pos).mul(scalingFactor)

			if i >= 2 {
				v1 := Vector2(prev2.Normpos).sub(prev1.Normpos)
				v2 := Vector2(obj.Normpos).sub(prev1.Normpos)
				dot := v1.dot(v2)
				det := v1.X * v2.Y - v1.Y * v2.X
				obj.Angle = math.Abs(math.Atan2(det, dot))
			}else {
				obj.Angle = math.NaN()
			}

			prev2 = prev1
			prev1 = obj
		}
	}

	/* speed and aim stars */
	aimvals := d.calcIndividual(DiffAim)
	d.Aim = aimvals.Difficulty
	d.AimDifficulty = aimvals.Total
	d.AimLengthBonus = lengthBonus(d.Aim, d.AimDifficulty)

	speedvals := d.calcIndividual(DiffSpeed)
	d.Speed = speedvals.Difficulty
	d.SpeedDifficulty = speedvals.Total
	d.SpeedLengthBonus = lengthBonus(d.Speed, d.SpeedDifficulty)

	d.Aim = math.Sqrt(d.Aim) * StarScalingFactor
	d.Speed = math.Sqrt(d.Speed) * StarScalingFactor

	if (mods & ModsTD) != 0 {
		d.Aim = pow(d.Aim, 0.8)
	}

	/* total stars */
	d.Total = d.Aim + d.Speed +
		math.Abs(d.Speed-d.Aim)*ExtremeScalingFactor

	/* singletap stats */
	for i := 1; i < len(d.Beatmap.Objects); i++ {
		prev := d.Beatmap.Objects[i-1]
		obj := d.Beatmap.Objects[i]

		if obj.IsSingle {
			d.NSingles++
		}

		if (obj.Type & (ObjCircle | ObjSlider)) == 0 {
			continue
		}

		interval := (obj.Time - prev.Time) / d.Speed

		if interval >= singletapThreshold {
			d.NSinglesThreshold++
		}
	}

	return *d
}

func (d *DiffCalc) calcWithMods(mods int) DiffCalc {
	return d.Calc(mods, DefaultSingletapThreshold)
}

func (d *DiffCalc) calc() DiffCalc {
	return d.Calc(ModsNOMOD, DefaultSingletapThreshold)
}

// sets beatmap field and calls
func (d *DiffCalc) calc3(beatmap Map, mods int,
	singletapThreshold float64) DiffCalc {
	d.Beatmap = beatmap
	return d.Calc(mods, singletapThreshold)
}

// sets beatmap field and calls
func (d *DiffCalc) calcMapWithMods(beatmap Map, mods int) DiffCalc {
	return d.calc3(beatmap, mods, DefaultSingletapThreshold)
}

// sets beatmap field and calls
func (d *DiffCalc) calcMap(beatmap Map) DiffCalc {
	return d.calc3(beatmap, ModsNOMOD, DefaultSingletapThreshold)
}
