package auth

import (
	"encoding/json"
	"net/http"
	"time"

	jwtResponse "github.com/merq-rodriguez/twitter-clone-backend-go/common/jwt/types"
	HttpStatus "github.com/merq-rodriguez/twitter-clone-backend-go/common/response/http"
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/users"
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/models"

	"github.com/merq-rodriguez/twitter-clone-backend-go/common/jwt"
)

/*
Signup function controller for create user
*/
func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error receive data user: "+err.Error(), HttpStatus.BAD_REQUEST)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email required", HttpStatus.BAD_REQUEST)
		return
	}

	if len(user.Password) < 8 {
		http.Error(w, "Password length more of 8 characters", HttpStatus.BAD_REQUEST)
		return
	}

	_, wanted, _ := users.UserAlreadyExist(user.Email)
	if wanted == true {
		http.Error(w, "A user already exists with this email", HttpStatus.BAD_REQUEST)
		return
	}

	_, status, err := users.CreateUser(user)
	if err != nil {
		http.Error(w, "User not created: "+err.Error(), HttpStatus.BAD_REQUEST)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", HttpStatus.BAD_REQUEST)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/*
Signin function controller for login user
*/
func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Credentials Invalid "+err.Error(), HttpStatus.UNAUTHORIZED)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email required", HttpStatus.BAD_REQUEST)
		return
	}

	document, exist := users.Signin(user.Email, user.Password)

	if exist == false {
		http.Error(w, "Email or password invalid", HttpStatus.BAD_REQUEST)
		return
	}

	jwtKey, err := jwt.CreateToken(document)
	if err != nil {
		http.Error(w, "Error creating token", HttpStatus.INTERNAL_SERVER_ERROR)
		return
	}

	response := jwtResponse.JsonWebToken{
		AccessToken: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expiresIn := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expiresIn,
	})
}
