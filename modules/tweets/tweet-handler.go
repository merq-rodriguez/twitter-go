package tweets

import (
	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/modules/tweets/controllers"
)

const (
	prefix = "/tweets"
)

func TweetHandler(e *echo.Echo) {
	tweets := &TweetController{}
	e.GET(prefix, tweets.GetTweetsByUserID)
	e.POST(prefix, tweets.CreateTweet)
}
