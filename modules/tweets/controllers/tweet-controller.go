package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"
	. "github.com/merq-rodriguez/twitter-go/modules/tweets/dto"
	. "github.com/merq-rodriguez/twitter-go/modules/tweets/models"
	tweetService "github.com/merq-rodriguez/twitter-go/modules/tweets/services"
)

type TweetController struct{}

/*
GetTweetsByUserID function controller
*/
func (t *TweetController) GetTweetsByUserID(c echo.Context) error {
	ID := c.QueryParam("id")
	page := c.QueryParam("page")

	if len(ID) < 1 {
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
	results, err := tweetService.GetTweetsByUserID(ID, pag)

	if err != nil {
		return BadRequestError(c, "Error read tweets", nil)
	}

	return c.JSON(http.StatusOK, results)
}

/*
CreateTweet function controller
*/
func (t TweetController) CreateTweet(c echo.Context) error {
	var tweet Tweet
	dto := CreateTweetDTO{}

	err := dto.Decoder(c.Request().Body)
	if err != nil {
		return BadRequestError(c, "Error fields data", err)
	}

	err = dto.Validate()
	if err != nil {
		return BadRequestError(c, "Fields required", err)
	}

	tweet = dto.ConvertToTweet()
	tweet.Timestamp = time.Now()

	result, err := tweetService.CreateTweet(tweet)
	if err != nil {
		return BadRequestError(c, "Tweet not created", err)
	}

	return c.JSON(http.StatusCreated, result)
}
