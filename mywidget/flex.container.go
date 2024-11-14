package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/myfyne"
)

// FlexContainer 是自定义的容器，支持横向和纵向的布局
type FlexContainer struct {
	widget.BaseWidget

	isHorizontal       bool
	mainAxisAlignment  myfyne.MainAxisAlignment
	crossAxisAlignment myfyne.CrossAxisAlignment
	items              []fyne.CanvasObject
}

// NewRow 创建一个新的横向布局容器
func NewRow(mainAxisalignment myfyne.MainAxisAlignment, crossAxisAlignment myfyne.CrossAxisAlignment, items ...fyne.CanvasObject) *FlexContainer {
	container := &FlexContainer{
		isHorizontal:       true,
		mainAxisAlignment:  mainAxisalignment,
		crossAxisAlignment: crossAxisAlignment,
		items:              items,
	}
	container.ExtendBaseWidget(container) // 确保事件处理正确
	return container
}

// NewColumn 创建一个新的纵向布局容器
func NewColumn(mainAxisalignment myfyne.MainAxisAlignment, crossAxisAlignment myfyne.CrossAxisAlignment, items ...fyne.CanvasObject) *FlexContainer {
	container := &FlexContainer{
		isHorizontal:       false,
		mainAxisAlignment:  mainAxisalignment,
		crossAxisAlignment: crossAxisAlignment,
		items:              items,
	}
	container.ExtendBaseWidget(container) // 确保事件处理正确
	return container
}

// CreateRenderer 实现 fyne.Widget 接口，用于创建 FlexContainer 的渲染器
func (f *FlexContainer) CreateRenderer() fyne.WidgetRenderer {
	return &flexContainerRenderer{
		container: f,
	}
}

// flexContainerRenderer 实现 fyne.WidgetRenderer 接口，用于绘制 FlexContainer
type flexContainerRenderer struct {
	container *FlexContainer
}

// Layout 实现 fyne.WidgetRenderer 接口，用于布局 FlexContainer
func (r *flexContainerRenderer) Layout(size fyne.Size) {
	if r.container.isHorizontal {
		r.layoutHorizontal(size)
	} else {
		r.layoutVertical(size)
	}
}

// layoutHorizontal 处理子控件的横向布局
func (r *flexContainerRenderer) layoutHorizontal(size fyne.Size) {
	var totalFixedWidth float32

	// 计算固定宽度
	for _, item := range r.container.items {
		totalFixedWidth += item.MinSize().Width
	}

	// 计算剩余的宽度
	remainingWidth := size.Width - totalFixedWidth
	if remainingWidth < 0 {
		remainingWidth = 0 // 防止出现负值
	}

	// 处理对齐方式
	var startX, spacing float32
	itemCount := len(r.container.items)

	switch r.container.mainAxisAlignment {
	case myfyne.MainAxisAlignSpaceBetween:
		if itemCount > 1 {
			spacing = remainingWidth / float32(itemCount-1)
		}
	case myfyne.MainAxisAlignSpaceAround:
		if itemCount > 0 {
			spacing = remainingWidth / float32(itemCount)
			startX = spacing / 2
		}
	case myfyne.MainAxisAlignSpaceEvenly:
		if itemCount > 0 {
			spacing = remainingWidth / float32(itemCount+1)
			startX = spacing
		}
	case myfyne.MainAxisAlignCenter:
		startX = remainingWidth / 2
	case myfyne.MainAxisAlignEnd:
		startX = remainingWidth
	case myfyne.MainAxisAlignStart:
		startX = 0
	default:
		startX = 0
	}

	// 处理交叉轴对齐方式
	var maxHeight float32
	for _, item := range r.container.items {
		if item.MinSize().Height > maxHeight {
			maxHeight = item.MinSize().Height
		}
	}

	// 重新布局所有子控件，应用对齐方式
	for _, item := range r.container.items {
		switch r.container.crossAxisAlignment {
		case myfyne.CrossAxisAlignStretch:
			item.Resize(fyne.NewSize(item.MinSize().Width, size.Height))
			item.Move(fyne.NewPos(startX, 0))
			break

		case myfyne.CrossAxisAlignStart:
			item.Resize(item.MinSize())
			item.Move(fyne.NewPos(startX, 0))
			break

		case myfyne.CrossAxisAlignEnd:
			item.Resize(item.MinSize())
			item.Move(fyne.NewPos(startX, size.Height-item.Size().Height))
			break

		case myfyne.CrossAxisAlignCenter:
			item.Resize(item.MinSize())
			item.Move(fyne.NewPos(startX, (size.Height-item.Size().Height)/2))
			break
		}

		startX += item.Size().Width + spacing
	}

}

// layoutVertical 处理子控件的纵向布局
func (r *flexContainerRenderer) layoutVertical(size fyne.Size) {
	var totalFixedHeight float32

	// 计算固定高度
	for _, item := range r.container.items {
		totalFixedHeight += item.MinSize().Height
	}

	// 计算剩余的高度
	remainingHeight := size.Height - totalFixedHeight
	if remainingHeight < 0 {
		remainingHeight = 0 // 防止出现负值
	}

	// 处理对齐方式
	var startY, spacing float32
	itemCount := len(r.container.items)

	switch r.container.mainAxisAlignment {
	case myfyne.MainAxisAlignSpaceBetween:
		if itemCount > 1 {
			spacing = remainingHeight / float32(itemCount-1)
		}
	case myfyne.MainAxisAlignSpaceAround:
		if itemCount > 0 {
			spacing = remainingHeight / float32(itemCount)
			startY = spacing / 2
		}
	case myfyne.MainAxisAlignSpaceEvenly:
		if itemCount > 0 {
			spacing = remainingHeight / float32(itemCount+1)
			startY = spacing
		}
	case myfyne.MainAxisAlignCenter:
		startY = remainingHeight / 2
	case myfyne.MainAxisAlignEnd:
		startY = remainingHeight
	case myfyne.MainAxisAlignStart:
		startY = 0
	default:
		startY = 0
	}

	// 重新布局所有子控件，应用对齐方式
	for _, item := range r.container.items {
		switch r.container.crossAxisAlignment {
		case myfyne.CrossAxisAlignStretch:
			item.Resize(fyne.NewSize(size.Width, item.MinSize().Height))
			item.Move(fyne.NewPos(0, startY))
			break

		case myfyne.CrossAxisAlignStart:
			item.Resize(item.MinSize())
			item.Move(fyne.NewPos(0, startY))
			break

		case myfyne.CrossAxisAlignEnd:
			item.Resize(item.MinSize())
			item.Move(fyne.NewPos(size.Width-item.Size().Width, startY))
			break

		case myfyne.CrossAxisAlignCenter:
			item.Resize(item.MinSize())
			item.Move(fyne.NewPos((size.Width-item.Size().Width)/2, startY))
			break
		}

		startY += item.Size().Height + spacing
	}
}

// MinSize 实现 fyne.WidgetRenderer 接口，用于获取 FlexContainer 的最小尺寸
func (r *flexContainerRenderer) MinSize() fyne.Size {
	var totalWidth, totalHeight float32
	var maxWidth, maxHeight float32

	for _, item := range r.container.items {
		minSize := item.MinSize()
		if r.container.isHorizontal {
			totalWidth += minSize.Width
			if minSize.Height > maxHeight {
				maxHeight = minSize.Height
			}
		} else {
			totalHeight += minSize.Height
			if minSize.Width > maxWidth {
				maxWidth = minSize.Width
			}
		}
	}

	if r.container.isHorizontal {
		return fyne.NewSize(totalWidth, maxHeight)
	}
	return fyne.NewSize(maxWidth, totalHeight)
}

// Refresh 实现 fyne.WidgetRenderer 接口，用于刷新 FlexContainer
func (r *flexContainerRenderer) Refresh() {
	r.container.Refresh()
}

// Objects 实现 fyne.WidgetRenderer 接口，用于获取 FlexContainer 中的所有 CanvasObject
func (r *flexContainerRenderer) Objects() []fyne.CanvasObject {
	return r.container.items
}

// Destroy 实现 fyne.WidgetRenderer 接口，用于销毁 FlexContainer
func (r *flexContainerRenderer) Destroy() {
	// No resources to clean up
}
