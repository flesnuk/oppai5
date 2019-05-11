package oppai

import (
	"log"
	"math"
	"strconv"
	"strings"
)

// Warnings changes whether or not to log verbosely
var Warnings = true

func info(s string) {
	if Warnings {
		// Verbose info should go to stderr
		log.Println(s)
	}
}

// https://git.sr.ht/~tdeo/pp/tree/master/parser.go
func keyValue(line string) (key string, value string) {
	split := strings.SplitN(line, ":", 2)
	if len(split) != 2 {
		return
	}

	key = split[0]
	value = strings.TrimLeft(split[1], " ")
	return
}

func parseFloat(s string) float64 {
	// we ignore the error since it'd be a
	// malformed beatmap anyway, so 0 is an okay value
	float, _ := strconv.ParseFloat(s, 64)
	return float
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
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

func round(x float64) int {
	return int(math.Round(x))
}

/*
func pow(a, b float64) float64 {
	if a == 1.52163 {
		return math.Exp(b * 0.4197821287029574)
	} else if a == 0.97 {
		return math.Exp(b * -0.030459207484708574)
	}

	return math.Exp(b * math.Log(a))
}
*/

var pow = math.Pow

// optimized reversing
// https://github.com/golang/go/wiki/SliceTricks
func reverseSortFloat64s(a []float64) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}
