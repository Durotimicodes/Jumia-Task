package main

import (
	"github.com/Durotimicodes/jumia-phone-number-task/database"
	"github.com/Durotimicodes/jumia-phone-number-task/handlers"
	"github.com/Durotimicodes/jumia-phone-number-task/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//func cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, HEAD, OPTIONS, GET")
//
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//
//		c.Next()
//	}
//}

func main() {

	db, err := database.SetUpDBConnection()
	if err != nil {
		log.Printf("Error setting up db connection %v", err)
		return
	}
	router := gin.Default()
	//router.Use(cors())

	handler := handlers.NewHandler(database.NewSqliteDb(db))

	er := routers.SetUpRouters(router, handler)
	if er != nil {
		log.Printf("Error setting up router %v", er)
		return
	}

	router.Run(":8081")
}
