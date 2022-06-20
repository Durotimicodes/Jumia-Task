package models

import "gorm.io/gorm"

type User struct {
	ID       int    `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
}

type Customer struct {
	ID       int    `json:"id,omitempty"`
	UserName string `json:"name,omitempty"`
	Mobile   string `json:"phone,omitempty"`
}

type ContactVerification struct {
	gorm.Model
	MobileNumber string `json:"mobile_number" binding:"required"`
	Country      string `json:"country" binding:"required"`
	CountryCode  string `json:"country_code" binding:"required"`
	IsValid      bool   `json:"valid" binding:"required"`
}

type ValidateCountryInfo struct {
	CountryName string
	CountryCode string
	RegularExp  string
}

var ValidateCountries = []ValidateCountryInfo{
	{CountryName: "Cameroon", CountryCode: "(237)", RegularExp: `\(237\)\ ?[2368]\d{7,8}$`},

	{CountryName: "Ethiopia", CountryCode: "(251)", RegularExp: `\(251\)\ ?[1-59]\d{8}$`},

	{CountryName: "Morocco", CountryCode: "(212)", RegularExp: `\(212\)\ ?[5-9]\d{8}$`},

	{CountryName: "Mozambique", CountryCode: "(258)", RegularExp: `\(258\)\ ?[28]\d{7,8}$`},

	{CountryName: "Uganda", CountryCode: "(256)", RegularExp: `\(256\)\ ?\d{9}$`},
}

// ConfigureMobileNumber are phone numbers configured into respective countries
type ConfigureMobileNumber struct {
	Country      string `json:"country" binding:"required"`
	State        string `json:"state" binding:"required"`
	CountryCode  string `json:"countryCode" binding:"required"`
	MobileNumber string `json:"mobileNumber" binding:"required"`
}
