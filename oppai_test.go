package oppai

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Benchmark(b *testing.B) {
	f, err := ioutil.ReadFile("examples/one/Halozy - Kikoku Doukoku Jigokuraku (Hollow Wings) [Notch Hell].osu")
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		var b = bytes.NewReader(f)

		bm, err := Parse(b)
		if err != nil {
			panic(err)
		}

		PPInfo(bm, nil)
	}
}
