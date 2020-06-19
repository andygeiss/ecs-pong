package components

import "time"

// Timeout should contain a creation time and a duration.
type Timeout struct {
	CreationTime time.Time
	Duration     time.Duration
}

// Name ...
func (t *Timeout) Name() string {
	return "timeout"
}
