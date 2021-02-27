package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"
)

type SearchUserDTO struct {
	Page       int64
	TextSearch string
}

func (s *SearchUserDTO) Bind(data SearchUser) {
	s.Page = data.Page
	s.TextSearch = data.TextSearch
}

func (s *SearchUserDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Page, validation.Required, is.Int),
		validation.Field(&s.TextSearch, validation.Required),
	)
}
