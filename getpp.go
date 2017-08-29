package main

import (
	"fmt"
	"io"
)

// ModeStd ...
const ModeStd = 0

// obj types constants
const (
	ObjCircle  int = 1 << 0
	ObjSlider  int = 1 << 1
	ObjSpinner int = 1 << 3
)

// DiffSpeed : strain index for speed
const DiffSpeed = 0

// DiffAim : strain index for aim
const DiffAim = 1

func getpp(f io.Reader, mods, n300, n100, n50, nmiss, combo int) PPv2 {

	parser := &Parser{}
	parser.Beatmap = &Map{}
	parser.Map(f)
	beatmap := parser.Beatmap

	diff := (&DiffCalc{}).calcMapWithMods(*beatmap, mods)

	pp := &PPv2{}
	pp.PPv2WithMods(diff.Aim, diff.Speed, beatmap, mods, n300, n100, n50, nmiss, combo)

	fmt.Printf("%s -  %s (%s - %s) [%s] mapped by %s\n\n",
		beatmap.Artist, beatmap.Title, beatmap.ArtistUnicode,
		beatmap.TitleUnicode, beatmap.Version, beatmap.Creator)
	fmt.Printf("%d circles, %d sliders, %d spinners\n",
		beatmap.NCircles, beatmap.NSliders, beatmap.NSpinners)
	fmt.Printf("%d/%dx\n", combo, beatmap.maxCombo())
	fmt.Printf("%d spacing singletaps (%.4f%%)\n\n", diff.NSingles, float64(diff.NSingles)/float64(len(beatmap.Objects))*100.0)
	fmt.Printf("%.2f stars (%.2f aim, %.2f speed)\n", diff.Total,
		diff.Aim, diff.Speed)
	fmt.Printf("%.2f%%\n", pp.ComputedAccuracy.value0()*100)
	fmt.Printf("%.2f aim pp\n%.2f speed pp\n%.2f acc pp\n\n",
		pp.Aim, pp.Speed, pp.Acc)

	fmt.Printf("%.2f pp\n", pp.Total)
	return *pp
}

