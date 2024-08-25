package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Icon struct {
	widget.Icon
	_tapped    func()
	_secTapped func()
}

func NewIcon(res fyne.Resource) *Icon {
	icon := &Icon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

func (self *Icon) SetTapped(f func()) *Icon {
	self._tapped = f
	return self
}

func (self *Icon) SetSecTapped(f func()) *Icon {
	self._secTapped = f
	return self
}

func (t *Icon) Tapped(_ *fyne.PointEvent) {
	if t._tapped != nil {
		t._tapped()
	}
}

func (t *Icon) TappedSecondary(_ *fyne.PointEvent) {
	if t._secTapped != nil {
		t._secTapped()
	}
}
