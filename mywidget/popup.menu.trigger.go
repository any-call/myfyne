package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type PopupMenuTrigger struct {
	widget.BaseWidget
	trigger fyne.CanvasObject // 触发器对象（任意 fyne.CanvasObject）
	menu    *fyne.Menu        // 弹出的菜单
}

// 创建新的 PopupMenuTrigger
func NewPopupMenuTrigger(trigger fyne.CanvasObject, menu *fyne.Menu) *PopupMenuTrigger {
	p := &PopupMenuTrigger{
		trigger: trigger,
		menu:    menu,
	}

	p.ExtendBaseWidget(p)
	return p
}

// 实现点击事件
func (p *PopupMenuTrigger) Tapped(e *fyne.PointEvent) {
	canvas := fyne.CurrentApp().Driver().CanvasForObject(p.trigger.(fyne.CanvasObject))
	widget.ShowPopUpMenuAtPosition(p.menu, canvas, e.AbsolutePosition)
}

// 渲染器实现
func (p *PopupMenuTrigger) CreateRenderer() fyne.WidgetRenderer {
	// 包装 trigger 的渲染器
	return widget.NewSimpleRenderer(p.trigger.(fyne.CanvasObject))
}
