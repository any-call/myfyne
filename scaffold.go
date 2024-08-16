package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Scaffold struct {
	widget.BaseWidget
	topBar, bottomBar, sideBar, content, floatingButton fyne.CanvasObject
}

func NewScaffold(topBar, bottomBar, sideBar, content, floatingButton fyne.CanvasObject) *Scaffold {
	s := &Scaffold{
		topBar:         topBar,
		bottomBar:      bottomBar,
		sideBar:        sideBar,
		content:        content,
		floatingButton: floatingButton,
	}
	s.ExtendBaseWidget(s) // 扩展基础组件
	return s
}

func (s *Scaffold) CreateRenderer() fyne.WidgetRenderer {
	objects := []fyne.CanvasObject{}
	if s.topBar != nil {
		objects = append(objects, s.topBar)
	}
	if s.sideBar != nil {
		objects = append(objects, s.sideBar)
	}
	if s.content != nil {
		objects = append(objects, s.content)
	}
	if s.bottomBar != nil {
		objects = append(objects, s.bottomBar)
	}
	if s.floatingButton != nil {
		objects = append(objects, s.floatingButton)
	}

	return &scaffoldRenderer{
		scaffold: s,
	}
}

// SetTopBar 设置顶部栏
func (s *Scaffold) SetTopBar(topBar fyne.CanvasObject) {
	s.topBar = topBar
	s.Refresh()
}

// GetTopBar 获取顶部栏
func (s *Scaffold) GetTopBar() fyne.CanvasObject {
	return s.topBar
}

// SetBottomBar 设置底部栏
func (s *Scaffold) SetBottomBar(bottomBar fyne.CanvasObject) {
	s.bottomBar = bottomBar
	s.Refresh()
}

// GetBottomBar 获取底部栏
func (s *Scaffold) GetBottomBar() fyne.CanvasObject {
	return s.bottomBar
}

// SetSideBar 设置侧边栏
func (s *Scaffold) SetSideBar(sideBar fyne.CanvasObject) {
	s.sideBar = sideBar
	s.Refresh()
}

// GetSideBar 获取侧边栏
func (s *Scaffold) GetSideBar() fyne.CanvasObject {
	return s.sideBar
}

// SetContent 设置内容区域
func (s *Scaffold) SetContent(content fyne.CanvasObject) {
	s.content = content
	s.Refresh()
}

// GetContent 获取内容区域
func (s *Scaffold) GetContent() fyne.CanvasObject {
	return s.content
}

// SetFloatingButton 设置浮动按钮
func (s *Scaffold) SetFloatingButton(floatingButton fyne.CanvasObject) {
	s.floatingButton = floatingButton
	s.Refresh()
}

// GetFloatingButton 获取浮动按钮
func (s *Scaffold) GetFloatingButton() fyne.CanvasObject {
	return s.floatingButton
}

type scaffoldRenderer struct {
	scaffold *Scaffold
}

func (r *scaffoldRenderer) Layout(size fyne.Size) {
	contentX, contentY := float32(0), float32(0)
	contentWidth, contentHeight := size.Width, size.Height

	// 布局顶栏
	if r.scaffold.topBar != nil {
		topBarHeight := r.scaffold.topBar.MinSize().Height
		r.scaffold.topBar.Resize(fyne.NewSize(size.Width, topBarHeight))
		r.scaffold.topBar.Move(fyne.NewPos(0, 0))
		contentY += topBarHeight
		contentHeight -= topBarHeight
	}

	// 布局底栏
	if r.scaffold.bottomBar != nil {
		bottomBarHeight := r.scaffold.bottomBar.MinSize().Height
		r.scaffold.bottomBar.Resize(fyne.NewSize(size.Width, bottomBarHeight))
		r.scaffold.bottomBar.Move(fyne.NewPos(0, size.Height-bottomBarHeight))
		contentHeight -= bottomBarHeight
	}

	// 布局侧边栏
	if r.scaffold.sideBar != nil {
		sideBarWidth := r.scaffold.sideBar.MinSize().Width
		sideBarHeight := contentHeight // 去掉 topBar 和 bottomBar 的高度
		r.scaffold.sideBar.Resize(fyne.NewSize(sideBarWidth, sideBarHeight))
		r.scaffold.sideBar.Move(fyne.NewPos(0, contentY))
		contentX += sideBarWidth
		contentWidth -= sideBarWidth
	}

	// 布局内容区域
	if r.scaffold.content != nil {
		r.scaffold.content.Resize(fyne.NewSize(contentWidth, contentHeight))
		r.scaffold.content.Move(fyne.NewPos(contentX, contentY))
	}

	// 布局浮动按钮
	if r.scaffold.floatingButton != nil {
		fbSize := r.scaffold.floatingButton.MinSize()
		r.scaffold.floatingButton.Resize(fbSize)
		r.scaffold.floatingButton.Move(fyne.NewPos(size.Width-fbSize.Width-16, size.Height-fbSize.Height-16))
	}

}

func (r *scaffoldRenderer) MinSize() fyne.Size {
	width, height := float32(0), float32(0)

	if r.scaffold.topBar != nil {
		height += r.scaffold.topBar.MinSize().Height
		width = fyne.Max(width, r.scaffold.topBar.MinSize().Width)
	}
	if r.scaffold.bottomBar != nil {
		height += r.scaffold.bottomBar.MinSize().Height
		width = fyne.Max(width, r.scaffold.bottomBar.MinSize().Width)
	}
	if r.scaffold.sideBar != nil {
		width += r.scaffold.sideBar.MinSize().Width
		height = fyne.Max(height, r.scaffold.sideBar.MinSize().Height)
	}
	if r.scaffold.content != nil {
		width = fyne.Max(width, r.scaffold.content.MinSize().Width)
		height = fyne.Max(height, r.scaffold.content.MinSize().Height)
	}
	return fyne.NewSize(width, height)
}

func (r *scaffoldRenderer) Refresh() {
	// 刷新所有子组件
	if r.scaffold.topBar != nil {
		r.scaffold.topBar.Refresh()
	}
	if r.scaffold.bottomBar != nil {
		r.scaffold.bottomBar.Refresh()
	}
	if r.scaffold.sideBar != nil {
		r.scaffold.sideBar.Refresh()
	}
	if r.scaffold.content != nil {
		r.scaffold.content.Refresh()
	}
	if r.scaffold.floatingButton != nil {
		r.scaffold.floatingButton.Refresh()
	}
}

func (r *scaffoldRenderer) Objects() []fyne.CanvasObject {
	objects := []fyne.CanvasObject{}

	if r.scaffold.topBar != nil {
		objects = append(objects, r.scaffold.topBar)
	}
	if r.scaffold.sideBar != nil {
		objects = append(objects, r.scaffold.sideBar)
	}
	if r.scaffold.content != nil {
		objects = append(objects, r.scaffold.content)
	}
	if r.scaffold.bottomBar != nil {
		objects = append(objects, r.scaffold.bottomBar)
	}
	if r.scaffold.floatingButton != nil {
		objects = append(objects, r.scaffold.floatingButton)
	}

	return objects
}

func (r *scaffoldRenderer) Destroy() {}
