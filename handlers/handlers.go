package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/merq-rodriguez/twitter-clone-backend-go/middlewares"
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/auth"
	"github.com/rs/cors"
)

/*
RunHandlers function: run handdlers with controllers enpoints
*/
func RunHandlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middlewares.CheckDB(auth.Signup)).Methods("POST")
	router.HandleFunc("/signin", middlewares.CheckDB(auth.Signin)).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
