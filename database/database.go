package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//var db *sql.DB

func SetUpDBConnection() error {

	//Open connection, sql is hard to be bug so golang uses err handle to manage this
	db, err := sql.Open("sqlite3", "user-contact.db")
	if err != nil {
		panic(err.Error())
	}

	query := "CREATE TABLE IF NOT EXISTS contact_verification_info (" +
		"id INTEGER NOT NULL PRIMARY KEY," +
		"time DATETIME NOT NULL," +
		"mobile_number INTEGER NOT NULL," +
		"country TEXT NOT NULL," +
		"country_code TEXT NOT NULL," +
		"is_valid INTEGER NOT NULL)"

	//execute query
	create, err := db.Exec(query)
	if err != nil {
		log.Printf("failed executing query %v", err)
		panic(err.Error())
	}

	log.Println(create)

	defer db.Close()

	return nil
}
