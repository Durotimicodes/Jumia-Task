package database

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func SetUpDBConnection() (*gorm.DB, error) {

	//instantiating a db of type pointer to gorm.db
	var db *gorm.DB

	log.Println("Connecting Sqlite3 DB")

	//Open database connection
	db, err := gorm.Open(sqlite.Open("mobileNumber.db"), &gorm.Config{})

	if err != nil {

		log.Printf("Failed to connect database %v", err)

		return nil, err
	}

	/*Migrate Schema
	TODO : Only uncomment to migrate when starting with an empty db else an error will occur
	*/
	/*er := db.AutoMigrate(&models.User{})
	if er != nil {
		log.Printf("Failed to migrate schema %v", er)
		return nil, er
	}*/

	return db, nil

}
