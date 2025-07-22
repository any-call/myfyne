package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"math"
	"time"
)

// ClickableContainer 是一个自定义容器控件，支持单击和双击事件。
type TapContainer struct {
	widget.BaseWidget
	child       fyne.CanvasObject // 容器内的子控件
	onTapped    func()            // 单击回调
	onDoubleTap func()

	lastTapTime time.Time
	lastTapPos  fyne.Position
	clickTimer  *time.Timer
}

func NewTapContainer(child fyne.CanvasObject, onTap func()) *TapContainer {
	c := &TapContainer{
		child:    child,
		onTapped: onTap,
	}
	c.ExtendBaseWidget(c)
	return c
}

func NewTapContainerWithDoubleTap(child fyne.CanvasObject, onTap func()) *TapContainer {
	c := &TapContainer{
		child:       child,
		onDoubleTap: onTap,
	}
	c.ExtendBaseWidget(c)
	return c
}

func (c *TapContainer) SetOnTap(onFn func()) {
	c.onTapped = onFn
}

func (c *TapContainer) SetOnDoubleTap(onFn func()) {
	c.onDoubleTap = onFn
}

func (c *TapContainer) GetChild() fyne.CanvasObject {
	return c.child
}

// Tapped 处理单击和双击事件。
func (c *TapContainer) Tapped(ev *fyne.PointEvent) {
	now := time.Now()
	pos := ev.Position

	if now.Sub(c.lastTapTime) < 300*time.Millisecond &&
		c.distance(c.lastTapPos, pos) < 10 {
		// 双击：取消单击计时器，触发双击
		if c.clickTimer != nil {
			c.clickTimer.Stop()
			c.clickTimer = nil
		}
		if c.onDoubleTap != nil {
			c.onDoubleTap()
		}
	} else {
		// 单击：启动延迟回调，防止双击误触
		if c.clickTimer != nil {
			c.clickTimer.Stop()
		}
		c.clickTimer = time.AfterFunc(300*time.Millisecond, func() {
			if c.onTapped != nil {
				c.onTapped()
			}
		})
	}

	c.lastTapTime = now
	c.lastTapPos = pos
}

// TappedSecondary 用于处理右键点击事件，但此处未使用。
func (c *TapContainer) TappedSecondary(_ *fyne.PointEvent) {}

// CreateRenderer 创建容器的渲染器。
func (c *TapContainer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.child)
}

func (c *TapContainer) distance(p1, p2 fyne.Position) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	return math.Hypot(dx, dy)
}
