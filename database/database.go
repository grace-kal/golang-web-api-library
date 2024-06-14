package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func GetDb() *sql.DB {
	return db
}

func InitDb() {

	var err error
	db, err = sql.Open("sqlite", "local.db")
	if err != nil {

		fmt.Println(err)
		panic("Database connection lost")
	}
	db.SetMaxOpenConns(10)

	createTableQuery := ` 
	 CREATE TABLE IF NOT EXISTS books(
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 title TEXT,
	 author TEXT,
	 isbn TEXT,
	 release INTEGER
	)`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
