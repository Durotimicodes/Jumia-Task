package database

import (
	"github.com/Durotimicodes/jumia-phone-number-task/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func SetUpDBConnection() error {

	log.Println("Connecting Sqlite3 DB")

	//Open connection
	db, err := gorm.Open(sqlite.Open("number-verification.db"), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect database %v", err)
		return err
	}

	//Migrate Schema
	er := db.AutoMigrate(&models.ContactVerification{})
	if er != nil {
		log.Printf("Failed to migrate schema %v", er)
		return er
	}

	return nil
}
