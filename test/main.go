// use gentest.py to generate test_suite.json file and then run
// go run main.go test_suite.go
// the file test_suite.go should be the output from gentest.py
package main

import (
	"fmt"
	"math"
	"os"

	oppai "github.com/flesnuk/oppai5"
)

const (
	jsonFile     = "test_suite.json"
	urlPrefix    = "https://osu.ppy.sh/osu/"
	errorMargin  = 0.02
	start        = 0
	msBetweenReq = 0
	oppaiBin     = false // set to true if you have the oppai binary inside the test folder
)

type Map struct {
	Maxcombo    string
	Pp          string
	Countmiss   string
	BeatmapID   string `json:"beatmap_id"`
	EnabledMods string `json:"enabled_mods"`
	Count300    string
	Count100    string
	Count50     string
}

type Maps struct {
	M []Map
}

func main() {
	var bmap *oppai.Map

	// var combo, mods, miss, n50, n100, n300 int
	var f *os.File
	// var realPP, margin, aimPP, spePP, accPP, calcPP, marg float64
	var margin, marg float64
	var mypp oppai.PPv2
	for _, key := range scores {

		f, _ = os.Open("test_suite/" + key.BeatmapID + ".osu")
		bmap = oppai.Parse(f)
		f.Close()

		margin = key.Pp * errorMargin
		if key.Pp < 100 {
			margin *= 3
		} else if key.Pp < 200 {
			margin *= 2
		} else if key.Pp < 300 {
			margin *= 1.5
		}

		mypp = oppai.PPInfo(bmap, &oppai.Parameters{
			Combo:  key.Maxcombo,
			Mods:   key.EnabledMods,
			N300:   key.Count300,
			N100:   key.Count100,
			N50:    key.Count50,
			Misses: key.Countmiss,
		}).PP

		marg = math.Abs(mypp.Total - key.Pp)
		if marg >= margin {

			aimPP := mypp.Aim
			spePP := mypp.Speed
			accPP := mypp.Acc
			calcPP := mypp.Total
			/*
				cmd := "curl " + urlPrefix + key.BeatmapID + " | ./oppai -" + " +" + oppai.ModsStr(mods) + " " + fmt.Sprintf("%.2f", mypp.ComputedAccuracy.Value()*100) + "% " + key.Maxcombo + "x"
				out, _ := exec.Command("bash", "-c", cmd).Output()
				fmt.Println(cmd)

				lines := strings.Split(string(out), "\n")
				loliAimS := lines[len(lines)-6]
				loliSpeS := lines[len(lines)-5]
				loliAccS := lines[len(lines)-4]
				loliTotS := lines[len(lines)-2]

				loliAimS = loliAimS[:len(loliAimS)-7]
				loliSpeS = loliSpeS[:len(loliSpeS)-9]
				loliAccS = loliAccS[:len(loliAccS)-7]
				loliTotS = loliTotS[:len(loliTotS)-3]

				loliAim, _ := strconv.ParseFloat(loliAimS, 64)
				loliSpe, _ := strconv.ParseFloat(loliSpeS, 64)
				loliAcc, _ := strconv.ParseFloat(loliAccS, 64)
				loliTot, _ := strconv.ParseFloat(loliTotS, 64) */

			fmt.Printf("Real PP: %.4f calc PP: %.4f Aim: %.2f Spe: %.2f Acc: %.2f BID: %-7s\n",
				key.Pp, calcPP, aimPP, spePP, accPP, key.BeatmapID)
			// fmt.Printf("OPPAI-NG %.4f ----------------- Aim: %.2f Spe: %.2f Acc: %.2f\n",
			// 	loliTot, loliAim, loliSpe, loliAcc)
		}

	}
}
