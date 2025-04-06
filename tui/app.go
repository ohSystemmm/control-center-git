package tui

import (
	"github.com/rivo/tview"
)

func SetupUI() *tview.Application {
	app := tview.NewApplication()
	return app.SetRoot(MasterLayout(app), true)
}
