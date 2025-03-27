package actions

import (
	"control-center-git/objective"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	ROWS = []int{1, 1, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 0, 3, 1}
	COLS = []int{2, 25, 0, 25, 2, 25, 0, 25, 2}
)

func ActionSection(app *tview.Application) *tview.Grid {
	buttons := make([]*tview.Button, 0) // 0 -> Empty slice
	buttons = append(buttons, button("Quit", objective.TempFunc(app)))

	element := tview.NewGrid().
		SetRows(ROWS...).
		SetColumns(COLS...)

	//element.AddItem(button(), 19, 7, 1, 1, 0, 0, true)

	return element
}

func button(label string, action func()) *tview.Button {
	return tview.NewButton(label).SetSelectedFunc(action)
}

func eventHandler(app *tview.Application, buttons [][]*tview.Button, quitButton *tview.Button) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		var row, col int
		found := false

		for r := range buttons {
			for c := range buttons[r] {
				if app.GetFocus() == buttons[r][c] {
					row, col = r, c
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		switch event.Key() {
		case tcell.KeyUp:
			if row > 0 {
				app.SetFocus(buttons[row-1][col])
			}
		case tcell.KeyDown:
			if row < len(buttons)-1 {
				app.SetFocus(buttons[row+1][col])
			} else {
				app.SetFocus(quitButton)
			}
		case tcell.KeyRight:
			if col < len(buttons[row])-1 {
				app.SetFocus(buttons[row][col+1])
			}
		case tcell.KeyLeft:
			if col > 0 {
				app.SetFocus(buttons[row][col-1])
			}
		case tcell.KeyRune:
			if event.Rune() == 'q' {
				app.SetFocus(quitButton)
			}
		default:
			return event
		}
		return nil
	}
}
