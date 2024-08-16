package myfyne

import (
	"fyne.io/fyne/v2"
	"math"
)

type (
	EdgeInset struct {
		Top, Right, Bottom, Left float32
	}

	Position int

	Page interface {
		Content() fyne.CanvasObject
		WinTitle() string
		WinID() int
		WinWidth() float32
		WinHeight() float32
	}
)

const (
	TopLeft Position = iota
	TopCenter
	TopRight
	CenterLeft
	Center
	CenterRight
	BottomLeft
	BottomCenter
	BottomRight
)

const (
	Infinity float32 = math.MaxFloat32 //代表无穷大，一般表示可以尽可能的占用父类的空间
)

func GetWindow(obj fyne.CanvasObject) fyne.Window {
	listWindow := fyne.CurrentApp().Driver().AllWindows()

	for _, win := range listWindow {
		if containsObject(win.Content(), obj) {
			return win
		}
	}

	return nil
}

func ChildPosition(position Position, parentSize, childSize fyne.Size) fyne.Position {
	var x, y float32

	switch position {
	case TopLeft:
		x, y = 0, 0
		break

	case TopCenter:
		x = (parentSize.Width - childSize.Width) / 2
		y = 0
		break

	case TopRight:
		x = parentSize.Width - childSize.Width
		y = 0
		break

	case CenterLeft:
		x = 0
		y = (parentSize.Height - childSize.Height) / 2
		break

	case Center:
		x = (parentSize.Width - childSize.Width) / 2
		y = (parentSize.Height - childSize.Height) / 2
		break

	case CenterRight:
		x = parentSize.Width - childSize.Width
		y = (parentSize.Height - childSize.Height) / 2
		break

	case BottomLeft:
		x = 0
		y = parentSize.Height - childSize.Height
		break

	case BottomCenter:
		x = (parentSize.Width - childSize.Width) / 2
		y = parentSize.Height - childSize.Height
		break

	case BottomRight:
		x = parentSize.Width - childSize.Width
		y = parentSize.Height - childSize.Height
		break
	}

	return fyne.NewPos(x, y)
}

// 递归查找 obj 是否在 root 中
func containsObject(root, obj fyne.CanvasObject) bool {
	if root == obj {
		return true
	}

	if container, ok := root.(*fyne.Container); ok {
		for _, child := range container.Objects {
			if containsObject(child, obj) {
				return true
			}
		}
	}
	return false
}
