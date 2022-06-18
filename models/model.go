package models

import "gorm.io/gorm"'



type ContactVerification struct {
	gorm.Model
	MobileNumber string `json:"mobile_number" binding:"required"`
	Country      string `json:"country" binding:"required"`
	CountryCode  string `json:"country_code" binding:"required"`
	IsValid      bool   `json:"valid" binding:"required"`
}

type IndividualCountry struct {
	Country     string
	CountryCode string
	RegularExp  string
}

var Countries = []IndividualCountry{
	{Country: "Cameroon", CountryCode: "237", RegularExp: `\(237\)\ ?[2368]\d{7,8}$`},

	{Country: "Ethiopia", CountryCode: "251", RegularExp: `\(251\)\ ?[1-59]\d{8}$`},

	{Country: "Morocco", CountryCode: "212", RegularExp: `\(212\)\ ?[5-9]\d{8}$`},

	{Country: "Mozambique", CountryCode: "258", RegularExp: `\(258\)\ ?[28]\d{7,8}$`},

	{Country: "Uganda", CountryCode: "256", RegularExp: `\(256\)\ ?\d{9}$`},
}
