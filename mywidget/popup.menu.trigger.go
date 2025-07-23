package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MenuDefine struct {
	Name  string
	OnTap func()
}

type PopupMenuTrigger struct {
	widget.BaseWidget
	trigger           fyne.CanvasObject // 触发器对象（任意 fyne.CanvasObject）
	menu              *fyne.Menu        // 弹出的菜单
	isSecondaryTapped bool
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

func NewSimplePopupMenu(trigger fyne.CanvasObject, listMenus []MenuDefine) *PopupMenuTrigger {
	menu := fyne.NewMenu("")
	for i, _ := range listMenus {
		menu.Items = append(menu.Items, fyne.NewMenuItem(listMenus[i].Name, listMenus[i].OnTap))
	}
	return NewPopupMenuTrigger(trigger, menu)
}

func (p *PopupMenuTrigger) SetSecondaryTapped(flag bool) {
	p.isSecondaryTapped = flag
}

// 实现点击事件
func (p *PopupMenuTrigger) Tapped(e *fyne.PointEvent) {
	if p.isSecondaryTapped == false {
		canvas := fyne.CurrentApp().Driver().CanvasForObject(p.trigger.(fyne.CanvasObject))
		widget.ShowPopUpMenuAtPosition(p.menu, canvas, e.AbsolutePosition)
	}
}

func (p *PopupMenuTrigger) TappedSecondary(e *fyne.PointEvent) {
	if p.isSecondaryTapped {
		canvas := fyne.CurrentApp().Driver().CanvasForObject(p.trigger.(fyne.CanvasObject))
		widget.ShowPopUpMenuAtPosition(p.menu, canvas, e.AbsolutePosition)
	}
}

func (p *PopupMenuTrigger) GetMenu() *fyne.Menu {
	return p.menu
}

// 渲染器实现
func (p *PopupMenuTrigger) CreateRenderer() fyne.WidgetRenderer {
	// 包装 trigger 的渲染器
	return widget.NewSimpleRenderer(p.trigger.(fyne.CanvasObject))
}
