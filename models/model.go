package models

type User struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Contact   ContactVerification
}

type ContactVerification struct {
	MobileNumber string `json:"mobile_number" binding:"required"`
	Country      string `json:"country" binding:"required"`
	CountryCode  uint   `json:"country_code" binding:"required"`
	Valid        bool   `json:"valid" binding:"required"`
	LocalFormat  string `json:"local_format" binding:"required"`
}
