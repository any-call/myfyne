package mylayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

func isVerticalSpacer(obj fyne.CanvasObject) bool {
	spacer, ok := obj.(layout.SpacerObject)
	return ok && spacer.ExpandVertical()
}

func isHorizontalSpacer(obj fyne.CanvasObject) bool {
	spacer, ok := obj.(layout.SpacerObject)
	return ok && spacer.ExpandHorizontal()
}
