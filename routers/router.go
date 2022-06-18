package routers

import (
	"github.com/Durotimicodes/jumia-phone-number-task/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouters(router *gin.Engine, h *handlers.Handler) error {

	apirouter := router.Group("api/v1")

	apirouter.GET("/ping", handlers.HealthCheck)
	apirouter.GET("/all/contacts", h.GetAllContacts)

	return nil

}
