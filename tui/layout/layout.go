package layout

import (
	"control-center-git/tui/actions"
	"github.com/rivo/tview"
)

var (
	ROWS    = []int{3, 0, 3}
	COLUMNS = []int{8, 0, 0, 8}
)

func MasterLayout(app *tview.Application) *tview.Grid {

	masterLayout := tview.NewGrid().
		SetRows(ROWS...).
		SetColumns(COLUMNS...).
		SetBorders(true)

	asdf := PlaceHolder("Asdf")

	masterLayout.
		AddItem(asdf, 1, 1, 1, 1, 0, 0, false).
		AddItem(actions.ActionSection(app), 1, 2, 1, 1, 0, 0, true)

	return masterLayout
}

func PlaceHolder(placeholder string) *tview.TextView {
	return tview.NewTextView().SetText(placeholder).SetTextAlign(tview.AlignCenter)
}
