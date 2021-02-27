package decorators

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	. "github.com/merq-rodriguez/twitter-go/modules/authorization/jwt/models"
)

func GetCurrentUser(c echo.Context) *JWTCustomClaim {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaim)
	return claims
}
