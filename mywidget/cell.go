package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Cell struct {
	widget.BaseWidget
	container *fyne.Container
}

func NewCell() *Cell {
	cell := &Cell{}
	cell.container = container.NewStack()
	cell.ExtendBaseWidget(cell)
	return cell
}

func (c *Cell) CreateRenderer() fyne.WidgetRenderer {
	return &cellRenderer{cell: c}
}

// 泛型方法，获取或者创建特定类型的控件
func GetChildByCell[T fyne.CanvasObject](cell *Cell, creator func() T) T {
	for _, child := range cell.container.Objects {
		if canvasObj, ok := child.(T); ok {
			child.Show()
			return canvasObj
		}
	}

	// 如果没有找到相应类型的对象，通过回调创建
	newChild := creator()
	cell.container.Add(newChild)
	cell.Refresh()
	return newChild
}

type cellRenderer struct {
	cell *Cell
}

func (r *cellRenderer) Layout(size fyne.Size) {
	r.cell.container.Resize(size)
}

func (r *cellRenderer) MinSize() fyne.Size {
	return r.cell.container.MinSize()
}

func (r *cellRenderer) Refresh() {
	r.cell.container.Refresh()
}

func (r *cellRenderer) Objects() []fyne.CanvasObject {
	return r.cell.container.Objects
}

func (r *cellRenderer) Destroy() {
	// 清理操作
}