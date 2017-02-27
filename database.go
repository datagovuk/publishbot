package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type AdapterRecord struct {
	Label string
}

var database sql.DB

func open_db() {
	database, err := sql.Open("sqlite3",
		"./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal("Unable to open the database")
	}
}

func close_db() {
	database.Close()
}
