package oppai

import (
	"io"
	"fmt"
)

// Parameters for the PPInfo function
type Parameters struct {
	N300   uint16
	N100   uint16
	N50    uint16
	Misses uint16
	Combo  uint16
	Mods   uint32
}

// PP info to return
type PP struct {
	PP     PPv2
	Stats  MapStats
	Diff   DiffCalc
	StepPP MultiAccPP
}

// PPInfo ..
func PPInfo(beatmap *Map, p *Parameters) PP {
	var pp PP
	if p != nil {
		pp = getPP(beatmap, int(p.Mods), int(p.N300), int(p.N100), int(p.N50), int(p.Misses), int(p.Combo))
	} else {
		pp = getPP(beatmap, 0, -1, 0, 0, 0, -1)
	}
	maxcombo := pp.Diff.Beatmap.MaxCombo
	pp.Diff.Beatmap = Map{MaxCombo: maxcombo}
	return pp

}

func getPP(beatmap *Map, mods, n300, n100, n50, nmiss, combo int) PP {

	diff := (&DiffCalc{}).CalcMapWithMods(*beatmap, mods)

	pp := &PPv2{}
	pp.PPv2WithMods(diff.Aim, diff.Speed, beatmap, mods, n300, n100, n50, nmiss, combo)
	diff.Beatmap.MaxCombo = beatmap.MaxCombo
	// fmt.Printf("%s -  %s (%s - %s) [%s] mapped by %s\n\n",
	// 	beatmap.Artist, beatmap.Title, beatmap.ArtistUnicode,
	// 	beatmap.TitleUnicode, beatmap.Version, beatmap.Creator)
	// fmt.Printf("%d circles, %d sliders, %d spinners\n",
	//	beatmap.NCircles, beatmap.NSliders, beatmap.NSpinners)
	// fmt.Printf("%d/%dx\n", combo, beatmap.maxCombo())
	// fmt.Printf("%d spacing singletaps (%.4f%%)\n\n", diff.NSingles, float64(diff.NSingles)/float64(len(beatmap.Objects))*100.0)
	// fmt.Printf("%.2f stars (%.2f aim, %.2f speed)\n", diff.Total,
	//	diff.Aim, diff.Speed)
	// fmt.Printf("%.2f%%\n", pp.ComputedAccuracy.Value()*100)
	// fmt.Printf("%.2f aim pp\n%.2f speed pp\n%.2f acc pp\n\n",
	//	pp.Aim, pp.Speed, pp.Acc)

	fmt.Printf("%.2f pp\n", pp.Total)
	return PP{
		PP:     *pp,
		Stats:  *diff.mapStats,
		Diff:   diff,
		StepPP: (&PPv2{}).PPv2StepWithMods(diff.Aim, diff.Speed, beatmap, mods),
	}
}

// Parse parses the file f and returns a beatmap
func Parse(f io.Reader) *Map {
	parser := &Parser{}
	parser.Beatmap = &Map{}
	parser.Map(f)
	return parser.Beatmap
}

// Parse parses the file f and returns a beatmap with a specified number of objects parsed
func ParsebyNum(f io.Reader, objnum int) *Map {
	parser := &Parser{}
	parser.Beatmap = &Map{}
	parser.MapbyNum(f, objnum)
	return parser.Beatmap
}
