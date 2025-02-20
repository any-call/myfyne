package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"time"
)

// CopyableContainer 泛型，支持任意 fyne.CanvasObject
type CopyableContainer[T fyne.CanvasObject] struct {
	widget.BaseWidget
	content T      // 内部内容对象
	text    string // 内部文本
}

// NewCopyableContainer 创建实例
func NewCopyableContainer[T fyne.CanvasObject](content T, text string) *CopyableContainer[T] {
	c := &CopyableContainer[T]{
		content: content,
		text:    text,
	}
	c.ExtendBaseWidget(c)
	return c
}

// CreateRenderer 让容器大小适应内容
func (c *CopyableContainer[T]) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.content)
}

// GetText 获取当前复制文本
func (c *CopyableContainer[T]) GetCopyText() string {
	return c.text
}

// SetText 设置复制文本
func (c *CopyableContainer[T]) SetCopyText(newText string) {
	c.text = newText
}

// GetContent 获取内部组件
func (c *CopyableContainer[T]) GetContent() T {
	return c.content
}

// Tapped 处理点击复制
func (c *CopyableContainer[T]) Tapped(*fyne.PointEvent) {
	clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
	clipboard.SetContent(c.text)
	c.showToast("已复制")
}

// showToast 显示 "已复制" 提示
func (c *CopyableContainer[T]) showToast(msg string) {
	win := fyne.CurrentApp().Driver().AllWindows()[0] // 自动获取当前窗口
	if win == nil {
		return
	}

	toast := canvas.NewText(msg, theme.Color(theme.ColorNameForeground))
	toast.TextSize = 14
	toast.Alignment = fyne.TextAlignCenter

	overlay := container.NewWithoutLayout(toast)
	overlay.Move(fyne.NewPos((win.Canvas().Size().Width-toast.MinSize().Width)/2, 0)) // 顶部居中
	win.Canvas().Overlays().Add(overlay)

	go func() {
		time.Sleep(1 * time.Second) // 显示 1 秒
		win.Canvas().Overlays().Remove(overlay)
	}()
}
