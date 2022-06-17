package handlers

import (
	"github.com/Durotimicodes/jumia-phone-number-task/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

// SqLite3Db implements the DB interface
type SqLite3Db struct {
	DB *gorm.DB
}

func (Sq *SqLite3Db) PrePolulateTable(c *gin.Context) {

	err := Sq.DB.AutoMigrate(&models.ContactVerification{})

	if err != nil {
		log.Printf("Migration Error %v", err)
		return
	}

	allContactInfo := []models.ContactVerification{}

}
