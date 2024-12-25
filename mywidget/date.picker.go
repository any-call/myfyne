package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
	"time"
)

// DatePicker 继承 widget.BaseWidget，定义日期选择器组件
type DatePicker struct {
	widget.BaseWidget
	selectedDate  time.Time
	calendarPop   *widget.PopUp
	displayButton *widget.Button
	onChanged     func(t time.Time)
}

// NewDatePicker 创建一个新的 DatePicker 控件
func NewDatePicker(t time.Time, zeroTimePlace string, changed func(t time.Time)) *DatePicker {
	dp := &DatePicker{
		onChanged:    changed,
		selectedDate: t,
	}

	dp.displayButton = widget.NewButton("", func() {
		dp.showCalendar(fyne.CurrentApp().Driver().AllWindows()[0])
	})
	dp.displayButton.SetIcon(theme.Icon(theme.IconNameArrowDropDown))
	dp.displayButton.IconPlacement = widget.ButtonIconTrailingText

	if dp.selectedDate.IsZero() {
		if zeroTimePlace != "" {
			dp.displayButton.SetText(zeroTimePlace)
		} else {
			dp.displayButton.SetText("请选择日期")
		}
	} else {
		dp.displayButton.SetText(dp.selectedDate.Format("2006-01-02"))
	}

	dp.ExtendBaseWidget(dp)
	return dp
}

// CreateRenderer 实现自定义渲染器
func (dp *DatePicker) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(dp.displayButton)
}

func (dp *DatePicker) GetDate() time.Time {
	return dp.selectedDate
}

// showCalendar 弹出日历选择框，显示在 DatePicker 的下方
func (dp *DatePicker) showCalendar(w fyne.Window) {
	calendar := xwidget.NewCalendar(dp.selectedDate, func(t time.Time) {
		oldTime := dp.selectedDate
		dp.selectedDate = t
		dp.displayButton.SetText(dp.selectedDate.Format("2006-01-02"))
		dp.calendarPop.Hide()
		if dp.onChanged != nil && !oldTime.Equal(t) {
			dp.onChanged(t)
		}
	})

	// 获取 DatePicker 的相对位置
	pos := dp.popupPos()

	// 使用弹窗显示 Calendar，显示在 DatePicker 的下方
	dp.calendarPop = widget.NewPopUp(container.NewVBox(calendar), w.Canvas())

	// 显示日历的弹窗，设置位置在 DatePicker 下方
	//dp.calendarPop.ShowAtPosition(fyne.NewPos(pos.X, pos.Y+dp.Size().Height))
	dp.calendarPop.ShowAtPosition(pos)
}

func (db *DatePicker) popupPos() fyne.Position {
	buttonPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(db)
	return buttonPos.Add(fyne.NewPos(0, db.Size().Height-db.Theme().Size(theme.SizeNameInputBorder)))
}
