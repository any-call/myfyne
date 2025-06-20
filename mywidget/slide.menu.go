package mywidget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/myfyne"
	"strconv"
	"strings"
)

// SideMenu 定义，继承 fyne.Widget
type SideMenu struct {
	widget.BaseWidget
	menuItems []myfyne.MenuItemModel
	onSelect  func(item myfyne.MenuItemModel)
	tree      *widget.Tree
}

// 创建一个新的 SideMenu 实例
func NewSideMenu(menuItems []myfyne.MenuItemModel, onSelect func(item myfyne.MenuItemModel)) *SideMenu {
	menu := &SideMenu{
		menuItems: menuItems,
		onSelect:  onSelect,
	}

	// 初始化 tree 的逻辑
	menu.tree = widget.NewTree(
		// 获取子节点
		func(uid string) (children []string) {
			if uid == "" {
				children := make([]string, 0, len(menuItems))
				for i, item := range menuItems {
					if !item.IsHidden { // 过滤掉隐藏的菜单项
						children = append(children, fmt.Sprintf("%d", i))
					}
				}

				return children
			}

			uids, err := menu.parseUID(uid)
			if err != nil {
				return []string{}
			}

			item := menu.findMenuItemByUID(menuItems, uids)
			if item == nil {
				return []string{}
			}

			children = make([]string, 0, len(item.SubItems))
			for i, subItem := range item.SubItems {
				if !subItem.IsHidden { // 过滤掉隐藏的子菜单项
					children = append(children, fmt.Sprintf("%s-%d", uid, i))
				}
			}

			return children
		},
		// 判断是否为分支节点
		func(uid string) bool {
			if uid == "" {
				return true
			}

			uids, err := menu.parseUID(uid)
			if err != nil {
				return false
			}

			item := menu.findMenuItemByUID(menuItems, uids)
			return item != nil && len(item.SubItems) > 0
		},
		// 创建节点，包含图标和文本
		func(branch bool) fyne.CanvasObject {
			hbox := container.NewHBox(
				widget.NewIcon(nil), // 图标
				widget.NewLabel(""), // 文本
			)
			return hbox
		},
		// 更新节点内容
		func(uid string, branch bool, node fyne.CanvasObject) {
			hbox := node.(*fyne.Container)
			icon := hbox.Objects[0].(*widget.Icon)
			label := hbox.Objects[1].(*widget.Label)
			if uid == "" {
				label.SetText("")
				return
			}

			uids, err := menu.parseUID(uid)
			if err != nil {
				return
			}

			item := menu.findMenuItemByUID(menuItems, uids)
			if item == nil {
				return
			}

			// 只在有图标时设置
			if item.Icon != nil {
				icon.SetResource(item.Icon)
				icon.Show() // 显示图标
			} else {
				icon.Hide() // 隐藏图标
			}

			label.SetText(item.Name)
		},
	)

	// 设置根节点
	menu.tree.Root = ""

	// 监听点击事件
	menu.tree.OnSelected = func(uid string) {
		if uid == "" {
			return
		}

		uids, err := menu.parseUID(uid)
		if err != nil {
			return
		}

		item := menu.findMenuItemByUID(menuItems, uids)
		if item != nil && item.OnTapCb != nil {
			item.OnTapCb(item.Name) //先调用menu自带的回调
		}

		if item != nil && menu.onSelect != nil {
			menu.onSelect(*item) // 触发点击回调
		}
	}

	menu.tree.OpenAllBranches()
	menu.ExtendBaseWidget(menu) // 注册自定义控件
	return menu
}

// CreateRenderer 实现自定义控件的渲染
func (m *SideMenu) CreateRenderer() fyne.WidgetRenderer {
	scroll := container.NewScroll(m.tree)
	return widget.NewSimpleRenderer(scroll)
}

// 辅助函数：递归查找节点
func (m *SideMenu) findMenuItemByUID(items []myfyne.MenuItemModel, uids []int) *myfyne.MenuItemModel {
	if len(uids) == 0 {
		return nil
	}

	item := &items[uids[0]]
	if len(uids) == 1 {
		return item
	}

	return m.findMenuItemByUID(item.SubItems, uids[1:])
}

// 辅助函数：将 uid 转换为整数 slice
func (m *SideMenu) parseUID(uid string) ([]int, error) {
	parts := strings.Split(uid, "-")
	uids := make([]int, len(parts))
	var err error
	for i, part := range parts {
		uids[i], err = strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
	}
	return uids, nil
}

func (m *SideMenu) GetTree() *widget.Tree {
	return m.tree
}

func (m *SideMenu) SelectNames(names ...string) {
	if node, b := m.FindNodeID(names...); b {
		m.tree.Select(node)
	}
}

func (m *SideMenu) FindNodeID(names ...string) (string, bool) {
	var path []int

	var search func(items []myfyne.MenuItemModel, depth int) bool
	search = func(items []myfyne.MenuItemModel, depth int) bool {
		if depth >= len(names) {
			return false
		}
		for i, item := range items {
			if item.Name == names[depth] {
				path = append(path, i)
				if depth == len(names)-1 {
					return true
				}
				if search(item.SubItems, depth+1) {
					return true
				}
				// 回溯
				path = path[:len(path)-1]
			}
		}
		return false
	}

	if search(m.menuItems, 0) {
		// 转换为 "0-1-2" 形式
		var parts []string
		for _, idx := range path {
			parts = append(parts, strconv.Itoa(idx))
		}
		return strings.Join(parts, "-"), true
	}
	return "", false
}
