package mywidget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

type DatePicker struct {
	widget.BaseWidget
	yearSel  *widget.Select
	monthSel *widget.Select
	daySel   *widget.Select
	current  time.Time
}

// NewDatePicker 创建日期选择器，如果未设置日期，则默认显示当前日期
func NewDatePicker() *DatePicker {
	picker := &DatePicker{
		current: time.Now(), // 设置当前日期为默认值
	}

	// 初始化年份、月份、日期选择
	picker.initSelectors()

	// 设置默认日期显示
	picker.SetDate(picker.current)

	// 初始化控件
	picker.ExtendBaseWidget(picker)
	return picker
}

// 初始化选择框
func (d *DatePicker) initSelectors() {
	// 年份选择
	years := []string{}
	for i := 1970; i <= d.current.Year(); i++ {
		years = append(years, strconv.Itoa(i))
	}
	d.yearSel = widget.NewSelect(years, nil)

	// 月份选择
	months := []string{}
	for i := 1; i <= 12; i++ {
		months = append(months, strconv.Itoa(i))
	}
	d.monthSel = widget.NewSelect(months, nil)

	// 日期选择，根据当前月份确定最大天数
	days := d.getDaysInMonth(d.current.Year(), int(d.current.Month()))
	d.daySel = widget.NewSelect(days, nil)

	// 年、月、日选择事件，动态更新日期
	d.yearSel.OnChanged = func(_ string) {
		d.updateDaysInMonth()
	}
	d.monthSel.OnChanged = func(_ string) {
		d.updateDaysInMonth()
	}
}

// 获取指定年份和月份的天数
func (d *DatePicker) getDaysInMonth(year, month int) []string {
	if month < 1 || month > 12 {
		return []string{}
	}

	// 计算下个月的第一天
	nextMonth := time.Month(month)%12 + 1
	nextYear := year
	if nextMonth == time.January {
		nextYear++
	}

	// 使用 time.Date 计算当月的最后一天
	firstDayOfNextMonth := time.Date(nextYear, nextMonth, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfNextMonth.AddDate(0, 0, -1)

	// 获取当月的天数
	totalDays := lastDayOfMonth.Day()
	// 创建包含当月所有天数的切片
	days := make([]string, totalDays)
	for i := 1; i <= totalDays; i++ {
		days[i-1] = fmt.Sprintf("%d", i)
	}

	return days
}

// 根据选择的年份和月份更新日期选择框
func (d *DatePicker) updateDaysInMonth() {
	year, _ := strconv.Atoi(d.yearSel.Selected)
	month, _ := strconv.Atoi(d.monthSel.Selected)
	if year == 0 || month == 0 {
		return
	}

	days := d.getDaysInMonth(year, month)
	selectedDay := d.daySel.Selected

	// 重新设置日期选择
	d.daySel.Options = days

	// 保留当前选择的日期，如果超出范围，则自动选最后一天
	tmpSelectDay, _ := strconv.Atoi(selectedDay)
	if selectedDay == "" || tmpSelectDay > len(days) {
		d.daySel.SetSelected(days[len(days)-1])
	} else {
		d.daySel.SetSelected(selectedDay)
	}
}

// SetDate 设置选择器的日期
func (d *DatePicker) SetDate(t time.Time) {
	d.yearSel.SetSelected(strconv.Itoa(t.Year()))
	d.monthSel.SetSelected(strconv.Itoa(int(t.Month())))
	d.daySel.SetSelected(strconv.Itoa(t.Day()))
}

// GetDate 获取当前选择的日期
func (d *DatePicker) GetDate() time.Time {
	year, _ := strconv.Atoi(d.yearSel.Selected)
	month, _ := strconv.Atoi(d.monthSel.Selected)
	day, _ := strconv.Atoi(d.daySel.Selected)

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

// CreateRenderer 实现 fyne.WidgetRenderer 接口
func (d *DatePicker) CreateRenderer() fyne.WidgetRenderer {
	// 将选择框组件排列成水平容器
	content := container.NewHBox(NewWidthBox(75, d.yearSel), canvas.NewText("年", nil),
		NewWidthBox(50, d.monthSel), canvas.NewText("月", nil),
		NewWidthBox(58, d.daySel), canvas.NewText("日", nil))
	return widget.NewSimpleRenderer(content)
}
