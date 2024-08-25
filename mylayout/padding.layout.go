package mylayout

import (
	"fyne.io/fyne/v2"
	"github.com/any-call/myfyne/mybase"
)

type LayoutPadding struct {
	padding mybase.EdgeInset
}

func NewHPaddingLayout(padding float32) *LayoutPadding {
	return &LayoutPadding{
		padding: mybase.EdgeInset{
			Top:    0,
			Right:  padding,
			Bottom: 0,
			Left:   padding,
		},
	}
}

func NewVPaddingLayout(padding float32) *LayoutPadding {
	return &LayoutPadding{
		padding: mybase.EdgeInset{
			Top:    padding,
			Right:  0,
			Bottom: padding,
			Left:   0,
		},
	}
}

func NewPaddingLayout(top, bottom, left, right float32) *LayoutPadding {
	return &LayoutPadding{
		padding: mybase.EdgeInset{
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
