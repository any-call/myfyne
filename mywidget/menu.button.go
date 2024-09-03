package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/myfyne"
	"image/color"
)

// MenuButton 定义了一个自定义的按钮控件
type MenuButton struct {
	widget.BaseWidget
	Text              string
	OnClick           func()
	isSelected        bool
	isHovered         bool
	textColor         color.Color
	hoverTextColor    color.Color
	selectedTextColor color.Color
	bgColor           color.Color
	hoverBgColor      color.Color
	selectedBgColor   color.Color
	textSize          *float32
	textAlign         fyne.TextAlign
	padding           myfyne.EdgeInset // 新增的 padding 属性
}

// NewMenuButton 创建一个新的 MenuButton
func NewMenuButton(text string, onClick func()) *MenuButton {
	button := &MenuButton{
		Text:      text,
		OnClick:   onClick,
		textAlign: fyne.TextAlignCenter, // 默认对齐方式
	}
	button.ExtendBaseWidget(button)
	return button
}

// SetTextColor 设置正常状态的字体颜色
func (b *MenuButton) SetTextColor(c color.Color) *MenuButton {
	b.textColor = c
	b.Refresh()
	return b
}

// SetBgColor 设置正常状态的背景颜色
func (b *MenuButton) SetBgColor(c color.Color) *MenuButton {
	b.bgColor = c
	b.Refresh()
	return b
}

// SetHoverTextColor 设置 Hover 状态的字体颜色
func (b *MenuButton) SetHoverTextColor(c color.Color) *MenuButton {
	b.hoverTextColor = c
	b.Refresh()
	return b
}

// SetHoverBgColor 设置 Hover 状态的背景颜色
func (b *MenuButton) SetHoverBgColor(c color.Color) *MenuButton {
	b.hoverBgColor = c
	b.Refresh()
	return b
}

// SetSelectedTextColor 设置选中状态的字体颜色
func (b *MenuButton) SetSelectedTextColor(c color.Color) *MenuButton {
	b.selectedTextColor = c
	b.Refresh()
	return b
}

// SetSelectedBgColor 设置选中状态的背景颜色
func (b *MenuButton) SetSelectedBgColor(c color.Color) *MenuButton {
	b.selectedBgColor = c
	b.Refresh()
	return b
}

// SetTextSize 设置文本的字体大小
func (b *MenuButton) SetTextSize(size float32) *MenuButton {
	b.textSize = &size
	b.Refresh()
	return b
}

// SetTextAlign 设置文本对齐方式
func (b *MenuButton) SetTextAlign(align fyne.TextAlign) *MenuButton {
	b.textAlign = align
	b.Refresh()
	return b
}

// SetPadding 设置按钮内边距
func (b *MenuButton) SetPadding(p myfyne.EdgeInset) *MenuButton {
	b.padding = p
	b.Refresh()
	return b
}

// GetPadding 获取按钮内边距
func (b *MenuButton) GetPadding() myfyne.EdgeInset {
	return b.padding
}

// SetIsSelected 设置按钮的选中状态
func (b *MenuButton) SetIsSelected(selected bool) *MenuButton {
	b.isSelected = selected
	b.Refresh()
	return b
}

// GetIsSelected 获取按钮的选中状态
func (b *MenuButton) GetIsSelected() bool {
	return b.isSelected
}

// GetTextColor 获取正常状态的字体颜色
func (b *MenuButton) GetTextColor() color.Color {
	if b.textColor != nil {
		return b.textColor
	}
	return theme.Color(theme.ColorNameForeground)
}

// GetHoverTextColor 获取 Hover 状态的字体颜色
func (b *MenuButton) GetHoverTextColor() color.Color {
	if b.hoverTextColor != nil {
		return b.hoverTextColor
	}
	return theme.Color(theme.ColorNameForeground)
}

// GetSelectedTextColor 获取选中状态的字体颜色
func (b *MenuButton) GetSelectedTextColor() color.Color {
	if b.selectedTextColor != nil {
		return b.selectedTextColor
	}
	return theme.Color(theme.ColorNameForeground)
}

func (b *MenuButton) GetBgColor() color.Color {
	if b.bgColor != nil {
		return b.bgColor
	}
	return theme.Color(theme.ColorNameBackground)
}

func (b *MenuButton) GetSelectedBgColor() color.Color {
	if b.selectedBgColor != nil {
		return b.selectedBgColor
	}
	return theme.Color(theme.ColorNameBackground)
}

func (b *MenuButton) GetHoverBgColor() color.Color {
	if b.hoverBgColor != nil {
		return b.hoverBgColor
	}
	return theme.Color(theme.ColorNameBackground)
}

// GetTextSize 获取文本的字体大小
func (b *MenuButton) GetTextSize() float32 {
	if b.textSize != nil {
		return *b.textSize
	}
	return theme.TextSize()
}

// GetTextAlign 获取文本的对齐方式
func (b *MenuButton) GetTextAlign() fyne.TextAlign {
	return b.textAlign
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
	label.TextSize = b.GetTextSize()
	label.Alignment = b.GetTextAlign()

	objects := []fyne.CanvasObject{background, label}

	return &menuButtonRenderer{
		button:     b,
		background: background,
		label:      label,
		objects:    objects,
	}
}

func (b *MenuButton) getTextColor() color.Color {
	if b.isSelected {
		return b.GetSelectedTextColor()
	} else if b.isHovered {
		return b.GetHoverTextColor()
	}
	return b.GetTextColor()
}

func (b *MenuButton) getBackgroundColor() color.Color {
	if b.isSelected {
		return b.GetSelectedBgColor()
	} else if b.isHovered {
		return b.GetHoverBgColor()
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
	padding := r.button.GetPadding()
	labelSize := r.label.MinSize()

	// Adjust label size considering padding
	labelSize.Width += padding.Left + padding.Right
	labelSize.Height += padding.Top + padding.Bottom

	r.background.Resize(size)
	r.label.Resize(labelSize)
	r.label.Move(fyne.NewPos(padding.Left, padding.Top))
}

func (r *menuButtonRenderer) MinSize() fyne.Size {
	padding := r.button.GetPadding()
	labelSize := r.label.MinSize()

	// Calculate min size including padding
	return fyne.NewSize(
		labelSize.Width+padding.Left+padding.Right,
		labelSize.Height+padding.Top+padding.Bottom,
	)
}

func (r *menuButtonRenderer) Refresh() {
	r.background.FillColor = r.button.getBackgroundColor()
	r.label.Color = r.button.getTextColor()
	r.label.TextSize = r.button.GetTextSize()
	r.label.Alignment = r.button.GetTextAlign()
	canvas.Refresh(r.button)
}

func (r *menuButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *menuButtonRenderer) Destroy() {}
