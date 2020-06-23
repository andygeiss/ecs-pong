package components

const (
	TextAlignBottom = 0
	TextAlignCenter = 1
	TextAlignLeft   = 2
	TextAlignRight  = 3
	TextAlignTop    = 4
)

// Text ...
type Text struct {
	Align     int
	Color     uint32
	Content   string
	FontSize  int32
	IsEnabled bool
}

// Mask ...
func (i *Text) Mask() uint64 {
	return MaskText
}
