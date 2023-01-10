package main

import (
	"awesomeProject1/src/internal/database"
	"fmt"
	"log"
)

func execDatabase(p *database.Postgres, done chan string) {
	rows, err := p.Db.Query("SELECT current_database();")

	if err != nil {
		log.Fatal(err)
	}

	var databaseName string

	for rows.Next() {
		rows.Scan(&databaseName)
	}

	rows.Close()

	done <- databaseName
}

func main() {
	done := make(chan string)
	p := database.PostgresInstance()
	go execDatabase(p, done)
	fmt.Println(<-done)

}
