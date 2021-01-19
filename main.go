package main

import (
	"log"

	"github.com/merq-rodriguez/twitter-clone-backend-go/common/database"
	"github.com/merq-rodriguez/twitter-clone-backend-go/handlers"
)

func main() {
	if database.CheckConnection() == false {
		log.Fatal("Database not connected")
		return
	}
	handlers.RunHandlers()
}
