package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// ReadonlyEntry 是一个不可编辑但可选中复制、外观类似Label的控件
type ReadonlyEntry struct {
	widget.Entry
}

func NewReadonlyEntry(text string) *ReadonlyEntry {
	e := &ReadonlyEntry{}
	e.ExtendBaseWidget(e)
	e.SetText(text)
	e.Wrapping = fyne.TextWrapWord
	e.TextStyle = fyne.TextStyle{} // 可设置加粗等
	return e
}

// 捕获按键事件
func (e *ReadonlyEntry) TypedRune(_ rune) {
	// 阻止输入任何字符（包括中文）
	return
}

func (e *ReadonlyEntry) TypedKey(k *fyne.KeyEvent) {
	switch k.Name {
	case fyne.KeyBackspace, fyne.KeyDelete, fyne.KeyReturn, fyne.KeyEnter:
		// 阻止删除、换行
	default:
		// 允许选中 + Ctrl+C 等复制操作
		e.Entry.TypedKey(k)
	}
}

// 禁用粘贴
func (e *ReadonlyEntry) TypedShortcut(sc fyne.Shortcut) {
	switch sc.(type) {
	case *fyne.ShortcutPaste:
		// 忽略粘贴
	default:
		e.Entry.TypedShortcut(sc)
	}
}

// 让其成为 desktop.Focusable，避免显示光标
func (e *ReadonlyEntry) FocusGained() {
	// 不显示光标
	e.Entry.FocusLost()
}
