package commander

import "github.com/gdamore/tcell"

type ListView interface {
	MaxSizeWidget
	Rows() []Row
	SelectedRowId() int
	SelectedRow() Row
	SetStyler(styler ListViewStyler)
}

type ListViewStyler func(list ListView, rowId int, row Row) tcell.Style

type Row []string

type ResourceListView interface {
	ListView
	Resource() *Resource
}

type MenuListView interface {
	ListView
	Items() []MenuItem
	SelectedItem() MenuItem
}

type MenuItem interface {
	Title() string
	Widget() Widget
}