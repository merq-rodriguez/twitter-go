package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"
	accountService "github.com/merq-rodriguez/twitter-go/modules/account/services"
	. "github.com/merq-rodriguez/twitter-go/modules/authorization/jwt/models"
)

type AccountController struct{}

/*
GetProfile controller for get profile user
*/
func (h *AccountController) GetProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaim)
	username := claims.Username

	profile, err := accountService.GetProfile(username)
	if err != nil {
		return NotFoundError(c, "Profile not exists", nil)
	}

	return c.JSON(http.StatusOK, profile)
}
