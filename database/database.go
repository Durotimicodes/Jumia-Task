package database

import (
	"database/sql"
	"fmt"
	"log"
)

func SetUpDBConnection() error {

	db, err := sql.Open("sqlite3", "phonenumbervalidation.db")
	if err != nil {
		log.Printf("Error in connecting to database %v", err)
		return err
	}

	if db == nil {
		return fmt.Errorf("database was not initialize")
	}
	return nil
}
