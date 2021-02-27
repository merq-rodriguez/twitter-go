package context

import (
	"context"
	"log"
	"net/http"
	"time"
)

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}

/*AddContext for add context to handlers routes */
func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("username")

		if cookie != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
