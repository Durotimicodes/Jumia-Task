package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) GetAllContacts(c *gin.Context) {

	err := h.repository.PrePolulateTable() ///
	if err != nil {
		log.Printf("Error prepopulating %v", err)
		return
	}

	contacts, err := h.repository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request, could not get all user contacts",
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"Message": "Successfully retrieved data",
		"data":    contacts,
	})

}

// Handler abstracts methods implemented for rest requests
//type Handler interface {
//	// GetCountries get a list of all countries
//	GetCountries() gin.HandlerFunc
//
//	// GetPhonenumbers gets all phone numbers
//	GetPhonenumbers() gin.HandlerFunc
//}

//type impl struct {
//	uc usecases.CustomersUsecase
//}
//
//// New creates a new instance of the Handler interface
//func New(repo database.Repository) Handler {
//	uc := usecases.New(repo)
//	return &impl{
//		uc: uc,
//	}
//}
//
//func (h *impl) GetCountries() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.JSON(http.StatusOK, models.AllCountries)
//	}
//}
//
