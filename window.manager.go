package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"sync"
)

// windowManager 是 WindowManager 的单例实现
type windowManager struct {
	app     fyne.App
	windows map[int]fyne.Window
	mutex   sync.Mutex
}

var instance *windowManager
var once sync.Once

// WinManagerIns 获取 WindowManager 的单例
func winManagerIns() *windowManager {
	once.Do(func() {
		instance = &windowManager{
			windows: make(map[int]fyne.Window),
		}
	})
	return instance
}

// SetApp 设置 fyne.App 实例
func (wm *windowManager) SetApp(app fyne.App) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	if wm.app == nil {
		wm.app = app
	}
}

// ShowPage 显示页面，如果窗口不存在，则创建并显示
func (wm *windowManager) ShowPage(page Page, centerOnScreen bool, fixedSize bool) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	if wm.app == nil {
		panic("App instance not set. Use SetApp to initialize.")
	}

	windowID := page.WinID()

	// 检查窗口是否已经存在
	window, exists := wm.windows[windowID]
	if !exists {
		// 创建新窗口
		window = wm.app.NewWindow("")
		wm.windows[windowID] = window
	}

	// 设置页面内容并调整窗口大小
	window.SetContent(page.Content())
	window.SetTitle(page.WinTitle())

	winSize := page.WinSize()
	if winSize.Width <= 0 {
		winSize.Width = page.Content().MinSize().Width
	}

	if winSize.Height <= 0 {
		winSize.Height = page.Content().MinSize().Height
	}

	window.Resize(winSize)
	if centerOnScreen {
		window.CenterOnScreen()
	}

	window.SetFixedSize(fixedSize)

	window.Show()
}

// ClosePage 关闭页面对应的窗口
func (wm *windowManager) ClosePage(page Page) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	if wm.app == nil {
		panic("App instance not set. Use SetApp to initialize.")
	}

	windowID := page.WinID()
	if window, exists := wm.windows[windowID]; exists {
		window.Close()
		delete(wm.windows, windowID)
	}
}

// GetWindow 获取页面对应的窗口
func (wm *windowManager) GetWindow(page Page) fyne.Window {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	return wm.windows[page.WinID()]
}

// ShowWindow 显示指定页面的窗口
func (wm *windowManager) ShowWindow(windowId int) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	if wm.app == nil {
		panic("App instance not set. Use SetApp to initialize.")
	}

	if window, ok := wm.windows[windowId]; ok {
		window.Show()
	}
}

// HideWindow 隐藏指定页面的窗口
func (wm *windowManager) HideWindow(windowId int) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	if wm.app == nil {
		panic("App instance not set. Use SetApp to initialize.")
	}

	if window, ok := wm.windows[windowId]; ok {
		window.Hide()
	}

}

func SetApp(app fyne.App) {
	winManagerIns().SetApp(app)
}

func ShowPage(page Page, centerOnScreen bool, fixedSize bool) {
	winManagerIns().ShowPage(page, centerOnScreen, fixedSize)
}

func ClosePage(page Page) {
	winManagerIns().ClosePage(page)
}

func GetWindows(p Page) fyne.Window {
	return winManagerIns().GetWindow(p)
}

func ShowWindow(windowId int) {
	winManagerIns().ShowWindow(windowId)
}

func HideWindow(windowId int) {
	winManagerIns().HideWindow(windowId)
}

func ShowError(err error, page Page) {
	dialog.ShowError(err, winManagerIns().GetWindow(page))
}

func ShowInfo(title, message string, page Page) {
	dialog.ShowInformation(title, message, winManagerIns().GetWindow(page))
}
