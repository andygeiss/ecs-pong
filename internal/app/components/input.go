package components

// Input stores the current user input.
type Input struct {
	Down bool
	Up   bool
}

// Mask ...
func (i *Input) Mask() uint64 {
	return MaskInput
}
