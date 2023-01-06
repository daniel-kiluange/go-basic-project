package main

import (
	"awesomeProject1/src/internal/database"
	"fmt"
	"log"
)

func main() {
	p := database.PostgresInstance()
	rows, err := p.Db.Query("SELECT current_database();")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var (
		currentDatabase string
	)
	fmt.Println(rows.Columns())

	for rows.Next() {
		rows.Scan(&currentDatabase)
	}

	fmt.Println(currentDatabase)
}
