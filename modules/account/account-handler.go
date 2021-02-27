package account

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	. "github.com/merq-rodriguez/twitter-go/modules/account/controllers"
	"github.com/merq-rodriguez/twitter-go/modules/authorization/jwt/constants"
	. "github.com/merq-rodriguez/twitter-go/modules/authorization/jwt/models"
)

const (
	prefix = "/accounts"
)

func AccountHandler(e *echo.Echo) {
	account := &AccountController{}

	r := e.Group(prefix)
	secretKey := []byte(constants.SecretKey)

	config := middleware.JWTConfig{
		Claims:     &JWTCustomClaim{},
		SigningKey: []byte(secretKey),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("/profile", account.GetProfile)
	r.PUT("/profile", account.UpdateAvatar)

}
