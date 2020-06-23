package components

// Score contains the game score value for the enemy and player.
type Score struct {
	Enemy  int32
	Player int32
}

// Mask ...
func (i *Score) Mask() uint64 {
	return MaskScore
}
