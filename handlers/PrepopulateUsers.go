package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) PrepopulateUser(c *gin.Context) error {
	err := h.repository.PrePolulateTable()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad request, could populate number verification table with values",
			"value":   err,
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"Message": "Status Ok, Successfully populated number verification table",
	})
	return nil
}
