package oppai

import (
	"github.com/flesnuk/oppai5/vector"
)

// ModeStd is the osu! standard gamemode
const ModeStd = 0

// obj types constants
const (
	ObjCircle  int = 1 << 0
	ObjSlider  int = 1 << 1
	ObjSpinner int = 1 << 3
)

// DiffSpeed : strain index for speed
const DiffSpeed = 0

// DiffAim : strain index for aim
const DiffAim = 1

// Circle ...
type Circle struct {
	pos vector.Vector2
}

// Slider ...
type Slider struct {
	pos         vector.Vector2
	distance    float64 ///distance travelled by one repetition.
	repetitions int     /// 1 = no repeats
}

// HitObject : circle, slider or spinner of the map
type HitObject struct {
	Time      float64 //start time in milliseconds
	Type      int     //an instance of Circle or Slider or null
	Data      interface{}
	Normpos   vector.Vector2
	Angle     float64
	Strains   []float64
	IsSingle  bool
	DeltaTime float64
	DDistance float64
}

// Timing is used in Map for timing points
type Timing struct {
	Time      float64 // start time in milliseconds
	MsPerBeat float64
	Change    bool // if false, ms_per_beat is -100 * bpm_multiplier
}
