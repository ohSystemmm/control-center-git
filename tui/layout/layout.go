package layout

import (
	"control-center-git/tui/actions"
	"github.com/rivo/tview"
)

var (
	ROWS    = []int{3, 0, 3}
	COLUMNS = []int{8, 0, 8}
)

func MasterLayout(app *tview.Application) *tview.Grid {

	masterLayout := tview.NewGrid().
		SetRows(ROWS...).
		SetColumns(COLUMNS...).
		SetBorders(true)

	masterLayout.
		AddItem(actions.ActionSection(app), 1, 1, 1, 1, 0, 0, true)

	return masterLayout
}