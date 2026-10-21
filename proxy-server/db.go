package main

import (
	"database/sql"
	"log"
)

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite", "api_debugger.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}

	return db
}


func CreateAppTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS application (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	return err
}



