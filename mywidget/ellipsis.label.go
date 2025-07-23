package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// EllipsisLabel 是自定义组件，支持设置颜色、省略显示、点击复制原文
type EllipsisLabel struct {
	widget.BaseWidget
	fullText       string
	display        string
	left, right    int
	textObj        *canvas.Text
	textColor      color.Color
	textSize       float32
	textStyle      fyne.TextStyle
	enableTooltip  bool
	tooltip        *widget.PopUp
	hovering       bool
	tooltipContent *canvas.Text
}

// NewEllipsisLabel 创建一个 EllipsisLabel
func NewEllipsisLabel(text string, left, right int, color color.Color) *EllipsisLabel {
	l := &EllipsisLabel{
		fullText:  text,
		left:      left,
		right:     right,
		textColor: color,
		textSize:  theme.TextSize(),
		textStyle: fyne.TextStyle{},
	}
	l.textObj = canvas.NewText("", color)
	l.updateDisplay()
	l.ExtendBaseWidget(l)
	return l
}

// CreateRenderer 实现渲染器
func (e *EllipsisLabel) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.textObj)
}

// SetText 重新设置文本
func (e *EllipsisLabel) SetText(text string) {
	e.fullText = text
	e.Refresh()
}

func (e *EllipsisLabel) GetText() string {
	return e.fullText
}

func (e *EllipsisLabel) SetTextColor(c color.Color) {
	e.textColor = c
	e.textObj.Color = c
	canvas.Refresh(e.textObj)
}
func (e *EllipsisLabel) SetTextSize(size float32) {
	e.textSize = size
	e.textObj.TextSize = size
	canvas.Refresh(e.textObj)
}

func (e *EllipsisLabel) SetTextStyle(style fyne.TextStyle) {
	e.textStyle = style
	e.textObj.TextStyle = style
	canvas.Refresh(e.textObj)
}

func (e *EllipsisLabel) SetTooltipEnabled(enabled bool) {
	e.enableTooltip = enabled
}

func (e *EllipsisLabel) IsTooltipEnabled() bool {
	return e.enableTooltip
}

func (e *EllipsisLabel) MinSize() fyne.Size {
	return e.textObj.MinSize()
}

// updateDisplay 根据组件尺寸决定是否省略
func (e *EllipsisLabel) updateDisplay() {
	txt := []rune(e.fullText)
	if len(txt) <= e.left+e.right {
		e.display = e.fullText
	} else {
		e.display = string(txt[:e.left]) + "..." + string(txt[len(txt)-e.right:])
	}
	e.textObj.Text = e.display
	e.textObj.TextStyle = e.textStyle
	e.textObj.TextSize = e.textSize
	e.textObj.Color = e.textColor
	canvas.Refresh(e.textObj)
}

// 支持 desktop 鼠标 hover 状态
func (e *EllipsisLabel) MouseIn(ev *desktop.MouseEvent) {
	e.hovering = true
	if !e.enableTooltip {
		return
	}
	// 可选：立即显示 tooltip，而不是等 MouseMoved()
	if e.tooltip == nil {
		text := canvas.NewText(e.fullText, e.textColor)
		text.TextSize = e.textSize
		text.TextStyle = e.textStyle

		e.tooltipContent = text
		e.tooltip = widget.NewPopUp(text, fyne.CurrentApp().Driver().CanvasForObject(e))
	}
	pos := ev.Position
	e.tooltip.Move(fyne.NewPos(pos.X+10, pos.Y+10))
	e.tooltip.Show()

}

func (e *EllipsisLabel) MouseMoved(ev *desktop.MouseEvent) {
	if !e.enableTooltip {
		return
	}

	if e.tooltip == nil {
		text := canvas.NewText(e.fullText, e.textColor)
		text.TextSize = e.textSize
		text.TextStyle = e.textStyle

		e.tooltipContent = text
		e.tooltip = widget.NewPopUp(text, fyne.CurrentApp().Driver().CanvasForObject(e))
		e.tooltip.Show() // 只初始化时 Show 一次
	}

	// 只更新位置，不重复调用 Show()
	if e.tooltip.Visible() {
		e.tooltip.Move(fyne.NewPos(ev.Position.X+10, ev.Position.Y+10))
	}
}

func (e *EllipsisLabel) MouseOut() {
	e.hovering = false
	if e.tooltip != nil {
		e.tooltip.Hide()
	}
}

func (e *EllipsisLabel) showTooltipAt(pos fyne.Position) {
	if e.tooltip == nil {
		return
	}
	e.tooltip.ShowAtPosition(fyne.NewPos(pos.X+10, pos.Y+20))
}
