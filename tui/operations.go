package tui

import (
	"os"

	"github.com/rivo/tview"
)

func TempFunc(app *tview.Application) func() {
	return func() {
		app.Stop()
		os.Exit(0)
	}
}

func ApplyWallpaper(app *tview.Application) func() {
	return func() {
	}
}

func QuitApp(app *tview.Application) func() {
	return func() {
		app.Stop()
		os.Exit(0)
	}
}

func ChangeDesign(app *tview.Application) func() {
	return func() {
	}
}

func WallpaperType(app *tview.Application) func() {
	return func() {
	}
}

func DiashowFolder(app *tview.Application) func() {
	return func() {
	}
}

func UpdateFastfetch(app *tview.Application) func() {
	return func() {
	}
}

func ChangeAvatar(app *tview.Application) func() {
	return func() {
	}
}

func UpdateGrubTheme(app *tview.Application) func() {
	return func() {
	}
}

func UpdateSddmTheme(app *tview.Application) func() {
	return func() {
	}
}

func BlueLightFilter(app *tview.Application) func() {
	return func() {
	}
}
func KeyboardLayout(app *tview.Application) func() {
	return func() {
	}
}
func UnlockDatabase(app *tview.Application) func() {
	return func() {
	}
}
func RegeneratePywal(app *tview.Application) func() {
	return func() {
	}
}
func XDGSettings(app *tview.Application) func() {
	return func() {
	}
}

func ChangeShell(app *tview.Application) func() {
	return func() {
	}
}
