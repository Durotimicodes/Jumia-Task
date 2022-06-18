package handlers

import (
	"github.com/Durotimicodes/jumia-phone-number-task/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllCountries(c *gin.Context) {
	allCountries := models.ValidateCountries
	c.IndentedJSON(http.StatusOK, allCountries)
}
