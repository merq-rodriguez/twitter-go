package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"
	. "github.com/merq-rodriguez/twitter-go/modules/users/dto"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"

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

func (u *UserController) SearchUsers(c echo.Context) error {
	_p, _ := strconv.Atoi(c.QueryParam("page"))
	page := int64(_p)

	textSearch := c.QueryParam("textSearch")
	dto := SearchUserDTO{}

	param := SearchUser{
		TextSearch: textSearch,
		Page:       page,
	}

	dto.Bind(param)
	err := dto.Validate()

	if err != nil {
		return BadRequestError(c, "Fields required", err)
	}

	results, err := userService.SearchUsers(param)

	if err != nil {
		return InternalServerError(c, "Not load users", err)
	}

	return c.JSON(http.StatusOK, results)
}
