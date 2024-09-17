package mywidget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PageTable 是基于 widget.BaseWidget 的自定义控件
type PageTable struct {
	widget.BaseWidget
	table          *widget.Table
	placeholder    *fyne.Container
	placeholderMsg *widget.Label
	onRefresh      func()                // 刷新数据的回调
	onConfig       func(t *widget.Table) // 配置 Table 的回调
}

// NewPageTable 创建一个新的 PageTable
// onRefresh 用于无数据时显示的刷新操作，onConfig 用于配置 Table 的回调
func NewPageTable(onRefresh func(), onConfig func(t *widget.Table)) *PageTable {
	pt := &PageTable{
		onRefresh: onRefresh,
		onConfig:  onConfig,
	}

	// 创建表格
	pt.table = widget.NewTable(
		func() (int, int) { return 0, 0 }, // 默认没有行和列
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, o fyne.CanvasObject) {})

	// 使用回调配置 Table
	if pt.onConfig != nil {
		pt.onConfig(pt.table)
	}

	// 必须调用 .ExtendBaseWidget 来确保控件能够正确渲染
	pt.ExtendBaseWidget(pt)
	return pt
}

// CreateRenderer 用于渲染 PageTable 控件
func (pt *PageTable) CreateRenderer() fyne.WidgetRenderer {
	// 占位界面
	pt.placeholderMsg = widget.NewLabel("没有数据")
	pt.placeholderMsg.Alignment = fyne.TextAlignCenter
	refreshButton := widget.NewButton("点击刷新", func() {
		// 点击刷新时调用外部提供的刷新回调
		if pt.onRefresh != nil {
			pt.onRefresh()
		}
	})

	pt.placeholder = container.NewVBox(pt.placeholderMsg, refreshButton)
	stack := container.NewStack(pt.placeholder, pt.table)

	// 初始显示占位界面
	pt.ShowPlaceholder("")

	return widget.NewSimpleRenderer(stack)
}

// ShowPlaceholder 显示占位界面，动态更新提示信息
func (pt *PageTable) ShowPlaceholder(message string) {
	if message == "" {
		message = "没有数据"
	}
	pt.placeholderMsg.SetText(message) // 动态更新提示信息
	pt.placeholder.Show()              // 显示占位界面
	pt.table.Hide()                    // 隐藏表格
}

// ShowTable 显示表格
func (pt *PageTable) ShowTable() {
	pt.placeholder.Hide() // 隐藏占位界面
	pt.table.Show()       // 显示表格
	pt.table.Refresh()    // 刷新表格内容
}

func (pt *PageTable) Show() {
	fmt.Println("enter page table show")
	pt.BaseWidget.Show()
}

func (pt *PageTable) Hide() {
	fmt.Println("enter page table hide")
	pt.BaseWidget.Hide()
}
