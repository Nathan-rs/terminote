package main

import (
	"log"
	"terminote/internal/connection"
	"terminote/internal/ui"
)

func main() {
	db, err := connection.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ui.Show(db)
}
