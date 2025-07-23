package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// EllipsisLabel 是自定义组件，支持设置颜色、省略显示、点击复制原文
type EllipsisLabel struct {
	widget.BaseWidget

	fullText      string      // 完整文本
	leftKeep      int         // 左边保留字符数
	rightKeep     int         // 右边保留字符数
	textColor     color.Color // 字体颜色
	textSize      float32
	textStyle     fyne.TextStyle
	display       string        // 当前显示的文本（可能省略）
	hovered       bool          // 是否鼠标悬停，用于变色或提示
	tooltip       *widget.PopUp // Tooltip popup
	tooltipLabel  *canvas.Text
	enableTooltip bool // 私有字段
}

// NewEllipsisLabel 创建一个 EllipsisLabel
func NewEllipsisLabel(text string, left, right int, color color.Color) *EllipsisLabel {
	l := &EllipsisLabel{
		fullText:  text,
		leftKeep:  left,
		rightKeep: right,
		textColor: color,
		textSize:  theme.TextSize(),
		textStyle: fyne.TextStyle{},
	}
	l.ExtendBaseWidget(l)
	return l
}

// SetText 重新设置文本
func (e *EllipsisLabel) SetText(text string) {
	e.fullText = text
	e.Refresh()
}

func (e *EllipsisLabel) SetTextColor(c color.Color) {
	e.textColor = c
	e.Refresh()
}
func (e *EllipsisLabel) SetTextSize(size float32) {
	e.textSize = size
	e.Refresh()
}

func (e *EllipsisLabel) SetTextStyle(style fyne.TextStyle) {
	e.textStyle = style
	e.Refresh()
}

func (e *EllipsisLabel) SetTooltipEnabled(enabled bool) {
	e.enableTooltip = enabled
}

func (e *EllipsisLabel) IsTooltipEnabled() bool {
	return e.enableTooltip
}

// CreateRenderer 实现渲染器
func (e *EllipsisLabel) CreateRenderer() fyne.WidgetRenderer {
	e.updateDisplay()

	textObj := canvas.NewText(e.display, e.textColor)
	textObj.TextSize = e.textSize
	textObj.TextStyle = e.textStyle

	return &ellipsisLabelRenderer{
		label:   e,
		textObj: textObj,
		objects: []fyne.CanvasObject{textObj},
	}
}

// updateDisplay 根据组件尺寸决定是否省略
func (e *EllipsisLabel) updateDisplay() {
	textLen := len([]rune(e.fullText))
	if textLen <= e.leftKeep+e.rightKeep+3 {
		e.display = e.fullText
		return
	}

	runes := []rune(e.fullText)
	left := ""
	right := ""
	if e.leftKeep > 0 {
		left = string(runes[:e.leftKeep])
	}
	if e.rightKeep > 0 {
		right = string(runes[textLen-e.rightKeep:])
	}
	e.display = left + "..." + right
}

// 鼠标点击事件：复制原文
func (e *EllipsisLabel) Tapped(_ *fyne.PointEvent) {
	fyne.CurrentApp().Clipboard().SetContent(e.fullText)
}

// 支持 desktop 鼠标 hover 状态
func (e *EllipsisLabel) MouseIn(ev *desktop.MouseEvent) {
	e.hovered = true
	if e.enableTooltip && e.tooltip == nil {
		e.createTooltip()
	}
	e.showTooltipAt(ev.Position)
}

func (e *EllipsisLabel) MouseMoved(ev *desktop.MouseEvent) {
	if e.tooltip != nil && e.enableTooltip {
		e.showTooltipAt(ev.Position)
	}
}

func (e *EllipsisLabel) MouseOut() {
	e.hovered = false
	e.hideTooltip()
}

func (e *EllipsisLabel) TappedSecondary(ev *fyne.PointEvent) {
	menu := fyne.NewMenu("",
		fyne.NewMenuItem("复制全部内容", func() {
			fyne.CurrentApp().Clipboard().SetContent(e.fullText)
		}),
	)
	widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().AllWindows()[0].Canvas(), ev.Position)
}

func (e *EllipsisLabel) createTooltip() {
	if e.tooltip != nil {
		return
	}

	win := fyne.CurrentApp().Driver().AllWindows()[0]
	e.tooltipLabel = canvas.NewText(e.fullText, theme.ForegroundColor())
	e.tooltipLabel.TextSize = e.textSize
	e.tooltipLabel.TextStyle = fyne.TextStyle{Italic: true}

	bg := container.NewMax(e.tooltipLabel)
	e.tooltip = widget.NewPopUp(bg, win.Canvas())
}

func (e *EllipsisLabel) showTooltipAt(pos fyne.Position) {
	if e.tooltip == nil {
		return
	}
	e.tooltip.ShowAtPosition(fyne.NewPos(pos.X+10, pos.Y+20))
}

func (e *EllipsisLabel) hideTooltip() {
	if e.tooltip != nil {
		e.tooltip.Hide()
		e.tooltip = nil
		e.tooltipLabel = nil
	}
}

type ellipsisLabelRenderer struct {
	label   *EllipsisLabel
	textObj *canvas.Text
	objects []fyne.CanvasObject
}

func (r *ellipsisLabelRenderer) Layout(size fyne.Size) {
	r.textObj.Resize(size)
}

func (r *ellipsisLabelRenderer) MinSize() fyne.Size {
	return r.textObj.MinSize()
}

func (r *ellipsisLabelRenderer) Refresh() {
	r.label.updateDisplay()

	r.textObj.Text = r.label.display
	r.textObj.Color = r.label.textColor
	r.textObj.TextSize = r.label.textSize
	r.textObj.TextStyle = r.label.textStyle

	canvas.Refresh(r.textObj)
}

func (r *ellipsisLabelRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *ellipsisLabelRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *ellipsisLabelRenderer) Destroy() {}
