package main

import (
	"github.com/Durotimicodes/jumia-phone-number-task/database"
	"github.com/Durotimicodes/jumia-phone-number-task/routers"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	err := database.SetUpDBConnection()
	if err != nil {
		log.Printf("Error setting up db connection %v", err)
		return
	}

	er := routers.SetUpRouters()
	if er != nil {
		log.Printf("Error setting up router %v", er)
		return
	}

}
