package oppai

// ObjScanner splits input line into chunks separated with , character
// It is used to go through a slice of bytes and iterate each field with
// GetField method.
// Code is based on: https://github.com/sirkon/snippets/blob/master/main.go
type ObjScanner struct {
	rest []byte

	i    int
	char byte
	seen bool
}

// SetSource ...
func (s *ObjScanner) SetSource(line []byte) {
	s.rest = line
}

// GetField ...
func (s *ObjScanner) GetField() (field []byte) {
	s.i = -1
	s.seen = false
	for s.i, s.char = range s.rest {
		if s.char == ',' {
			s.seen = true
			break
		}
	}
	if s.seen {
		field = s.rest[:s.i]
		s.rest = s.rest[s.i+1:]
	} else {
		field = s.rest[:s.i+1]
	}
	return
}
