package oppai

import (
	"bufio"
	"io"
	"strings"

	"github.com/flesnuk/oppai5/vector"
)

// https://git.sr.ht/~tdeo/pp/tree/master/parser.go

func parse(r io.Reader, objnum int) (*Map, error) {
	var b = &Map{}

	s := bufio.NewScanner(r)

	var section string
	var arFound bool
	for s.Scan() {
		t := s.Text()

		if strings.HasPrefix(t, " ") || strings.HasPrefix(t, "_") {
			continue
		}

		t = strings.TrimSpace(t)

		if strings.HasPrefix(t, "//") {
			continue
		}

		if strings.HasPrefix(t, "[") && strings.HasSuffix(t, "]") {
			section = t[1 : len(t)-1]
			continue
		}

		switch section {
		case "Metadata":
			switch k, v := keyValue(t); k {
			case "Title":
				b.Title = v
			case "TitleUnicode":
				b.TitleUnicode = v
			case "Artist":
				b.Artist = v
			case "ArtistUnicode":
				b.ArtistUnicode = v
			case "Creator":
				b.Creator = v
			case "Version":
				b.Version = v
			}

		case "General":
			k, v := keyValue(t)
			switch k {
			case "Mode":
				switch v {
				case "0":
					b.Mode = ModeStd
				default:
					return nil, ErrGamemodeUnsupported

					/* Kept for future reference

					case "1":
						b.Mode = ModeTaiko
					case "2":
						b.Mode = ModeCatchTheBeat
					case "3":
						b.Mode = ModeMania

					*/
				}
			}

		case "Difficulty":
			k, v := keyValue(t)
			switch k {
			case "HPDrainRate":
				b.HP = parseFloat(v)
			case "CircleSize":
				b.CS = parseFloat(v)
			case "OverallDifficulty":
				b.OD = parseFloat(v)
			case "ApproachRate":
				b.AR = parseFloat(v)
				arFound = true
			case "SliderMultiplier":
				b.SV = parseFloat(v)
			case "SliderTickRate":
				b.TickRate = parseFloat(v)
			}

		case "TimingPoints":
			split := strings.Split(t, ",")
			if len(split) < 2 {
				continue
			}

			t := &Timing{
				Time:      parseFloat(split[0]),
				MsPerBeat: parseFloat(split[1]),
			}

			if len(split) >= 7 {
				t.Change = split[6] != "0"
			}

			b.TPoints = append(b.TPoints, t)

		case "HitObjects":
			split := strings.Split(t, ",")

			if len(split) < 4 {
				continue
			}

			o := &HitObject{
				Time:    parseFloat(split[2]),
				Strains: []float64{0.0, 0.0},

				// ignoring error because it wouldn't match anyway
				Type: parseInt(split[3]),
			}

			// this would be & 0b1011 if Go had binary literals
			switch o.Type & 11 {
			case ObjCircle:
				b.NCircles++
				o.Data = Circle{
					pos: vector.New(
						parseFloat(split[0]),
						parseFloat(split[1]),
					),
				}

			case ObjSpinner:
				b.NSpinners++

			case ObjSlider:
				if len(split) < 8 {
					continue
				}

				b.NSliders++
				o.Data = Slider{
					pos: vector.New(
						parseFloat(split[0]),
						parseFloat(split[1]),
					),
					repetitions: parseInt(split[6]),
					distance:    parseFloat(split[7]),
				}
			}

			b.Objects = append(b.Objects, o)

			// if objnum is fed
			if objnum > -1 {
				objnum--
				if objnum == 0 {
					return b, nil
				}
			}
		}
	}

	if !arFound {
		b.AR = b.OD
	}

	return b, nil
}
