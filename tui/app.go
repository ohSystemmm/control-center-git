package tui

import (
	"control-center-git/tui/layout"
	"github.com/rivo/tview"
)

func SetupUI() *tview.Application {
	app := tview.NewApplication()
	masterGrid := layout.MasterLayout(app)
	app.SetRoot(masterGrid, true)
	return app
}
