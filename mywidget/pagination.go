package mywidget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Pagination struct {
	widget.BaseWidget
	grid          *widget.GridWrap
	totalPages    int            //总页数
	currentPage   int            //当前页
	onPageChange  func(page int) // 页码改变的回调函数
	totalCellSize fyne.Size
}

func NewPagination(onPageChange func(page int)) *Pagination {
	p := &Pagination{
		totalPages:   0,
		currentPage:  0,
		onPageChange: onPageChange,
	}
	p.grid = widget.NewGridWrap(p.getTotalPages, p.getPageItemObject, p.updatePageItem)
	p.grid.OnSelected = p.selectPage
	p.ExtendBaseWidget(p)
	return p
}

func (p *Pagination) UpdatePaginationData(newTotalPages, newCurrentPage int) {
	if newTotalPages < 0 {
		newTotalPages = 0
	}

	if newCurrentPage < 0 {
		newCurrentPage = 1
	}

	if newCurrentPage > newTotalPages {
		newCurrentPage = newTotalPages
	}

	p.totalPages = newTotalPages
	p.currentPage = newCurrentPage
	p.grid.Refresh() // 刷新分页显示
}

func (p *Pagination) CreateRenderer() fyne.WidgetRenderer {
	// 创建分页网格
	return &paginationRenderer{pagination: p}
}

func (p *Pagination) getTotalPages() int {
	p.totalCellSize = fyne.NewSize(0, 0)
	return p.totalPages
}

func (p *Pagination) getPageItemObject() fyne.CanvasObject {
	// 创建分页按钮对象
	cell := NewCell()
	_ = SetCellChild(cell, func() *SizedBox {
		lab := widget.NewLabel("")
		lab.Alignment = fyne.TextAlignCenter
		return NewSizedBox(fyne.NewSize(40, 40), nil, lab)
	})

	return cell
}

func (p *Pagination) updatePageItem(id widget.GridWrapItemID, object fyne.CanvasObject) {
	if box, ok := GetCellChild[*SizedBox](object.(*Cell)); ok {
		lab := box.GetChild().(*widget.Label)
		lab.Text = fmt.Sprintf(" %d ", id+1)
		if int(id+1) == p.currentPage {
			lab.Importance = widget.HighImportance // 高亮当前页
		} else {
			lab.Importance = widget.MediumImportance
		}
		lab.Refresh()
	}
}

func (p *Pagination) selectPage(id widget.GridWrapItemID) {
	// 选择某一页时的处理
	p.currentPage = int(id + 1)
	p.onPageChange(p.currentPage) // 调用页码改变的回调函数
	p.Refresh()                   // 刷新组件
}

type paginationRenderer struct {
	pagination *Pagination
}

func (r *paginationRenderer) MinSize() fyne.Size {
	return r.pagination.grid.MinSize()
}

func (r *paginationRenderer) Layout(size fyne.Size) {
	r.pagination.grid.Resize(size)
}

func (r *paginationRenderer) Refresh() {
	r.pagination.grid.Refresh()
}

func (r *paginationRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.pagination.grid}
}

func (r *paginationRenderer) Destroy() {}
