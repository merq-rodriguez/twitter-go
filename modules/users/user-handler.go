package users

import (
	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/modules/users/controllers"
)

const (
	prefix = "/users"
)

func UserHandler(e *echo.Echo) {
	users := &UserController{}
	e.PUT(prefix, users.UpdateUser)
}
