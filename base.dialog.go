package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ShowDialogWithCallback(win fyne.Window, content DialogContent, onClose func(param any)) dialog.Dialog {
	// 通过 content 的 Content() 方法获取展示的 fyne.CanvasObject
	dlg := dialog.NewCustomWithoutButtons(content.Title(), content.Content(), win)
	// 设置 content 的 dialog 对象
	content.SetDialog(dlg)
	content.SetWindow(win)
	dlg.SetOnClosed(func() {
		if onClose != nil {
			onClose(content.GetCloseParam()) // 通过回调传递关闭时的参数
		}
	})
	dlg.Show() // 显示 dialog
	return dlg
}
