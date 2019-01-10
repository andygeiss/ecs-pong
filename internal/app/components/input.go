package components

// Input stores the current user input.
type Input struct {
	Down  bool
	Up    bool
}

// Name ...
func (i *Input) Name() string {
	return "input"
}
