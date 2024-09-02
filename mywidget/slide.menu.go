package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/myfyne"
	"image/color"
)

// SideMenu 定义自定义组件
type SideMenu struct {
	widget.BaseWidget
	menuItems       []myfyne.MenuItem
	onItemSelected  func(myfyne.MenuItem)
	alignment       fyne.TextAlign
	padding         float32
	accordion       *widget.Accordion
	textColor       color.Color
	selectTextColor color.Color
	hoverTextColor  color.Color
}

// NewSideMenu 创建一个新的 SideMenu 控件
func NewSideMenu(menuItems []myfyne.MenuItem, onItemSelected func(myfyne.MenuItem)) *SideMenu {
	sideMenu := &SideMenu{
		menuItems:      menuItems,
		alignment:      fyne.TextAlignLeading,
		padding:        8,
		onItemSelected: onItemSelected,
	}
	sideMenu.ExtendBaseWidget(sideMenu)
	sideMenu.buildMenu()
	return sideMenu
}

// CreateRenderer 创建组件的渲染器
func (sm *SideMenu) CreateRenderer() fyne.WidgetRenderer {
	sidebar := container.NewVBox(
		widget.NewLabelWithStyle("主菜单", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		sm.accordion,
	)
	return widget.NewSimpleRenderer(sidebar)
}

// buildMenu 构建菜单的内容
func (sm *SideMenu) buildMenu() {
	sm.accordion = widget.NewAccordion()
	sm.refreshMenuItems()
}

// refreshMenuItems 刷新菜单项显示
func (sm *SideMenu) refreshMenuItems() {
	sm.accordion.Items = nil
	for _, item := range sm.menuItems {
		subMenu := sm.createSubMenu(item, 0)
		accordionItem := widget.NewAccordionItem(item.Name, subMenu)
		sm.accordion.Append(accordionItem)
	}
	sm.Refresh()
}

// createSubMenu 创建子菜单，支持无限嵌套并设置左侧缩进
func (sm *SideMenu) createSubMenu(item myfyne.MenuItem, level int) *fyne.Container {
	subMenuList := container.NewVBox()

	for _, subItem := range item.SubItems {
		subItemCopy := subItem // 避免闭包引用错误
		btn := NewMenuButton(subItem.Name, func() {
			if sm.onItemSelected != nil {
				sm.onItemSelected(subItemCopy)
			}
		})

		btn.SetTextColor(sm.textColor).SetSelectTextColor(sm.selectTextColor).SetHoverTextColor(sm.hoverTextColor)

		// 根据 alignment 设置对齐方式，并增加 left padding
		leftPadding := NewFixedWidthBox(sm.padding, nil, nil)
		paddingContainer := container.NewHBox(leftPadding, btn)

		switch sm.alignment {
		case fyne.TextAlignCenter:
			paddingContainer = container.NewHBox(layout.NewSpacer(), btn, layout.NewSpacer())
		case fyne.TextAlignTrailing:
			paddingContainer = container.NewHBox(layout.NewSpacer(), btn)
		default: // 默认左对齐并增加左侧缩进
			paddingContainer = container.NewHBox(leftPadding, btn)
		}

		subMenuList.Add(paddingContainer)

		// 如果存在子菜单，递归创建，并且保持缩进
		if len(subItem.SubItems) > 0 {
			nestedSubMenu := sm.createSubMenu(subItem, level+1)
			accordionItem := widget.NewAccordionItem(subItem.Name, nestedSubMenu)
			accordion := widget.NewAccordion(accordionItem)

			// 包装 accordion 以保持缩进和对齐
			paddedAccordion := container.NewVBox(
				container.NewHBox(leftPadding, accordion),
			)
			subMenuList.Add(paddedAccordion)
		}
	}

	return subMenuList
}

func (sm *SideMenu) SetAlignment(alignment fyne.TextAlign) *SideMenu {
	sm.alignment = alignment
	sm.refreshMenuItems()
	return sm
}

func (sm *SideMenu) SetLeftPadding(padding float32) *SideMenu {
	sm.padding = padding
	sm.refreshMenuItems()
	return sm
}

func (sm *SideMenu) SetTextColor(c color.Color) *SideMenu {
	sm.textColor = c
	sm.refreshMenuItems()
	return sm
}

func (sm *SideMenu) SetHoverTextColor(c color.Color) *SideMenu {
	sm.hoverTextColor = c
	sm.refreshMenuItems()
	return sm
}

func (sm *SideMenu) SetSelectTextColor(c color.Color) *SideMenu {
	sm.selectTextColor = c
	sm.refreshMenuItems()
	return sm
}

// AddMenuItem 动态增加一个菜单项
func (sm *SideMenu) AddMenu(item myfyne.MenuItem) {
	sm.menuItems = append(sm.menuItems, item)
	sm.refreshMenuItems()
}

func (sm *SideMenu) AddSubMenu(parItem myfyne.MenuItem, subItem myfyne.MenuItem) {
	for i, item := range sm.menuItems {
		if item.Name == parItem.Name {
			sm.menuItems[i].SubItems = append(sm.menuItems[i].SubItems, subItem)
			break
		}
	}
	sm.refreshMenuItems()
}

// RemoveMenuItem 动态删除一个菜单项
func (sm *SideMenu) RemoveMenu(itemName string) {
	for i, item := range sm.menuItems {
		if item.Name == itemName {
			sm.menuItems = append(sm.menuItems[:i], sm.menuItems[i+1:]...)
			break
		}
	}
	sm.refreshMenuItems()
}
