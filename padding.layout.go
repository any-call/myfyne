package myfyne

import "fyne.io/fyne/v2"

type LayoutPadding struct {
	padding EdgeInset
}

func NewPaddingLayout(top, bottom, left, right float32) *LayoutPadding {
	return &LayoutPadding{
		padding: EdgeInset{
			Top:    top,
			Right:  right,
			Bottom: bottom,
			Left:   left,
		},
	}
}

func (s *LayoutPadding) Layout(c []fyne.CanvasObject, containerSize fyne.Size) {
	x := s.padding.Left
	y := s.padding.Top
	width := containerSize.Width - s.padding.Left - s.padding.Right
	height := containerSize.Height - s.padding.Top - s.padding.Bottom
	for _, obj := range c {
		obj.Resize(fyne.NewSize(width, height))
		obj.Move(fyne.NewPos(x, y))
	}
}

func (s *LayoutPadding) MinSize(c []fyne.CanvasObject) fyne.Size {
	minWidth := s.padding.Left + s.padding.Right
	minHeight := s.padding.Top + s.padding.Bottom

	for _, obj := range c {
		minSize := obj.MinSize()
		minWidth += minSize.Width
		minHeight += minSize.Height
	}

	return fyne.NewSize(minWidth, minHeight)
}
