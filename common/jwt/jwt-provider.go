package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	constant "github.com/merq-rodriguez/twitter-clone-backend-go/common/jwt/constants"
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/models"
)

/*
CreateToken function for generate Json Web Token
*/
func CreateToken(user models.User) (string, error) {
	var expiresin = time.Duration(constant.EXPIRES_IN)
	secretKey := []byte(constant.SECRET_KEY)

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
