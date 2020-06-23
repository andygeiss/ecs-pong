package components

// Size contains the width and height of an entity.
type Size struct {
	Width  float32
	Height float32
}

// Mask ...
func (i *Size) Mask() uint64 {
	return MaskSize
}
