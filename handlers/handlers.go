package handlers

import "github.com/Durotimicodes/jumia-phone-number-task/database"

type Handler struct {
	repository database.DBInterface
}

//factory method
func NewHandler(repository database.DBInterface) *Handler {
	return &Handler{
		repository: repository,
	}
}
