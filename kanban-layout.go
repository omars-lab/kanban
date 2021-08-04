package main

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func l(title string, app *tview.Application) *tview.List {
	l := tview.NewList().
		SetSelectedBackgroundColor(tcell.ColorLightBlue).
		SetSelectedFocusOnly(true).
		AddItem(title+" item 0", "Some explanatory text", '*', nil).
		AddItem(title+" item 1", "Some explanatory text", '*', nil).
		AddItem(title+" item 2", "Some explanatory text", '*', nil).
		AddItem(title+" item 3", "Some explanatory text", '*', nil).
		AddItem(title+" item 4", "Some explanatory text", '*', nil).
		AddItem(title+" item 5", "Some explanatory text", '*', nil).
		AddItem(title+" item 6", "Some explanatory text", '*', nil).
		AddItem(title+" item 7", "Some explanatory text", '*', nil).
		AddItem(title+" item 8", "Some explanatory text", '*', nil).
		AddItem(title+" item 9", "Some explanatory text", '*', nil)
	l.SetBorder(true).SetTitle(title)
	return l
}

func configureList(app *tview.Application, previousList *tview.List, list *tview.List, nextList *tview.List) {
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// fmt.Println(event.Rune())

		if event.Modifiers()&tcell.ModShift != 0 {
			if event.Key() == tcell.KeyRight {
				app.SetFocus(nextList)
				return tcell.NewEventKey(tcell.KeyNUL, 0, tcell.ModNone)
			}
			if event.Key() == tcell.KeyLeft {
				app.SetFocus(previousList)
				return tcell.NewEventKey(tcell.KeyNUL, 0, tcell.ModNone)
			}
			if event.Key() == tcell.KeyUp {
				list.SetCurrentItem(0)
				return tcell.NewEventKey(tcell.KeyNUL, 0, tcell.ModNone)
			}
			if event.Key() == tcell.KeyDown {
				list.SetCurrentItem(-1)
				return tcell.NewEventKey(tcell.KeyNUL, 0, tcell.ModNone)
			}

		} else if event.Modifiers()&tcell.ModAlt != 0 {
		} else {
			// Shift + ',' = '<'
			if event.Rune() == '<' && list.GetItemCount() > 0 {
				x := list.GetCurrentItem()
				pText, sText := list.GetItemText(x)
				list.RemoveItem(x)
				previousList.AddItem(pText, sText, '*', nil)
			}
			// Shift + '.' = '>'
			if event.Rune() == '>' && list.GetItemCount() > 0 {
				x := list.GetCurrentItem()
				pText, sText := list.GetItemText(x)
				list.RemoveItem(x)
				nextList.AddItem(pText, sText, '*', nil)
			}
		}

		return event
	})
}

func main() {

	app := tview.NewApplication()
	menu := tview.NewBox().SetBorder(true).SetTitle("Menu")

	backlog := l("Backlog", app)
	ready := l("Ready", app)
	doing := l("Doing", app)
	done := l("Done", app)

	// Handling Inputs per list so that they are navigatable ...
	configureList(app, done, backlog, ready)
	configureList(app, backlog, ready, doing)
	configureList(app, ready, doing, done)
	configureList(app, doing, done, backlog)

	details := tview.NewBox().SetBorder(true).SetTitle("Details")
	metrics := tview.NewBox().SetBorder(true).SetTitle("Metrics")

	// backlog.()

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().AddItem(menu, 0, 1, false),
			0, 1, false).
		AddItem(
			tview.NewFlex().
				AddItem(backlog, 0, 1, false).
				AddItem(ready, 0, 1, false).
				AddItem(doing, 0, 1, true).
				AddItem(done, 0, 1, false),
			0, 7, false).
		AddItem(
			tview.NewFlex().
				AddItem(details, 0, 1, false).
				AddItem(metrics, 0, 1, false),
			0, 4, false)

	if err := app.SetRoot(flex, true).SetFocus(doing).Run(); err != nil {
		panic(err)
	}
}
