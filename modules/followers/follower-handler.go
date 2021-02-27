package followers

import (
	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/modules/followers/controllers"
)

const prefix = "followers"

func FollowerHandler(e *echo.Echo) {
	follower := &FollowerController{}
	e.POST(prefix, follower.AddFollower)
	e.GET(prefix, follower.GetFollowersByUser)
}
