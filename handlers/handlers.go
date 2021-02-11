package handlers

import (
	"log"
	"strconv"

	"github.com/labstack/echo"
	"github.com/merq-rodriguez/twitter-go/common/config"

	. "github.com/merq-rodriguez/twitter-go/modules/account"
	. "github.com/merq-rodriguez/twitter-go/modules/auth"
	. "github.com/merq-rodriguez/twitter-go/modules/tweets"
	. "github.com/merq-rodriguez/twitter-go/modules/users"
)

/*
RunHandlers function: run handdlers with controllers enpoints
*/
func RunHandlers() {
	e := echo.New()
	viper, err := config.Settings()
	port := viper.GetInt("port")

	if port == 0 {
		port = 8080
	}

	if err != nil {
		log.Fatal("Error loading configuration")
	}

	UserHandler(e)
	AuthHandler(e)
	TweetHandler(e)
	AccountHandler(e)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}
