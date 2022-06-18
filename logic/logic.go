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
func GetMatchingNumber(countryCode string, user []*models.User) ([]*models.User, error) {

	var person []*models.User
	for _, individualUser := range user {

		newCountryCode := fmt.Sprintf(`^\(%s\)`, countryCode)
		checkMatched, err := regexp.Match(newCountryCode, []byte(individualUser.Contact.MobileNumber))

		if err == nil && checkMatched {
			person = append(person, individualUser)
		}
	}
	return person, nil
}

//Config Numb returns numbers that are configured according to country code and an error if ir exist
func ConfigNumb(users []*models.User) ([]*models.ConfigureMobileNumber, error) {

	var configuredNumber []*models.ConfigureMobileNumber

	for _, value := range users {

		country, err := FindCountry(value.Contact.MobileNumber)
		if err != nil {
			log.Println("Failed to find country with the phone number")
			return nil, nil
		}

		regularExpr := regexp.MustCompile(`\S+$`)

		mobileNumber := regularExpr.FindStringSubmatch(value.Contact.MobileNumber)[0]

		countryCode := fmt.Sprintf("+%s", country.CountryCode)

		result := &models.ConfigureMobileNumber{
			Country:      country.CountryName,
			State:        countryCode,
			MobileNumber: mobileNumber,
		}

		checkMatched, err := regexp.Match(country.RegularExp, []byte(value.Contact.MobileNumber))

		if err != nil || !checkMatched {
			result.State = "NOK"
		} else {
			result.State = "OK"
		}
		configuredNumber = append(configuredNumber, result)
	}

	return configuredNumber, nil
}
