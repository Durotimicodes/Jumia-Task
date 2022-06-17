package models

import "time"

type User struct {
	Id        uint   `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Contact   ContactVerification
}

type ContactVerification struct {
	ID           uint      `json:"id" binding:"required"`
	Time         time.Time `json:"time" binding:"required"`
	MobileNumber uint      `json:"mobile_number" binding:"required"`
	Country      string    `json:"country" binding:"required"`
	CountryCode  uint      `json:"country_code" binding:"required"`
	IsValid      bool      `json:"valid" binding:"required"`
}
