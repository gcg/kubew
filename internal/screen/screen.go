package screen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	App     *tview.Application
	Main    *tview.Box
	Sidebar *tview.Box
	Layout  *tview.Flex
)

func Init() *tview.Application {
	App = tview.NewApplication()

	drawSideBar()
	drawMainBox()
	drawLayout()
	registerShortcuts()

	App.SetRoot(Layout, true).SetFocus(Layout)
	return App
}

func drawSideBar() {
	Sidebar = tview.NewBox().
		SetTitle("Side Bar").
		SetBorder(true).
		SetBackgroundColor(tcell.ColorFuchsia)
}

func drawMainBox() {
	Main = tview.NewBox().
		SetBorder(true).
		SetBackgroundColor(tcell.ColorBisque).
		SetTitle("Main Box")
}

func drawLayout() {
	Layout = tview.NewFlex().
		AddItem(Sidebar, 30, 1, true).
		AddItem(Main, 0, 2, false)
}

func registerShortcuts() {

	Layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		return event
	})
}
