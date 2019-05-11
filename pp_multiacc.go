package oppai

// MultiAccPP returns the PP for every acc step
type MultiAccPP struct {
	P95   float64
	P98   float64
	P99   float64
	P99p5 float64
	P100  float64
}

var accuracies = []float64{
	0.95, 0.98, 0.99, 0.995, 1.0,
}

// PPv2StepWithMods calculates the pp of the map with the mods passed and step acc
func (pp *PPv2) PPv2StepWithMods(aimStars, speedStars float64, b *Map, mods int) (*MultiAccPP, error) {
	var score [5]*PPv2

	for i := 0; i < 5; i++ {
		pp, err := pp.ppv2x(
			aimStars, speedStars, -1, b.NSliders, b.NCircles,
			len(b.Objects), b.Mode, mods, -1, -1, 0, 0, 0,
			b.AR, b.OD, 1, b, accuracies[i],
		)

		if err != nil {
			return nil, err
		}

		score[i] = pp
	}

	return &MultiAccPP{
		P95:   score[0].Total,
		P98:   score[1].Total,
		P99:   score[2].Total,
		P99p5: score[3].Total,
		P100:  score[4].Total,
	}, nil
}
