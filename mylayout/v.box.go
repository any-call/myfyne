package mylayout

import (
	"fyne.io/fyne/v2"
)

type VBoxLayoutWithSpacing struct {
	paddingSpace float32
}

func NewVBoxWithSpacing(space float32) fyne.Layout {
	return VBoxLayoutWithSpacing{paddingSpace: space}
}

func (v VBoxLayoutWithSpacing) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	spacers := 0
	visibleObjects := 0
	// Size taken up by visible objects
	total := float32(0)

	for _, child := range objects {
		if !child.Visible() {
			continue
		}
		if isHorizontalSpacer(child) {
			spacers++
			continue
		}

		visibleObjects++
		total += child.MinSize().Height
	}

	padding := v.paddingSpace

	// Amount of space not taken up by visible objects and inter-object padding
	extra := size.Height - total - (padding * float32(visibleObjects-1))

	// Spacers split extra space equally
	spacerSize := float32(0)
	if spacers > 0 {
		spacerSize = extra / float32(spacers)
	}

	x, y := float32(0), float32(0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		if isVerticalSpacer(child) {
			child.Move(fyne.NewPos(x, y))
			child.Resize(fyne.NewSize(size.Width, spacerSize))
			y += spacerSize
			continue
		}
		child.Move(fyne.NewPos(x, y))

		height := child.MinSize().Height
		y += padding + height
		child.Resize(fyne.NewSize(size.Width, height))
	}
}

func (v VBoxLayoutWithSpacing) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	addPadding := false
	padding := v.paddingSpace
	for _, child := range objects {
		if !child.Visible() || isVerticalSpacer(child) {
			continue
		}

		childMin := child.MinSize()
		minSize.Width = fyne.Max(childMin.Width, minSize.Width)
		minSize.Height += childMin.Height
		if addPadding {
			minSize.Height += padding
		}
		addPadding = true
	}
	return minSize
}
