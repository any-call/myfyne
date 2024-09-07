package mywidget

import "fyne.io/fyne/v2"

type BasePage struct {
	title string
	id    int
}

func (self *BasePage) Content() fyne.CanvasObject {
	panic("Content() method must be implemented by the concrete page")
}

func (self *BasePage) WinTitle() string {
	return self.title
}

func (self *BasePage) WinID() int {
	return self.id
}

func (self *BasePage) WinSize() fyne.Size {
	panic("WinSize() method must be implemented by the concrete page")
}

func (self *BasePage) SetWinID(id int) {
	self.id = id
	return
}

func (self *BasePage) SetTitle(title string) {
	self.title = title
	return
}

func (self *BasePage) WinWillClose() {
	return
}

func (self *BasePage) WinClosed() {
	return
}
