package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetMobileNumbers(c *gin.Context) {

	countryCode := c.Param("code")
	state := c.Param("state")

	mobileNumber, err := h.repository.GetMobileNumbers(c, countryCode, state)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad request, unable to get mobile number",
			"data":    err,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"Message": "Successfully got mobile number",
		"data":    mobileNumber,
	})

}
