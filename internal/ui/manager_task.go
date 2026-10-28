package ui

import (
	"database/sql"
	"terminote/internal/repository"
	"terminote/internal/types"
	"terminote/internal/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func newModal(form *tview.Form) *tview.Flex {
	modal := tview.NewFlex().
		SetDirection(tview.FlexRow)

	modal.AddItem(form, 0, 1, true)

	// Centralizar horizontalmente
	horizontal := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(modal, 0, 3, false).
		AddItem(nil, 0, 1, false)

	// Centralizar verticalmente
	vertical := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(horizontal, 0, 4, true).
		AddItem(nil, 0, 1, false)

	return vertical
}

func CreateTask(app *tview.Application, pages *tview.Pages, db *sql.DB) {
	var task types.Task

	form := tview.NewForm()

	// counties, err := repository.GetAllCounties(db)
	// if err != nil {
	// 	return
	// }

	// difficulties, err := repository.GetAllDifficulties(db)
	// if err != nil {
	// 	return
	// }

	// countyOptions := make([]string, len(counties))
	// for i, county := range counties {
	// 	countyOptions[i] = county.Name
	// }

	// difficultyOptions := make([]string, len(difficulties))
	// for i, difficulty := range difficulties {
	// 	difficultyOptions[i] = difficulty.Name
	// }

	form.AddTextArea(
		"Descrição",
		"",
		0,
		2,
		0,
		func(value string) {
			task.Description = value
		},
	)

	form.AddTextArea(
		"Problema",
		"",
		0,
		2,
		0,
		func(value string) {
			task.Problem = value
		},
	)

	form.AddTextArea(
		"Solução",
		"",
		0,
		2,
		0,
		func(value string) {
			task.Solution = value
		},
	)

	// form.AddDropDown(
	// 	"Municipio",
	// 	countyOptions,
	// 	0,
	// 	func(option string, optionIndex int) {
	// 		task.CountyID = counties[optionIndex].ID
	// 	},
	// )

	// form.AddDropDown(
	// 	"Dificuldate",
	// 	difficultyOptions,
	// 	0,
	// 	func(option string, optionIndex int) {
	// 		task.DifficultyID = difficulties[optionIndex].ID
	// 	},
	// )

	form.AddInputField(
		"Data de inicio",
		"",
		0,
		nil,
		nil,
	)

	form.AddInputField(
		"Data de termino",
		"",
		0,
		nil,
		nil,
	)

	form.AddButton("Salvar", func() {
		// salvar no banco

		// repository.SaveTask(, db)

		pages.HidePage("create-task")
	})

	form.AddButton("Cancelar", func() {
		// fechar modal
		pages.HidePage(utils.PageCreateTask)
	})

	form.SetBorder(true)
	form.SetBorderColor(tcell.NewHexColor(0xD4D4D4))
	form.SetTitle(" Nova Task ")
	form.SetTitleColor(tcell.NewHexColor(0x00A693))

	modal := newModal(form)

	pages.AddPage(utils.PageCreateTask, modal, true, true)

	app.SetFocus(form)
}

func EditTask(app *tview.Application, pages *tview.Pages, db *sql.DB, id uint32) {
	task, err := repository.GetTaskById(db, id)
	if err != nil {
		return
	}

	form := tview.NewForm()

	form.AddTextArea(
		"Descrição",
		task.Description,
		0,
		2,
		0,
		func(value string) {
			task.Description = value
		},
	)

	form.AddTextArea(
		"Problema",
		task.Problem,
		0,
		2,
		0,
		func(value string) {
			task.Problem = value
		},
	)

	form.AddTextArea(
		"Solução",
		task.Solution,
		0,
		2,
		0,
		func(value string) {
			task.Solution = value
		},
	)

	form.AddInputField(
		"Data de inicio",
		task.DateStart.Format("02/01/2006"),
		10,
		nil,
		nil,
	)

	dateEnd := ""
	if task.DateEnd != nil {
		dateEnd = task.DateEnd.Format("02/01/2006")
	}

	form.AddInputField(
		"Data de termino",
		dateEnd,
		0,
		nil,
		nil,
	)

	form.AddButton("Salvar", func() {
		// salvar no banco

		// repository.SaveTask(, db)

		pages.HidePage(utils.PageEditTask)

	})

	form.AddButton("Cancelar", func() {
		// fechar modal
		pages.HidePage(utils.PageEditTask)
	})

	form.SetBorder(true)
	form.SetTitle(" Editar Task ")
	form.SetTitleColor(tcell.NewHexColor(0x00A693))

	modal := newModal(form)

	pages.AddPage(utils.PageEditTask, modal, true, true)

	app.SetFocus(form)

}
