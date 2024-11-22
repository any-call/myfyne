package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PageList 是基于 widget.BaseWidget 的自定义控件
type PageGrid struct {
	widget.BaseWidget
	gridWrap       *widget.GridWrap
	placeholder    *fyne.Container
	placeholderMsg *widget.Label
	onRefresh      func()                   // 刷新数据的回调
	onConfig       func(t *widget.GridWrap) // 配置 Table 的回调
}

// NewPageGrid 创建一个新的 PageGrid
// onRefresh 用于无数据时显示的刷新操作，onConfig 用于配置 Table 的回调
func NewPageGrid(onRefresh func(), onConfig func(t *widget.GridWrap)) *PageGrid {
	pt := &PageGrid{
		onRefresh: onRefresh,
		onConfig:  onConfig,
	}

	// 创建表格
	pt.gridWrap = widget.NewGridWrap(
		func() int { return 0 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.GridWrapItemID, object fyne.CanvasObject) {},
	)

	// 使用回调配置 Table
	if pt.onConfig != nil {
		pt.onConfig(pt.gridWrap)
	}

	// 必须调用 .ExtendBaseWidget 来确保控件能够正确渲染
	pt.ExtendBaseWidget(pt)
	return pt
}

// CreateRenderer 用于渲染 PageGrid 控件
func (pt *PageGrid) CreateRenderer() fyne.WidgetRenderer {
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
	stack := container.NewStack(pt.placeholder, pt.gridWrap)

	// 初始显示占位界面
	pt.ShowPlaceholder("")

	return widget.NewSimpleRenderer(stack)
}

// ShowPlaceholder 显示占位界面，动态更新提示信息
func (pt *PageGrid) ShowPlaceholder(message string) {
	if message == "" {
		message = "没有数据"
	}
	pt.placeholderMsg.SetText(message) // 动态更新提示信息
	pt.placeholder.Show()              // 显示占位界面
	pt.gridWrap.Hide()                 // 隐藏表格
}

// ShowTable 显示表格
func (pt *PageGrid) ShowGrid() {
	pt.placeholder.Hide() // 隐藏占位界面
	pt.gridWrap.Show()    // 显示表格
	pt.gridWrap.Refresh() // 刷新表格内容
}

func (pt *PageGrid) GetGrid() *widget.GridWrap {
	return pt.gridWrap
}
