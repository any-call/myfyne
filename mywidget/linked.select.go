package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type LinkedSelect[M comparable] struct {
	sel   *widget.Select
	items []SelectItemModel[M]

	keyIndex map[string]int // Key -> index
	modelMap map[M]string   // Model -> Key

	silent   bool //程序自己改值时，不要执行 onChange
	onChange func(M)
}

func NewLinkedSelect[M comparable](
	onChange func(M),
	configFn func(*widget.Select),
) *LinkedSelect[M] {

	ls := &LinkedSelect[M]{
		keyIndex: make(map[string]int),
		modelMap: make(map[M]string),
		onChange: onChange,
	}

	ls.sel = widget.NewSelect(nil, func(label string) {

		if ls.silent {
			return
		}

		if idx, ok := ls.keyIndex[label]; ok {
			if ls.onChange != nil {
				ls.onChange(ls.items[idx].Model)
			}
		}
	})

	if configFn != nil {
		configFn(ls.sel)
	}

	return ls
}

func (ls *LinkedSelect[M]) Update(
	defaultModel *M,
	list []SelectItemModel[M],
) {

	ls.silent = true
	defer func() { ls.silent = false }()

	ls.items = list
	ls.keyIndex = make(map[string]int, len(list))
	ls.modelMap = make(map[M]string, len(list))

	options := make([]string, 0, len(list))
	selectIndex := -1

	for i, item := range list {

		options = append(options, item.DisplayName)

		ls.keyIndex[item.DisplayName] = i
		ls.modelMap[item.Model] = item.DisplayName

		if defaultModel != nil && item.Model == *defaultModel {
			selectIndex = i
		}
	}

	ls.sel.SetOptions(options)

	if selectIndex >= 0 {
		ls.sel.SetSelectedIndex(selectIndex)
	} else {
		ls.sel.ClearSelected()
	}
}

func (ls *LinkedSelect[M]) SetSelectedByModel(m M) {

	if label, ok := ls.modelMap[m]; ok {
		ls.silent = true
		ls.sel.SetSelected(label)
		ls.silent = false
	}
}

func (ls *LinkedSelect[M]) GetSelected() (M, bool) {
	label := ls.sel.Selected
	if label == "" {
		var zero M
		return zero, false
	}

	if idx, ok := ls.keyIndex[label]; ok {
		return ls.items[idx].Model, true
	}

	var zero M
	return zero, false
}

func (ls *LinkedSelect[M]) Object() fyne.CanvasObject {
	return ls.sel
}
