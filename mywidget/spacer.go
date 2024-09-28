package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/myfyne"
	"image/color"
)

type FixedSpacer struct {
	widget.BaseWidget
	width  float32
	height float32
}

func NewFixedSpacer(width, height float32) *FixedSpacer {
	spacer := &FixedSpacer{
		width:  width,
		height: height,
	}
	spacer.ExtendBaseWidget(spacer)
	return spacer
}

func NewWidthSpacer(widget float32) *FixedSpacer {
	return NewFixedSpacer(widget, myfyne.Infinity)
}

func NewHeightSpacer(height float32) *FixedSpacer {
	return NewFixedSpacer(myfyne.Infinity, height)
}

// CreateRenderer 定义了组件的渲染
func (s *FixedSpacer) CreateRenderer() fyne.WidgetRenderer {
	rect := canvas.NewRectangle(color.Transparent) // 占位的透明矩形
	return &fixedSpacerRenderer{spacer: s, rect: rect}
}

// MinSize 返回组件的最小尺寸，如果宽度或高度为 Infinity 则占据父容器空间
func (s *FixedSpacer) MinSize() fyne.Size {
	parentWidth := fyne.Size{Width: s.width, Height: s.height}
	if s.width == myfyne.Infinity {
		parentWidth.Width = 0 // 如果是 Infinity，最小尺寸设置为 0，实际会由布局决定
	}
	if s.height == myfyne.Infinity {
		parentWidth.Height = 0 // 如果是 Infinity，最小尺寸设置为 0，实际会由布局决定
	}
	return parentWidth
}

// Renderer 负责组件的渲染
type fixedSpacerRenderer struct {
	spacer *FixedSpacer
	rect   *canvas.Rectangle
}

func (r *fixedSpacerRenderer) Layout(size fyne.Size) {
	if r.spacer.width == myfyne.Infinity {
		size.Width = size.Width // 占据父容器剩余的宽度
	}
	if r.spacer.height == myfyne.Infinity {
		size.Height = size.Height // 占据父容器剩余的高度
	}
	r.rect.Resize(size)
}

func (r *fixedSpacerRenderer) MinSize() fyne.Size {
	return r.spacer.MinSize()
}

func (r *fixedSpacerRenderer) Refresh() {
	canvas.Refresh(r.spacer)
}

func (r *fixedSpacerRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.rect}
}

func (r *fixedSpacerRenderer) Destroy() {}
