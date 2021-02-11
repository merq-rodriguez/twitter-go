package account

import (
	"github.com/labstack/echo"
	. "github.com/merq-rodriguez/twitter-go/modules/account/controllers"
)

const (
	prefix = "/accounts"
)

func AccountHandler(e *echo.Echo) {
	account := &AccountController{}
	e.GET(prefix+"/profile", account.GetProfile)
}
