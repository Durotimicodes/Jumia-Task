package database

import (
	"github.com/Durotimicodes/jumia-phone-number-task/models"
	"gorm.io/gorm"
	"log"
	"time"
)

// SqLite3Db implements the DB interface
type SqLite3Db struct {
	DB *gorm.DB
}

func NewSqliteDb(DB *gorm.DB) *SqLite3Db {
	return &SqLite3Db{
		DB: DB,
	}
}

func (Sq *SqLite3Db) PrePolulateTable() error {

	ContactInfo := []*models.ContactVerification{
		{Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, MobileNumber: "8181338386", Country: "Nigeria", CountryCode: "+234", IsValid: true},

		{Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, MobileNumber: "8181338386", Country: "Nigeria", CountryCode: "+234", IsValid: false},

		{Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, MobileNumber: "23682233", Country: "Cameroon", CountryCode: "+237", IsValid: true},

		{Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, MobileNumber: "23682633", Country: "Cameroon", CountryCode: "+237", IsValid: true},

		{Model: gorm.Model{
			ID:        5,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, MobileNumber: "2368233", Country: "Cameroon", CountryCode: "+237", IsValid: false},
		//{MobileNumber: "", Country: "Ethiopia", CountryCode: "+251", IsValid: true},
		//{MobileNumber: "08181338386", Country: "Ethiopia", CountryCode: "+251", IsValid: false},
		//{MobileNumber: "08181338386", Country: "Morocco", CountryCode: "+212", IsValid: true},
		//{MobileNumber: "08181338386", Country: "Mozambique", CountryCode: "+258", IsValid: false},
		//{MobileNumber: "08181338386", Country: "Mozambique", CountryCode: "+258", IsValid: true},
		//{MobileNumber: "08181338386", Country: "Mozambique", CountryCode: "+258", IsValid: false},
		//{MobileNumber: "08181338386", Country: "Uganda", CountryCode: "+256", IsValid: true},
		//{MobileNumber: "08181338386", Country: "Uganda", CountryCode: "+256", IsValid: true},
	}

	log.Println("Before result")
	err := Sq.DB.Create(&ContactInfo)
	if err != nil {
		log.Printf("Error PREPOPULATING %v", err)
	}
	//check := result.RowsAffected

	//fmt.Println(check)

	return nil

}

func (Sq *SqLite3Db) GetAllUsers() ([]models.ContactVerification, error) {
	Contacts := []models.ContactVerification{}

	if er := Sq.DB.Find(&Contacts).Error; er != nil {
		log.Printf("Failed to find all contacts %v", er)
		return nil, er
	}
	return Contacts, nil
}

/*
	Cameroon| Country code = +237 | regex = \(237\)\ ?[2368]\d{7,8}$
		Ethiopia | Country code = +251 | regex = \(251\)\ ?[1-59]\d{8}$ |
		Morocco|Country code = +212 | regex = \(212\)\ ?[5-9]\d{8}$
		 Mozambique| Country code = +258 | regex = \(258\)\ ?[28]\d{7,8}$
	Uganda| Country code = +256 | regex = \(256\)\ ?\d{9}$
	Nigeria| Country code = +234 | regex = \(234\)\ ?\d{8}$
*/
