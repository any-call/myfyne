package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/myfyne/myfyneKit"
	"image/color"
)

// SizedBox 是一个自定义控件，用于设置固定大小的容器，并提供背景色
type SizedBox struct {
	widget.BaseWidget
	width      float32
	height     float32
	padding    myfyneKit.EdgeInset
	background color.Color
	child      fyne.CanvasObject
}

func NewFixedWidthBox(width float32, background color.Color, child fyne.CanvasObject) *SizedBox {
	return NewSizedBox(fyne.NewSize(width, myfyneKit.Infinity), background, child)
}

func NewFixedHeightBox(height float32, background color.Color, child fyne.CanvasObject) *SizedBox {
	return NewSizedBox(fyne.NewSize(myfyneKit.Infinity, height), background, child)
}

// NewSizedBox 创建一个新的 SizedBox 实例
func NewSizedBox(size fyne.Size, background color.Color, child fyne.CanvasObject) *SizedBox {
	box := &SizedBox{
		width:      size.Width,
		height:     size.Height,
		background: background,
		child:      child,
	}
	box.ExtendBaseWidget(box)
	return box
}

// SetWidth 设置宽度
func (b *SizedBox) SetWidth(width float32) *SizedBox {
	b.width = width
	b.Refresh()
	return b
}

// GetWidth 获取宽度
func (b *SizedBox) GetWidth() float32 {
	if b.width >= 0 {
		return b.width
	}

	return myfyneKit.Infinity
}

func (b *SizedBox) SetPadding(padding myfyneKit.EdgeInset) *SizedBox {
	b.padding = padding
	b.Refresh()
	return b
}

func (b *SizedBox) GetPadding() myfyneKit.EdgeInset {
	return b.padding
}

// SetHeight 设置高度
func (b *SizedBox) SetHeight(height float32) *SizedBox {
	b.height = height
	b.Refresh()
	return b
}

// GetHeight 获取高度
func (b *SizedBox) GetHeight() float32 {
	if b.height >= 0 {
		return b.height
	}

	return myfyneKit.Infinity
}

// SetBackgroundColor 设置背景色
func (b *SizedBox) SetBackgroundColor(color color.Color) *SizedBox {
	b.background = color
	b.Refresh()
	return b
}

// GetBackgroundColor 获取背景色
func (b *SizedBox) GetBackgroundColor() color.Color {
	// 如果背景色为 nil，返回主题色
	if b.background == nil {
		return theme.Color(theme.ColorNameBackground)
	}
	return b.background
}

// SetChild 设置子控件
func (b *SizedBox) SetChild(child fyne.CanvasObject) *SizedBox {
	b.child = child
	b.Refresh()
	return b
}

// GetChild 获取子控件
func (b *SizedBox) GetChild() fyne.CanvasObject {
	return b.child
}

// CreateRenderer 创建控件的渲染器
func (b *SizedBox) CreateRenderer() fyne.WidgetRenderer {
	backgroundRect := canvas.NewRectangle(b.GetBackgroundColor())

	return &sizedBoxRenderer{
		box:        b,
		background: backgroundRect,
	}
}

// sizedBoxRenderer 负责渲染 SizedBox 控件
type sizedBoxRenderer struct {
	box        *SizedBox
	background *canvas.Rectangle
}

func (r *sizedBoxRenderer) Layout(size fyne.Size) {
	boxSize := r.MinSize()
	if r.box.width == myfyneKit.Infinity {
		boxSize.Width = size.Width
	}

	if r.box.height == myfyneKit.Infinity {
		boxSize.Height = size.Height
	}

	boxPosition := myfyneKit.ChildPosition(myfyneKit.PositionCenter, size, boxSize)
	r.background.Resize(boxSize)
	r.background.Move(boxPosition)

	//fmt.Println("Size :", size)
	//fmt.Println("box Size :", boxSize, boxPosition)
	if r.box.child != nil {
		childSize := boxSize
		childSize.Width -= r.box.padding.Left + r.box.padding.Right
		childSize.Height -= r.box.padding.Top + r.box.padding.Bottom

		childSizePosition := fyne.NewPos(boxPosition.X+r.box.padding.Left, boxPosition.Y+r.box.padding.Top)
		r.box.child.Resize(childSize)
		r.box.child.Move(childSizePosition)
		//fmt.Println("child Size :", childSize, childSizePosition)
	}
}

func (r *sizedBoxRenderer) MinSize() fyne.Size {
	ret := fyne.NewSize(r.box.GetWidth(), r.box.GetHeight())
	if ret.Width == myfyneKit.Infinity {
		ret.Width = r.box.padding.Left + r.box.padding.Right
		if r.box.child != nil {
			ret.Width += r.box.child.MinSize().Width
		}
	}

	if ret.Height == myfyneKit.Infinity {
		ret.Height = r.box.padding.Top + r.box.padding.Bottom
		if r.box.child != nil {
			ret.Height += r.box.child.MinSize().Height
		}
	}

	//fmt.Println("MinSize:....", ret)
	return ret
}

func (r *sizedBoxRenderer) Refresh() {
	r.background.FillColor = r.box.GetBackgroundColor()
	r.background.Refresh()
	canvas.Refresh(r.box)
}

func (r *sizedBoxRenderer) Objects() []fyne.CanvasObject {
	if r.box.child == nil {
		return []fyne.CanvasObject{r.background}
	}

	return []fyne.CanvasObject{r.background, r.box.child}
}

func (r *sizedBoxRenderer) Destroy() {}
