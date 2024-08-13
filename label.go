package myfyne

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/gobase/frame/myctrl"
	"image/color"
)

type Label struct {
	widget.BaseWidget
	title           string
	foregroundColor color.Color
	backgroundColor color.Color
	fontSize        float32
	alignment       fyne.TextAlign
	fixedWidth      float32 //<=0 则说明是自适应的宽
	fixedHeight     float32 //<=0 则说明是自适应的高
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
	text := canvas.NewText(c.title, myctrl.ObjFun(func() color.Color {
		if c.foregroundColor == nil {
			return theme.Color(theme.ColorNameForeground)
		}

		return c.foregroundColor
	}))

	text.Alignment = c.alignment

	if c.foregroundColor != nil {
		text.Color = c.foregroundColor
	}

	if c.fontSize > 0 {
		text.TextSize = c.fontSize
	} else {
		text.TextSize = theme.TextSize() // 使用默认字体大小
	}

	background := canvas.NewRectangle(myctrl.ObjFun(func() color.Color {
		if c.backgroundColor == nil {
			return theme.Color(theme.ColorNameBackground)
		}

		return c.backgroundColor
	}))

	if c.backgroundColor != nil {
		background.FillColor = c.backgroundColor
	}

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

func (c *Label) SetColor(color color.Color) *Label {
	c.foregroundColor = color
	c.Refresh()
	return c
}

func (c *Label) SetBackgroundColor(color color.Color) *Label {
	c.backgroundColor = color
	c.Refresh()
	return c
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

type labelRenderer struct {
	label      *Label
	text       *canvas.Text
	background *canvas.Rectangle
}

func (r *labelRenderer) Layout(size fyne.Size) {
	fmt.Println("enter  layout：", r.label.alignment)
	// 根据固定宽度或高度调整文本大小
	textSize := r.text.MinSize()

	if r.label.fixedWidth > 0 {
		textSize.Width = r.label.fixedWidth
	} else {
		// 自动调整宽度以适应文本
		textSize.Width = r.text.MinSize().Width
	}

	if r.label.fixedHeight > 0 {
		textSize.Height = r.label.fixedHeight
	} else {
		// 自动调整高度以适应文本
		textSize.Height = r.text.MinSize().Height
	}

	r.text.Resize(textSize)
	r.background.Resize(size)
	switch r.label.alignment {
	case fyne.TextAlignCenter:
		r.text.Move(fyne.NewPos((size.Width-textSize.Width)/2, (size.Height-textSize.Height)/2))
		break
	case fyne.TextAlignTrailing:
		r.text.Move(fyne.NewPos(size.Width-textSize.Width, (size.Height-textSize.Height)/2))
		break
	default: // fyne.TextAlignLeading
		r.text.Move(fyne.NewPos(0, (size.Height-textSize.Height)/2))
		break
	}
}

func (r *labelRenderer) MinSize() fyne.Size {
	textSize := r.text.MinSize()

	// 如果设置了固定宽度或高度，则使用固定值
	if r.label.fixedWidth > 0 {
		textSize.Width = r.label.fixedWidth
	} else {
		// 自适应宽度
		textSize.Width = r.text.MinSize().Width
	}

	if r.label.fixedHeight > 0 {
		textSize.Height = r.label.fixedHeight
	} else {
		// 自适应高度
		textSize.Height = r.text.MinSize().Height
	}

	return fyne.NewSize(textSize.Width+10, textSize.Height+10) // 增加一些填充以适应背景
}

func (r *labelRenderer) Refresh() {
	fmt.Println("enter refresh..labelRenderer..")
	r.text.Text = r.label.title
	if r.label.foregroundColor != nil {
		r.text.Color = r.label.foregroundColor
	} else {
		r.text.Color = theme.Color(theme.ColorNameBackground)
	}

	if r.label.backgroundColor != nil {
		r.background.FillColor = r.label.backgroundColor
	} else {
		r.background.FillColor = theme.Color(theme.ColorNameBackground)
	}

	if r.label.fontSize > 0 {
		r.text.TextSize = r.label.fontSize
	} else {
		r.text.TextSize = theme.TextSize()
	}

	r.text.Alignment = r.label.alignment
	r.background.Refresh()
	r.text.Refresh()
}

func (r *labelRenderer) Destroy() {}

func (r *labelRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.text}
}
