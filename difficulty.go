package main

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
	StreamSpacing float64 = 110.0
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

func dSpacingWeight(Type int, distance float64) float64 {
	switch Type {
	case DiffAim:
		return math.Pow(distance, 0.99)
	case DiffSpeed:
		if distance > SingleSpacing {
			return 2.5
		} else if distance > StreamSpacing {
			return 1.6 + 0.9*(distance-StreamSpacing)/
				(SingleSpacing-StreamSpacing)
		} else if distance > AlmostDiameter {
			return 1.2 + 0.4*(distance-AlmostDiameter)/
				(StreamSpacing-AlmostDiameter)
		} else if distance > (AlmostDiameter / 2.0) {
			return 0.95 + 0.25*
				(distance-AlmostDiameter/2.0)/
				(AlmostDiameter/2.0)
		}
		return 0.95
	}
	panic("this diff type does not exist")
}

/**
* calculates the strain for one difficulty type and stores it in
* obj. this assumes that normpos is already computed.
* this also sets is_single if type is DIFF_SPEED
 */
func dStrain(Type int, obj HitObject, prev HitObject, speedMul float64) {
	var value float64
	timeElapsed := (obj.Time - prev.Time) / speedMul
	var decay = math.Pow(DecayBase[Type], timeElapsed/1000.0)

	if (obj.Type & (ObjSlider | ObjCircle)) != 0 {
		var distance = obj.Normpos.sub(prev.Normpos).len()

		if Type == DiffSpeed {
			obj.IsSingle = distance > SingleSpacing
		}

		value = dSpacingWeight(Type, distance)
		value *= WeightScaling[Type]
	}
	value /= math.Max(timeElapsed, 50.0)
	obj.Strains[Type] = prev.Strains[Type]*decay + value
}

// DiffCalc difficulty calculator
type DiffCalc struct {
	Total float64 // star rating
	Aim   float64 // aim stars
	Speed float64 // speed stars
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
}

func (d *DiffCalc) calcIndividual(Type int) float64 {

	var strainStep = StrainStep * d.speedMul
	var intervalEnd = strainStep
	var maxStrain float64

	// calculate all strains
	for i := 0; i < len(d.Beatmap.Objects); i++ {
		var obj = d.Beatmap.Objects[i]

		var prev *HitObject
		if i > 0 {
			prev = &d.Beatmap.Objects[i-1]
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

	/* weight the top strains sorted from highest to lowest */
	weight := 1.0
	var difficulty float64

	reverseSortFloat64s(d.strains)

	for _, strain := range d.strains {
		difficulty += strain * weight
		weight *= DecayWeight
	}

	return difficulty
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
	}
	ModsApply(mods, mapstats, ApplyCS)
	d.speedMul = float64(mapstats.speed)

	radius := (PlayfieldWidth / 16.0) *
		(1.0 - 0.7*(float64(mapstats.CS)-5.0)/5.0)

	/* positions are normalized on circle radius so that we can
	calc as if everything was the same circlesize */
	scalingFactor := 52.0 / radius

	if radius < CirclesizeBuffThreshold {
		scalingFactor *= 1.0 +
			math.Min(CirclesizeBuffThreshold-radius, 5.0)/50.0
	}

	normalizedCenter := PlayfieldCenter.mul(scalingFactor)

	/* calculate normalized positions */
	for i := 0; i < len(d.Beatmap.Objects); i++ {
		obj := &d.Beatmap.Objects[i]
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
		}
	}

	/* speed and aim stars */
	d.Speed = d.calcIndividual(DiffSpeed)
	d.Aim = d.calcIndividual(DiffAim)

	d.Speed = math.Sqrt(d.Speed) * StarScalingFactor
	d.Aim = math.Sqrt(d.Aim) * StarScalingFactor

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
