package myfyne

import "fyne.io/fyne/v2"

func CopyToClipboard(w fyne.Window, text string) {
	if text == "" {
		return
	}

	if w == nil {
		w = fyne.CurrentApp().Driver().AllWindows()[0]
	}

	w.Clipboard().SetContent(text)
}
