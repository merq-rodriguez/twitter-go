package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"
	. "github.com/merq-rodriguez/twitter-go/modules/followers/dto"
	. "github.com/merq-rodriguez/twitter-go/modules/followers/models"
	followerService "github.com/merq-rodriguez/twitter-go/modules/followers/services"
)

type FollowerController struct{}

func (f *FollowerController) GetFollowersByUser(c echo.Context) error {
	userID := c.QueryParam("id")
	page := c.QueryParam("page")

	if len(userID) < 1 {
		return BadRequestError(c, "Id parameter is required", nil)
	}
	if len(page) < 1 {
		return BadRequestError(c, "Page parameter is required", nil)
	}

	_page, err := strconv.Atoi(page)
	if err != nil {
		return NotAcceptableError(c, "The page parameter must be greater than 0", nil)
	}

	pag := int64(_page)
	results, err := followerService.GetFollowersByUser(userID, pag)

	if err != nil {
		return BadRequestError(c, "Error read followers", nil)
	}

	return c.JSON(http.StatusOK, results)
}

func (f *FollowerController) AddFollower(c echo.Context) error {
	var follower UserFollower
	dto := AddFollowerDTO{}

	err := dto.Bind(c.Request().Body)
	if err != nil {
		return NotAcceptableError(c, "", err)
	}

	err = dto.Validate()
	if err != nil {
		return NotAcceptableError(c, "", err)
	}

	follower = dto.ConvertToFollower()
	result, err := followerService.AddFollower(follower)

	if err != nil {
		return BadRequestError(c, "Follower not created", err)
	}

	return c.JSON(http.StatusCreated, result)
}
