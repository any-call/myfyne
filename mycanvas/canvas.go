package mycanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// Canvas is where objects are drawn into
type Canvas struct {
	*fyne.Container
	width  float64
	height float64
}

// NewCanvas makes a new canvas
func NewCanvas(w, h int) *Canvas {
	c := Canvas{
		Container: container.NewWithoutLayout(),
		width:     float64(w),
		height:    float64(h),
	}
	return &c
}

// CreateRenderer 让容器大小适应内容
func (c *Canvas) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.Container)
}

// MinSize 实现 CanvasObject 接口
func (c *Canvas) MinSize() fyne.Size {
	return fyne.NewSize(float32(c.width), float32(c.height))
}

// TextWidth returns the width of a string
func (c *Canvas) matchTextSize(s string, size float32, style fyne.TextStyle) fyne.Size {
	return fyne.MeasureText(s, size, style)
}

func (c *Canvas) DrawText(align Align, xOffset, yOffset float32,
	size float32, text string, style fyne.TextStyle, col color.Color,
) {
	if text == "" || size <= 0 {
		return
	}

	tSize := c.matchTextSize(text, size, style)
	t := &canvas.Text{Text: text, Color: col, TextSize: size}
	var px, py float32
	// 水平基准点
	switch align {
	case AlignTopLeft:
		px = 0
		py = 0
		break
	case AlignTopCenter:
		px = (float32(c.width) - tSize.Width) / 2
		py = 0
		break
	case AlignTopRight:
		px = float32(c.width) - tSize.Width
		py = 0
		break
	case AlignCenterLeft:
		px = 0
		py = (float32(c.height) - tSize.Height) / 2
		break
	case AlignCenter:
		px = (float32(c.width) - tSize.Width) / 2
		py = (float32(c.height) - tSize.Height) / 2
		break
	case AlignCenterRight:
		px = float32(c.width) - tSize.Width
		py = (float32(c.height) - tSize.Height) / 2
		break
	case AlignBottomLeft:
		px = 0
		py = float32(c.height) - tSize.Height
		break
	case AlignBottomCenter:
		px = (float32(c.width) - tSize.Width) / 2
		py = float32(c.height) - tSize.Height
		break
	case AlignBottomRight:
		px = float32(c.width) - tSize.Width
		py = float32(c.height) - tSize.Height
		break
	}

	t.Move(fyne.Position{
		X: px + xOffset,
		Y: py + yOffset,
	})
	c.Container.Add(t)
}

func (c *Canvas) DrawCircle(align Align, xOffset, yOffset float32, radius float32, col color.Color) {
	if radius <= 0 {
		return
	}

	var p1, p2 fyne.Position
	// 水平基准点
	switch align {
	case AlignTopLeft:
		p1 = fyne.NewPos(0, 0)
		p2 = fyne.NewPos(radius*2, radius*2)
		break
	case AlignTopCenter:
		p1 = fyne.NewPos((float32(c.width)-radius*2)/2, 0)
		p2 = fyne.NewPos(((float32(c.width)-radius*2)/2)+radius*2, radius*2)
		break
	case AlignTopRight:
		p1 = fyne.NewPos((float32(c.width) - 2*radius), 0)
		p2 = fyne.NewPos(float32(c.width), radius*2)
		break
	case AlignCenterLeft:
		p1 = fyne.NewPos(0, (float32(c.height)-radius*2)/2)
		p2 = fyne.NewPos(radius*2, ((float32(c.height)-radius*2)/2)+radius*2)
		break
	case AlignCenter:
		p1 = fyne.NewPos((float32(c.width)-radius*2)/2, (float32(c.height)-radius*2)/2)
		p2 = fyne.NewPos(((float32(c.width)-radius*2)/2)+radius*2, ((float32(c.height)-radius*2)/2)+radius*2)
		break
	case AlignCenterRight:
		p1 = fyne.NewPos((float32(c.width) - 2*radius), (float32(c.height)-radius*2)/2)
		p2 = fyne.NewPos((float32(c.width)-2*radius)+radius*2, ((float32(c.height)-radius*2)/2)+radius*2)
		break
	case AlignBottomLeft:
		p1 = fyne.NewPos(0, float32(c.height)-radius*2)
		p2 = fyne.NewPos(radius*2, float32(c.height))
		break
	case AlignBottomCenter:
		p1 = fyne.NewPos((float32(c.width)-radius*2)/2, float32(c.height)-radius*2)
		p2 = fyne.NewPos(((float32(c.width)-radius*2)/2)+radius*2, float32(c.height))
		break
	case AlignBottomRight:
		p1 = fyne.NewPos((float32(c.width) - 2*radius), float32(c.height)-radius*2)
		p2 = fyne.NewPos((float32(c.width)-2*radius)+radius*2, float32(c.height))
		break
	}

	t := &canvas.Circle{FillColor: col,
		Position1: fyne.NewPos(p1.X+xOffset, p1.Y+yOffset),
		Position2: fyne.NewPos(p2.X+xOffset, p2.Y+yOffset)}
	c.Container.Add(t)
}

func (c *Canvas) DrawRect(align Align, xOffset, yOffset float32, radius float32, width, height float32, col color.Color) {
	if width <= 0 || height <= 0 {
		return
	}

	var px, py float32
	// 水平基准点
	switch align {
	case AlignTopLeft:
		px = 0
		py = 0
		break
	case AlignTopCenter:
		px = (float32(c.width) - width) / 2
		py = 0
		break
	case AlignTopRight:
		px = float32(c.width) - width
		py = 0
		break
	case AlignCenterLeft:
		px = 0
		py = (float32(c.height) - height) / 2
		break
	case AlignCenter:
		px = (float32(c.width) - width) / 2
		py = (float32(c.height) - height) / 2
		break
	case AlignCenterRight:
		px = float32(c.width) - width
		py = (float32(c.height) - height) / 2
		break
	case AlignBottomLeft:
		px = 0
		py = float32(c.height) - height
		break
	case AlignBottomCenter:
		px = (float32(c.width) - width) / 2
		py = float32(c.height) - height
		break
	case AlignBottomRight:
		px = float32(c.width) - width
		py = float32(c.height) - height
		break
	}

	t := &canvas.Rectangle{
		FillColor:    col,
		CornerRadius: radius,
	}
	t.Move(fyne.Position{X: px + xOffset, Y: py + yOffset})
	t.Resize(fyne.NewSize(width, height))
	c.Container.Add(t)
}

func (c *Canvas) DrawImage(align Align, xOffset, yOffset float32, imagePath string, imageW, imageH float32) {
	if len(imagePath) <= 0 || imageW <= 0 || imageH <= 0 {
		return
	}

	var px, py float32
	// 水平基准点
	switch align {
	case AlignTopLeft:
		px = 0
		py = 0
		break
	case AlignTopCenter:
		px = (float32(c.width) - imageW) / 2
		py = 0
		break
	case AlignTopRight:
		px = float32(c.width) - imageW
		py = 0
		break
	case AlignCenterLeft:
		px = 0
		py = (float32(c.height) - imageH) / 2
		break
	case AlignCenter:
		px = (float32(c.width) - imageW) / 2
		py = (float32(c.height) - imageH) / 2
		break
	case AlignCenterRight:
		px = float32(c.width) - imageW
		py = (float32(c.height) - imageH) / 2
		break
	case AlignBottomLeft:
		px = 0
		py = float32(c.height) - imageH
		break
	case AlignBottomCenter:
		px = (float32(c.width) - imageW) / 2
		py = float32(c.height) - imageH
		break
	case AlignBottomRight:
		px = float32(c.width) - imageW
		py = float32(c.height) - imageH
		break
	}

	t := canvas.NewImageFromFile(imagePath)
	t.Move(fyne.Position{X: px + xOffset, Y: py + yOffset})
	t.Resize(fyne.NewSize(imageW, imageH))
	c.Container.Add(t)
}

func (c *Canvas) DrawLine(x1, y1, x2, y2 float32, lineSize float32, lineCol color.Color) {
	p1 := fyne.Position{X: float32(x1), Y: float32(y1)}
	p2 := fyne.Position{X: float32(x2), Y: float32(y2)}
	c.Container.Add(&canvas.Line{StrokeColor: lineCol, StrokeWidth: lineSize, Position1: p1, Position2: p2})
}
