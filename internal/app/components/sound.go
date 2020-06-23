package components

// Sound contains the filename of the sound which is currently playing (empty = nothing).
type Sound struct {
	EventFilename map[string]string
	Filename      string
	IsEnabled     bool
	Volume        float32
}

// Mask ...
func (i *Sound) Mask() uint64 {
	return MaskSound
}
