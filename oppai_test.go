package oppai

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/k0kubun/pp"
)

func BenchmarkPP(b *testing.B) {
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

func TestPP(t *testing.T) {
	f, err := os.Open("examples/one/Halozy - Kikoku Doukoku Jigokuraku (Hollow Wings) [Notch Hell].osu")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	b, err := Parse(f)
	if err != nil {
		t.Fatal(err)
	}

	pp.Println(PPInfo(b, nil))
}
