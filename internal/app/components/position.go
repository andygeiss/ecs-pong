package components

// Position contains the 2D X and Y coordinate.
type Position struct {
	X float32
	Y float32
}

// Mask ...
func (i *Position) Mask() uint64 {
	return MaskPosition
}
