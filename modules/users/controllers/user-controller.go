package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	HttpStatus "github.com/merq-rodriguez/twitter-clone-backend-go/common/response/http"
	. "github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/models"
	userService "github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/services"
)

/*
GetProfile controller for get profile user
*/
func GetProfile(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if len(username) < 1 {
		http.Error(w, "Username paremeter not provided", HttpStatus.BAD_REQUEST)
	}

	profile, err := userService.GetProfile(username)
	if err != nil {
		http.Error(w, "Profile not exists", HttpStatus.NOT_FOUND)
		return
	}

	w.WriteHeader(HttpStatus.CREATED)
	json.NewEncoder(w).Encode(profile)
}

/*
UpdateUser controller for update user profile
*/
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)

	if err != nil {
		http.Error(w, "Invalid data user"+err.Error(), HttpStatus.BAD_REQUEST)
	}

	var status bool

	status, err = userService.UpdateUser(ID, u)
	if err != nil {
		http.Error(w, "User not update"+err.Error(), HttpStatus.BAD_REQUEST)
		return
	}

	if status == false {
	}

	w.WriteHeader(HttpStatus.CREATED)
}
