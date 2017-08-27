package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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

func main() {
	flag.Parse()
	var f io.Reader
	var err error
	switch flag.NArg() {
	case 0:
		f = os.Stdin
		check(err)
	case 1:
		f, err = os.Open(flag.Arg(0))
		check(err)
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}

	parser := &Parser{}
	parser.Beatmap = &Map{}
	parser.Map(f)
	beatmap := parser.Beatmap
	mods := ModsHD | ModsDT | ModsHR
	mapstats := ModsApply(mods, &MapStats{}, 0xFF)

	diff := (&DiffCalc{}).calcMapWithMods(*beatmap, mods)

	pp := &PPv2{}
	pp.PPv2ssWithMods(diff.Aim, diff.Speed, beatmap, mods)

	fmt.Printf("%s -  %s (%s - %s) [%s] mapped by %s\n\n",
		beatmap.Artist, beatmap.Title, beatmap.ArtistUnicode,
		beatmap.TitleUnicode, beatmap.Version, beatmap.Creator)
	fmt.Printf("+%s AR%.2f OD%.2f CS%.2f HP%.2f\n", ModsStr(mods),
		mapstats.AR, mapstats.OD, mapstats.CS, mapstats.HP)
	fmt.Printf("%d circles, %d sliders, %d spinners\n",
		beatmap.NCircles, beatmap.NSliders, beatmap.NSpinners)
	fmt.Printf("%dx\n", beatmap.maxCombo())
	fmt.Printf("%.2f stars (%.2f aim, %.2f speed)\n", diff.Total,
		diff.Aim, diff.Speed)
	fmt.Printf("%.0f%%\n", pp.ComputedAccuracy.value0()*100)
	fmt.Printf("%.2f aim pp\n%.2f speed pp\n%.2f acc pp\n\n",
		pp.Aim, pp.Speed, pp.Acc)

	fmt.Printf("%.2f pp\n", pp.Total)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
