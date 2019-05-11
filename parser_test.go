package oppai

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

func TestParser(t *testing.T) {
	f, err := os.Open("examples/one/Halozy - Kikoku Doukoku Jigokuraku (Hollow Wings) [Notch Hell].osu")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	p, err := Parse(f)
	if err != nil {
		t.Fatal(err)
	}

	deep.FloatPrecision = 3

	if diff := deep.Equal(p, kikokuMap); diff != nil {
		t.Error(diff)
	}
}
