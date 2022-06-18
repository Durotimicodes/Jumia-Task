package models

import "gorm.io/gorm"

type User struct {
	Id        uint   `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Contact   ContactVerification
}

type ContactVerification struct {
	gorm.Model
	MobileNumber string `json:"mobile_number" binding:"required"`
	Country      string `json:"country" binding:"required"`
	CountryCode  string `json:"country_code" binding:"required"`
	IsValid      bool   `json:"valid" binding:"required"`
}
