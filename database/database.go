package database

import (
	"github.com/Durotimicodes/jumia-phone-number-task/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func SetUpDBConnection() (*gorm.DB, error) {

	var DB *gorm.DB

	log.Println("Connecting Sqlite3 DB")

	//Open connection
	db, err := gorm.Open(sqlite.Open("Num_verify.db"), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect database %v", err)
		return nil, err
	}

	//Migrate Schema
	er := db.AutoMigrate(&models.ContactVerification2{})
	if er != nil {
		log.Printf("Failed to migrate schema %v", er)
		return nil, er
	}

	return DB, nil
}
