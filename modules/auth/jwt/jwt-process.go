package jwt

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	constant "github.com/merq-rodriguez/twitter-clone-backend-go/common/jwt/constants"
	userService "github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/services"

	. "github.com/merq-rodriguez/twitter-clone-backend-go/common/jwt/types"
)

var Email string
var UserID string

/*validateFormatToken function*/
func validateFormatToken(bearerToken string, claims *Claim) (string, bool) {
	splitToken := strings.Split(bearerToken, "Bearer")
	if len(splitToken) != 2 {
		return string(""), false
	}
	tk := strings.TrimSpace(splitToken[1])
	return tk, true
}

/*decodeToken function*/
func decodeToken(tk string, claims *Claim) (*jwt.Token, error) {
	secretKey := []byte(constant.SECRET_KEY)
	return jwt.ParseWithClaims(
		tk,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
}

/*ValidateToken function for process jwt*/
func ValidateToken(bearerToken string) (*Claim, bool, error) {
	claims := &Claim{}
	token, status := validateFormatToken(bearerToken, claims)

	if status == false {
		return claims, status, errors.New("Format token invalid")
	}

	tkn, err := decodeToken(token, claims)

	if err == nil {
		_, wanted, _ := userService.UserAlreadyExist(claims.Email)
		if wanted == true {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, wanted, nil
	}

	if !tkn.Valid {
		return claims, false, errors.New("Invalid Token")
	}
	return claims, false, err
}
