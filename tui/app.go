package main

import (
	"control-center-git/tui"
	"github.com/rivo/tview"
)

func SetupUI() *tview.Application {
	app := tview.NewApplication()
	return app.SetRoot(tui.MasterLayout(app), true)
}
