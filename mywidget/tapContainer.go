package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// ClickableContainer 是一个自定义容器控件，支持单击和双击事件。
type TapContainer struct {
	widget.BaseWidget
	Child    fyne.CanvasObject // 容器内的子控件
	OnTapped func()            // 单击回调
}

func NewTapContainer(child fyne.CanvasObject, onTap func()) *TapContainer {
	c := &TapContainer{
		Child:    child,
		OnTapped: onTap,
	}
	c.ExtendBaseWidget(c)
	return c
}

// Tapped 处理单击和双击事件。
func (c *TapContainer) Tapped(ev *fyne.PointEvent) {
	if c.OnTapped != nil {
		c.OnTapped()
	}
}

// TappedSecondary 用于处理右键点击事件，但此处未使用。
func (c *TapContainer) TappedSecondary(_ *fyne.PointEvent) {}

// CreateRenderer 创建容器的渲染器。
func (c *TapContainer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.Child)
}
