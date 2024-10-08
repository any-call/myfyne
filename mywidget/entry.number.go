package mywidget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type EntryNumber struct {
	widget.Entry
}

func NewEntryNumber() *EntryNumber {
	entry := &EntryNumber{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func NewEntryByInt(i int) *EntryNumber {
	entry := NewEntryNumber()
	entry.Text = fmt.Sprintf("%d", i)
	return entry
}

func NewEntryByFloat(f float64) *EntryNumber {
	entry := NewEntryNumber()
	entry.Text = fmt.Sprintf("%f", f)
	return entry
}

func (self *EntryNumber) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' {
		self.Entry.TypedRune(r)
	}
}

func (self *EntryNumber) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		self.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		self.Entry.TypedShortcut(shortcut)
	}
}

func (self *EntryNumber) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
