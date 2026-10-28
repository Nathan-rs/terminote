package connection

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "internal/database/tasks.db")
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir banco: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("Erro ao conectar ao banco: %w", err)
	}

	return db, nil

}
