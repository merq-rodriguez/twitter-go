package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/models"
)

/*
CreateToken function for generate Json Web Token
*/
func CreateToken(user models.User) (string, error) {
	secretKey := []byte("XDXDXDDDDDDDDDDDDDDDD")

	payload := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"_id":   user.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	tokenStr, err := token.SignedString(secretKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
