package components

// Texture contains the filename of the current texture.
type Texture struct {
	Filename  string
	IsEnabled bool
}

// Mask ...
func (i *Texture) Mask() uint64 {
	return MaskTexture
}
