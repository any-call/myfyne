package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// MenuButton 定义了一个自定义的按钮控件
type MenuButton struct {
	widget.BaseWidget
	Text              string
	OnClick           func()
	isSelected        bool
	isHovered         bool
	textSize          float32
	textColor         *color.Color
	hoverTextColor    *color.Color
	selectedTextColor *color.Color
	bgColor           *color.Color
	hoverBgColor      *color.Color
	selectedBgColor   *color.Color
}

// NewMenuButton 创建一个新的 MenuButton
func NewMenuButton(text string, onClick func()) *MenuButton {
	button := &MenuButton{
		Text:    text,
		OnClick: onClick,
	}
	button.ExtendBaseWidget(button)
	return button
}

// SetNormalColors 设置正常状态的字体颜色和背景颜色
func (b *MenuButton) SetTextColor(textColor color.Color) *MenuButton {
	b.textColor = &textColor
	b.Refresh()
	return b
}

func (b *MenuButton) SetHoverTextColor(textColor color.Color) *MenuButton {
	b.hoverTextColor = &textColor
	b.Refresh()
	return b
}

func (b *MenuButton) SetSelectTextColor(textColor color.Color) *MenuButton {
	b.selectedTextColor = &textColor
	b.Refresh()
	return b
}

func (b *MenuButton) SetBgColor(textColor color.Color) *MenuButton {
	b.bgColor = &textColor
	b.Refresh()
	return b
}

func (b *MenuButton) SetHoverBgColor(textColor color.Color) *MenuButton {
	b.hoverBgColor = &textColor
	b.Refresh()
	return b
}

func (b *MenuButton) SetSelectBgColor(textColor color.Color) *MenuButton {
	b.selectedBgColor = &textColor
	b.Refresh()
	return b
}

func (b *MenuButton) SetTextSize(size float32) *MenuButton {
	b.textSize = size
	b.Refresh()
	return b
}

func (b *MenuButton) SetSelectState(flag bool) *MenuButton {
	b.isSelected = flag
	b.Refresh()
	return b
}

func (b *MenuButton) GetTextColor() color.Color {
	if b.textColor != nil {
		return *b.textColor
	}

	return theme.Color(theme.ColorNameForeground)
}

func (b *MenuButton) GetHoverTextColor() color.Color {
	if b.hoverTextColor != nil {
		return *b.hoverTextColor
	}
	return theme.Color(theme.ColorNameHover)
}

func (b *MenuButton) GetSelectedTextColor() color.Color {
	if b.selectedTextColor != nil {
		return *b.selectedTextColor
	}
	return theme.Color(theme.ColorNameSelection)
}

func (b *MenuButton) GetBgColor() color.Color {
	if b.bgColor != nil {
		return *b.bgColor
	}

	return theme.Color(theme.ColorNameBackground)
}

func (b *MenuButton) GetHoverBgColor() color.Color {
	if b.hoverBgColor != nil {
		return *b.hoverBgColor
	}

	return theme.Color(theme.ColorNameBackground)
}

func (b *MenuButton) GetSelectBgColor() color.Color {
	if b.selectedBgColor != nil {
		return *b.selectedBgColor
	}

	return theme.Color(theme.ColorNameBackground)
}

func (b *MenuButton) GetTextSize() float32 {
	if b.textSize > 0 {
		return b.textSize
	}
	return theme.TextSize()
}

func (b *MenuButton) GetSelectState() bool {
	return b.isSelected
}

// Tapped 实现点击事件
func (b *MenuButton) Tapped(*fyne.PointEvent) {
	if b.OnClick != nil {
		b.OnClick()
	}
	b.Refresh()
}

// MouseIn 处理鼠标进入事件
func (b *MenuButton) MouseIn(*desktop.MouseEvent) {
	b.isHovered = true
	b.Refresh()
}

// MouseOut 处理鼠标离开事件
func (b *MenuButton) MouseOut() {
	b.isHovered = false
	b.Refresh()
}

// MouseMoved 处理鼠标移动事件
func (b *MenuButton) MouseMoved(*desktop.MouseEvent) {}

// CreateRenderer 定义控件的绘制逻辑
func (b *MenuButton) CreateRenderer() fyne.WidgetRenderer {
	background := canvas.NewRectangle(b.getBackgroundColor())
	label := canvas.NewText(b.Text, b.getTextColor())
	label.Alignment = fyne.TextAlignLeading

	objects := []fyne.CanvasObject{background, label}

	return &menuButtonRenderer{
		button:     b,
		background: background,
		label:      label,
		objects:    objects,
	}
}

func (b *MenuButton) getTextColor() color.Color {
	if b.isHovered {
		return b.GetHoverTextColor()
	} else if b.isSelected {
		return b.GetSelectedTextColor()
	}
	return b.getTextColor()
}

func (b *MenuButton) getBackgroundColor() color.Color {
	if b.isHovered {
		return b.GetHoverBgColor()
	} else if b.isSelected {
		return b.GetSelectBgColor()
	}
	return b.GetBgColor()
}

// menuButtonRenderer 实现控件的渲染器
type menuButtonRenderer struct {
	button     *MenuButton
	background *canvas.Rectangle
	label      *canvas.Text
	objects    []fyne.CanvasObject
}

func (r *menuButtonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.label.Resize(size)
	r.label.Move(fyne.NewPos(0, size.Height/2-r.label.MinSize().Height/2))
}

func (r *menuButtonRenderer) MinSize() fyne.Size {
	return r.label.MinSize()
}

func (r *menuButtonRenderer) Refresh() {
	r.background.FillColor = r.button.getBackgroundColor()
	r.label.Color = r.button.getTextColor()
	r.label.TextSize = r.button.GetTextSize()
	canvas.Refresh(r.button)
}

func (r *menuButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *menuButtonRenderer) Destroy() {}
