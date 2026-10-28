package ui

import (
	"database/sql"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ListTasks(db *sql.DB) *tview.Table {
	table := tview.NewTable()

	table.SetBorders(false) // true deixa como linhas de tabela, false desativa
	table.SetSelectable(true, false)

	headers := []string{
		"ID",
		"Data inicio",
		"Dificuldade",
		"Descrição",
		"Municipio",
		"Estado",
	}

	for col, header := range headers {

		cell := tview.NewTableCell(header).
			SetTextColor(tcell.NewHexColor(0x00A693)).
			SetSelectable(false).
			SetExpansion(1)

		cell.SetAlign(tview.AlignLeft)

		table.SetCell(0, col, cell)

	}

	rows, err := db.Query(`
		SELECT
			tasks.id,
			DATE(tasks.date_start),
			difficulty.name,
			tasks.description,
			county.name,
			county.state
		FROM
			tasks
		INNER JOIN
			difficulty
		ON tasks.difficulty_id = difficulty.id
		INNER JOIN 
			county
		ON tasks.county_id = county.id
		ORDER BY
			tasks.id DESC
	`)
	if err != nil {
		table.SetCell(1, 0, tview.NewTableCell(fmt.Sprintf("Error: %v", err)))
		return table
	}
	defer rows.Close()

	row := 1

	for rows.Next() {
		var (
			id          int
			dateStart   string
			dif         string
			description string
			county      string
			state       string
		)

		err := rows.Scan(
			&id,
			&dateStart,
			&dif,
			&description,
			&county,
			&state,
		)

		if err != nil {
			continue
		}

		table.SetCell(row, 0, tview.NewTableCell(fmt.Sprintf("%d", id)))
		table.SetCell(row, 1, tview.NewTableCell(dateStart))
		table.SetCell(row, 2, tview.NewTableCell(dif))
		table.SetCell(row, 3, tview.NewTableCell(description))
		table.SetCell(row, 4, tview.NewTableCell(county))
		table.SetCell(row, 5, tview.NewTableCell(state))

		row++
	}

	if err := rows.Err(); err != nil {
		table.SetCell(
			row,
			0,
			tview.NewTableCell(fmt.Sprintf("Erro ao ler tarefas: %v", err)),
		)
	}

	return table
}
