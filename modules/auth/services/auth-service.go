package services

import (
	"errors"

	. "github.com/merq-rodriguez/twitter-go/modules/auth/dto"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"
	userService "github.com/merq-rodriguez/twitter-go/modules/users/services"
	"golang.org/x/crypto/bcrypt"
)

/*
Signin function for login user
**/
func Signin(email string, password string) (User, error) {
	user, err := userService.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, err
	}

	return user, nil
}

func Signup(dto SignUpDTO) (User, error) {
	user, err := userService.FindUserByEmail(dto.Email)
	if err != nil {
		return user, err
	}

	if user.IsEmpty() {
		return user, errors.New("User exists")
	}

	user = dto.ConvertToUser()
	result, err := userService.CreateUser(user)
	if err != nil {
		return result, err
	}

	return result, nil
}
