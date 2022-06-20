package routers

import (
	"github.com/Durotimicodes/jumia-phone-number-task/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetUpRouters(router *gin.Engine, h *handlers.Handler) error {

	//CORs setup
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	//APi routes
	apirouter := router.Group("api/v1")

	apirouter.GET("/ping", handlers.HealthCheck)
	apirouter.GET("/user/countries", h.GetAllCountryInfo)
	apirouter.GET("/user/mobile/:code/:state", h.GetMobileNumbers)

	return nil

}
