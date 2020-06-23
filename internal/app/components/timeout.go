package components

import "time"

// Timeout should contain a creation time and a duration.
type Timeout struct {
	CreationTime time.Time
	Duration     time.Duration
}

// Mask ...
func (i *Timeout) Mask() uint64 {
	return MaskTimeout
}
