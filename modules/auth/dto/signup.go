package dto

import (
	"encoding/json"
	"io"
	. "time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"
)

type SignUpDTO struct {
	Email     string
	Password  string
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

func (a SignUpDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
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
func (a SignUpDTO) ConvertToUser() User {
	return User{
		Email:     a.Email,
		Password:  a.Password,
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

func (a *SignUpDTO) Decoder(data io.Reader) error {
	err := json.NewDecoder(data).Decode(a)
	if err != nil {
		return err
	}
	return nil
}
