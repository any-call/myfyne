package myfyne

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
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

func ShowDialogBySize(win fyne.Window, size fyne.Size, content DialogContent, onClose func(param any)) dialog.Dialog {
	// 通过 content 的 Content() 方法获取展示的 fyne.CanvasObject
	dlg := dialog.NewCustomWithoutButtons(content.Title(), container.NewGridWrap(size, content.Content()), win)
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

func ShowToast(window fyne.Window, message string, duration time.Duration) {
	content := widget.NewLabel(message)
	dlg := dialog.NewCustomWithoutButtons("", container.NewCenter(content), window)
	dlg.Show()
	time.AfterFunc(duration, dlg.Hide)
	return
}

func SendNotificationMsg(title, content string) {
	if title == "" {
		title = "提示信息"
	}
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   title,
		Content: content,
	})
}
