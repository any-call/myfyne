package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PageTable 是基于 widget.BaseWidget 的自定义控件
type PageTree struct {
	widget.BaseWidget
	tree           *widget.Tree
	placeholder    *fyne.Container
	placeholderMsg *widget.Label
	onRefresh      func()               // 刷新数据的回调
	onConfig       func(t *widget.Tree) // 配置 Table 的回调
}

// NewPageTable 创建一个新的 PageTable
// onRefresh 用于无数据时显示的刷新操作，onConfig 用于配置 Table 的回调
func NewPageTree(onRefresh func(), onConfig func(t *widget.Tree)) *PageTree {
	pt := &PageTree{
		onRefresh: onRefresh,
		onConfig:  onConfig,
	}

	// 创建表格
	pt.tree = widget.NewTreeWithStrings(nil)
	pt.tree.HideSeparators = false
	// 使用回调配置 Table
	if pt.onConfig != nil {
		pt.onConfig(pt.tree)
	}

	// 必须调用 .ExtendBaseWidget 来确保控件能够正确渲染
	pt.ExtendBaseWidget(pt)
	return pt
}

// CreateRenderer 用于渲染 PageTable 控件
func (pt *PageTree) CreateRenderer() fyne.WidgetRenderer {
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
	stack := container.NewStack(pt.placeholder, pt.tree)

	// 初始显示占位界面
	pt.ShowPlaceholder("")

	return widget.NewSimpleRenderer(stack)
}

// ShowPlaceholder 显示占位界面，动态更新提示信息
func (pt *PageTree) ShowPlaceholder(message string) {
	if message == "" {
		message = "没有数据"
	}
	pt.placeholderMsg.SetText(message) // 动态更新提示信息
	pt.placeholder.Show()              // 显示占位界面
	pt.tree.Hide()                     // 隐藏表格
}

// ShowTable 显示表格
func (pt *PageTree) ShowTree() {
	pt.placeholder.Hide() // 隐藏占位界面
	pt.tree.Show()        // 显示表格
	pt.tree.Refresh()     // 刷新表格内容
}

func (pt *PageTree) GetTree() *widget.Tree {
	return pt.tree
}
