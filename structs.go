package oppai

// Circle ...
type Circle struct {
	pos Vector2
}

// Slider ...
type Slider struct {
	pos         Vector2
	distance    float64 ///distance travelled by one repetition.
	repetitions int     /// 1 = no repeats
}

// HitObject : circle, slider or spinner of the map
type HitObject struct {
	Time     float64 //start time in milliseconds
	Type     int     //an instance of Circle or Slider or null
	Data     interface{}
	Normpos  Vector2
	Angle 	 float64
	Strains  []float64
	IsSingle bool
	DeltaTime	float64
	DDistance	float64
}

// Timing is used in Map for timing points
type Timing struct {
	Time      float64 // start time in milliseconds
	MsPerBeat float64
	Change    bool // if false, ms_per_beat is -100 * bpm_multiplier
}
