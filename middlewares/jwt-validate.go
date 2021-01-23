package middlewares

import (
	"fmt"
	"net/http"

	jwt "github.com/merq-rodriguez/twitter-clone-backend-go/modules/auth/jwt"
)

/*
JWTValidate function for validate Json Web Token
*/
func JWTValidate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		fmt.Println(token)
		_, _, _, err := jwt.ProcessToken(token)
		if err != nil {
			http.Error(w, "Error in token "+err.Error(), http.StatusBadRequest)
		}
	}
}
