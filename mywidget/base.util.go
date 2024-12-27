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

func CreateSearchBox(placeHolder string, width float32, configFn func(entry *widget.Entry), isSearchBtn bool) fyne.CanvasObject {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder(placeHolder)
	if configFn != nil {
		configFn(searchEntry)
	}

	if isSearchBtn {
		searchBtn := widget.NewButton("查询", func() {
			if searchEntry.OnSubmitted != nil {
				searchEntry.OnSubmitted(searchEntry.Text)
			} else if searchEntry.OnChanged != nil {
				searchEntry.OnChanged(searchEntry.Text)
			}
		})

		return container.NewHBox(
			NewWidthBox(width, searchEntry),
			NewWidthBox(100, searchBtn),
		)
	}

	return NewWidthBox(width, searchEntry)
}

func CreateSelect[M any](selectFn func(item M), configFn func(entry *widget.Select), listItemFn func() (defaultV *SelectItemModel[M], list []SelectItemModel[M])) fyne.CanvasObject {
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

	if configFn != nil {
		configFn(selectObject)
	}

	if listItemFn != nil {
		go func(cbFun func() (*SelectItemModel[M], []SelectItemModel[M])) {
			var defV *SelectItemModel[M]
			defV, list = cbFun()
			if list != nil || len(list) > 0 {
				var selectIndex int = 0
				items := []string{}
				for i, _ := range list {
					items = append(items, list[i].DisplayName)
					if defV != nil {
						if defV.DisplayName == list[i].DisplayName {
							selectIndex = i
						}
					}
				}

				selectObject.SetOptions(items)
				selectObject.SetSelectedIndex(selectIndex)
				if selectFn != nil {
					selectFn(list[selectIndex].Model)
				}
			}

		}(listItemFn)
	}

	return selectObject
}
