package oppai

import "math"

// Map : the bare minimum beatmap data for difficulty calculation
type Map struct {
	Mode                  int
	Title, TitleUnicode   string
	Artist, ArtistUnicode string

	Creator string // mapper name
	Version string // difficulty name

	NCircles, NSliders, NSpinners int
	HP, CS, OD, AR                float32
	SV, TickRate                  float32
	MaxCombo                      int

	Objects []*HitObject
	TPoints []*Timing
}

func (m *Map) maxCombo() int {
	res := 0
	tindex := -1
	tnext := math.Inf(-1)
	var pxPerBeat float64

	for i := 0; i < len(m.Objects); i++ {
		if (m.Objects[i].Type & ObjSlider) == 0 {
			res++ // non-sliders add 1 combo
			continue
		}

		/* keep track of the current timing point without
		   looping through all of them for every object */
		for m.Objects[i].Time >= tnext {
			tindex++

			if len(m.TPoints) > tindex+1 {
				tnext = m.TPoints[tindex+1].Time
			} else {
				tnext = math.Inf(1)
			}

			t := m.TPoints[tindex]
			var svMultiplier float64
			svMultiplier = 1.0

			if !t.Change && t.MsPerBeat < 0 {
				svMultiplier = -100.0 / t.MsPerBeat
			}

			pxPerBeat = float64(m.SV) * 100.0 * svMultiplier
		}

		/* slider, we need to calculate slider ticks */
		sl := m.Objects[i].Data.(Slider)

		var numBeats float64
		numBeats = (sl.distance * float64(sl.repetitions)) / pxPerBeat

		ticks := int(math.Ceil((numBeats - 0.1) /
			float64(sl.repetitions) * float64(m.TickRate)))

		ticks--
		ticks *= sl.repetitions
		ticks += sl.repetitions + 1

		res += max(0, ticks)
	}

	m.MaxCombo = res
	return res

}
