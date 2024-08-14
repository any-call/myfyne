package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type BaseWindow struct {
	app    fyne.App
	window fyne.Window
}

func NewBaseWindow(app fyne.App) *BaseWindow {
	return &BaseWindow{
		app:    app,
		window: nil,
	}
}

func (self *BaseWindow) GetApp() fyne.App {
	if self.app == nil {
		self.app = fyne.CurrentApp()
	}
	return self.app
}

func (self *BaseWindow) GetWindow() fyne.Window {
	if self.window == nil {
		self.window = self.GetApp().NewWindow("")
	}

	return self.window
}

func (self *BaseWindow) Quit() {
	self.GetApp().Quit()
}

// 显示对话框提示信息
func (self *BaseWindow) ShowInfo(title string, message string) {
	if title == "" {
		title = "Info"
	}

	dialog.ShowInformation(title, message, self.GetWindow())
}

func (self *BaseWindow) ShowErr(err error) {
	dialog.ShowError(err, self.GetWindow())
}
