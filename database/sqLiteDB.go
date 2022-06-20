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

//NewSqlitedB Instanstiate and returns an Sqlist3db
func NewSqliteDb(DB *gorm.DB) *SqLite3Db {

	return &SqLite3Db{
		DB: DB,
	}
}

//GetAllMobileUser get all the values in the User table
func (Sq *SqLite3Db) GetAllMobileUsers(c context.Context) ([]*models.User, error) {
	var user []*models.User

	if err := Sq.DB.Raw("SELECT * FROM users").Scan(&user).Error; err != nil {

		log.Printf("Failed to find all users  %v", err)

		return nil, err
	}

	return user, nil
}

//GetMobileNumber get mobile number with the matching the code and state
func (Sq *SqLite3Db) GetMobileNumbers(c context.Context, countryCode, state string) ([]*models.ConfigureMobileNumber, error) {

	var mobileNum []*models.ConfigureMobileNumber

	users, err := Sq.GetAllMobileUsers(c)
	if err != nil {

		log.Printf("Failed getting all mobile numbers %v", err)

		return nil, err
	}

	if countryCode != "all" && countryCode != "" {

		users = logic.GetMatchingNumber(countryCode, users)
	}

	configNumb, er := logic.ConfigNumb(users)
	CheckError(er)

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

//CheckError is a Helper function
func CheckError(err error) {

	if err != nil {

		log.Printf("Failed to execute %v", err)

		return
	}
}
