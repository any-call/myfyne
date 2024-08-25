package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/myfyne/myfynebase"
	"image/color"
)

type Label struct {
	widget.BaseWidget
	title           string
	foregroundColor color.Color
	backgroundColor color.Color
	fontSize        float32
	alignment       fyne.TextAlign
	textStyle       fyne.TextStyle
	fixedWidth      float32              //<=0 则说明是自适应的宽
	fixedHeight     float32              //<=0 则说明是自适应的高
	padding         myfynebase.EdgeInset //定义内间距
}

func NewLabel(text string) *Label {
	label := &Label{
		title:     text,
		fontSize:  theme.TextSize(),
		alignment: fyne.TextAlignLeading,
	}
	label.ExtendBaseWidget(label)

	return label
}

func (c *Label) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(c.title, c.Color())
	text.Alignment = c.alignment
	text.TextSize = c.FontSize()
	text.TextStyle = c.textStyle
	background := canvas.NewRectangle(c.BackgroundColor())
	return &labelRenderer{
		label:      c,
		text:       text,
		background: background,
	}
}

func (c *Label) SetFontSize(size float32) *Label {
	if size != c.fontSize {
		c.fontSize = size
		c.Refresh()
	}
	return c
}

func (c *Label) FontSize() float32 {
	if c.fontSize <= 0 {
		return theme.TextSize()
	}

	return c.fontSize
}

func (c *Label) SetPadding(p myfynebase.EdgeInset) *Label {
	c.padding = p
	c.Refresh()
	return c
}

func (c *Label) GetPadding() myfynebase.EdgeInset {
	return c.padding
}

func (c *Label) SetColor(color color.Color) *Label {
	c.foregroundColor = color
	c.Refresh()
	return c
}

func (c *Label) Color() color.Color {
	if c.foregroundColor == nil {
		return theme.Color(theme.ColorNameForeground)
	}

	return c.foregroundColor
}

func (c *Label) SetBackgroundColor(color color.Color) *Label {
	c.backgroundColor = color
	c.Refresh()
	return c
}

func (c *Label) BackgroundColor() color.Color {
	if c.backgroundColor == nil {
		return theme.Color(theme.ColorNameBackground)
	}

	return c.backgroundColor
}

func (c *Label) SetFixedSize(fixedWidth, fixedHeight float32) *Label {
	c.fixedWidth = fixedWidth
	c.fixedHeight = fixedHeight
	c.Refresh()

	return c
}

func (c *Label) SetAlign(align fyne.TextAlign) *Label {
	c.alignment = align
	c.Refresh()
	return c
}

func (c *Label) Alignment() fyne.TextAlign {
	return c.alignment
}

func (c *Label) SetTextStyle(style fyne.TextStyle) *Label {
	c.textStyle = style
	c.Refresh()
	return c
}

type labelRenderer struct {
	label      *Label
	text       *canvas.Text
	background *canvas.Rectangle
}

func (r *labelRenderer) Layout(size fyne.Size) {
	// 根据固定宽度或高度调整文本大小
	textSize := r.text.MinSize() //文本实际大小
	if r.label.fixedWidth > 0 {
		textSize.Width = r.label.fixedWidth
	}

	if r.label.fixedHeight > 0 {
		textSize.Height = r.label.fixedHeight
	}

	r.text.Resize(textSize)
	r.background.Resize(size)
	switch r.label.alignment {
	case fyne.TextAlignCenter:
		r.text.Move(myfynebase.ChildPosition(myfynebase.PositionCenter, size, textSize))
		break
	case fyne.TextAlignTrailing:
		r.text.Move(myfynebase.ChildPosition(myfynebase.PositionCenterRight, size, textSize))
		break
	default: // fyne.TextAlignLeading
		r.text.Move(myfynebase.ChildPosition(myfynebase.PositionCenterLeft, size, textSize))
		break
	}
}

func (r *labelRenderer) MinSize() fyne.Size {
	textSize := r.text.MinSize()

	// 如果设置了固定宽度或高度，则使用固定值
	if r.label.fixedWidth > 0 {
		textSize.Width = r.label.fixedWidth
	} else {
		textSize.Width += r.label.padding.Left + r.label.padding.Right
	}

	if r.label.fixedHeight > 0 {
		textSize.Height = r.label.fixedHeight
	} else {
		textSize.Height += r.label.padding.Top + r.label.padding.Bottom
	}

	return textSize
}

func (r *labelRenderer) Refresh() {
	r.text.Text = r.label.title
	r.text.Color = r.label.Color()
	r.text.TextSize = r.label.FontSize()
	r.text.Alignment = r.label.Alignment()
	r.text.TextStyle = r.label.textStyle
	r.background.FillColor = r.label.BackgroundColor()

	r.background.Refresh()
	r.text.Refresh()
}

func (r *labelRenderer) Destroy() {}

func (r *labelRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.text}
}
