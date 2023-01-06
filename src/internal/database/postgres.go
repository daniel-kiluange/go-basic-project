package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Postgres struct {
	Db *sql.DB
}

var postgresInstance *Postgres

var lock = &sync.Mutex{}

func PostgresInstance() *Postgres {
	lock.Lock()
	lock.Unlock()
	if postgresInstance == nil {
		connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		postgresInstance = &Postgres{db}
	}
	return postgresInstance
}
