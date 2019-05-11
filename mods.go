package oppai

import "strings"

/* ------------------------------------------------------------- */
/* mods utils                                                    */

// bitmask mods
const (
	ModsNOMOD int = 0

	ModsNF int = 1 << 0
	ModsEZ int = 1 << 1
	ModsTD int = 1 << 2
	ModsHD int = 1 << 3
	ModsHR int = 1 << 4
	ModsDT int = 1 << 6
	ModsHT int = 1 << 8
	ModsNC int = 1 << 9
	ModsFL int = 1 << 10
	ModsSO int = 1 << 12

	ModsSpeedChanging = ModsDT | ModsHT | ModsNC
	ModsMapChanging   = ModsHR | ModsEZ | ModsSpeedChanging
)

// ModsStr returns a string representation of the mods, such as HDDT
func ModsStr(mods int) string {
	var s strings.Builder

	if mods == 0 {
		return ""
	}

	if (mods & ModsNF) != 0 {
		s.WriteString("NF")
	}
	if (mods & ModsEZ) != 0 {
		s.WriteString("EZ")
	}
	if (mods & ModsTD) != 0 {
		s.WriteString("TD")
	}
	if (mods & ModsHD) != 0 {
		s.WriteString("HD")
	}
	if (mods & ModsHR) != 0 {
		s.WriteString("HR")
	}
	if (mods & ModsDT) != 0 {
		s.WriteString("DT")
	}
	if (mods & ModsHT) != 0 {
		s.WriteString("HT")
	}
	if (mods & ModsNC) != 0 {
		s.WriteString("NC")
	}
	if (mods & ModsFL) != 0 {
		s.WriteString("FL")
	}
	if (mods & ModsSO) != 0 {
		s.WriteString("SO")
	}

	return s.String()
}
