package handlers

import "github.com/Durotimicodes/jumia-phone-number-task/database"

type Handler struct {
	Repository database.DBInterface
}

//using factory method
func NewHandler(repository database.DBInterface) *Handler {
	return &Handler{
		Repository: repository,
	}
}
