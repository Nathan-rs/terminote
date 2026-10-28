package repository

import (
	"database/sql"
	"terminote/internal/types"
)

func GetAllCounties(db *sql.DB) ([]types.County, error) {
	rows, err := db.Query(`
		SELECT id, name FROM county ORDER BY id ASC	
	`)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var counties []types.County

	for rows.Next() {
		var county types.County

		if err := rows.Scan(
			&county.ID,
			&county.Name,
		); err != nil {
			return nil, err
		}

		counties = append(counties, county)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return counties, nil
}
