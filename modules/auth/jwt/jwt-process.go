package jwt

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	constant "github.com/merq-rodriguez/twitter-clone-backend-go/common/jwt/constants"
	userService "github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/services"

	"github.com/merq-rodriguez/twitter-clone-backend-go/common/jwt/types"
)

var Email string

var UserID string

/*
ProcessToken function for process jwt
*/
func ProcessToken(token string) (*types.Claim, bool, string, error) {
	secretKey := []byte(constant.SECRET_KEY)
	claims := &types.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Format token invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(
		token,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

	if err == nil {
		_, wanted, _ := userService.UserAlreadyExist(claims.Email)
		if wanted == true {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, wanted, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}
	return claims, false, string(""), err
}
