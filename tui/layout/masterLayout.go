package layout

import (
	"github.com/rivo/tview"
)

func MasterLayout(app *tview.Application) *tview.Grid {

	rows := []int{3, 0, 3}
	columns := []int{8, 0, 8}

	masterLayout := tview.NewGrid().
		SetRows(rows...).
		SetColumns(columns...).
		SetBorders(true)

	masterLayout.
		AddItem(buttonLayout(app), 1, 1, 1, 1, 0, 0, true)

	return masterLayout
}
