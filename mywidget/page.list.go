package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PageList 是基于 widget.BaseWidget 的自定义控件
type PageList struct {
	widget.BaseWidget
	list           *widget.List
	placeholder    *fyne.Container
	placeholderMsg *widget.Label
	onRefresh      func()               // 刷新数据的回调
	onConfig       func(t *widget.List) // 配置 Table 的回调
}

// NewPageList 创建一个新的 PageList
// onRefresh 用于无数据时显示的刷新操作，onConfig 用于配置 Table 的回调
func NewPageList(onRefresh func(), onConfig func(t *widget.List)) *PageList {
	pt := &PageList{
		onRefresh: onRefresh,
		onConfig:  onConfig,
	}

	// 创建表格
	pt.list = widget.NewList(
		func() int { return 0 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, object fyne.CanvasObject) {})

	pt.list.HideSeparators = false
	// 使用回调配置 Table
	if pt.onConfig != nil {
		pt.onConfig(pt.list)
	}

	// 必须调用 .ExtendBaseWidget 来确保控件能够正确渲染
	pt.ExtendBaseWidget(pt)
	return pt
}

// CreateRenderer 用于渲染 PageList 控件
func (pt *PageList) CreateRenderer() fyne.WidgetRenderer {
	// 占位界面
	pt.placeholderMsg = widget.NewLabel("没有数据")
	pt.placeholderMsg.Alignment = fyne.TextAlignCenter
	pt.placeholderMsg.Wrapping = fyne.TextWrapWord
	refreshButton := widget.NewButton("点击刷新", func() {
		// 点击刷新时调用外部提供的刷新回调
		if pt.onRefresh != nil {
			pt.onRefresh()
		}
	})

	pt.placeholder = container.NewVBox(pt.placeholderMsg, refreshButton)
	stack := container.NewStack(pt.placeholder, pt.list)

	// 初始显示占位界面
	pt.ShowPlaceholder("")

	return widget.NewSimpleRenderer(stack)
}

// ShowPlaceholder 显示占位界面，动态更新提示信息
func (pt *PageList) ShowPlaceholder(message string) {
	if message == "" {
		message = "没有数据"
	}
	pt.placeholderMsg.SetText(message) // 动态更新提示信息
	pt.placeholder.Show()              // 显示占位界面
	pt.list.Hide()                     // 隐藏表格
}

// ShowTable 显示表格
func (pt *PageList) ShowList() {
	pt.placeholder.Hide() // 隐藏占位界面
	pt.list.Show()        // 显示表格
	pt.list.Refresh()     // 刷新表格内容
}

func (pt *PageList) GetList() *widget.List {
	return pt.list
}
