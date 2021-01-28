package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/merq-rodriguez/twitter-clone-backend-go/common/config"
	"github.com/merq-rodriguez/twitter-clone-backend-go/middlewares"
	authController "github.com/merq-rodriguez/twitter-clone-backend-go/modules/auth/controllers"
	tweetController "github.com/merq-rodriguez/twitter-clone-backend-go/modules/tweets/controllers"
	userController "github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/controllers"
	"github.com/rs/cors"
)

/*
RunHandlers function: run handdlers with controllers enpoints
*/
func RunHandlers() {
	viper, err := config.Settings()
	router := mux.NewRouter()
	port := viper.GetInt("port")

	if port == 0 {
		port = 8080
	}

	if err != nil {
		log.Fatal("Error loading configuration")
	}

	router.HandleFunc(
		"/signup",
		middlewares.CheckDB(authController.Signup),
	).Methods("POST")

	router.HandleFunc(
		"/signin",
		middlewares.CheckDB(authController.Signin),
	).Methods("POST")

	router.HandleFunc(
		"/profile",
		middlewares.CheckDB(
			/* 	middlewares.JWTValidate( */ userController.GetProfile, /* ) */
		),
	).Methods("GET")

	router.HandleFunc(
		"/profile",
		middlewares.CheckDB(
			userController.UpdateUser,
		),
	).Methods("PUT")

	router.HandleFunc(
		"/tweets",
		middlewares.CheckDB(
			tweetController.CreateTweet,
		),
	).Methods("POST")

	handler := cors.AllowAll().Handler(router)
	log.Println("App running in port: " + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handler))

}
