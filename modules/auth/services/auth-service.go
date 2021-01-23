package services

import (
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/models"
	userService "github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/services"
	"golang.org/x/crypto/bcrypt"
)

/*
Signin function for login user
**/
func Signin(email string, password string) (models.User, bool) {
	user, wanted, _ := userService.UserAlreadyExist(email)
	if wanted == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
