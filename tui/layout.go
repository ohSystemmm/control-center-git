package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MasterLayout(app *tview.Application) *tview.Grid {
	return tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(8, 0, 8).
		SetBorders(true).
		AddItem(buttonLayout(app), 1, 1, 1, 1, 0, 0, true)
}

func button(label string, action func()) *tview.Button {
	return tview.NewButton(label).SetSelectedFunc(action)
}

func buttonLayout(app *tview.Application) *tview.Grid {
	buttonLabels := []string{
		"Appearance", "Wallpaper", "Cleanup",
		"Fastfetch Image", "Lockscreen Avatar", "GTK Settings",
		"GRUB2 Theme", "SDDM Theme", "Night Mode",
		"Keyboard Layout", "Unlock Database", "Regenerate Pywal",
		"XDG", "Change Shell", "Temp Function", "Quit",
	}

	buttonActions := []func(){
		func() { ChangeDesign(app) }, func() { WallpaperType(app) }, func() { ApplyWallpaper(app) },
		func() { DiashowFolder(app) }, func() { UpdateFastfetch(app) }, func() { ChangeAvatar(app) },
		func() { UpdateGrubTheme(app) }, func() { UpdateSddmTheme(app) }, func() { BlueLightFilter(app) },
		func() { KeyboardLayout(app) }, func() { UnlockDatabase(app) }, func() { RegeneratePywal(app) },
		func() { XDGSettings(app) }, func() { ChangeShell(app) }, func() { TempFunc(app) }, func() { QuitApp(app) },
	}

	buttons := make([]*tview.Button, len(buttonLabels))
	buttonPositions := make(map[*tview.Button][2]int)

	grid := tview.NewGrid().
		SetRows(1, 3, 0, 5, 1, 5, 1, 5, 1, 5, 1, 5, 0).
		SetColumns(0, 30, 2, 30, 2, 30, 2, 30, 0)

	for i, label := range buttonLabels {
		buttons[i] = button(label, buttonActions[i])
		row, col := 3+(i/4)*2, 1+(i%4)*2
		buttonPositions[buttons[i]] = [2]int{row, col}
		grid.AddItem(buttons[i], row, col, 1, 1, 0, 0, true)
	}

	app.SetInputCapture(eventHandler(buttons, buttonPositions, app))

	return grid
}

func eventHandler(buttons []*tview.Button, buttonPositions map[*tview.Button][2]int,
	app *tview.Application) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		currentItem := app.GetFocus()
		var currentButton *tview.Button

		for _, btn := range buttons {
			if btn == currentItem {
				currentButton = btn
				break
			}
		}

		if currentButton == nil {
			return event
		}

		currentPos := buttonPositions[currentButton]
		currentRow, currentCol := currentPos[0], currentPos[1]

		var nextButton *tview.Button
		switch event.Key() {
		case tcell.KeyRight:
			nextButton = findButtonAt(currentRow, currentCol+2, buttonPositions)
		case tcell.KeyDown:
			nextButton = findButtonAt(currentRow+1, currentCol, buttonPositions)
		case tcell.KeyLeft:
			nextButton = findButtonAt(currentRow, currentCol-2, buttonPositions)
		case tcell.KeyUp:
			nextButton = findButtonAt(currentRow-1, currentCol, buttonPositions)
		}

		if nextButton != nil {
			app.SetFocus(nextButton)
			return nil
		}

		return event
	}
}

func findButtonAt(row, col int, buttonPositions map[*tview.Button][2]int) *tview.Button {
	for btn, pos := range buttonPositions {
		if pos[0] == row && pos[1] == col {
			return btn
		}
	}
	return nil
}
