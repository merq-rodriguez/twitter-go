package dto

import (
	"encoding/json"
	"io"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	. "github.com/merq-rodriguez/twitter-go/modules/followers/models"
)

type AddFollowerDTO struct {
	UserID     string `bson:"userId" json:"userId"`
	FollowerID string `bson:"followerId" json:"followerId"`
}

func (t AddFollowerDTO) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.UserID, validation.Required),
		validation.Field(&t.FollowerID, validation.Required),
	)
}

func (t *AddFollowerDTO) Bind(data io.Reader) error {
	err := json.NewDecoder(data).Decode(t)
	if err != nil {
		return err
	}
	return nil
}

func (t AddFollowerDTO) ConvertToFollower() UserFollower {
	return UserFollower{
		UserID:         t.UserID,
		UserFollowerID: t.FollowerID,
	}
}
