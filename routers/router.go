package routers

import (
	"github.com/Durotimicodes/jumia-phone-number-task/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouters() error {

	router := gin.Default()

	apirouter := router.Group("api/v1")

	apirouter.GET("/ping", handlers.HealthCheck)

	router.Run() //listen and serve on localhost port 8080
	return nil

}
