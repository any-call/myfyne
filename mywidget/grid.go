package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Grid struct {
	widget.BaseWidget
	children    []fyne.CanvasObject
	columns     int
	borderWidth float32
	borderColor color.Color
}

func NewGrid(listItem []fyne.CanvasObject, columns int, borderWidth float32, borderColor color.Color) *Grid {
	grid := &Grid{
		children:    listItem,
		columns:     columns,
		borderWidth: borderWidth,
		borderColor: borderColor,
	}
	grid.ExtendBaseWidget(grid)
	return grid
}

func (g *Grid) CreateRenderer() fyne.WidgetRenderer {
	border := &canvas.Rectangle{
		FillColor:    color.Transparent,
		StrokeColor:  g.borderColor,
		StrokeWidth:  g.borderWidth,
		CornerRadius: 2,
	}
	lines := []fyne.CanvasObject{}

	// Add internal lines for grid
	for i := 0; i < len(g.children); i++ {
		lines = append(lines, canvas.NewRectangle(g.borderColor)) // Horizontal and vertical lines
	}

	objects := append(lines, g.children...)
	objects = append(objects, border) // Add the border to the render objects

	return &gridRenderer{
		grid:    g,
		border:  border,
		lines:   lines,
		objects: objects,
	}
}

type gridRenderer struct {
	grid       *Grid
	background *canvas.Rectangle
	border     *canvas.Rectangle
	lines      []fyne.CanvasObject
	objects    []fyne.CanvasObject
}

func (r *gridRenderer) Layout(size fyne.Size) {
	cellWidth := (size.Width - float32(r.grid.columns-1)*r.grid.borderWidth) / float32(r.grid.columns)
	cellHeight := (size.Height - float32(len(r.grid.children)/r.grid.columns-1)*r.grid.borderWidth) / float32(len(r.grid.children)/r.grid.columns)

	// Layout children and lines
	lineIndex := 0
	for i, child := range r.grid.children {
		col := i % r.grid.columns
		row := i / r.grid.columns

		x := float32(col) * (cellWidth + r.grid.borderWidth)
		y := float32(row) * (cellHeight + r.grid.borderWidth)

		if child != nil {
			// Position child
			child.Resize(fyne.NewSize(cellWidth, cellHeight))
			child.Move(fyne.NewPos(x, y))
		}

		// Add horizontal lines
		if row > 0 && col == 0 {
			r.lines[lineIndex].Resize(fyne.NewSize(size.Width, r.grid.borderWidth))
			r.lines[lineIndex].Move(fyne.NewPos(0, y-r.grid.borderWidth/2))
			lineIndex++
		}
		// Add vertical lines
		if col > 0 {
			r.lines[lineIndex].Resize(fyne.NewSize(r.grid.borderWidth, cellHeight-r.grid.borderWidth))
			r.lines[lineIndex].Move(fyne.NewPos(x-r.grid.borderWidth/2, y))
			lineIndex++
		}
	}

	// Draw the outer border
	r.border.Resize(size)
	r.border.Move(fyne.NewPos(0, 0))
}

func (r *gridRenderer) MinSize() fyne.Size {
	// 计算每个单元格的最小宽度和高度
	var cellWidth, cellHeight float32
	if len(r.grid.children) > 0 {
		cellWidth = r.grid.children[0].MinSize().Width
		cellHeight = r.grid.children[0].MinSize().Height
	}

	// 计算总宽度：列数 * 单元格宽度 + (列数-1) * 边框宽度
	totalWidth := float32(r.grid.columns)*cellWidth + float32(r.grid.columns-1)*r.grid.borderWidth
	// 计算总高度：行数 * 单元格高度 + (行数-1) * 边框宽度
	// 假设每行有 r.grid.Columns 个元素
	totalHeight := float32(len(r.grid.children)/r.grid.columns)*cellHeight + float32(len(r.grid.children)/r.grid.columns-1)*r.grid.borderWidth

	// 考虑边框的宽度
	totalWidth += 2 * r.grid.borderWidth  // 左右边框
	totalHeight += 2 * r.grid.borderWidth // 上下边框

	return fyne.NewSize(totalWidth, totalHeight)
}

func (r *gridRenderer) Refresh() {
	canvas.Refresh(r.grid)
}

func (r *gridRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *gridRenderer) Destroy() {}
