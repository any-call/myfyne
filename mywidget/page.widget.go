package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PageWidget 是基于 widget.BaseWidget 的自定义控件
type PageWidget[T fyne.CanvasObject] struct {
	widget.BaseWidget
	content        T
	placeholder    *fyne.Container
	placeholderMsg *widget.Label
	onRefresh      func() // 刷新数据的回调
}

// NewPageList 创建一个新的 PageList
// onRefresh 用于无数据时显示的刷新操作，onConfig 用于配置 Table 的回调
func NewPageWidget[T fyne.CanvasObject](onRefresh func(), w T) *PageWidget[T] {
	pt := &PageWidget[T]{
		onRefresh: onRefresh,
		content:   w,
	}

	// 必须调用 .ExtendBaseWidget 来确保控件能够正确渲染
	pt.ExtendBaseWidget(pt)
	return pt
}

// CreateRenderer 用于渲染 PageList 控件
func (pt *PageWidget[T]) CreateRenderer() fyne.WidgetRenderer {
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
	stack := container.NewStack(pt.placeholder, pt.content)

	// 初始显示占位界面
	pt.ShowPlaceholder("")

	return widget.NewSimpleRenderer(stack)
}

// ShowPlaceholder 显示占位界面，动态更新提示信息
func (pt *PageWidget[T]) ShowPlaceholder(message string) {
	if message == "" {
		message = "没有数据"
	}
	pt.placeholderMsg.SetText(message) // 动态更新提示信息
	pt.placeholder.Show()              // 显示占位界面
	pt.content.Hide()                  // 隐藏表格
}

// ShowTable 显示表格
func (pt *PageWidget[T]) ShowContent() {
	pt.placeholder.Hide() // 隐藏占位界面
	pt.content.Show()     // 显示表格
	pt.content.Refresh()  // 刷新表格内容
}

func (pt *PageWidget[T]) GetContent() T {
	return pt.content
}
