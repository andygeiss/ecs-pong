package components

// Score contains the game score value, winning text and its position on the screen.
type Score struct {
	Text  string
	Value int
	X     int32
	Y     int32
}

// Name ...
func (s *Score) Name() string {
	return "score"
}
