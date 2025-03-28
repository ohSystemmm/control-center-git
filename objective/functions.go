package objective

import (
	"github.com/rivo/tview"
	"os"
)

func TempFunc(app *tview.Application) func() {
	return func() {
		app.Stop()
		os.Exit(0)
	}
}

func QuitApp(app *tview.Application) func() {
	return func() {
		app.Stop()
		os.Exit(0)
	}
}
