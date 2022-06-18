package database

import (
	"context"
	"github.com/Durotimicodes/jumia-phone-number-task/logic"
	"github.com/Durotimicodes/jumia-phone-number-task/models"
	"gorm.io/gorm"
	"log"
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

	ContactInfo := &[]models.ContactVerification{
		{MobileNumber: "8181338386", Country: "Nigeria", CountryCode: "+234", IsValid: true},

		{MobileNumber: "8181338386", Country: "Nigeria", CountryCode: "+234", IsValid: false},

		{MobileNumber: "23682233", Country: "Cameroon", CountryCode: "+237", IsValid: true},

		{MobileNumber: "23682633", Country: "Cameroon", CountryCode: "+237", IsValid: true},

		{MobileNumber: "2368233", Country: "Cameroon", CountryCode: "+237", IsValid: false},

		{MobileNumber: "23451334", Country: "Ethiopia", CountryCode: "+251", IsValid: true},

		{MobileNumber: "0234566", Country: "Ethiopia", CountryCode: "+251", IsValid: false},

		{MobileNumber: "55667788", Country: "Morocco", CountryCode: "+212", IsValid: true},

		{MobileNumber: "55679758", Country: "Morocco", CountryCode: "+212", IsValid: true},

		{MobileNumber: "56967768", Country: "Morocco", CountryCode: "+212", IsValid: true},

		{MobileNumber: "08181338386", Country: "Mozambique", CountryCode: "+258", IsValid: false},

		{MobileNumber: "08181338386", Country: "Mozambique", CountryCode: "+258", IsValid: true},

		{MobileNumber: "08181338386", Country: "Mozambique", CountryCode: "+258", IsValid: false},

		{MobileNumber: "08181338386", Country: "Uganda", CountryCode: "+256", IsValid: true},

		{MobileNumber: "08181338386", Country: "Uganda", CountryCode: "+256", IsValid: true},
	}

	log.Println("Before result")
	err := Sq.DB.Create(ContactInfo)
	if err != nil {
		log.Printf("Error PREPOPULATING %v", err)
	}
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

func (Sq *SqLite3Db) GetAllMobileNumb(c context.Context) ([]*models.User, error) {
	var user []*models.User

	if err := Sq.DB.Find(&user).Error; err != nil {
		log.Printf("Failed to find all users  %v", err)
		return nil, err
	}

	return user, nil
}

func (Sq *SqLite3Db) GetMobileNumbers(c context.Context, countryCode, state string) ([]*models.ConfigureMobileNumber, error) {

	var mobileNum []*models.ConfigureMobileNumber
	users, err := Sq.GetAllMobileNumb(c)
	if err != nil {
		log.Printf("Failed getting all mobile numbers %numb", err)
		return nil, err
	}

	// if country countryCode not all and not empty
	if countryCode != "all" && countryCode != "" {
		users, err = logic.GetMatchingNumber(countryCode, users)
		CheckError(err)
	}

	configNumb, er := logic.ConfigNumb(users)
	CheckError(er)

	// get numbers based on state
	if state == "all" || state == "" {
		mobileNum = configNumb
		return nil, nil
	}

	for _, numb := range configNumb {
		number := *numb
		if number.State == state {
			mobileNum = append(mobileNum, &number)
		}
	}
	return mobileNum, nil
}

func CheckError(err error) {
	if err != nil {
		log.Printf("Failed to execute %v", err)
		return
	}
}
