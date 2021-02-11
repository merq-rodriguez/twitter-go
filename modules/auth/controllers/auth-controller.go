package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/merq-rodriguez/twitter-go/common/jwt"
	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/merq-rodriguez/twitter-go/common/jwt/types"
	. "github.com/merq-rodriguez/twitter-go/common/response/errors"

	. "github.com/merq-rodriguez/twitter-go/modules/auth/dto"
	authService "github.com/merq-rodriguez/twitter-go/modules/auth/services"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"
	userService "github.com/merq-rodriguez/twitter-go/modules/users/services"
)

type AuthController struct{}

func (h *AuthController) SignIn(c echo.Context) error {
	dto := SignInDTO{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	err := dto.Validate()
	if err != nil {
		return BadRequestError(c, "Fields required", err)
	}

	user, err := authService.Signin(dto.Email, dto.Password)
	if err != nil {
		return UnauthorizedError(c, "Email or password invalid", nil)
	}

	jwtKey, err := jwt.CreateToken(user)
	if err != nil {
		return InternalServerError(c, "Error creating token", nil)
	}

	response := new(JsonWebToken)
	response.AccessToken = jwtKey

	return c.JSON(http.StatusOK, map[string]string{
		"token": response.AccessToken,
	})
}

func (h *AuthController) SignUp(c echo.Context) error {
	var user User
	dto := SignUpDTO{}
	err := dto.Decoder(c.Request().Body)

	log.Println(dto)
	if err != nil {
		return BadRequestError(c, "Fields required", err)
	}

	user, err = userService.FindUserByEmail(dto.Email)
	if !user.IsEmpty() {
		return BadRequestError(c, "User exists with email", nil)
	}

	if err != mongo.ErrNoDocuments {
		return BadRequestError(c, "", err)
	}

	user = dto.ConvertToUser()
	userCreated, err := userService.CreateUser(user)

	if err != nil {
		return BadRequestError(c, "User not created", err)
	}

	return c.JSON(http.StatusCreated, userCreated)
}
