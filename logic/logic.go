package logic

import (
	"fmt"
	"github.com/Durotimicodes/jumia-phone-number-task/models"
	"log"
	"regexp"
)

//Find the country that matches the regular expression
func FindCountry(mobileNumber string) (*models.ValidateCountryInfo, error) {

	for _, ValCountry := range models.ValidateCountries {

		newCountryCode := fmt.Sprintf(`^\(%s\)`, ValCountry.CountryCode)

		checkMatched, err := regexp.Match(newCountryCode, []byte(mobileNumber))

		if err != nil {
			log.Println("Country code regexp did not match mobile number")
			return nil, err
		}

		if checkMatched {
			return &ValCountry, nil
		}

	}
	return nil, nil
}

//Get numbers that match country code
func GetMatchingNumber(countryCode string, users []*models.User) (response []*models.User) {

	for _, individualUser := range users {

		user := *individualUser

		countryCodeRegex := fmt.Sprintf(`^\(%s\)`, countryCode)

		match, err := regexp.Match(countryCodeRegex, []byte(user.Mobile))

		if err == nil && match {
			response = append(response, &user)
		}

	}
	return
}

//Config Numb returns numbers that are configured according to country code and an error if ir exist
func ConfigNumb(users []*models.User) ([]*models.ConfigureMobileNumber, error) {

	var configuredNumber []*models.ConfigureMobileNumber

	for _, value := range users {

		country, err := FindCountry(value.Mobile)
		if err != nil {
			log.Println("Failed to find country with the phone number")
			return nil, nil
		}

		regularExpr := regexp.MustCompile(`\S+$`)

		mobileNumber := regularExpr.FindStringSubmatch(value.Mobile)[0]

		countryCode := fmt.Sprintf("+%s", country.CountryCode)

		result := &models.ConfigureMobileNumber{
			Country:      country.CountryName,
			CountryCode:  countryCode,
			MobileNumber: mobileNumber,
		}

		checkMatched, err := regexp.Match(country.RegularExp, []byte(value.Mobile))

		if err != nil || !checkMatched {
			result.State = "NOK"
		} else {
			result.State = "OK"
		}
		configuredNumber = append(configuredNumber, result)
	}

	return configuredNumber, nil
}
