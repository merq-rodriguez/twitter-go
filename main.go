package main

import (
	"log"

	"github.com/merq-rodriguez/twitter-go/common/database"
	"github.com/merq-rodriguez/twitter-go/handlers"
)

func main() {
	if database.CheckConnection() == false {
		log.Fatal("Database not connected")
		return
	}

	handlers.RunHandlers()
}
