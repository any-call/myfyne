package mywidget

import "fyne.io/fyne/v2"

// RunOnMain 提交一个 UI 操作到主线程执行（同步阻塞）
func RunOnMain(fn func()) {
	if fn != nil {
		fyne.DoAndWait(fn)
	}
}

// RunOnMainAsync 提交一个 UI 操作到主线程执行（异步不阻塞）
func RunOnMainAsync(fn func()) {
	if fn != nil {
		fyne.Do(fn)
	}
}
