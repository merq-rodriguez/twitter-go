package auth

import (
	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/modules/auth/controllers"
)

const (
	prefix = "/auth"
)

func AuthHandler(e *echo.Echo) {
	auth := &AuthController{}
	e.POST(prefix+"/signin", auth.SignIn)
	e.POST(prefix+"/signup", auth.SignUp)
}
