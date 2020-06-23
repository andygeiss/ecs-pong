package components

// AI contains the current decision by the AI system.
type AI struct {
	Down bool
	Up   bool
}

// Mask ...
func (i *AI) Mask() uint64 {
	return MaskAI
}
