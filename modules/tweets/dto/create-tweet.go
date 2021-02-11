package dto

import (
	"encoding/json"
	"io"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	. "github.com/merq-rodriguez/twitter-go/modules/tweets/models"
)

type CreateTweetDTO struct {
	UserID  string `bson:"userId" json:"userId, omitempty"`
	Message string `bson:"message" json:"message, omitempty"`
}

func (t CreateTweetDTO) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Message, validation.Required),
		validation.Field(&t.UserID, validation.Required),
	)
}

func (t *CreateTweetDTO) Decoder(data io.Reader) error {
	err := json.NewDecoder(data).Decode(t)
	if err != nil {
		return err
	}
	return nil
}

func (t *CreateTweetDTO) ConvertToTweet() Tweet {
	return Tweet{
		UserID:  t.UserID,
		Message: t.Message,
	}
}
