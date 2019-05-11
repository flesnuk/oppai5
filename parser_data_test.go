package oppai

import "github.com/flesnuk/oppai5/vector"

/*
	Test suite generated at Sat May 11 12:12:02 PDT 2019
	Repository: https://github.com/flesnuk/oppai5
	Commit hash: eaa6cda1f50905c50fb33220586f694535ad8d3b

	Reproducing:

		func TestParser(t *testing.T) {
			f, err := os.Open("examples/one/Halozy - Kikoku Doukoku Jigokuraku (Hollow Wings) [Notch Hell].osu")
			if err != nil {
				t.Fatal(err)
			}

			parser := &Parser{}
			parser.Beatmap = &Map{}
			parser.Map(f)

			pp.BufferFoldThreshold = 5745348957
			pp.ColoringEnabled = false
			clipboard.WriteAll(pp.Sprint(parser.Beatmap))
		}

	Pasted data goes underneath, along with these regexes executed:
		s/Vector2/vector.Vector2
		s/oppai\.//g
*/
var kikokuMap = &Map{
	Mode:          0,
	Title:         "Kikoku Doukoku Jigokuraku",
	TitleUnicode:  "鬼哭慟哭地獄楽",
	Artist:        "Halozy",
	ArtistUnicode: "Halozy",
	Creator:       "Hollow Wings",
	Version:       "Notch Hell",
	NCircles:      471,
	NSliders:      578,
	NSpinners:     2,
	HP:            4.000000,
	CS:            4.000000,
	OD:            8.000000,
	AR:            9.500000,
	SV:            3.100000,
	TickRate:      1.000000,
	MaxCombo:      0,
	Objects: []*HitObject{
		&HitObject{
			Time: 1590.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 219.000000,
					Y: 226.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 1687.000000,
			Type: 12,
			Data: nil,
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 12429.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 438.000000,
					Y: 348.000000,
				},
				distance:    1085.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 13977.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 137.000000,
					Y: 186.000000,
				},
				distance:    116.250000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 14364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 294.000000,
					Y: 272.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 14655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 181.000000,
					Y: 298.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 14848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 136.000000,
					Y: 364.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 15139.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 189.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 15332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 86.000000,
					Y: 249.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 15526.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 252.000000,
					Y: 345.000000,
				},
				distance:    116.250000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 15913.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 444.000000,
					Y: 356.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 16204.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 481.000000,
					Y: 94.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 16397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 490.000000,
					Y: 215.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 16688.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 373.000000,
					Y: 31.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 16881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 356.000000,
					Y: 139.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 17074.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 198.000000,
					Y: 60.000000,
				},
				distance:    116.250000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 17461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 292.000000,
					Y: 226.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 17752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 101.000000,
					Y: 364.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 17945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 157.000000,
					Y: 279.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 18236.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 180.000000,
					Y: 350.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 18429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 174.000000,
					Y: 201.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 18623.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 23.000000,
					Y: 25.000000,
				},
				distance:    116.250000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 19010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 96.000000,
					Y: 202.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 19106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 79.000000,
					Y: 277.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 19301.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 179.000000,
					Y: 349.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 19493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 287.000000,
					Y: 330.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 19687.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 420.000000,
					Y: 256.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 19784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 420.000000,
					Y: 256.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 19978.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 398.000000,
					Y: 356.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 20171.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 377.000000,
					Y: 192.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 20461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 189.000000,
					Y: 244.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 20558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 189.000000,
					Y: 244.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 20752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 177.000000,
					Y: 104.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 20849.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 177.000000,
					Y: 104.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 21042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 94.000000,
					Y: 200.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 21235.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 32.000000,
					Y: 70.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 21333.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 32.000000,
					Y: 70.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 21719.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 181.000000,
					Y: 359.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 22010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 228.000000,
					Y: 170.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 22107.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 228.000000,
					Y: 170.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 22301.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 268.000000,
					Y: 296.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 22397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 268.000000,
					Y: 296.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 22591.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 396.000000,
					Y: 344.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 22784.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 435.000000,
					Y: 227.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 22881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 435.000000,
					Y: 227.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 23268.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 346.000000,
					Y: 133.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 23559.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 20.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 23656.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 20.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 23848.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 264.000000,
					Y: 103.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 23945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 264.000000,
					Y: 103.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 24139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 323.000000,
					Y: 216.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 24332.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 243.000000,
					Y: 332.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 24429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 243.000000,
					Y: 332.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 24816.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 73.000000,
					Y: 241.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 25107.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 152.000000,
					Y: 232.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 25204.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 152.000000,
					Y: 232.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 25397.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 10.000000,
					Y: 142.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 25493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 10.000000,
					Y: 142.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 25688.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 213.000000,
					Y: 125.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 25881.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 409.000000,
					Y: 203.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 25977.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 409.000000,
					Y: 203.000000,
				},
				distance:    174.375007,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 26364.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 388.000000,
					Y: 286.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 26655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 291.000000,
					Y: 139.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 26752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 291.000000,
					Y: 139.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 26945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 100.000000,
					Y: 67.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 27042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 100.000000,
					Y: 67.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 27235.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 246.000000,
					Y: 80.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 27429.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 113.000000,
					Y: 140.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 27526.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 113.000000,
					Y: 140.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 27719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 147.000000,
					Y: 287.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 27913.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 302.000000,
					Y: 353.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 28204.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 388.000000,
					Y: 286.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 28301.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 388.000000,
					Y: 286.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 28494.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 358.000000,
					Y: 179.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 28591.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 358.000000,
					Y: 179.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 28784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 254.000000,
					Y: 2.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 28978.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 112.000000,
					Y: 140.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 29075.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 112.000000,
					Y: 140.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 29268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 268.000000,
					Y: 174.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 29461.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 72.000000,
					Y: 307.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 29752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 209.000000,
					Y: 234.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 29849.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 209.000000,
					Y: 234.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 30042.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 468.000000,
					Y: 142.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 30139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 468.000000,
					Y: 142.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 30332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 268.000000,
					Y: 97.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 30526.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 83.000000,
					Y: 28.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 30623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 83.000000,
					Y: 28.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 30816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 340.000000,
					Y: 62.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 31010.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 275.000000,
					Y: 282.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 31301.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 278.000000,
					Y: 203.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 31397.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 331.000000,
					Y: 144.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 31493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 331.000000,
					Y: 144.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 31591.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 278.000000,
					Y: 203.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 31688.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 278.000000,
					Y: 203.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 31881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 132.000000,
					Y: 253.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 32075.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 476.000000,
					Y: 211.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 32172.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 476.000000,
					Y: 211.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 32365.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 395.000000,
					Y: 345.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 32558.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 132.000000,
					Y: 253.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 32849.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 27.000000,
					Y: 95.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 32946.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 27.000000,
					Y: 95.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 33139.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 181.000000,
					Y: 143.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 33236.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 181.000000,
					Y: 143.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 33429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 37.000000,
					Y: 171.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 33623.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 343.000000,
					Y: 12.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 33720.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 343.000000,
					Y: 12.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 33913.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 332.000000,
					Y: 165.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 34107.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 471.000000,
					Y: 351.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 34398.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 248.000000,
					Y: 322.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 34493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 248.000000,
					Y: 322.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 34688.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 332.000000,
					Y: 165.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 34785.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 332.000000,
					Y: 165.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 34978.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 429.000000,
					Y: 270.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 35172.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 490.000000,
					Y: 46.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 35268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 490.000000,
					Y: 46.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 35462.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 449.000000,
					Y: 122.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 35655.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 399.000000,
					Y: 378.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 35946.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 248.000000,
					Y: 322.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 36043.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 248.000000,
					Y: 322.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 36236.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 318.000000,
					Y: 279.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 36333.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 318.000000,
					Y: 279.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 36526.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 134.000000,
					Y: 266.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 36720.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 271.000000,
					Y: 115.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 36817.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 271.000000,
					Y: 115.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 37010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 167.000000,
					Y: 343.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 37204.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 72.000000,
					Y: 118.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 37493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 135.000000,
					Y: 163.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 37590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 194.000000,
					Y: 113.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 37687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 194.000000,
					Y: 113.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 37784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 135.000000,
					Y: 163.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 37881.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 135.000000,
					Y: 163.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 38074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 234.000000,
					Y: 43.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 38364.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 376.000000,
					Y: 75.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 38461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 376.000000,
					Y: 75.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 38558.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 414.000000,
					Y: 145.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 38655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 414.000000,
					Y: 145.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 38752.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 477.000000,
					Y: 191.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 39139.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 167.000000,
					Y: 343.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 39526.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 150.000000,
					Y: 264.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 39719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 300.000000,
					Y: 265.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 40300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 89.000000,
					Y: 316.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 40493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 167.000000,
					Y: 343.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 40881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 325.000000,
					Y: 351.000000,
				},
				distance:    155.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 41848.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 300.000000,
					Y: 265.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 42235.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 300.000000,
					Y: 265.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 42622.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 100.000000,
					Y: 205.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 43010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 223.000000,
					Y: 251.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 43202.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 148.000000,
					Y: 275.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 43397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 113.000000,
					Y: 124.000000,
				},
				distance:    155.000000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 44752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 218.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 44848.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 218.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 44945.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 218.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 45526.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 232.000000,
					Y: 100.000000,
				},
				distance:    387.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 46881.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 218.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 47268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 314.000000,
					Y: 229.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 47461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 170.000000,
					Y: 264.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 47655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 218.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 48042.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 356.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 48623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 233.000000,
					Y: 198.000000,
				},
				distance:    310.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 49590.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 318.000000,
					Y: 37.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 49784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 322.000000,
					Y: 176.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 49977.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 322.000000,
					Y: 176.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 50558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 478.000000,
					Y: 290.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 50752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 384.000000,
					Y: 345.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 50945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 356.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 51139.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 154.000000,
					Y: 279.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 51430.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 27.000000,
					Y: 371.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 51527.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 27.000000,
					Y: 371.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 51720.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 256.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 51817.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 256.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 52010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 370.000000,
					Y: 206.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 52204.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 298.000000,
					Y: 95.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 52301.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 298.000000,
					Y: 95.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 52494.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 192.000000,
					Y: 190.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 52688.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 41.000000,
					Y: 204.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 52977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 211.000000,
					Y: 107.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 53074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 211.000000,
					Y: 107.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 53269.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 230.000000,
					Y: 255.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 53364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 230.000000,
					Y: 255.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 53558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 290.000000,
					Y: 17.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 53752.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 446.000000,
					Y: 38.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 53848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 446.000000,
					Y: 38.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 54042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 371.000000,
					Y: 124.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 54236.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 306.000000,
					Y: 349.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 54527.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 438.000000,
					Y: 276.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 54623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 438.000000,
					Y: 276.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 54817.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 372.000000,
					Y: 123.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 54914.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 372.000000,
					Y: 123.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 55107.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 229.000000,
					Y: 254.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 55301.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 152.000000,
					Y: 120.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 55397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 152.000000,
					Y: 120.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 55591.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 77.000000,
					Y: 349.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 55784.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 234.000000,
					Y: 108.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 56074.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 306.000000,
					Y: 81.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 56171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 359.000000,
					Y: 19.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 56268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 359.000000,
					Y: 19.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 56364.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 306.000000,
					Y: 81.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 56461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 306.000000,
					Y: 81.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 56655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 448.000000,
					Y: 137.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 56848.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 151.000000,
					Y: 119.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 56945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 151.000000,
					Y: 119.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 57139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 28.000000,
					Y: 198.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 57332.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 306.000000,
					Y: 163.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 57624.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 415.000000,
					Y: 373.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 57719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 415.000000,
					Y: 373.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 57913.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 320.000000,
					Y: 308.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 58010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 320.000000,
					Y: 308.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 58203.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 208.000000,
					Y: 236.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 58397.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 133.000000,
					Y: 44.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 58493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 133.000000,
					Y: 44.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 58688.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 230.000000,
					Y: 163.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 58881.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 25.000000,
					Y: 356.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 59171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 243.000000,
					Y: 323.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 59268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 243.000000,
					Y: 323.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 59461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 298.000000,
					Y: 120.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 59558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 298.000000,
					Y: 120.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 59752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 359.000000,
					Y: 234.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 59945.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 484.000000,
					Y: 304.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 60042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 484.000000,
					Y: 304.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 60235.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 349.000000,
					Y: 377.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 60430.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 367.000000,
					Y: 157.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 60719.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 217.000000,
					Y: 182.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 60816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 217.000000,
					Y: 182.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 61011.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 99.000000,
					Y: 270.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 61106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 99.000000,
					Y: 270.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 61301.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 156.000000,
					Y: 123.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 61493.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 2.000000,
					Y: 21.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 61590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 2.000000,
					Y: 21.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 61785.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 237.000000,
					Y: 102.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 61977.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 388.000000,
					Y: 229.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 62268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 392.000000,
					Y: 81.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 62364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 458.000000,
					Y: 34.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 62461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 458.000000,
					Y: 34.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 62558.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 392.000000,
					Y: 81.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 62655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 392.000000,
					Y: 81.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 62848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 258.000000,
					Y: 24.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 63042.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 118.000000,
					Y: 58.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 63139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 118.000000,
					Y: 58.000000,
				},
				distance:    232.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 63526.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 237.000000,
					Y: 102.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 63913.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 280.000000,
					Y: 324.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 64300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 397.000000,
					Y: 277.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 64493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 380.000000,
					Y: 193.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 65074.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 317.000000,
					Y: 256.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 65267.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 256.000000,
					Y: 203.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 65655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 148.000000,
					Y: 201.000000,
				},
				distance:    155.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 66622.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 65.000000,
					Y: 188.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 67009.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 317.000000,
					Y: 256.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 67396.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 376.000000,
					Y: 201.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 67784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 244.000000,
					Y: 223.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 67976.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 170.000000,
					Y: 242.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 68171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 403.000000,
					Y: 271.000000,
				},
				distance:    155.000000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 69526.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 353.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 69622.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 353.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 69719.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 353.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 70300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 320.000000,
					Y: 192.000000,
				},
				distance:    387.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 71655.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 420.000000,
					Y: 354.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 72042.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 485.000000,
					Y: 101.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 72235.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 320.000000,
					Y: 192.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 72429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 420.000000,
					Y: 354.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 72816.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 124.000000,
					Y: 318.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 73397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 179.000000,
					Y: 142.000000,
				},
				distance:    310.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 74364.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 38.000000,
					Y: 190.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 74461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 38.000000,
					Y: 190.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 74655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 222.000000,
					Y: 220.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 74752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 222.000000,
					Y: 220.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 74945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 330.000000,
					Y: 356.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 75042.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 330.000000,
					Y: 356.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 75139.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 330.000000,
					Y: 356.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 75332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 350.000000,
					Y: 258.000000,
				},
				distance:    54.249998,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 75526.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 382.000000,
					Y: 141.000000,
				},
				distance:    69.749998,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 75719.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 424.000000,
					Y: 6.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 75913.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 409.000000,
					Y: 316.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 76106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 368.000000,
					Y: 62.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 76300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 493.000000,
					Y: 131.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 76590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 289.000000,
					Y: 33.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 76687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 289.000000,
					Y: 33.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 76881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 86.000000,
					Y: 108.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 77074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 161.000000,
					Y: 137.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 77461.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 135.000000,
					Y: 368.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 77945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 258.000000,
					Y: 212.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 78042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 258.000000,
					Y: 212.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 78623.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 367.000000,
					Y: 266.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 78816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 474.000000,
					Y: 46.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 79010.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 291.000000,
					Y: 284.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 79203.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 383.000000,
					Y: 117.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 79397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 20.000000,
					Y: 169.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 79590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 128.000000,
					Y: 95.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 79784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 263.000000,
					Y: 78.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 79977.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 200.000000,
					Y: 264.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 80364.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 21.000000,
					Y: 348.000000,
				},
				distance:    116.250004,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 80945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 226.000000,
					Y: 355.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 81429.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 503.000000,
					Y: 310.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 81526.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 503.000000,
					Y: 310.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 81719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 226.000000,
					Y: 355.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 82106.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 316.000000,
					Y: 105.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 82300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 418.000000,
					Y: 314.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 82493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 472.000000,
					Y: 158.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 83074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 59.000000,
					Y: 54.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 83461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 432.000000,
					Y: 11.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 83655.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 357.000000,
					Y: 242.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 84235.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 357.000000,
					Y: 242.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 85010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 134.000000,
					Y: 143.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 85203.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 472.000000,
					Y: 225.000000,
				},
				distance:    697.500027,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 86558.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 293.000000,
					Y: 367.000000,
				},
				distance:    813.750031,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 88106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 16.000000,
					Y: 337.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 88300.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 216.000000,
					Y: 367.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 88493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 98.000000,
					Y: 377.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 88687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 315.000000,
					Y: 296.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 88881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 264.000000,
					Y: 204.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 89074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 104.000000,
					Y: 271.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 89848.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 346.000000,
					Y: 60.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 90429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 380.000000,
					Y: 161.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 91203.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 83.000000,
					Y: 182.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 91397.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 265.000000,
					Y: 208.000000,
				},
				distance:    232.500009,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 92752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 496.000000,
					Y: 273.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 92945.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 487.000000,
					Y: 134.000000,
				},
				distance:    813.750031,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 94493.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 28.000000,
					Y: 220.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 94687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 150.000000,
					Y: 261.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 94881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 265.000000,
					Y: 208.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 95268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 468.000000,
					Y: 355.000000,
				},
				distance:    116.250004,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 95848.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 490.000000,
					Y: 158.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 96042.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 267.000000,
					Y: 90.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 96816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 228.000000,
					Y: 335.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 97590.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 325.000000,
					Y: 34.000000,
				},
				distance:    930.000035,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 99332.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 228.000000,
					Y: 335.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 99526.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 386.000000,
					Y: 117.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 99623.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 386.000000,
					Y: 117.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 99816.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 215.000000,
					Y: 257.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 100010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 443.000000,
					Y: 196.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 100106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 443.000000,
					Y: 196.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 100300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 234.000000,
					Y: 177.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 100493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 474.000000,
					Y: 291.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 100687.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 474.000000,
					Y: 291.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 101268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 176.000000,
					Y: 261.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 101461.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 228.000000,
					Y: 335.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 101558.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 228.000000,
					Y: 335.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 101655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 215.000000,
					Y: 257.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 102139.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 55.000000,
					Y: 211.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 102235.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 55.000000,
					Y: 211.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 102816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 133.000000,
					Y: 93.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 103009.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 132.000000,
					Y: 180.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 103106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 132.000000,
					Y: 180.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 103203.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 133.000000,
					Y: 93.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 103687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 291.000000,
					Y: 114.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 103784.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 291.000000,
					Y: 114.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 104364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 343.000000,
					Y: 316.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 104558.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 382.000000,
					Y: 383.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 104655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 382.000000,
					Y: 383.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 104752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 343.000000,
					Y: 316.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 105235.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 424.000000,
					Y: 316.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 105332.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 424.000000,
					Y: 316.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 105913.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 165.000000,
					Y: 139.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 106106.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 210.000000,
					Y: 57.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 106203.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 210.000000,
					Y: 57.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 106300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 165.000000,
					Y: 139.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 106784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 466.000000,
					Y: 7.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 106881.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 466.000000,
					Y: 7.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 107074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 405.000000,
					Y: 135.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 107268.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 336.000000,
					Y: 298.000000,
				},
				distance:    58.125002,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 107654.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 380.000000,
					Y: 365.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 107751.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 380.000000,
					Y: 365.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 107848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 300.000000,
					Y: 369.000000,
				},
				distance:    58.125002,
				repetitions: 4,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 108332.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 135.000000,
					Y: 311.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 108429.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 135.000000,
					Y: 311.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 108623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 291.000000,
					Y: 198.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 108816.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 116.000000,
					Y: 175.000000,
				},
				distance:    58.125002,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 109202.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 77.000000,
					Y: 250.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 109299.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 77.000000,
					Y: 250.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 109396.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 35.000000,
					Y: 177.000000,
				},
				distance:    58.125002,
				repetitions: 4,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 109880.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 63.000000,
					Y: 41.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 109977.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 63.000000,
					Y: 41.000000,
				},
				distance:    38.750000,
				repetitions: 7,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 110364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 189.000000,
					Y: 52.000000,
				},
				distance:    38.750000,
				repetitions: 7,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 110752.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 281.000000,
					Y: 129.000000,
				},
				distance:    38.750000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 110945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 248.000000,
					Y: 219.000000,
				},
				distance:    38.750000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 111139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 270.000000,
					Y: 306.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 111332.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 65.000000,
					Y: 279.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 111429.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 147.000000,
					Y: 287.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 111526.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 98.000000,
					Y: 356.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 112106.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 253.000000,
					Y: 166.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 112300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 267.000000,
					Y: 351.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 112493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 366.000000,
					Y: 194.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 112687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 453.000000,
					Y: 358.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 112881.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 482.000000,
					Y: 127.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 113074.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 126.000000,
					Y: 36.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 114816.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 222.000000,
					Y: 267.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 115010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 29.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 115106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 29.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 115300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 247.000000,
					Y: 96.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 115493.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 205.000000,
					Y: 33.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 115590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 205.000000,
					Y: 33.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 115784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 255.000000,
					Y: 343.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 115977.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 404.000000,
					Y: 62.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 116171.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 112.000000,
					Y: 229.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 117913.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 403.000000,
					Y: 296.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 118106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 121.000000,
					Y: 352.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 118203.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 121.000000,
					Y: 352.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 118397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 272.000000,
					Y: 340.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 118589.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 411.000000,
					Y: 379.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 118686.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 411.000000,
					Y: 379.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 118881.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 121.000000,
					Y: 352.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 118978.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 121.000000,
					Y: 352.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 119171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 294.000000,
					Y: 109.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 119268.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 294.000000,
					Y: 109.000000,
				},
				distance:    1550.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 120429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 235.000000,
					Y: 180.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 120623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 336.000000,
					Y: 370.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 120816.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 439.000000,
					Y: 30.000000,
				},
				distance:    930.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 121493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 371.000000,
					Y: 237.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 121590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 371.000000,
					Y: 237.000000,
				},
				distance:    620.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 122074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 124.000000,
					Y: 175.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 122267.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 240.000000,
					Y: 23.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 122364.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 240.000000,
					Y: 23.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 124106.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 125.000000,
					Y: 176.000000,
				},
				distance:    620.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 124590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 15.000000,
					Y: 250.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 124687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 15.000000,
					Y: 250.000000,
				},
				distance:    155.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 125074.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 15.000000,
					Y: 250.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 125171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 15.000000,
					Y: 250.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 125268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 132.000000,
					Y: 287.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 125364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 132.000000,
					Y: 287.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 125461.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 247.000000,
					Y: 293.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 127203.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 44.000000,
					Y: 43.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 127397.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 351.000000,
					Y: 89.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 127493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 351.000000,
					Y: 89.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 127687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 284.000000,
					Y: 136.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 127880.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 247.000000,
					Y: 293.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 127977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 247.000000,
					Y: 293.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 128171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 44.000000,
					Y: 43.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 128364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 413.000000,
					Y: 19.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 128558.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 504.000000,
					Y: 40.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 130300.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 188.000000,
					Y: 17.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 130494.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 324.000000,
					Y: 319.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 130590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 324.000000,
					Y: 319.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 130784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 148.000000,
					Y: 357.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 130977.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 324.000000,
					Y: 319.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 131074.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 324.000000,
					Y: 319.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 131268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 204.000000,
					Y: 184.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 131364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 204.000000,
					Y: 184.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 131558.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 148.000000,
					Y: 357.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 131655.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 98.000000,
					Y: 271.000000,
				},
				distance:    1550.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 132816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 317.000000,
					Y: 11.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 133010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 371.000000,
					Y: 377.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 133203.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 441.000000,
					Y: 12.000000,
				},
				distance:    930.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 133880.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 412.000000,
					Y: 261.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 133977.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 371.000000,
					Y: 376.000000,
				},
				distance:    620.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 134461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 55.000000,
					Y: 341.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 134654.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 61.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 134752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 30.000000,
					Y: 212.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 136493.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 349.000000,
					Y: 188.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 136687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 61.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 136784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 61.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 136977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 440.000000,
					Y: 10.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 137171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 30.000000,
					Y: 212.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 137268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 30.000000,
					Y: 212.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 137461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 363.000000,
					Y: 82.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 137655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 117.000000,
					Y: 339.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 137752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 117.000000,
					Y: 339.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 137848.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 102.000000,
					Y: 255.000000,
				},
				distance:    1162.500044,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 139010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 349.000000,
					Y: 188.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 139397.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 30.000000,
					Y: 212.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 139881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 245.000000,
					Y: 347.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 140171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 459.000000,
					Y: 39.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 140558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 459.000000,
					Y: 39.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 140945.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 289.000000,
					Y: 281.000000,
				},
				distance:    930.000035,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 142687.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 117.000000,
					Y: 286.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 142881.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 364.000000,
					Y: 314.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 142977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 364.000000,
					Y: 314.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 143171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 195.000000,
					Y: 269.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 143364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 450.000000,
					Y: 324.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 143461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 450.000000,
					Y: 324.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 143655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 254.000000,
					Y: 206.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 144042.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 318.000000,
					Y: 252.000000,
				},
				distance:    697.500027,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 144816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 61.000000,
					Y: 68.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 145300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 318.000000,
					Y: 252.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 145590.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 467.000000,
					Y: 175.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 146074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 175.000000,
					Y: 348.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 146364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 52.000000,
					Y: 365.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 146848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 272.000000,
					Y: 159.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 147139.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 410.000000,
					Y: 237.000000,
				},
				distance:    1627.500062,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 148687.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 294.000000,
					Y: 147.000000,
				},
				distance:    1627.500062,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 150235.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 480.000000,
					Y: 78.000000,
				},
				distance:    2480.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 156429.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 65.000000,
					Y: 169.000000,
				},
				distance:    174.375007,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 156816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 267.000000,
					Y: 160.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 157106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 388.000000,
					Y: 255.000000,
				},
				distance:    116.250004,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 157590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 307.000000,
					Y: 241.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 157977.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 65.000000,
					Y: 168.000000,
				},
				distance:    174.375007,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 158364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 167.000000,
					Y: 355.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 158654.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 203.000000,
					Y: 286.000000,
				},
				distance:    116.250004,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 159138.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 307.000000,
					Y: 242.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 159526.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 496.000000,
					Y: 155.000000,
				},
				distance:    174.375007,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 159913.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 307.000000,
					Y: 242.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 160203.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 129.000000,
					Y: 328.000000,
				},
				distance:    116.250004,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 160687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 129.000000,
					Y: 240.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 161074.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 391.000000,
					Y: 52.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 161171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 391.000000,
					Y: 52.000000,
				},
				distance:    58.125002,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 161461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 259.000000,
					Y: 88.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 161558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 259.000000,
					Y: 88.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 161751.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 381.000000,
					Y: 133.000000,
				},
				distance:    116.250004,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 162235.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 259.000000,
					Y: 88.000000,
				},
				distance:    174.375007,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 162623.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 19.000000,
					Y: 106.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 162914.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 195.000000,
					Y: 174.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 163010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 195.000000,
					Y: 174.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 163203.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 323.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 163301.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 323.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 163494.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 49.000000,
					Y: 299.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 163688.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 18.000000,
					Y: 183.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 163785.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 18.000000,
					Y: 183.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 163978.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 138.000000,
					Y: 26.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 164171.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 330.000000,
					Y: 123.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 164461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 475.000000,
					Y: 59.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 164558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 475.000000,
					Y: 59.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 164753.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 368.000000,
					Y: 195.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 164848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 368.000000,
					Y: 195.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 165042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 502.000000,
					Y: 289.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 165235.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 325.000000,
					Y: 367.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 165332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 325.000000,
					Y: 367.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 165526.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 444.000000,
					Y: 215.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 165719.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 283.000000,
					Y: 302.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 166010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 158.000000,
					Y: 134.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 166106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 158.000000,
					Y: 134.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 166301.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 123.000000,
					Y: 294.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 166397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 123.000000,
					Y: 294.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 166591.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 242.000000,
					Y: 128.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 166785.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 158.000000,
					Y: 373.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 166881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 158.000000,
					Y: 373.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 167075.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 324.000000,
					Y: 365.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 167268.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 123.000000,
					Y: 294.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 167558.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 209.000000,
					Y: 281.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 167655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 242.000000,
					Y: 204.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 167752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 242.000000,
					Y: 204.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 167848.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 209.000000,
					Y: 281.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 167945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 209.000000,
					Y: 281.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 168139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 401.000000,
					Y: 372.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 168332.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 331.000000,
					Y: 333.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 168429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 331.000000,
					Y: 333.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 168623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 123.000000,
					Y: 294.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 168816.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 27.000000,
					Y: 177.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 169106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 123.000000,
					Y: 294.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 169203.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 123.000000,
					Y: 294.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 169397.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 165.000000,
					Y: 191.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 169493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 165.000000,
					Y: 191.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 169687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 391.000000,
					Y: 167.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 169881.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 305.000000,
					Y: 155.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 169977.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 305.000000,
					Y: 155.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 170171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 429.000000,
					Y: 16.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 170364.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 229.000000,
					Y: 31.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 170655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 229.000000,
					Y: 140.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 170752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 229.000000,
					Y: 140.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 170945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 165.000000,
					Y: 191.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 171042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 165.000000,
					Y: 191.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 171235.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 329.000000,
					Y: 232.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 171429.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 469.000000,
					Y: 357.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 171526.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 469.000000,
					Y: 357.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 171719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 385.000000,
					Y: 163.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 171913.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 321.000000,
					Y: 309.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 172203.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 414.000000,
					Y: 233.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 172300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 414.000000,
					Y: 233.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 172493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 498.000000,
					Y: 98.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 172590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 498.000000,
					Y: 98.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 172784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 399.000000,
					Y: 307.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 172977.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 329.000000,
					Y: 232.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 173074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 329.000000,
					Y: 232.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 173268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 243.000000,
					Y: 16.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 173461.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 243.000000,
					Y: 161.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 173752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 252.000000,
					Y: 235.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 173847.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 214.000000,
					Y: 304.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 173944.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 214.000000,
					Y: 304.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 174042.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 252.000000,
					Y: 235.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 174139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 252.000000,
					Y: 235.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 174332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 80.000000,
					Y: 346.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 174526.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 158.000000,
					Y: 148.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 174623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 158.000000,
					Y: 148.000000,
				},
				distance:    232.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 175010.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 82.000000,
					Y: 171.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 175397.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 421.000000,
					Y: 342.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 175784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 422.000000,
					Y: 199.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 175977.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 337.000000,
					Y: 198.000000,
				},
				distance:    77.500000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 176558.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 422.000000,
					Y: 199.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 176752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 460.000000,
					Y: 275.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 177139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 341.000000,
					Y: 340.000000,
				},
				distance:    155.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 178106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 421.000000,
					Y: 342.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 178493.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 134.000000,
					Y: 277.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 178880.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 288.000000,
					Y: 279.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 179268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 134.000000,
					Y: 277.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 179460.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 75.000000,
					Y: 345.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 179655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 152.000000,
					Y: 191.000000,
				},
				distance:    155.000000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 181010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 218.000000,
					Y: 33.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 181106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 218.000000,
					Y: 33.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 181203.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 218.000000,
					Y: 33.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 181784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 305.000000,
					Y: 237.000000,
				},
				distance:    387.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 183139.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 291.000000,
					Y: 339.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 183526.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 229.000000,
					Y: 162.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 183719.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 305.000000,
					Y: 237.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 183913.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 291.000000,
					Y: 339.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 184300.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 123.000000,
					Y: 157.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 184881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 46.000000,
					Y: 94.000000,
				},
				distance:    232.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 185848.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 139.000000,
					Y: 52.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 185945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 139.000000,
					Y: 52.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 186139.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 81.000000,
					Y: 237.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 186236.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 81.000000,
					Y: 237.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 186429.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 177.000000,
					Y: 370.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 186526.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 177.000000,
					Y: 370.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 186623.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 177.000000,
					Y: 370.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 186816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 157.000000,
					Y: 272.000000,
				},
				distance:    54.249998,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 187010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 125.000000,
					Y: 155.000000,
				},
				distance:    69.749998,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 187203.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 71.000000,
					Y: 23.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 187397.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 65.000000,
					Y: 311.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 187590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 139.000000,
					Y: 52.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 187784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 338.000000,
					Y: 120.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 188074.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 230.000000,
					Y: 243.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 188171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 230.000000,
					Y: 243.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 188365.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 288.000000,
					Y: 353.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 188558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 465.000000,
					Y: 273.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 188945.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 453.000000,
					Y: 181.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 189429.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 307.000000,
					Y: 47.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 189526.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 307.000000,
					Y: 47.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 190107.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 338.000000,
					Y: 120.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 190300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 52.000000,
					Y: 36.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 190494.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 339.000000,
					Y: 201.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 190687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 156.000000,
					Y: 286.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 190881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 469.000000,
					Y: 347.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 191074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 453.000000,
					Y: 181.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 191268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 287.000000,
					Y: 316.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 191461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 358.000000,
					Y: 353.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 191848.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 156.000000,
					Y: 286.000000,
				},
				distance:    116.250004,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 192429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 358.000000,
					Y: 354.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 192913.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 261.000000,
					Y: 235.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 193010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 261.000000,
					Y: 235.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 193203.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 25.000000,
					Y: 155.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 193590.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 377.000000,
					Y: 149.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 193784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 261.000000,
					Y: 235.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 193977.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 120.000000,
					Y: 44.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 194558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 440.000000,
					Y: 292.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 194945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 358.000000,
					Y: 21.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 195139.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 141.000000,
					Y: 125.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 195719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 341.000000,
					Y: 284.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 196494.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 211.000000,
					Y: 236.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 196687.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 144.000000,
				},
				distance:    697.500027,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 198042.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 159.000000,
					Y: 373.000000,
				},
				distance:    813.750031,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 199590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 457.000000,
					Y: 252.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 199784.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 419.000000,
					Y: 42.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 199977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 430.000000,
					Y: 174.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 200171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 400.000000,
					Y: 307.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 200558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 48.000000,
					Y: 279.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 201332.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 133.000000,
					Y: 338.000000,
				},
				distance:    348.750013,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 202687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 221.000000,
					Y: 239.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 202881.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 404.000000,
					Y: 234.000000,
				},
				distance:    697.500027,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 204235.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 404.000000,
					Y: 234.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 204429.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 459.000000,
					Y: 93.000000,
				},
				distance:    813.750031,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 205977.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 121.000000,
					Y: 305.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 206171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 190.000000,
					Y: 219.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 206364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 96.000000,
					Y: 101.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 206655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 262.000000,
					Y: 29.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 206752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 262.000000,
					Y: 29.000000,
				},
				distance:    116.250004,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 207235.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 299.000000,
					Y: 181.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 207332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 299.000000,
					Y: 181.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 207526.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 473.000000,
					Y: 51.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 208203.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 345.000000,
					Y: 264.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 208300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 345.000000,
					Y: 264.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 208977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 156.000000,
					Y: 212.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 209074.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 156.000000,
					Y: 212.000000,
				},
				distance:    930.000035,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 210816.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 156.000000,
					Y: 212.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 211010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 340.000000,
					Y: 74.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 211106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 340.000000,
					Y: 74.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 211300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 101.000000,
					Y: 279.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 211493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 184.000000,
					Y: 98.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 211590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 184.000000,
					Y: 98.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 211784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 228.000000,
					Y: 266.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 211977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 52.000000,
					Y: 180.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 212171.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 317.000000,
					Y: 248.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 212752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 340.000000,
					Y: 74.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 212946.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 270.000000,
					Y: 125.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 213042.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 270.000000,
					Y: 125.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 213139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 340.000000,
					Y: 74.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 213623.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 425.000000,
					Y: 107.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 213719.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 425.000000,
					Y: 107.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 214300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 270.000000,
					Y: 125.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 214493.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 186.000000,
					Y: 112.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 214590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 186.000000,
					Y: 112.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 214687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 270.000000,
					Y: 125.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 215171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 109.000000,
					Y: 267.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 215268.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 109.000000,
					Y: 267.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 215848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 68.000000,
					Y: 339.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 216042.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 25.000000,
					Y: 266.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 216139.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 25.000000,
					Y: 266.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 216236.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 68.000000,
					Y: 339.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 216719.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 276.000000,
					Y: 284.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 216816.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 276.000000,
					Y: 284.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 217397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 244.000000,
					Y: 173.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 217590.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 166.000000,
					Y: 208.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 217687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 166.000000,
					Y: 208.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 217784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 243.000000,
					Y: 172.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 218268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 470.000000,
					Y: 164.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 218365.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 470.000000,
					Y: 164.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 218558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 417.000000,
					Y: 225.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 218752.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 499.000000,
					Y: 367.000000,
				},
				distance:    58.125002,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 219138.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 395.000000,
					Y: 358.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 219235.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 395.000000,
					Y: 358.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 219332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 328.000000,
					Y: 318.000000,
				},
				distance:    58.125002,
				repetitions: 4,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 219816.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 195.000000,
					Y: 314.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 219913.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 195.000000,
					Y: 314.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 220107.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 19.000000,
					Y: 273.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 220300.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 192.000000,
					Y: 206.000000,
				},
				distance:    58.125002,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 220686.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 268.000000,
					Y: 171.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 220783.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 268.000000,
					Y: 171.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 220881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 252.000000,
					Y: 89.000000,
				},
				distance:    58.125002,
				repetitions: 4,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 221364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 175.000000,
					Y: 53.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 221461.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 175.000000,
					Y: 53.000000,
				},
				distance:    38.750000,
				repetitions: 7,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 221848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 131.000000,
					Y: 138.000000,
				},
				distance:    38.750000,
				repetitions: 7,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 222236.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 164.000000,
					Y: 208.000000,
				},
				distance:    38.750000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 222429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 268.000000,
					Y: 171.000000,
				},
				distance:    38.750000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 222623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 381.000000,
					Y: 193.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 222816.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 300.000000,
					Y: 372.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 222913.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 334.000000,
					Y: 296.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 223010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 383.000000,
					Y: 365.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 223590.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 268.000000,
					Y: 171.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 223784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 110.000000,
					Y: 293.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 223977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 30.000000,
					Y: 127.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 224171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 251.000000,
					Y: 304.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 224364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 352.000000,
					Y: 37.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 224558.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 383.000000,
					Y: 365.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 226300.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 210.000000,
					Y: 207.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 226493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 406.000000,
					Y: 271.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 226590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 406.000000,
					Y: 271.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 226784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 473.000000,
					Y: 233.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 226977.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 407.000000,
					Y: 192.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 227074.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 407.000000,
					Y: 192.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 227268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 254.000000,
					Y: 354.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 227461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 406.000000,
					Y: 271.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 227655.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 437.000000,
					Y: 39.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 229397.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 210.000000,
					Y: 207.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 229590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 448.000000,
					Y: 338.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 229687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 448.000000,
					Y: 338.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 229881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 364.000000,
					Y: 342.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 230073.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 287.000000,
					Y: 375.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 230170.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 287.000000,
					Y: 375.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 230365.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 476.000000,
					Y: 201.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 230461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 476.000000,
					Y: 201.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 230655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 288.000000,
					Y: 225.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 230752.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 288.000000,
					Y: 225.000000,
				},
				distance:    1240.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 231719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 210.000000,
					Y: 207.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 231913.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 49.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 232106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 359.000000,
					Y: 150.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 232300.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 18.000000,
					Y: 348.000000,
				},
				distance:    930.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 232977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 492.000000,
					Y: 224.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 233074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 492.000000,
					Y: 224.000000,
				},
				distance:    620.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 233558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 299.000000,
					Y: 286.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 233751.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 32.000000,
					Y: 213.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 233848.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 32.000000,
					Y: 213.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 235590.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 304.000000,
					Y: 55.000000,
				},
				distance:    620.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 236074.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 11.000000,
					Y: 323.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 236171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 11.000000,
					Y: 323.000000,
				},
				distance:    155.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 236558.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 11.000000,
					Y: 323.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 236655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 11.000000,
					Y: 323.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 236752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 32.000000,
					Y: 213.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 236848.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 32.000000,
					Y: 213.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 236945.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 90.000000,
					Y: 116.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 238687.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 11.000000,
					Y: 323.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 238881.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 286.000000,
					Y: 261.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 238977.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 286.000000,
					Y: 261.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 239171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 239.000000,
					Y: 196.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 239364.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 319.000000,
					Y: 185.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 239461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 319.000000,
					Y: 185.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 239655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 11.000000,
					Y: 323.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 239848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 399.000000,
					Y: 148.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 240042.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 46.000000,
					Y: 52.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 241784.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 319.000000,
					Y: 184.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 241978.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 372.000000,
					Y: 6.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 242074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 372.000000,
					Y: 6.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 242268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 194.000000,
					Y: 170.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 242461.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 372.000000,
					Y: 6.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 242558.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 372.000000,
					Y: 6.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 242752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 426.000000,
					Y: 244.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 242848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 426.000000,
					Y: 244.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 243042.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 372.000000,
					Y: 6.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 243139.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 372.000000,
					Y: 6.000000,
				},
				distance:    1550.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 244300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 165.000000,
					Y: 384.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 244494.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 16.000000,
					Y: 32.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 244687.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 345.000000,
					Y: 307.000000,
				},
				distance:    930.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 245364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 227.000000,
					Y: 198.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 245461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 232.000000,
					Y: 338.000000,
				},
				distance:    620.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 245945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 147.000000,
					Y: 83.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 246138.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 465.000000,
					Y: 266.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 246235.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 345.000000,
					Y: 307.000000,
				},
				distance:    1240.000000,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 247977.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 30.000000,
					Y: 97.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 248171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 345.000000,
					Y: 307.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 248268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 345.000000,
					Y: 307.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 248461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 246.000000,
					Y: 13.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 248655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 117.000000,
					Y: 307.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 248752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 117.000000,
					Y: 307.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 248945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 149.000000,
					Y: 82.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 249139.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 232.000000,
					Y: 338.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 249236.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 232.000000,
					Y: 338.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 249332.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 280.000000,
					Y: 274.000000,
				},
				distance:    1162.500044,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 250493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 226.000000,
					Y: 134.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 250881.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 25.000000,
					Y: 367.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 251365.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 84.000000,
					Y: 11.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 251655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 223.000000,
					Y: 364.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 252042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 324.000000,
					Y: 349.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 252429.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 45.000000,
					Y: 227.000000,
				},
				distance:    930.000035,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 254171.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 324.000000,
					Y: 349.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 254365.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 26.000000,
					Y: 368.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 254461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 26.000000,
					Y: 368.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 254655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 223.000000,
					Y: 364.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 254848.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 45.000000,
					Y: 227.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 254945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 45.000000,
					Y: 227.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 255139.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 127.000000,
					Y: 328.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 255526.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 298.000000,
					Y: 55.000000,
				},
				distance:    697.500027,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 256300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 120.000000,
					Y: 52.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 256784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 134.000000,
					Y: 205.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 257074.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 234.000000,
					Y: 335.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 257558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 473.000000,
					Y: 141.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 257848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 238.000000,
					Y: 132.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 258332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 189.000000,
					Y: 23.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 258623.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 49.000000,
					Y: 128.000000,
				},
				distance:    1627.500062,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 260171.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 460.000000,
					Y: 55.000000,
				},
				distance:    1627.500062,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 261719.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 30.000000,
					Y: 116.000000,
				},
				distance:    2480.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 267719.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 212.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 267816.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 212.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 267913.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 282.000000,
					Y: 212.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 268300.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 194.000000,
					Y: 100.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 268687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 242.000000,
					Y: 285.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 269074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 209.000000,
					Y: 180.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 269461.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 339.000000,
					Y: 35.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 269590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 339.000000,
					Y: 35.000000,
				},
				distance:    51.666667,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 269848.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 340.000000,
					Y: 113.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 270235.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 242.000000,
					Y: 285.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 270623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 508.000000,
					Y: 375.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 271010.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 456.000000,
					Y: 325.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 271397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 228.000000,
					Y: 361.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 271784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 169.000000,
					Y: 262.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 272171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 422.000000,
					Y: 254.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 272558.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 432.000000,
					Y: 25.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 272655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 432.000000,
					Y: 25.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 272752.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 376.000000,
					Y: 78.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 272848.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 376.000000,
					Y: 78.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 272945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 358.000000,
					Y: 154.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 273139.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 152.000000,
					Y: 188.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 273236.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 152.000000,
					Y: 188.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 273333.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 169.000000,
					Y: 262.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 273429.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 169.000000,
					Y: 262.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 273526.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 224.000000,
					Y: 317.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 273719.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 60.000000,
					Y: 313.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 273913.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 74.000000,
					Y: 161.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 274106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 122.000000,
					Y: 37.000000,
				},
				distance:    348.750013,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 274784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 60.000000,
					Y: 85.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 274881.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 60.000000,
					Y: 85.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 275268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 224.000000,
					Y: 317.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 275461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 411.000000,
					Y: 285.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 275558.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 473.000000,
					Y: 343.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 275655.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 473.000000,
					Y: 343.000000,
				},
				distance:    348.750013,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 276332.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 411.000000,
					Y: 285.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 276429.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 411.000000,
					Y: 285.000000,
				},
				distance:    348.750013,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 277106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 380.000000,
					Y: 208.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 277203.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 380.000000,
					Y: 208.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 277493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 439.000000,
					Y: 102.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 277784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 197.000000,
					Y: 45.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 277977.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 100.000000,
					Y: 111.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 278268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 346.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 278558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 265.000000,
					Y: 335.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 278752.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 157.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 278848.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 157.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 278945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 273.000000,
					Y: 86.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 279041.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 273.000000,
					Y: 86.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 279139.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 313.000000,
					Y: 158.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 279235.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 313.000000,
					Y: 158.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 279332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 356.000000,
					Y: 87.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 279526.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 407.000000,
					Y: 20.000000,
				},
				distance:    348.750013,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 280300.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 42.000000,
					Y: 61.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 280590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 157.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 280687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 231.000000,
					Y: 157.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 280881.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 109.000000,
					Y: 311.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 280977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 109.000000,
					Y: 311.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 281074.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 191.000000,
					Y: 301.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 281364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 80.000000,
					Y: 166.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 281461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 80.000000,
					Y: 166.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 281655.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 224.000000,
					Y: 35.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 281751.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 224.000000,
					Y: 35.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 281848.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 308.000000,
					Y: 42.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 282138.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 424.000000,
					Y: 168.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 282235.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 424.000000,
					Y: 168.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 282429.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 393.000000,
					Y: 21.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 282525.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 393.000000,
					Y: 21.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 282623.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 369.000000,
					Y: 101.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 282913.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 309.000000,
					Y: 186.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 283010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 302.000000,
					Y: 267.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 283204.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 423.000000,
					Y: 333.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 283300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 423.000000,
					Y: 333.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 283397.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 481.000000,
					Y: 274.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 283590.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 203.000000,
					Y: 278.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 283784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 244.000000,
					Y: 209.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 283880.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 244.000000,
					Y: 209.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 283977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 196.000000,
					Y: 144.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 284074.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 196.000000,
					Y: 144.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 284268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 129.000000,
					Y: 31.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 284364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 129.000000,
					Y: 31.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 284558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 164.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 284752.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 125.000000,
					Y: 256.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 284945.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 12.000000,
					Y: 349.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 285139.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 289.000000,
					Y: 250.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 285332.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 42.000000,
					Y: 260.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 285429.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 42.000000,
					Y: 260.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 285623.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 382.000000,
					Y: 262.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 285816.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 164.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 285913.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 164.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 286106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 203.000000,
					Y: 278.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 286300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 12.000000,
					Y: 349.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 286493.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 267.000000,
					Y: 333.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 287074.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 486.000000,
					Y: 207.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 287268.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 495.000000,
					Y: 285.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 287364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 495.000000,
					Y: 285.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 287461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 486.000000,
					Y: 207.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 287945.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 274.000000,
					Y: 157.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 288042.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 274.000000,
					Y: 157.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 288622.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 442.000000,
					Y: 98.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 288815.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 366.000000,
					Y: 79.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 288912.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 366.000000,
					Y: 79.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 289010.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 442.000000,
					Y: 98.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 289493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 495.000000,
					Y: 285.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 289590.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 495.000000,
					Y: 285.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 290171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 118.000000,
					Y: 242.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 290364.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 61.000000,
					Y: 295.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 290461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 61.000000,
					Y: 295.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 290558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 118.000000,
					Y: 242.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 291041.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 280.000000,
					Y: 313.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 291139.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 280.000000,
					Y: 313.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 291719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 195.000000,
					Y: 44.000000,
				},
				distance:    38.750000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 291912.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 266.000000,
					Y: 14.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 292009.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 266.000000,
					Y: 14.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 292106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 195.000000,
					Y: 44.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 292590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 287.000000,
					Y: 232.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 292687.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 287.000000,
					Y: 232.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 292880.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 335.000000,
					Y: 52.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 293074.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 415.000000,
					Y: 62.000000,
				},
				distance:    58.125002,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 293460.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 383.000000,
					Y: 136.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 293557.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 383.000000,
					Y: 136.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 293654.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 444.000000,
					Y: 193.000000,
				},
				distance:    58.125002,
				repetitions: 4,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 294138.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 421.000000,
					Y: 272.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 294235.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 421.000000,
					Y: 272.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 294429.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 292.000000,
					Y: 316.000000,
				},
				distance:    58.125002,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 294622.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 244.000000,
					Y: 192.000000,
				},
				distance:    58.125002,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 295010.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 159.000000,
					Y: 203.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 295105.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 159.000000,
					Y: 203.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 295203.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 92.000000,
					Y: 149.000000,
				},
				distance:    58.125002,
				repetitions: 4,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 295686.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 53.000000,
					Y: 221.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 295784.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 53.000000,
					Y: 221.000000,
				},
				distance:    38.750000,
				repetitions: 7,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 296171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 110.000000,
					Y: 286.000000,
				},
				distance:    38.750000,
				repetitions: 7,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 296558.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 221.000000,
					Y: 312.000000,
				},
				distance:    38.750000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 296751.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 318.000000,
					Y: 251.000000,
				},
				distance:    38.750000,
				repetitions: 3,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 296945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 356.000000,
					Y: 139.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 297138.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 167.000000,
					Y: 105.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 297235.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 244.000000,
					Y: 150.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 297332.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 247.000000,
					Y: 58.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 297912.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 352.000000,
					Y: 377.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 298106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 24.000000,
					Y: 361.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 298299.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 318.000000,
					Y: 251.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 298493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 110.000000,
					Y: 286.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 298687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 432.000000,
					Y: 176.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 298881.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 147.000000,
					Y: 170.000000,
				},
				distance:    3100.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 301010.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 353.000000,
					Y: 377.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 301203.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 352.000000,
					Y: 376.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 301397.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 471.000000,
					Y: 346.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 301590.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 476.000000,
					Y: 192.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 301784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 325.000000,
					Y: 40.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 301977.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 74.000000,
					Y: 34.000000,
				},
				distance:    3100.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 304106.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 494.000000,
					Y: 120.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 304300.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 494.000000,
					Y: 120.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 304493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 379.000000,
					Y: 168.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 304687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 221.000000,
					Y: 149.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 304881.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 74.000000,
					Y: 34.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 305074.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 109.000000,
					Y: 263.000000,
				},
				distance:    1550.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 306235.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 74.000000,
					Y: 33.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 306429.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 474.000000,
					Y: 254.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 306623.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 328.000000,
					Y: 108.000000,
				},
				distance:    930.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 307397.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 244.000000,
					Y: 54.000000,
				},
				distance:    620.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 307881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 456.000000,
					Y: 372.000000,
				},
				distance:    310.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 308171.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 448.000000,
					Y: 21.000000,
				},
				distance:    3100.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 310300.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 2.000000,
					Y: 126.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 310493.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 2.000000,
					Y: 126.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 310687.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 68.000000,
					Y: 312.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 310881.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 283.000000,
					Y: 270.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 311074.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 206.000000,
					Y: 42.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 311268.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 376.000000,
					Y: 205.000000,
				},
				distance:    3100.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 313397.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 68.000000,
					Y: 312.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 313590.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 68.000000,
					Y: 312.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 313784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 148.000000,
					Y: 180.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 313977.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 333.000000,
					Y: 160.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 314171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 356.000000,
					Y: 375.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 314364.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 132.000000,
					Y: 281.000000,
				},
				distance:    3100.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 316493.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 337.000000,
					Y: 24.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 316687.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 337.000000,
					Y: 24.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 316881.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 313.000000,
					Y: 145.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 317074.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 175.000000,
					Y: 213.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 317268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 14.000000,
					Y: 121.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 317461.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 215.000000,
					Y: 23.000000,
				},
				distance:    1550.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 318623.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 412.000000,
					Y: 320.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 318816.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 257.000000,
					Y: 205.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 319010.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 163.000000,
					Y: 361.000000,
				},
				distance:    930.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 319784.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 12.000000,
					Y: 302.000000,
				},
				distance:    620.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 320268.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 307.000000,
					Y: 367.000000,
				},
				distance:    310.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 320558.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 477.000000,
					Y: 26.000000,
				},
				distance:    2480.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 322300.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 248.000000,
					Y: 99.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 322494.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 46.000000,
					Y: 159.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 322591.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 46.000000,
					Y: 159.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 322784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 352.000000,
					Y: 15.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 322978.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 371.000000,
					Y: 365.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 323075.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 371.000000,
					Y: 365.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 323268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 477.000000,
					Y: 27.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 323462.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 256.000000,
					Y: 187.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 323559.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 256.000000,
					Y: 187.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 323655.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 179.000000,
					Y: 163.000000,
				},
				distance:    1162.500044,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 324816.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 325.000000,
					Y: 64.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 325203.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 46.000000,
					Y: 352.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 325687.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 237.000000,
					Y: 104.000000,
				},
				distance:    116.250004,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 325978.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 172.000000,
					Y: 264.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 326364.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 237.000000,
					Y: 104.000000,
				},
				distance:    348.750013,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 326752.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 180.000000,
					Y: 173.000000,
				},
				distance:    930.000035,
				repetitions: 2,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 328494.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 388.000000,
					Y: 35.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 328688.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 368.000000,
					Y: 228.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 328784.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 368.000000,
					Y: 228.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 328978.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 180.000000,
					Y: 173.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 329171.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 46.000000,
					Y: 352.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 329268.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 46.000000,
					Y: 352.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 329461.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 13.000000,
					Y: 81.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 329848.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 442.000000,
					Y: 364.000000,
				},
				distance:    697.500027,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 330623.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 429.000000,
					Y: 21.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 331106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 317.000000,
					Y: 373.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 331397.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 179.000000,
					Y: 300.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 331881.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 156.000000,
					Y: 102.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 332171.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 429.000000,
					Y: 21.000000,
				},
				distance:    465.000018,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 332655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 370.000000,
					Y: 140.000000,
				},
				distance:    232.500009,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 332945.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 495.000000,
					Y: 240.000000,
				},
				distance:    1627.500062,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 334493.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 92.000000,
					Y: 355.000000,
				},
				distance:    1627.500062,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 336042.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 476.000000,
					Y: 273.000000,
				},
				distance:    2480.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 342235.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 396.000000,
					Y: 231.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 342622.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 164.000000,
					Y: 262.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 343009.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 447.000000,
					Y: 304.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 343396.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 255.000000,
					Y: 172.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 343783.000000,
			Type: 5,
			Data: Circle{
				pos: vector.Vector2{
					X: 71.000000,
					Y: 66.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 343912.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 71.000000,
					Y: 66.000000,
				},
				distance:    51.666667,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 344170.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 140.000000,
					Y: 105.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 344558.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 396.000000,
					Y: 231.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 344945.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 104.000000,
					Y: 181.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 345332.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 19.000000,
					Y: 244.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 345719.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 259.000000,
					Y: 301.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 346106.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 182.000000,
					Y: 313.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 346493.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 327.000000,
					Y: 339.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 346881.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 415.000000,
					Y: 207.000000,
				},
				distance:    155.000000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 347364.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 327.000000,
					Y: 339.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 347461.000000,
			Type: 1,
			Data: Circle{
				pos: vector.Vector2{
					X: 327.000000,
					Y: 339.000000,
				},
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 347655.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 338.000000,
					Y: 160.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 348042.000000,
			Type: 2,
			Data: Slider{
				pos: vector.Vector2{
					X: 78.000000,
					Y: 148.000000,
				},
				distance:    77.500000,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 348429.000000,
			Type: 6,
			Data: Slider{
				pos: vector.Vector2{
					X: 84.000000,
					Y: 364.000000,
				},
				distance:    1395.000053,
				repetitions: 1,
			},
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
		&HitObject{
			Time: 349687.000000,
			Type: 12,
			Data: nil,
			Normpos: vector.Vector2{
				X: 0.000000,
				Y: 0.000000,
			},
			Angle: 0.000000,
			Strains: []float64{
				0.000000,
				0.000000,
			},
			IsSingle:  false,
			DeltaTime: 0.000000,
			DDistance: 0.000000,
		},
	},
	TPoints: []*Timing{
		&Timing{
			Time:      -1506.000000,
			MsPerBeat: 387.096774,
			Change:    true,
		},
		&Timing{
			Time:      13977.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      20171.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      26364.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      38752.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      51139.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      63526.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      70590.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      71171.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      73493.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      74074.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      75332.000000,
			MsPerBeat: -142.857143,
			Change:    false,
		},
		&Timing{
			Time:      75526.000000,
			MsPerBeat: -111.111111,
			Change:    false,
		},
		&Timing{
			Time:      75913.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      84526.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      84719.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      85493.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      85784.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      86074.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      86848.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      87816.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      88300.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      89364.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      89558.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      90719.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      90913.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      93235.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      93526.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      93913.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      94203.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      96332.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      96526.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      97106.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      97300.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      97881.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      98171.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      98558.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      98848.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      100687.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      106881.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      109977.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      111526.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      113074.000000,
			MsPerBeat: -50.000000,
			Change:    false,
		},
		&Timing{
			Time:      137848.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      150235.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      156429.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      162623.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      175010.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      182074.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      182655.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      185171.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      185364.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      186816.000000,
			MsPerBeat: -142.857143,
			Change:    false,
		},
		&Timing{
			Time:      187010.000000,
			MsPerBeat: -111.111111,
			Change:    false,
		},
		&Timing{
			Time:      187397.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      196010.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      196203.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      196977.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      197268.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      197558.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      198332.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      199300.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      199784.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      200848.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      201042.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      201623.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      201816.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      203171.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      203364.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      204719.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      204913.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      205493.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      205687.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      212171.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      218364.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      221461.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      223010.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      224558.000000,
			MsPerBeat: -50.000000,
			Change:    false,
		},
		&Timing{
			Time:      249332.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      261719.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      267913.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      272558.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      274106.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      286493.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      292687.000000,
			MsPerBeat: -133.333333,
			Change:    false,
		},
		&Timing{
			Time:      295784.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      297332.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      298881.000000,
			MsPerBeat: -50.000000,
			Change:    false,
		},
		&Timing{
			Time:      323655.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      336042.000000,
			MsPerBeat: -100.000000,
			Change:    false,
		},
		&Timing{
			Time:      342235.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      346881.000000,
			MsPerBeat: -200.000000,
			Change:    false,
		},
		&Timing{
			Time:      348429.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      351526.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      353074.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      354623.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
		&Timing{
			Time:      356171.000000,
			MsPerBeat: -66.666667,
			Change:    false,
		},
	},
}
