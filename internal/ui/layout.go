package ui

import (
	"terminote/internal/ui/components"

	"github.com/rivo/tview"
)

func CreateUI(tableList tview.Primitive) tview.Primitive {

	flexLayout := tview.NewFlex().SetDirection(tview.FlexRow)

	header := components.Header()
	listTasks := tableList
	footer := components.Footer()

	flexLayout.AddItem(header, 3, 0, false)
	flexLayout.AddItem(listTasks, 0, 1, true)
	flexLayout.AddItem(footer, 3, 0, false)

	return flexLayout
}
