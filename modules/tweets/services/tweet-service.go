package services

import (
	"context"
	"time"

	conn "github.com/merq-rodriguez/twitter-clone-backend-go/common/database"
	. "github.com/merq-rodriguez/twitter-clone-backend-go/modules/tweets/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = conn.MongoCN.Database("twitter")

func CreateTweet(t Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := db.Collection("tweet")

	registry := bson.M{
		"userid":    t.UserID,
		"message":   t.Message,
		"timestamp": t.Timestamp,
	}

	result, err := col.InsertOne(ctx, registry)
	if err != nil {
		return "", false, err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)
	return objectID.String(), true, nil

}
