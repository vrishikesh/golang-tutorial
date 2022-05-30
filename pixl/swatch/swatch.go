package swatch

import (
	"image/color"

	"fyne.io/fyne/v2/widget"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}
