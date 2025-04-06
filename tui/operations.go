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
