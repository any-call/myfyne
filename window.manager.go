package myfyne

import (
	"fyne.io/fyne/v2"
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
func WinManagerIns() *windowManager {
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
func (wm *windowManager) ShowPage(page Page, centerOnScreen bool) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	if wm.app == nil {
		panic("App instance not set. Use SetApp to initialize.")
	}

	windowID := page.WindowID()

	// 检查窗口是否已经存在
	window, exists := wm.windows[windowID]
	if !exists {
		// 创建新窗口
		window = wm.app.NewWindow(page.Title())
		if page.WindowWidth() > 0 && page.WindowHeight() > 0 {
			window.Resize(fyne.NewSize(page.WindowWidth(), page.WindowHeight()))
		} else {
			window.Resize(page.Content().MinSize())
		}
		wm.windows[windowID] = window
	}

	// 设置页面内容并调整窗口大小
	window.SetContent(page.Content())
	window.SetTitle(page.Title())

	if page.WindowWidth() > 0 && page.WindowHeight() > 0 {
		window.Resize(fyne.NewSize(page.WindowWidth(), page.WindowHeight()))
	} else {
		window.Resize(page.Content().MinSize())
	}

	if centerOnScreen {
		window.CenterOnScreen()
	}

	window.Show()
}

// ClosePage 关闭页面对应的窗口
func (wm *windowManager) ClosePage(page Page) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	if wm.app == nil {
		panic("App instance not set. Use SetApp to initialize.")
	}

	windowID := page.WindowID()
	if window, exists := wm.windows[windowID]; exists {
		window.Close()
		delete(wm.windows, windowID)
	}
}

// GetWindow 获取页面对应的窗口
func (wm *windowManager) GetWindow(page Page) fyne.Window {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()

	return wm.windows[page.WindowID()]
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
