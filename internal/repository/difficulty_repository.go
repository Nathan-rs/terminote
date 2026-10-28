package repository

import (
	"database/sql"
	"terminote/internal/types"
)

func GetAllDifficulties(db *sql.DB) ([]types.Difficulty, error) {
	rows, err := db.Query(`
		SELECT id, name FROM difficulty ORDER BY id ASC	
	`)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var difficulties []types.Difficulty

	for rows.Next() {
		var difficulty types.Difficulty

		if err := rows.Scan(
			&difficulty.ID,
			&difficulty.Name,
		); err != nil {
			return nil, err
		}

		difficulties = append(difficulties, difficulty)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return difficulties, nil
}
