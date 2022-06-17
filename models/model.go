package models

type ContactVerification struct {
	MobileNumber string `json:"mobile_number"`
	Country      string `json:"country"`
	CountryCode  uint   `json:"country_code"`
	Valid        bool   `json:"valid"`
	LocalFormat  string `json:"local_format"`
}
