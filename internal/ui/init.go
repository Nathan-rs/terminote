package ui

import (
	"database/sql"
	"terminote/internal/ui/components"

	"github.com/rivo/tview"
)

func Show(db *sql.DB) {
	// currentPage := utils.PageMain

	// app := tview.NewApplication()

	// flex := tview.NewFlex().SetDirection(tview.FlexRow)

	// header := components.Header()
	// listTasks := ListTasks(db)
	// footer := components.Footer()

	// flex.AddItem(header, 3, 0, false)
	// flex.AddItem(listTasks, 0, 1, true)
	// flex.AddItem(footer, 3, 0, false)

	// pages := tview.NewPages()

	// pages.AddPage(utils.PageMain, flex, true, true)

	// app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	return utils.HandleKeyboard(
	// 		currentPage,
	// 		utils.KeyActions{
	// 			New: func() {
	// 				currentPage = utils.PageCreateTask
	// 				CreateTask(app, pages, db)
	// 			},

	// 			Edit: func() {
	// 				row, _ := listTasks.GetSelection()

	// 				// header da tabela
	// 				if row == 0 {
	// 					return
	// 				}

	// 				cell := listTasks.GetCell(row, 0)

	// 				taskID, err := strconv.ParseUint(cell.Text, 10, 32)
	// 				if err != nil {
	// 					return
	// 				}

	// 				currentPage = utils.PageEditTask
	// 				EditTask(app, pages, db, uint32(taskID))
	// 			},

	// 			Escape: func() {
	// 				pages.HidePage(currentPage)
	// 				currentPage = utils.PageMain
	// 			},

	// 			Quit: func() {
	// 				app.Stop()
	// 			},
	// 		},
	// 	)(event)
	// })

	// if err := app.SetRoot(pages, true).Run(); err != nil {
	// 	panic(err)
	// }

	app := tview.NewApplication()

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Teste"), 0, 1, false)

	modal := components.ModalOverlay(tview.NewBox().SetBorder(true).SetTitle("Teste Modal"), 55, 11)

	flex.AddItem(modal, 0, 3, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}

}
