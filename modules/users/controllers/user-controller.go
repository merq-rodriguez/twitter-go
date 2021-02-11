package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"
	. "github.com/merq-rodriguez/twitter-go/modules/users/dto"

	userService "github.com/merq-rodriguez/twitter-go/modules/users/services"
)

type UserController struct{}

/*
UpdateUser controller for update user profile
*/
func (h *UserController) UpdateUser(c echo.Context) error {
	ID := c.QueryParam("id")
	var dto UpdateUserDTO

	err := dto.Validate()
	if err != nil {
		return BadRequestError(c, "Fields required", err)
	}

	u := dto.ConvertToUser()
	user, err := userService.UpdateUser(ID, u)
	if err != nil {
		return BadRequestError(c, "User not update", err)
	}

	return c.JSON(http.StatusCreated, user)
}
