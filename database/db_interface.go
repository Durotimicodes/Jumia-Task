package database

import (
	"context"
	"github.com/Durotimicodes/jumia-phone-number-task/models"
)

type DBInterface interface {
	PrePolulateTable() error
	GetAllUsers() ([]models.ContactVerification, error)
	GetAllMobileNumb(c context.Context) ([]*models.User, error)
	GetMobileNumbers(c context.Context, countryCode, state string) ([]*models.ConfigureMobileNumber, error)
}
