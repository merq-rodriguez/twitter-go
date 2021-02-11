package services

import (
	"context"
	"time"

	"github.com/merq-rodriguez/twitter-go/common/database"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"
	"go.mongodb.org/mongo-driver/bson"
)

var db = database.MongoCN.Database("twitter")

/*
GetProfile function for obtain profile data of user
@Params: ID string
*/
func GetProfile(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var user User
	col := db.Collection("users")
	query := bson.M{"username": username}

	err := col.FindOne(ctx, query).Decode(&user)
	user.Password = ""

	if err != nil {
		return user, err
	}
	return user, nil
}
