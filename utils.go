package oppai

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// Warnings ...
const Warnings = true

func info(s string) {
	if Warnings {
		fmt.Println(s)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func roundOppai(x float64) int { return int(math.Floor((x) + 0.5)) }

func parseFloat(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic("Failed to parse float")
	}

	return float32(f)
}

func parseDouble(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("Failed to parse double")
	}

	return f
}

func pow(a, b float64) float64 {
	if a == 1.52163 {
		return math.Exp(b * 0.4197821287029574)
	} else if a == 0.97 {
		return math.Exp(b * -0.030459207484708574)
	}
	return math.Exp(b * math.Log(a))
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic("Failed to parse int")
	}

	return int(i)
}

func reverseSortFloat64s(x []float64) {
	sort.Float64s(x)
	n := len(x)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		x[i], x[j] = x[j], x[i]
	}

}

/* ------------------------------------------------------------- */
/* mods utils                                                    */

// bitmask mods
const (
	ModsNOMOD int = 0

	ModsNF int = 1 << 0
	ModsEZ int = 1 << 1
	ModsTD int = 1 << 2
	ModsHD int = 1 << 3
	ModsHR int = 1 << 4
	ModsDT int = 1 << 6
	ModsHT int = 1 << 8
	ModsNC int = 1 << 9
	ModsFL int = 1 << 10
	ModsSO int = 1 << 12

	ModsSpeedChanging = ModsDT | ModsHT | ModsNC
	ModsMapChanging   = ModsHR | ModsEZ | ModsSpeedChanging
)

// ModsStr returns a string representation of the mods, such as HDDT
func ModsStr(mods int) string {
	var s string

	if mods == 0 {
		return "NOMOD"
	}

	if (mods & ModsNF) != 0 {
		s += "NF"
	}
	if (mods & ModsEZ) != 0 {
		s += "EZ"
	}
	if (mods & ModsTD) != 0 {
		s += "TD"
	}
	if (mods & ModsHD) != 0 {
		s += "HD"
	}
	if (mods & ModsHR) != 0 {
		s += "HR"
	}
	if (mods & ModsDT) != 0 {
		s += "DT"
	}
	if (mods & ModsHT) != 0 {
		s += "HT"
	}
	if (mods & ModsNC) != 0 {
		s += "NC"
	}
	if (mods & ModsFL) != 0 {
		s += "FL"
	}
	if (mods & ModsSO) != 0 {
		s += "SO"
	}
	return s
}
