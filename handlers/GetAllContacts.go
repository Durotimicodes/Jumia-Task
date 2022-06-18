package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) GetAllContacts(c *gin.Context) {

	err := h.repository.PrePolulateTable()
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
