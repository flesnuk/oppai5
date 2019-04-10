package main

import (
	"fmt"
	"os"

	"github.com/flesnuk/oppai5"
)

func main() {
	f, err := os.Open("Halozy - Kikoku Doukoku Jigokuraku (Hollow Wings) [Notch Hell].osu")
	if err != nil {
		panic(err)
	}

	beatmap := oppai.Parse(f)

	fmt.Printf("%+v\n", oppai.PPInfo(beatmap, nil).PP)

	fmt.Printf("%+v\n", oppai.PPInfo(beatmap, nil).StepPP)

	fmt.Printf("%+v\n", oppai.PPInfo(beatmap, nil).Stats)
}
