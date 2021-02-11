package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"
	accountService "github.com/merq-rodriguez/twitter-go/modules/account/services"
)

type AccountController struct{}

/*
GetProfile controller for get profile user
*/
func (h *AccountController) GetProfile(c echo.Context) error {
	username := c.QueryParam("username")

	if len(username) < 1 {
		return BadRequestError(c, "Username paremeter not provided", nil)
	}

	profile, err := accountService.GetProfile(username)
	if err != nil {
		return NotFoundError(c, "Profile not exists", nil)
	}

	return c.JSON(http.StatusOK, profile)
}
