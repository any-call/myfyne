package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MyIcon struct {
	widget.Icon
	_tapped    func()
	_secTapped func()
}

func NewIcon(res fyne.Resource) *MyIcon {
	icon := &MyIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

func (self *MyIcon) SetTapped(f func()) *MyIcon {
	self._tapped = f
	return self
}

func (self *MyIcon) SetSecTapped(f func()) *MyIcon {
	self._secTapped = f
	return self
}

func (t *MyIcon) Tapped(_ *fyne.PointEvent) {
	if t._tapped != nil {
		t._tapped()
	}
}

func (t *MyIcon) TappedSecondary(_ *fyne.PointEvent) {
	if t._secTapped != nil {
		t._secTapped()
	}
}
