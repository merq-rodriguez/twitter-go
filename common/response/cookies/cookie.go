package cookies

import (
	"net/http"
	"time"
)

/*AddCookie function for add cookie to response  */
func AddCookieToken(w http.ResponseWriter, jwtKey string, expiresIn time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expiresIn,
	})
}
