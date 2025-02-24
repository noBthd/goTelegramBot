package tgsql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB 
)

// db connection
func DBInit() {
	var err error
	
	connStr := "user=postgres password=1234 dbname=tgBot sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	
	if err != nil {
		log.Fatal(err)
	}

	log.Print("\n")
}