package main

import (
	"github.com/Durotimicodes/jumia-phone-number-task/database"
	"github.com/Durotimicodes/jumia-phone-number-task/handlers"
	"github.com/Durotimicodes/jumia-phone-number-task/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	db, err := database.SetUpDBConnection()
	if err != nil {
		log.Printf("Error setting up db connection %v", err)
		return
	}
	router := gin.Default()

	handler := handlers.NewHandler(database.NewSqliteDb(db))

	er := routers.SetUpRouters(router, handler)
	if er != nil {
		log.Printf("Error setting up router %v", er)
		return
	}

	router.Run(":8081")
}
