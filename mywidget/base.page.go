package mywidget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"golang.org/x/image/colornames"
	"image/color"
)

type BasePage struct {
	title string
	id    int
}

func (self *BasePage) ShowLoading(win fyne.Window, title string, col color.Color, radius float64) *dialog.CustomDialog {
	if title == "" {
		title = "Loading"
	}

	if col == nil {
		col = colornames.Blue
	}

	if radius <= 0 {
		radius = 30
	}

	loadingDlg := dialog.NewCustomWithoutButtons(title, NewLoadingDots(col, radius).Start(), win)
	loadingDlg.Show()
	return loadingDlg
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

func (self *BasePage) WinClosed() {
	fmt.Println("enter basePage WinClosed")
	return
}
