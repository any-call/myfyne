package mywidget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

type Pagination struct {
	widget.BaseWidget
	grid          *widget.GridWrap
	totalPages    int            //总页数
	currentPage   int            //当前页
	pagesPerBatch int            //快进多页值
	onPageChange  func(page int) // 页码改变的回调函数
}

func NewPagination(onPageChange func(page int)) *Pagination {
	p := &Pagination{
		totalPages:    0,
		currentPage:   0,
		pagesPerBatch: 5, //这个5固定下来
		onPageChange:  onPageChange,
	}
	p.grid = widget.NewGridWrap(p.getTotalItems, p.getPageItemObject, p.updatePageItem)
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

// 根据pagesPerBath 计算最大显示的cell数,包括 <  >
func (p *Pagination) matchMaxCells() int {
	return p.pagesPerBatch + 4 + 2
}

func (p *Pagination) getTotalItems() int {
	if p.totalPages <= p.pagesPerBatch+4 {
		return p.totalPages + 2
	}

	if p.currentPage <= 3 || p.currentPage > (p.totalPages-3) {
		return p.matchMaxCells() - 2
	} else if p.currentPage == 4 || p.currentPage == (p.totalPages-3) {
		return p.matchMaxCells() - 1
	} else {
		return p.matchMaxCells()
	}
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

func (p *Pagination) handlePagination(id widget.GridWrapItemID, firstItemCb func(), latestItemCb func(), noBatchCb func(), batchLoss2PrefixCb func(), batchLoss2PSuffixCb func(), batchLoss1PrefixCb func(), batchLoss1SuffixCb func(), batch2Cb func()) {
	switch id {
	case 0:
		if firstItemCb != nil {
			firstItemCb()
		}
		break
	case p.getTotalItems() - 1:
		if latestItemCb != nil {
			latestItemCb()
		}
		break
	default:
		{
			if p.totalPages <= p.pagesPerBatch+4 { //无batch
				if noBatchCb != nil {
					noBatchCb()
				}
			} else {
				if p.currentPage <= 3 || p.currentPage > (p.totalPages-3) { //有batch 短2
					if p.currentPage <= 3 {
						if batchLoss2PrefixCb != nil {
							batchLoss2PrefixCb()
						}
					} else {
						if batchLoss2PSuffixCb != nil {
							batchLoss2PSuffixCb()
						}
					}
				} else if p.currentPage == 4 || p.currentPage == (p.totalPages-3) { //有batch 短1
					if p.currentPage <= 4 { //前显
						if batchLoss1PrefixCb != nil {
							batchLoss1PrefixCb()
						}
					} else { //后显
						if batchLoss1SuffixCb != nil {
							batchLoss1SuffixCb()
						}
					}
				} else { //batch 显示两个
					if batch2Cb != nil {
						batch2Cb()
					}
				}
			}
		}
		break
	}

}

func (p *Pagination) updatePageItem(id widget.GridWrapItemID, object fyne.CanvasObject) {
	if box, ok := GetCellChild[*SizedBox](object.(*Cell)); ok {
		lab := box.GetChild().(*widget.Label)

		p.handlePagination(id,
			func() {
				lab.Text = "←"
			},
			func() {
				lab.Text = "→"
			},
			func() {
				lab.Text = fmt.Sprintf(" %d ", id)
			},
			func() {
				if id <= p.pagesPerBatch {
					lab.Text = fmt.Sprintf(" %d ", id)
				} else if id == p.pagesPerBatch+1 {
					lab.Text = ">>"
				} else {
					lab.Text = fmt.Sprintf(" %d ", p.totalPages)
				}
			},
			func() {
				if id == 1 {
					lab.Text = fmt.Sprintf(" %d ", id)
				} else if id == 2 {
					lab.Text = "<<"
				} else {
					lab.Text = fmt.Sprintf(" %d ", p.totalPages-p.pagesPerBatch+id-2)
				}
			},
			func() {
				if id < p.pagesPerBatch+2 {
					lab.Text = fmt.Sprintf(" %d ", id)
				} else if id == p.pagesPerBatch+2 { //>>
					lab.Text = ">>"
				} else {
					lab.Text = fmt.Sprintf(" %d ", p.totalPages)
				}
			},
			func() {
				if id == 1 {
					lab.Text = fmt.Sprintf(" %d ", id)
				} else if id == 2 { //<<
					lab.Text = "<<"
				} else {
					lab.Text = fmt.Sprintf(" %d ", p.totalPages-p.pagesPerBatch+id-3)
				}
			},
			func() {
				switch id {
				case 1: //1
					lab.Text = fmt.Sprintf(" %d ", id)
					break
				case 2: //<<
					lab.Text = "<<"
					break
				case 8: //>>
					lab.Text = ">>"
					break
				case 9: //最后一个
					lab.Text = fmt.Sprintf(" %d ", p.totalPages)
					break
				default:
					if id == 3 {
						lab.Text = fmt.Sprintf(" %d ", p.currentPage-2)
					} else if id == 4 {
						lab.Text = fmt.Sprintf(" %d ", p.currentPage-1)
					} else if id == 6 {
						lab.Text = fmt.Sprintf(" %d ", p.currentPage+1)
					} else if id == 7 {
						lab.Text = fmt.Sprintf(" %d ", p.currentPage+2)
					} else {
						lab.Text = fmt.Sprintf(" %d ", p.currentPage)
					}
					break
				}
			},
		)

		if intV, err := strconv.Atoi(strings.TrimSpace(lab.Text)); err == nil {
			if intV == p.currentPage {
				lab.Importance = widget.HighImportance // 高亮当前页
			} else {
				lab.Importance = widget.MediumImportance
			}
		} else {
			lab.Importance = widget.MediumImportance
		}
		lab.Refresh()
	}

}

func (p *Pagination) selectPage(id widget.GridWrapItemID) {
	p.handlePagination(id,
		func() { //first item
			if p.currentPage > 1 {
				p.currentPage--
			}
		},
		func() { //latest item
			if p.currentPage < p.totalPages {
				p.currentPage++
			}
		},
		func() { //no batch
			p.currentPage = id
		},
		func() { //batch loss 2 prefix
			if id == 6 { //点击>>
				p.currentPage += p.pagesPerBatch
			} else if id == 7 { //点击最后一个
				p.currentPage = p.totalPages
			} else {
				p.currentPage = id
			}
		},
		func() { ////batch loss 2 suffix
			if id == 2 { //点击<<
				p.currentPage -= p.pagesPerBatch
			} else if id == 1 { //点击首位
				p.currentPage = id
			} else {
				p.currentPage = p.totalPages - p.pagesPerBatch + id - 2
			}
		},
		func() { //batch loss 1 prefix
			if id == 7 { //点击>>
				p.currentPage += p.pagesPerBatch
			} else if id == 8 { //点击最后一位
				p.currentPage = p.totalPages
			} else {
				p.currentPage = id
			}
		},
		func() { //batch loss 1 suffix
			if id == 2 { //点击<<
				p.currentPage -= p.pagesPerBatch
			} else if id == 1 { //点击第一个
				p.currentPage = id
			} else {
				p.currentPage = p.totalPages - p.pagesPerBatch + id - 3
			}
		},
		func() { //batch full
			switch id {
			case 1: //点击1
				p.currentPage = id
				break
			case 2: //点击<<
				p.currentPage -= p.pagesPerBatch
				break
			case 8: //点击>>
				p.currentPage += p.pagesPerBatch
				break
			case 9: //点击最后一个
				p.currentPage = p.totalPages
				break
			default:
				if id == 3 {
					p.currentPage -= 2
				} else if id == 4 {
					p.currentPage -= 1
				} else if id == 6 {
					p.currentPage += 1
				} else if id == 7 {
					p.currentPage += 2
				} else {
					p.currentPage = p.currentPage
				}
				break
			}
		},
	)

	if p.currentPage > p.totalPages {
		p.currentPage = p.totalPages
	}

	if p.currentPage < 1 {
		p.currentPage = 1
	}

	p.onPageChange(p.currentPage) // 调用页码改变的回调函数
	p.grid.Unselect(id)
	p.Refresh() // 刷新组件
}

type paginationRenderer struct {
	pagination *Pagination
}

func (r *paginationRenderer) MinSize() fyne.Size {
	size := r.pagination.grid.MinSize()
	width := (size.Width + fyne.CurrentApp().Settings().Theme().Size("padding")) * float32(r.pagination.getTotalItems())
	return fyne.NewSize(width, size.Height)
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

func CreatePaginationSize(title string, limit int, fn func(limit int)) fyne.CanvasObject {
	enter := NewEntryByInt(limit)
	enter.OnSubmitted = func(s string) {
		if intV, err := strconv.Atoi(s); err == nil {
			if fn != nil {
				fn(intV)
			}
		}
	}
	if title == "" {
		title = "设置分页大小："
	}

	enter.SetText(fmt.Sprintf("%d", limit))
	formItem := widget.NewFormItem(title, enter)
	return widget.NewForm(formItem)
}
