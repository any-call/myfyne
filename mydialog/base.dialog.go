package mydialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type BaseDialogContent struct {
	title      string
	dialog     dialog.Dialog
	window     fyne.Window
	closeParam any
}

func (self *BaseDialogContent) Title() string {
	return self.title
}

func (self *BaseDialogContent) SetDialog(dlg dialog.Dialog) {
	self.dialog = dlg
}

func (self *BaseDialogContent) SetWindow(win fyne.Window) {
	self.window = win
}

func (self *BaseDialogContent) SetTitle(title string) {
	self.title = title
}

func (self *BaseDialogContent) GetWindow() fyne.Window {
	return self.window
}

func (self *BaseDialogContent) CloseDialog(param any) {
	self.closeParam = param
	if self.dialog != nil {
		self.dialog.Hide()
	}
}

func (self *BaseDialogContent) GetCloseParam() any {
	return self.closeParam
}

func (self *BaseDialogContent) Content() fyne.CanvasObject {
	panic("Content() method must be implemented by the concrete page")
}
