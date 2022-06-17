package database

import (
	"github.com/Durotimicodes/jumia-phone-number-task/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func SetUpDBConnection() error {

	//Open connection
	db, err := gorm.Open(sqlite.Open("user-contact.db"), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database %v", err)
		return err
	}

	//Migrate Schema
	er := db.AutoMigrate(&models.ContactVerification{})
	if er != nil {
		log.Printf("failed to migrate schema %v", er)
		return er
	}

	return nil
}
