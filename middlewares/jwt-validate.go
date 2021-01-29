package middlewares

import (
	"net/http"

	jwt "github.com/merq-rodriguez/twitter-clone-backend-go/modules/auth/jwt"
)

/*AuthJWT function for validate Json Web Token*/
func AuthJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, _, err := jwt.ValidateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
