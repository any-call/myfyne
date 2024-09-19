package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
	"time"
)

type LoadingDots struct {
	widget.BaseWidget
	color       color.Color
	radius      float64
	animRunning bool
	angle       float64
	ticker      *time.Ticker
	stopChan    chan bool
}

func NewLoadingDots(c color.Color, radius float64) *LoadingDots {
	l := &LoadingDots{
		color:    c,
		radius:   radius,
		stopChan: make(chan bool),
	}
	l.ExtendBaseWidget(l)
	return l
}

func (l *LoadingDots) CreateRenderer() fyne.WidgetRenderer {
	// 创建一组小圆点
	dots := make([]*canvas.Circle, 8)
	for i := 0; i < 8; i++ {
		dot := canvas.NewCircle(l.color)
		dots[i] = dot
	}

	return &loadingDotsRenderer{
		loading: l,
		dots:    dots,
		objects: make([]fyne.CanvasObject, len(dots)),
	}
}

func (l *LoadingDots) Start() *LoadingDots {
	if l.animRunning {
		return l
	}
	l.animRunning = true
	l.ticker = time.NewTicker(30 * time.Millisecond)
	go func() {
		for {
			select {
			case <-l.stopChan:
				l.ticker.Stop()
				return
			case <-l.ticker.C:
				l.angle += 0.1 // 每次旋转的角度
				if l.angle > 2*math.Pi {
					l.angle = 0
				}
				l.Refresh() // 刷新控件
			}
		}
	}()

	return l
}

func (l *LoadingDots) Stop() *LoadingDots {
	if !l.animRunning {
		return l
	}
	l.animRunning = false
	l.stopChan <- true
	return l
}

func (l *LoadingDots) Show() {
	l.BaseWidget.Show()
}

func (l *LoadingDots) Hide() {
	l.BaseWidget.Hide()
}

type loadingDotsRenderer struct {
	loading *LoadingDots
	dots    []*canvas.Circle
	objects []fyne.CanvasObject
}

func (r *loadingDotsRenderer) Layout(size fyne.Size) {
	centerX, centerY := size.Width/2, size.Height/2
	radius := float32(r.loading.radius)

	// 计算每个圆点的位置和大小
	for i, dot := range r.dots {
		angle := r.loading.angle + float64(i)*math.Pi/4 // 每个圆点的角度偏移
		offset := radius                                // 使用传入的半径
		dotSize := radius * 0.1                         // 根据半径调整圆点大小
		if i == 0 || i == 7 {                           // 两边的点较小
			dotSize *= 0.6
		} else if i == 3 || i == 4 { // 中间的点较大
			dotSize *= 1.4
		}
		dot.Resize(fyne.NewSize(dotSize, dotSize)) // 调整大小
		x := centerX + float32(math.Cos(angle))*offset
		y := centerY + float32(math.Sin(angle))*offset
		dot.Move(fyne.NewPos(x, y))
	}

	// 刷新对象
	for _, dot := range r.dots {
		dot.Refresh()
	}
}

func (r *loadingDotsRenderer) MinSize() fyne.Size {
	// 最小尺寸根据传入的半径来确定
	diameter := float32(r.loading.radius * 2)
	return fyne.NewSize(diameter, diameter)
}

func (r *loadingDotsRenderer) Refresh() {
	// 动态根据控件的实际大小进行布局
	size := r.loading.Size()
	r.Layout(size)
}

func (r *loadingDotsRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, len(r.dots))
	for i, dot := range r.dots {
		objects[i] = dot
	}
	return objects
}

func (r *loadingDotsRenderer) Destroy() {}
