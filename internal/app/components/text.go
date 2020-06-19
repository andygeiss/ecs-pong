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

// Name ...
func (t *Text) Name() string {
	return "text"
}
