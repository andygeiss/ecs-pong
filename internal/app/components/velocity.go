package components

// Velocity contains the speed in pixels on the X and Y axis.
type Velocity struct {
	X         float32
	Y         float32
	IsEnabled bool
}

// Mask ...
func (i *Velocity) Mask() uint64 {
	return MaskVelocity
}
