package dto

import (
	. "time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"
)

type UpdateUserDTO struct {
	Email     string
	Name      string
	Lastname  string
	BirthDate Time
	Username  string
	Avatar    string
	Banner    string
	Location  string
	Website   string
	Biography string
}

func (a UpdateUserDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Lastname, validation.Required),
		validation.Field(&a.BirthDate, validation.Required, validation.Date(a.BirthDate.String())),
		validation.Field(&a.Username, validation.Required),
		validation.Field(&a.Avatar, validation.Required),
		validation.Field(&a.Banner, validation.Required),
		validation.Field(&a.Location, validation.Required),
		validation.Field(&a.Website, validation.Required),
		validation.Field(&a.Biography, validation.Required),
	)
}

//ConvertToUser function for parsear dto to User struct
func (a UpdateUserDTO) ConvertToUser() User {
	return User{
		Email:     a.Email,
		Name:      a.Name,
		Lastname:  a.Lastname,
		BirthDate: a.BirthDate,
		Username:  a.Username,
		Avatar:    a.Avatar,
		Banner:    a.Banner,
		Location:  a.Location,
		Website:   a.Website,
		Biography: a.Biography,
	}
}
