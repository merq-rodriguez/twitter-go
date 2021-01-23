package middlewares

import (
	"net/http"

	"github.com/merq-rodriguez/twitter-clone-backend-go/common/database"
	HttpStatus "github.com/merq-rodriguez/twitter-clone-backend-go/common/response/http"
)

/*
HandlerRequest type for manage request error
*/
type HandlerRequest func(http.ResponseWriter, *http.Request) error

/*
CheckDB middleware for show status of database
*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == false {
			http.Error(w, "Failed connection with database", HttpStatus.INTERNAL_SERVER_ERROR)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (fn HandlerRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) //Call handler function
	if err == nil {
		return
	}

	if database.CheckConnection() == false {
		http.Error(w, "Failed connection with database", HttpStatus.INTERNAL_SERVER_ERROR)
		return
	}

	fn.ServeHTTP(w, r)
	//Error handling...

}
