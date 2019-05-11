package oppai

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/go-test/deep"
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

	pp, err := PPInfo(b, nil)
	if err != nil {
		panic(err)
	}

	deep.FloatPrecision = 3

	if diff := deep.Equal(pp.Total, 215.85800529265947); diff != nil {
		t.Error(diff)
	}

	fmt.Printf("Refactor: %.2fpp\n", pp.Total)
	fmt.Printf("Original: %.2fpp\n", 215.85800529265947)
}
