package mywidget

import (
	"fyne.io/fyne/v2"
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
