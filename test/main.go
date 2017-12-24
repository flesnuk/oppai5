// use gentest.py to generate test_suite.json file and then run this
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

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
	var client http.Client
	mapsBody, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic("failed read json")
	}
	var nfails = 0
	keys := make([]Map, 1000)
	unique := map[Map]bool{}

	json.Unmarshal(mapsBody, &keys)

	for j := 0; j < start; j++ {
		unique[keys[j]] = true
	}

	for i, key := range keys {
		if unique[key] {
			continue
		}
		unique[key] = true
		if i%25 == 0 {
			fmt.Println(i)
		}
		time.Sleep(time.Millisecond * msBetweenReq)
		combo, _ := strconv.Atoi(key.Maxcombo)
		mods, _ := strconv.Atoi(key.EnabledMods)
		miss, _ := strconv.Atoi(key.Countmiss)
		n50, _ := strconv.Atoi(key.Count50)
		n100, _ := strconv.Atoi(key.Count100)
		n300, _ := strconv.Atoi(key.Count300)

		resp, err := client.Get(urlPrefix + key.BeatmapID)
		if err != nil {
			panic("client get url fail")
		}
		bmap := oppai.Parse(resp.Body)
		resp.Body.Close()

		realPP, _ := strconv.ParseFloat(key.Pp, 64)
		margin := realPP * errorMargin
		if realPP < 100 {
			margin *= 3
		} else if realPP < 200 {
			margin *= 2
		} else if realPP < 300 {
			margin *= 1.5
		}

		mypp := oppai.PPInfo(bmap, &oppai.Parameters{
			Combo:  uint16(combo),
			Mods:   uint32(mods),
			N300:   uint16(n300),
			N100:   uint16(n100),
			N50:    uint16(n50),
			Misses: uint16(miss),
		}).PP

		aimPP := mypp.Aim
		spePP := mypp.Speed
		accPP := mypp.Acc
		calcPP := mypp.Total

		marg := math.Abs(calcPP - realPP)
		if marg >= margin {
			var loliAim, loliSpe, loliAcc, loliTot float64
			nfails++
			if oppaiBin {
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

				loliAim, _ = strconv.ParseFloat(loliAimS, 64)
				loliSpe, _ = strconv.ParseFloat(loliSpeS, 64)
				loliAcc, _ = strconv.ParseFloat(loliAccS, 64)
				loliTot, _ = strconv.ParseFloat(loliTotS, 64)

			}
			fmt.Printf("Real PP: %.4f calc PP: %.4f Aim: %.2f Spe: %.2f Acc: %.2f Mods: %-6s BID: %-7s\n",
				realPP, calcPP, aimPP, spePP, accPP, oppai.ModsStr(mods), key.BeatmapID)

			if oppaiBin {
				fmt.Printf("OPPAI-NG %.4f ----------------- Aim: %.2f Spe: %.2f Acc: %.2f\n",
					loliTot, loliAim, loliSpe, loliAcc)
			}
		}

	}
	fmt.Println(float32(nfails) / float32(len(keys)) * 100)
	fmt.Println(nfails)
	fmt.Println(len(keys))
}
