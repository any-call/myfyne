package mywidget

import (
	"context"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// LogConsole 高性能日志控件（RichText + 定时刷新 + 增量刷新）
type LogConsole struct {
	widget.BaseWidget
	mu         sync.Mutex
	rt         *widget.RichText
	scroll     *container.Scroll
	autoScroll bool
	buf        []string // 全量缓存，用于 CopyAll
	pending    []string // 待刷新新增行
	maxLines   int
	dirty      bool
	ctx        context.Context
	cancel     context.CancelFunc
}

// NewLogConsole 创建日志控件
func NewLogConsole(maxLines int) *LogConsole {
	rt := widget.NewRichText()
	rt.Wrapping = fyne.TextWrapWord

	scroll := container.NewScroll(rt)
	scroll.Direction = container.ScrollVerticalOnly

	ctx, cancel := context.WithCancel(context.Background())

	c := &LogConsole{
		rt:         rt,
		scroll:     scroll,
		autoScroll: true,
		maxLines:   maxLines,
		ctx:        ctx,
		cancel:     cancel,
	}

	c.ExtendBaseWidget(c)
	c.startFlush()
	return c
}

// Append 添加一行日志（增量刷新）
func (c *LogConsole) Append(line string) {
	c.mu.Lock()
	c.buf = append(c.buf, line)
	c.pending = append(c.pending, line)

	// 仅裁剪全量缓存 buf，pending 不裁剪
	if len(c.buf) > c.maxLines {
		overflow := len(c.buf) - c.maxLines
		c.buf = c.buf[overflow:]
	}
	c.dirty = true
	c.mu.Unlock()
}

// Clear 清空日志
func (c *LogConsole) Clear() {
	c.mu.Lock()
	c.buf = nil
	c.pending = nil
	c.dirty = false // pending 已清空，不用标记 dirty
	c.mu.Unlock()

	// UI 上立即清空
	fyne.Do(func() {
		c.rt.Segments = nil
		c.rt.Refresh()
		c.scroll.ScrollToTop()
	})
}

// CopyAll 全量复制日志到剪贴板
func (c *LogConsole) CopyAll() {
	c.mu.Lock()
	text := strings.Join(c.buf, "\n")
	c.mu.Unlock()

	if clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard(); clipboard != nil {
		clipboard.SetContent(text)
	}
}

// SetAutoScroll 设置自动滚动
func (c *LogConsole) SetAutoScroll(enable bool) {
	c.mu.Lock()
	c.autoScroll = enable
	c.mu.Unlock()
}

// GetAutoScroll 获取自动滚动状态
func (c *LogConsole) GetAutoScroll() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.autoScroll
}

func (c *LogConsole) SetMaxLines(num int) {
	c.mu.Lock()
	if num > 0 {
		c.maxLines = num
	}
	c.mu.Unlock()
}

// GetAutoScroll 获取自动滚动状态
func (c *LogConsole) GetMaxLines() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.maxLines
}

// startFlush 定时刷新 UI（高性能）
func (c *LogConsole) startFlush() {
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.flushIfNeeded()
			case <-c.ctx.Done():
				return
			}
		}
	}()
}

// flushIfNeeded 增量刷新 RichText，只刷新 pending
func (c *LogConsole) flushIfNeeded() {
	c.mu.Lock()
	if !c.dirty || len(c.pending) == 0 {
		c.mu.Unlock()
		return
	}

	// 拷贝 pending 并清空
	newLines := make([]string, len(c.pending))
	copy(newLines, c.pending)
	c.pending = nil
	c.dirty = false
	auto := c.autoScroll
	c.mu.Unlock()

	fyne.Do(func() {
		// 增量追加 segment
		for _, l := range newLines {
			c.rt.Segments = append(c.rt.Segments, &widget.TextSegment{
				Text: l,
				Style: widget.RichTextStyle{
					ColorName: theme.ColorNameForeground,
				},
			})
		}

		// 超过 maxLines 裁剪 RichText 头部 segment
		if len(c.rt.Segments) > c.maxLines {
			c.rt.Segments = c.rt.Segments[len(c.rt.Segments)-c.maxLines:]
		}

		c.rt.Refresh()
		if auto {
			c.scroll.ScrollToBottom()
		}
	})
}

// CreateRenderer 输出 CanvasObject
func (c *LogConsole) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.scroll)
}

// Destroy 停止刷新协程
func (c *LogConsole) Destroy() {
	c.cancel()
}
