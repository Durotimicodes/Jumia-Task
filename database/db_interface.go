package database

import "github.com/Durotimicodes/jumia-phone-number-task/models"

type DBInterface interface {
	PrePolulateTable() error
	GetAllUsers() ([]models.ContactVerification, error)
}
