package myfyne

import "fyne.io/fyne/v2"

type (
	EdgeInset struct {
		Top, Right, Bottom, Left float32
	}
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