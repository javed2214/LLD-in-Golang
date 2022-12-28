package models

type Snake struct {
	Start int
	End   int
}

func NewSnake() Snake {
	return Snake{}
}

func (s *Snake) GetStart() int {
	return s.Start
}

func (s *Snake) GetEnd() int {
	return s.End
}

func (s *Snake) SetStart(start int) {
	s.Start = start
}

func (s *Snake) SetEnd(end int) {
	s.End = end
}
