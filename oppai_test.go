package oppai

import (
	"os"
	"testing"
)

func Benchmark(b *testing.B) {

	for i := 0; i < b.N; i++ {
		f, err := os.Open("examples/one/Halozy - Kikoku Doukoku Jigokuraku (Hollow Wings) [Notch Hell].osu")
		if err != nil {
			panic(err)
		}
		bm := Parse(f)
		PPInfo(bm, nil)
	}
}
