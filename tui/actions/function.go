package actions

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	ROWS = []int{1, 1, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 0}
	COLS = []int{2, 10, 2, 10, 2, 0}
)

func ActionSection(app *tview.Application) *tview.Grid {

	title := tview.NewTextView().SetText("Actions").SetTextAlign(tview.AlignCenter)

	a := button()
	b := button()
	c := button()
	d := button()
	e := button()
	f := button()
	g := button()
	h := button()
	i := button()
	j := button()

	element := tview.NewGrid().
		SetRows(ROWS...).
		SetColumns(COLS...)

	element.
		AddItem(title, 1, 1, 1, 3, 0, 0, false).
		AddItem(a, 3, 1, 1, 1, 0, 0, true).
		AddItem(b, 3, 3, 1, 1, 0, 0, false).
		AddItem(c, 5, 1, 1, 1, 0, 0, false).
		AddItem(d, 5, 3, 1, 1, 0, 0, false).
		AddItem(e, 7, 1, 1, 1, 0, 0, false).
		AddItem(f, 7, 3, 1, 1, 0, 0, false).
		AddItem(g, 9, 1, 1, 1, 0, 0, false).
		AddItem(h, 9, 3, 1, 1, 0, 0, false).
		AddItem(i, 11, 1, 1, 1, 0, 0, false).
		AddItem(j, 11, 3, 1, 1, 0, 0, false)

	app.SetInputCapture(switchFocus(app, a, b, c, d))

	return element
}

func button() *tview.Button {
	return tview.NewButton("Actions")
}

func switchFocus(app *tview.Application, b1, b2, b3, b4 tview.Primitive) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if app.GetFocus() == b3 {
				app.SetFocus(b1)
			}
			if app.GetFocus() == b4 {
				app.SetFocus(b2)
			}

		case tcell.KeyDown:
			if app.GetFocus() == b1 {
				app.SetFocus(b3)
			}
			if app.GetFocus() == b2 {
				app.SetFocus(b4)
			}

		case tcell.KeyRight:
			if app.GetFocus() == b1 {
				app.SetFocus(b2)
			}
			if app.GetFocus() == b3 {
				app.SetFocus(b4)
			}

		case tcell.KeyLeft:
			if app.GetFocus() == b4 {
				app.SetFocus(b3)
			}
			if app.GetFocus() == b2 {
				app.SetFocus(b1)
			}
		}
		return event
	}
}
