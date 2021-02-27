package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/common/constants"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"
	accountService "github.com/merq-rodriguez/twitter-go/modules/account/services"
	. "github.com/merq-rodriguez/twitter-go/modules/authorization/jwt/decorators"
	"github.com/merq-rodriguez/twitter-go/modules/upload"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"
	userService "github.com/merq-rodriguez/twitter-go/modules/users/services"
)

type AccountController struct{}

/*
GetProfile controller for get profile user
*/
func (h *AccountController) GetProfile(c echo.Context) error {
	user := GetCurrentUser(c)
	profile, err := accountService.GetProfile(user.Username)
	if err != nil {
		return NotFoundError(c, "Profile not exists", nil)
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *AccountController) UpdateAvatar(c echo.Context) error {
	file, err := c.FormFile("avatar")
	user := GetCurrentUser(c)

	if err != nil {
		return BadRequestError(c, "", err)
	}

	err = upload.AddFileToStorage(c, user.ID.Hex(), file)
	log.Println(err)

	if err != nil {
		return err
	}

	var u User
	u.Avatar = GetURLApi() + user.ID.Hex()
	userService.UpdateUser(user.ID.Hex(), u)

	return c.JSON(http.StatusOK, echo.Map{
		"statusCode": http.StatusOK,
		"message":    "Avatar updated",
	})
}
