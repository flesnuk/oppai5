package oppai

import (
	"fmt"
	"io"
)

// Parameters for the PPInfo function
type Parameters struct {
	N300   uint16
	N100   uint16
	N50    uint16
	Misses uint16
	Combo  uint16
	Mods   uint32

	// When enabled, will calculate pp for the 4 accuracies
	CalculateMultiAccPP bool
}

// PP info to return
type PP struct {
	*PPv2
	Stats  *MapStats
	Diff   *DiffCalc
	StepPP *MultiAccPP
}

// PPInfo ..
func PPInfo(beatmap *Map, p *Parameters) (pp *PP, err error) {
	if p != nil {
		pp, err = getPP(
			beatmap, int(p.Mods),
			int(p.N300), int(p.N100), int(p.N50),
			int(p.Misses), int(p.Combo),
			p.CalculateMultiAccPP,
		)
	} else {
		pp, err = getPP(
			beatmap, 0, -1, 0, 0, 0, -1,
			false,
		)
	}

	pp.Diff.Beatmap = &Map{
		MaxCombo: pp.Diff.Beatmap.MaxCombo,
	}

	return
}

func getPP(beatmap *Map, mods, n300, n100, n50, nmiss, combo int, multiaccPP bool) (*PP, error) {
	diff := calcMapWithMods(beatmap, mods)
	diff.Beatmap.MaxCombo = beatmap.MaxCombo

	pp := &PP{
		PPv2:  PPv2WithMods(diff.Aim, diff.Speed, beatmap, mods, n300, n100, n50, nmiss, combo),
		Stats: diff.mapStats,
		Diff:  diff,
	}

	if multiaccPP {
		s, err := pp.PPv2StepWithMods(diff.Aim, diff.Speed, beatmap, mods)
		if err != nil {
			return nil, err
		}

		pp.StepPP = s
	}

	return pp, nil
}

func (pp *PP) String() string {
	return fmt.Sprintf(
		"%.2fpp %.2f%% +%s",
		pp.Total,
		pp.ComputedAccuracy.Value()*100,
		ModsStr(pp.Mods),
	)
}

// Parse returns a Beatmap parsed from an io.Reader.
func Parse(f io.Reader) (*Map, error) {
	return parse(f, -1)
}

// ParsebyNum parses the file f and returns a beatmap with a specified number of objects parsed
func ParsebyNum(f io.Reader, objnum int) (*Map, error) {
	return parse(f, objnum)
}
