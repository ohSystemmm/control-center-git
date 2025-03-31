package layout

import (
	"control-center-git/objective"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func buttonLayout(app *tview.Application) *tview.Grid {
	rows := []int{0, 3, 1, 3, 1, 3, 1, 3, 0, 3, 1}
	cols := []int{0, 25, 2, 25, 2, 25, 0}

	buttonLabels := []string{
		"Change Design", "Wallpaper Type", "Apply Wallpaper",
		"Set Daishow Folder", "Update Fastfetch", "Change Avatar",
		"Update Grub Theme", "Update SDDM Theme", "Blue Light Filter",
		"Exit",
	}

	buttonActions := []func(){
		objective.TempFunc(app), objective.TempFunc(app), objective.TempFunc(app),
		objective.TempFunc(app), objective.TempFunc(app), objective.TempFunc(app),
		objective.TempFunc(app), objective.TempFunc(app), objective.TempFunc(app),
		objective.QuitApp(app),
	}

	grid := tview.NewGrid().
		SetRows(rows...).
		SetColumns(cols...)

	var buttons []*tview.Button
	buttonPositions := make(map[*tview.Button][2]int)

	var quitButton, lastButton *tview.Button

	for i, label := range buttonLabels {
		btn := button(label, buttonActions[i])

		var row, col int
		if i == len(buttonLabels)-1 {
			row, col = 7, 3 // Absolute Position
			quitButton = btn
		} else {
			row = 1 + (i/3)*2 // 1 Row Threshold + Position
			col = 1 + (i%3)*2 // 3 Col Threshold + Position
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

// TODO: Fix Event Handling on last Row
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
