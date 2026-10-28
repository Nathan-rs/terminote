package repository

import (
	"database/sql"
	"fmt"
	"terminote/internal/types"
)

func SaveTask(task types.Task, db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO tasks (
			description,
			county_id,
			problem,
			solution,
			difficulty_id,
			date_start,
			date_end
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		task.Description,
		task.CountyID,
		task.Problem,
		task.Solution,
		task.DifficultyID,
		task.DateStart,
		task.DateEnd,
	)

	if err != nil {
		return fmt.Errorf("erro ao salvar task: %w", err)
	}

	return nil
}

func GetTaskById(db *sql.DB, id uint32) (types.Task, error) {
	var task types.Task

	// err := db.QueryRow(`
	// 	SELECT
	// 		id,
	// 		description,
	// 		county_id,
	// 		problem,
	// 		solution,
	// 		difficulty_id,
	// 		date_start,
	// 		date_end
	// 	FROM tasks
	// 	WHERE id = ?
	// `, id).Scan(
	// 	&task.ID,
	// 	&task.Description,
	// 	&task.CountyID,
	// 	&task.Problem,
	// 	&task.Solution,
	// 	&task.DifficultyID,
	// 	&task.DateStart,
	// 	&task.DateEnd,
	// )

	// if err != nil {
	// 	return task, fmt.Errorf("erro ao buscar task: %w", err)
	// }

	return task, nil
}
