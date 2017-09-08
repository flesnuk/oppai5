package oppai

import (
	"bufio"
	"io"
	"strings"
)

/* ------------------------------------------------------------- */
/* beatmap parser                                                */

// Parser for storing the status of parsing
type Parser struct {
	lastline string // last lines touched
	nline    int
	lastpos  string
	done     bool
	Beatmap  *Map
	section  string
}

func (p *Parser) setlastpos(v string) string {
	v = strings.TrimSpace(v)
	p.lastpos = v
	return v
}

func (p *Parser) property() []string {
	split := strings.SplitN(p.lastline, ":", 2)
	split[0] = p.setlastpos(split[0])
	if len(split) > 1 {
		split[1] = p.setlastpos(split[1])
	}
	return split
}

func (p *Parser) metadata() {
	pr := p.property()
	switch pr[0] {
	case "Title":
		p.Beatmap.Title = pr[1]
	case "TitleUnicode":
		p.Beatmap.TitleUnicode = pr[1]
	case "Artist":
		p.Beatmap.Artist = pr[1]
	case "ArtistUnicode":
		p.Beatmap.ArtistUnicode = pr[1]
	case "Creator":
		p.Beatmap.Creator = pr[1]
	case "Version":
		p.Beatmap.Version = pr[1]
	}
}

func (p *Parser) general() {
	pr := p.property()
	if pr[0] == "Mode" {
		p.Beatmap.Mode = parseInt(pr[1])

		if p.Beatmap.Mode != ModeStd {
			panic("this gamemode is not yet supported")
		}
	}
}

func (p *Parser) difficulty() {
	pr := p.property()

	switch pr[0] {
	case "CircleSize":
		p.Beatmap.CS = parseFloat(p.setlastpos(pr[1]))
	case "OverallDifficulty":
		p.Beatmap.OD = parseFloat(p.setlastpos(pr[1]))
	case "ApproachRate":
		p.Beatmap.AR = parseFloat(p.setlastpos(pr[1]))
	case "HPDrainRate":
		p.Beatmap.HP = parseFloat(p.setlastpos(pr[1]))
	case "SliderMultiplier":
		p.Beatmap.SV = parseFloat(p.setlastpos(pr[1]))
	case "SliderTickRate":
		p.Beatmap.TickRate = parseFloat(p.setlastpos(pr[1]))

	}
}

func (p *Parser) timing() {
	s := strings.Split(p.lastline, ",")

	if len(s) > 8 {
		info("timing point with trailing values")
	}

	t := Timing{
		Time:      parseDouble(p.setlastpos(s[0])),
		MsPerBeat: parseDouble(p.setlastpos(s[1])),
	}

	if len(s) >= 7 {
		t.Change = !(strings.TrimSpace(s[6]) == "0")
	}

	p.Beatmap.TPoints = append(p.Beatmap.TPoints, t)
}

func (p *Parser) objects() {
	s := strings.Split(p.lastline, ",")

	if len(s) > 11 {
		info("object with trailing values")
	}

	obj := HitObject{
		Time:    parseDouble(p.setlastpos(s[2])),
		Type:    parseInt(p.setlastpos(s[3])),
		Strains: []float64{0.0, 0.0},
	}

	if (obj.Type & ObjCircle) != 0 {
		p.Beatmap.NCircles++
		c := Circle{
			pos: Vector2{
				X: parseDouble(p.setlastpos(s[0])),
				Y: parseDouble(p.setlastpos(s[1])),
			},
		}
		obj.Data = c
	} else if (obj.Type & ObjSpinner) != 0 {
		p.Beatmap.NSpinners++
	} else if (obj.Type & ObjSlider) != 0 {
		p.Beatmap.NSliders++
		sli := Slider{
			pos: Vector2{
				X: parseDouble(p.setlastpos(s[0])),
				Y: parseDouble(p.setlastpos(s[1])),
			},
			repetitions: parseInt(p.setlastpos(s[6])),
			distance:    parseDouble(p.setlastpos(s[7])),
		}
		obj.Data = sli
	}
	p.Beatmap.Objects = append(p.Beatmap.Objects, obj)
}

// Map returns the beatmap info
func (p *Parser) Map(reader io.Reader) *Map {
	var line string

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line = scanner.Text()
		p.lastline = line
		p.nline++

		if strings.HasPrefix(line, " ") ||
			strings.HasPrefix(line, "_") {
			continue
		}

		p.lastline = strings.TrimSpace(line)
		line = p.lastline
		if len(line) <= 0 {
			continue
		}

		if strings.HasPrefix(line, "//") {
			continue
		}

		// [SectionName]
		if strings.HasPrefix(line, "[") {
			p.section = line[1 : len(line)-1]
			continue
		}

		switch p.section {
		case "Metadata":
			p.metadata()
		case "General":
			p.general()
		case "Difficulty":
			p.difficulty()
		case "TimingPoints":
			p.timing()
		case "HitObjects":
			p.objects()
		}

	}
	p.done = true
	return p.Beatmap
}
