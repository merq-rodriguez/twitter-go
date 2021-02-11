package jwt

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	constant "github.com/merq-rodriguez/twitter-go/common/jwt/constants"
	"github.com/merq-rodriguez/twitter-go/modules/users/models"
)

/*
CreateToken function for generate Json Web Token
*/
func CreateToken(user models.User) (string, error) {
	value, err := strconv.ParseInt(constant.ExpiresIn, 10, 64)
	var expiresin = time.Duration(value)
	secretKey := []byte(constant.SecretKey)

	payload := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"_id":   user.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * expiresin).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
