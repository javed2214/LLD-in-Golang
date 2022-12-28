package models

type Ladder struct {
	Start int
	End   int
}

func NewLadder() Ladder {
	return Ladder{}
}

func (s *Ladder) GetStart() int {
	return s.Start
}

func (s *Ladder) GetEnd() int {
	return s.End
}

func (s *Ladder) SetStart(start int) {
	s.Start = start
}

func (s *Ladder) SetEnd(end int) {
	s.End = end
}
