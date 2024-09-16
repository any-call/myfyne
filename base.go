package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"math"
)

type (
	EdgeInset struct {
		Top, Right, Bottom, Left float32
	}

	Page interface {
		Content() fyne.CanvasObject
		WinTitle() string
		WinID() int
		WinSize() fyne.Size
		WinClosed()
	}

	//自定义对话框
	DialogContent interface {
		Title() string
		Content() fyne.CanvasObject  // 返回用于展示的内容
		SetDialog(dlg dialog.Dialog) // 绑定 dialog 对象
		SetWindow(win fyne.Window)   // 绑定 window 对象
		CloseDialog(param any)       // 用于关闭 dialog，并传递关闭参数
		GetCloseParam() any          // 获取关闭时的参数
	}

	WinWillCloseFn func() bool
	// MenuItem 定义菜单项的结构
	MenuItemModel struct {
		Name     string
		Icon     fyne.Resource
		SubItems []MenuItemModel
		OnTapCb  func(name string)
	}
)

// MainAxisAlignment 定义了主轴对齐方式
type MainAxisAlignment int

// MainAxisAlignment 的可能取值
const (
	MainAxisAlignStart        MainAxisAlignment = iota // 将子控件在主轴方向上对齐到起始位置
	MainAxisAlignEnd                                   // 将子控件在主轴方向上对齐到结束位置
	MainAxisAlignCenter                                // 将子控件在主轴方向上居中对齐
	MainAxisAlignSpaceBetween                          // 在子控件之间均匀分配剩余空间
	MainAxisAlignSpaceAround                           // 在子控件周围均匀分配剩余空间
	MainAxisAlignSpaceEvenly                           // 在子控件之间和两端均匀分配剩余空间
)

// CrossAxisAlignment 定义了交叉轴对齐方式
type CrossAxisAlignment int

// CrossAxisAlignment 的可能取值
const (
	CrossAxisAlignStart   CrossAxisAlignment = iota // 子控件在交叉轴方向上对齐到起始位置
	CrossAxisAlignEnd                               // 子控件在交叉轴方向上对齐到结束位置
	CrossAxisAlignCenter                            // 子控件在交叉轴方向上居中对齐
	CrossAxisAlignStretch                           // 子控件在交叉轴方向上拉伸以填充可用空间
)

type Position int

const (
	PositionTopLeft Position = iota
	PositionTopCenter
	PositionTopRight
	PositionCenterLeft
	PositionCenter
	PositionCenterRight
	PositionBottomLeft
	PositionBottomCenter
	PositionBottomRight
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
	case PositionTopLeft:
		x, y = 0, 0
		break

	case PositionTopCenter:
		x = (parentSize.Width - childSize.Width) / 2
		y = 0
		break

	case PositionTopRight:
		x = parentSize.Width - childSize.Width
		y = 0
		break

	case PositionCenterLeft:
		x = 0
		y = (parentSize.Height - childSize.Height) / 2
		break

	case PositionCenter:
		x = (parentSize.Width - childSize.Width) / 2
		y = (parentSize.Height - childSize.Height) / 2
		break

	case PositionCenterRight:
		x = parentSize.Width - childSize.Width
		y = (parentSize.Height - childSize.Height) / 2
		break

	case PositionBottomLeft:
		x = 0
		y = parentSize.Height - childSize.Height
		break

	case PositionBottomCenter:
		x = (parentSize.Width - childSize.Width) / 2
		y = parentSize.Height - childSize.Height
		break

	case PositionBottomRight:
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
