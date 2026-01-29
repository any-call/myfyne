package mywidget

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

// DateTimePicker 日期时间选择器（精确到秒）
type DateTimePicker struct {
	widget.BaseWidget

	selectedTime  time.Time
	calendarPop   *widget.PopUp
	displayButton *widget.Button

	placeHolder string
	onChanged   func(t time.Time)
}

// NewDateTimePicker 创建 DateTimePicker
func NewDateTimePicker(t time.Time, zeroPlace string, changed func(time.Time)) *DateTimePicker {
	dp := &DateTimePicker{
		selectedTime: t,
		placeHolder:  zeroPlace,
		onChanged:    changed,
	}

	dp.displayButton = widget.NewButton("", func() {
		dp.showPopup(fyne.CurrentApp().Driver().AllWindows()[0])
	})
	dp.displayButton.SetIcon(theme.Icon(theme.IconNameArrowDropDown))
	dp.displayButton.IconPlacement = widget.ButtonIconTrailingText

	dp.refreshText()

	dp.ExtendBaseWidget(dp)
	return dp
}

// CreateRenderer
func (dp *DateTimePicker) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(dp.displayButton)
}

// GetTime 获取当前值
func (dp *DateTimePicker) GetTime() time.Time {
	return dp.selectedTime
}

// Clear 清空
func (dp *DateTimePicker) Clear() {
	dp.selectedTime = time.Time{}
	dp.refreshText()

	if dp.onChanged != nil {
		dp.onChanged(dp.selectedTime)
	}
}

// refreshText 刷新按钮显示
func (dp *DateTimePicker) refreshText() {
	if dp.selectedTime.IsZero() {
		if dp.placeHolder != "" {
			dp.displayButton.SetText(dp.placeHolder)
		} else {
			dp.displayButton.SetText("请选择时间")
		}
	} else {
		dp.displayButton.SetText(dp.selectedTime.Format("2006-01-02 15:04:05"))
	}
}

// showPopup 弹出 Calendar + Time Select
func (dp *DateTimePicker) showPopup(w fyne.Window) {

	showTime := dp.selectedTime
	if showTime.IsZero() {
		showTime = time.Now()
	}

	// 当前时分秒
	h := showTime.Hour()
	m := showTime.Minute()
	s := showTime.Second()

	// ===== 时间选择下拉框 =====
	hSel := widget.NewSelect(makeRange(0, 23), nil)
	mSel := widget.NewSelect(makeRange(0, 59), nil)
	sSel := widget.NewSelect(makeRange(0, 59), nil)

	hSel.SetSelected(fmt.Sprintf("%02d", h))
	mSel.SetSelected(fmt.Sprintf("%02d", m))
	sSel.SetSelected(fmt.Sprintf("%02d", s))

	// ===== 日历控件 =====
	calendar := xwidget.NewCalendar(showTime, func(t time.Time) {
		// 只更新日期部分，时间保持
		dp.selectedTime = time.Date(
			t.Year(), t.Month(), t.Day(),
			h, m, s,
			0, time.Local,
		)
		dp.refreshText()

		if dp.onChanged != nil {
			dp.onChanged(dp.selectedTime)
		}
	})

	// ===== 时间变化监听 =====
	updateTime := func() {
		if dp.selectedTime.IsZero() {
			dp.selectedTime = time.Now()
		}

		h = mustInt(hSel.Selected)
		m = mustInt(mSel.Selected)
		s = mustInt(sSel.Selected)

		t := dp.selectedTime
		dp.selectedTime = time.Date(
			t.Year(), t.Month(), t.Day(),
			h, m, s,
			0, time.Local,
		)

		dp.refreshText()

		if dp.onChanged != nil {
			dp.onChanged(dp.selectedTime)
		}
	}

	hSel.OnChanged = func(string) { updateTime() }
	mSel.OnChanged = func(string) { updateTime() }
	sSel.OnChanged = func(string) { updateTime() }

	// ===== 操作按钮 =====
	clearBtn := widget.NewButton("清空", func() {
		dp.Clear()
		dp.calendarPop.Hide()
	})

	okBtn := widget.NewButton("确定", func() {
		dp.calendarPop.Hide()
	})

	// ===== 布局 =====
	timeBar := container.NewHBox(
		widget.NewLabel("时间:"),
		hSel,
		widget.NewLabel(":"),
		mSel,
		widget.NewLabel(":"),
		sSel,
	)

	content := container.NewVBox(
		calendar,
		timeBar,
		container.NewHBox(clearBtn, okBtn),
	)

	// ===== 弹窗显示 =====
	pos := dp.popupPos()
	dp.calendarPop = widget.NewPopUp(content, w.Canvas())
	dp.calendarPop.ShowAtPosition(pos)
}

// popupPos 显示在按钮下方
func (dp *DateTimePicker) popupPos() fyne.Position {
	buttonPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(dp)
	return buttonPos.Add(
		fyne.NewPos(0, dp.Size().Height-dp.Theme().Size(theme.SizeNameInputBorder)),
	)
}

// ===== 工具函数 =====

// makeRange 生成 00~59
func makeRange(min, max int) []string {
	var list []string
	for i := min; i <= max; i++ {
		list = append(list, fmt.Sprintf("%02d", i))
	}
	return list
}

func mustInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
