package mywidget

import (
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ReadOnlyEntry struct {
	widget.Entry
}

func NewReadOnlyEntry() *ReadOnlyEntry {
	e := &ReadOnlyEntry{}
	e.MultiLine = true
	e.Wrapping = fyne.TextWrapWord
	e.ExtendBaseWidget(e)
	return e
}

// 禁止用户输入字符
func (e *ReadOnlyEntry) TypedRune(r rune) {}

// 禁止用户按键输入（回车、删除等）
func (e *ReadOnlyEntry) TypedKey(ev *fyne.KeyEvent) {}

// ===============================
// LogConsole Widget
// ===============================

type LogConsole struct {
	sync.Mutex
	widget.BaseWidget
	entry      *ReadOnlyEntry
	scroll     *container.Scroll
	autoScroll bool
}

// NewLogConsole 创建一个日志终端窗口（v2rayN 同款）
func NewLogConsole() *LogConsole {
	e := NewReadOnlyEntry()
	s := container.NewScroll(e)
	s.Direction = container.ScrollVerticalOnly

	c := &LogConsole{
		entry:      e,
		scroll:     s,
		autoScroll: true,
	}

	c.ExtendBaseWidget(c)
	return c
}

// Append 添加一行日志
func (c *LogConsole) Append(line string) {
	c.Lock()
	defer c.Unlock()
	c.entry.Append(line + "\n")
	c.entry.Refresh()
	//fyne.Do(func() {
	//	c.Lock()
	//	defer c.Unlock()
	//
	//	c.entry.Append(line + "\n")
	//	c.entry.Refresh()
	//
	//	if c.autoScroll {
	//		go func() {
	//			time.Sleep(20 * time.Millisecond)
	//			fyne.DoAndWait(func() {
	//				c.scroll.ScrollToBottom()
	//			})
	//		}()
	//	}
	//})
}

// Clear 清空日志
func (c *LogConsole) Clear() {
	fyne.Do(func() {
		c.Lock()
		defer c.Unlock()

		c.entry.SetText("")
		c.entry.Refresh()
	})
}

// SetAutoScroll 设置是否自动滚动
func (c *LogConsole) SetAutoScroll(enable bool) {
	c.autoScroll = enable
}
func (c *LogConsole) GetAutoScroll() bool {
	return c.autoScroll
}

// CanvasObject 输出（关键）
func (c *LogConsole) CreateRenderer() fyne.WidgetRenderer {
	return newLogConsoleRenderer(c)
}

// ===============================
// Renderer
// ===============================

type logConsoleRenderer struct {
	console *LogConsole

	bg      *canvas.Rectangle
	objects []fyne.CanvasObject
}

func newLogConsoleRenderer(c *LogConsole) *logConsoleRenderer {
	bg := canvas.NewRectangle(theme.Color(theme.ColorNameInputBackground))
	// 黑底效果（覆盖默认 Entry 背景）
	bg.FillColor = color.NRGBA{R: 15, G: 15, B: 15, A: 255}
	r := &logConsoleRenderer{
		console: c,
		bg:      bg,
	}

	r.objects = []fyne.CanvasObject{
		bg,
		c.scroll,
	}

	return r
}

func (r *logConsoleRenderer) Layout(size fyne.Size) {
	r.bg.Resize(size)
	r.console.scroll.Resize(size)
}

func (r *logConsoleRenderer) MinSize() fyne.Size {
	return fyne.NewSize(400, 200)
}

func (r *logConsoleRenderer) Refresh() {
	r.bg.Refresh()
	r.console.entry.Refresh()
}

func (r *logConsoleRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *logConsoleRenderer) Destroy() {}
