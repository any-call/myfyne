package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	SelectItemModel[M any] struct {
		DisplayName string
		Model       M
	}
)

func CreateSearchBox(placeHolder string, width float32, fn func(str string), isSearchBtn bool) fyne.CanvasObject {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder(placeHolder)
	if isSearchBtn {
		searchBtn := widget.NewButton("查询", func() {
			if fn != nil {
				fn(searchEntry.Text)
			}
		})

		return container.NewHBox(
			NewWidthBox(width, searchEntry),
			NewWidthBox(100, searchBtn),
		)
	} else {
		searchEntry.OnChanged = func(s string) {
			if fn != nil {
				fn(searchEntry.Text)
			}
		}
	}

	return NewWidthBox(width, searchEntry)
}

func CreateSelectObject[M any](selectFn func(item M), listItemFn func() []SelectItemModel[M]) fyne.CanvasObject {
	var list []SelectItemModel[M]
	selectObject := widget.NewSelect([]string{}, func(s string) {
		if selectFn != nil {
			if list != nil && len(list) > 0 {
				for i, _ := range list {
					if s == list[i].DisplayName {
						selectFn(list[i].Model)
					}
				}
			}
		}
	})

	if listItemFn != nil {
		go func(cbFun func() []SelectItemModel[M]) {
			list = cbFun()
			if list != nil || len(list) > 0 {
				items := []string{}
				for i, _ := range list {
					items = append(items, list[i].DisplayName)
				}

				selectObject.SetOptions(items)
				selectObject.SetSelectedIndex(0)
				if selectFn != nil {
					selectFn(list[0].Model)
				}
			}

		}(listItemFn)
	}

	return selectObject
}
