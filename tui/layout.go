package layout

import (
	"control-center-git/objective"
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

func buttonLayout(app *tview.Application) *tview.Grid {
	buttonLabels := []string{
		"Change Design", "Wallpaper Type", "Apply Wallpaper",
		"Set Daishow Folder", "Update Fastfetch", "Change Avatar",
		"Update Grub Theme", "Update SDDM Theme", "Blue Light Filter",
		"Exit",
	}

	buttonActions := []func(){
		objective.ChangeDesign(app), objective.WallpaperType(app), objective.ApplyWallpaper(app),
		objective.DiashowFolder(app), objective.UpdateFastfetch(app), objective.ChangeAvatar(app),
		objective.UpdateGrubTheme(app), objective.UpdateSddmTheme(app), objective.BlueLightFilter(app),
		objective.QuitApp(app),
	}

	grid := tview.NewGrid().
		SetRows(0, 3, 1, 3, 1, 3, 1, 3, 0).
		SetColumns(0, 25, 2, 25, 2, 25, 0)

	var buttons []*tview.Button
	buttonPositions := make(map[*tview.Button][2]int)

	var quitButton, lastButton *tview.Button

	for i, label := range buttonLabels {
		btn := button(label, buttonActions[i])

		var row, col int
		if i == len(buttonLabels)-1 {
			row, col = 7, 5
			quitButton = btn
		} else {
			row = 1 + (i/3)*2
			col = 1 + (i%3)*2
			lastButton = btn
		}

		grid.AddItem(btn, row, col, 1, 1, 0, 0, i == 0)
		buttons = append(buttons, btn)
		buttonPositions[btn] = [2]int{row, col}
	}

	grid.SetInputCapture(eventHandler(buttons, buttonPositions, quitButton, lastButton, app))

	return grid
}

func button(label string, action func()) *tview.Button {
	return tview.NewButton(label).SetSelectedFunc(action)
}

func eventHandler(buttons []*tview.Button, buttonPositions map[*tview.Button][2]int, quitButton, lastButton *tview.Button, app *tview.Application) func(event *tcell.EventKey) *tcell.EventKey {
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
			if currentButton == lastButton {
				nextButton = quitButton
			} else {
				nextButton = findButtonAt(currentRow, currentCol+2, buttonPositions)
			}
		case tcell.KeyDown:
			if currentButton == lastButton {
				nextButton = quitButton
			} else {
				nextButton = findButtonAt(currentRow+2, currentCol, buttonPositions)
			}
		case tcell.KeyLeft:
			if currentButton == quitButton {
				nextButton = lastButton
			} else {
				nextButton = findButtonAt(currentRow, currentCol-2, buttonPositions)
			}
		case tcell.KeyUp:
			if currentButton == quitButton {
				nextButton = lastButton
			} else {
				nextButton = findButtonAt(currentRow-2, currentCol, buttonPositions)
			}
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
