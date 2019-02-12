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
	Beatmap  *Map
	section  string
}

func trimSpace(s string) string { return strings.TrimSpace(s) }

func (p *Parser) property() []string {
	split := strings.SplitN(p.lastline, ":", 2)
	split[0] = trimSpace(split[0])
	if len(split) > 1 {
		split[1] = trimSpace(split[1])
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
		p.Beatmap.CS = parseFloat(trimSpace(pr[1]))
	case "OverallDifficulty":
		p.Beatmap.OD = parseFloat(trimSpace(pr[1]))
	case "ApproachRate":
		p.Beatmap.AR = parseFloat(trimSpace(pr[1]))
	case "HPDrainRate":
		p.Beatmap.HP = parseFloat(trimSpace(pr[1]))
	case "SliderMultiplier":
		p.Beatmap.SV = parseFloat(trimSpace(pr[1]))
	case "SliderTickRate":
		p.Beatmap.TickRate = parseFloat(trimSpace(pr[1]))

	}
}

func (p *Parser) timing() {
	s := strings.Split(p.lastline, ",")

	if len(s) > 8 {
		info("timing point with trailing values")
	}

	t := Timing{
		Time:      parseDouble(s[0]),
		MsPerBeat: parseDouble(s[1]),
	}

	if len(s) >= 7 {
		t.Change = !(s[6] == "0")
	}

	p.Beatmap.TPoints = append(p.Beatmap.TPoints, &t)
}

func (p *Parser) objects(s *ObjScanner) {
	t0 := parseDouble(unsafeByteToStr(s.GetField()))
	t1 := parseDouble(unsafeByteToStr(s.GetField()))
	t2 := parseDouble(unsafeByteToStr(s.GetField()))
	t3 := parseInt(unsafeByteToStr(s.GetField()))

	obj := HitObject{
		Time:    t2,
		Type:    t3,
		Strains: []float64{0.0, 0.0},
	}

	if (obj.Type & ObjCircle) != 0 {
		p.Beatmap.NCircles++
		obj.Data = Circle{
			pos: Vector2{
				X: t0,
				Y: t1,
			},
		}
	} else if (obj.Type & ObjSpinner) != 0 {
		p.Beatmap.NSpinners++
	} else if (obj.Type & ObjSlider) != 0 {
		s.GetField()
		s.GetField()
		p.Beatmap.NSliders++
		obj.Data = Slider{
			pos: Vector2{
				X: t0,
				Y: t1,
			},
			repetitions: parseInt(unsafeByteToStr(s.GetField())),
			distance:    parseDouble(unsafeByteToStr(s.GetField())),
		}
	}
	p.Beatmap.Objects = append(p.Beatmap.Objects, &obj)
}

// Map returns the beatmap info
func (p *Parser) Map(reader io.Reader) *Map {
	var line string
	objScanner := &ObjScanner{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line = scanner.Text()
		p.lastline = line

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
			objScanner.SetSource(scanner.Bytes())
			p.objects(objScanner)
		}

	}
	return p.Beatmap
}

// Map returns the beatmap info with a specified number of objects parsed
func (p *Parser) MapbyNum(reader io.Reader, objnum int) *Map {
	num := objnum
	var line string
	objScanner := &ObjScanner{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line = scanner.Text()
		p.lastline = line

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
			objScanner.SetSource(scanner.Bytes())
			p.objects(objScanner)
			num -= 1
			if num == 0 {
				return p.Beatmap
			}
		}
	}
	return p.Beatmap
}

