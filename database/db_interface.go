package database

import (
	"context"
	"github.com/Durotimicodes/jumia-phone-number-task/models"
)

//DBInterface housing DB methods
type DBInterface interface {
	GetAllMobileUsers(c context.Context) ([]*models.User, error)

	GetMobileNumbers(c context.Context, countryCode, state string) ([]*models.ConfigureMobileNumber, error)
}
