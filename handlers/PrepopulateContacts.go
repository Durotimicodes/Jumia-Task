package handlers

import (
	"database/sql"
	"log"
)

var db *sql.DB

func PrePopulateDB() error {

	statement, err := db.Prepare("INSERT INTO contact_verification_info " +
		"(mobile_number, country, country_code, is_valid)" + " VALUES(?,?,?)")
	if err != nil {
		log.Printf("Failed to insert values to db table %v", err)
		return err
	}

	statement.Exec("8081338386", "Nigeria", "+234", "false")

	return nil

}
