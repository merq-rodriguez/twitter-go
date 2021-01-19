package middlewares

import (
	"net/http"

	"github.com/merq-rodriguez/twitter-clone-backend-go/common/database"
)

/*
CheckDB middleware for show status of database
*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == false {
			http.Error(w, "Failed connection with database", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
