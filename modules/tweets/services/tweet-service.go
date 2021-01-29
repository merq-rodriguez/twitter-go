package services

import (
	"context"
	"log"
	"time"

	conn "github.com/merq-rodriguez/twitter-clone-backend-go/common/database"
	. "github.com/merq-rodriguez/twitter-clone-backend-go/modules/tweets/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = conn.MongoCN.Database("twitter")

/*
GetTweetsByUserID function for obtain tweets by user
*/
func GetTweetsByUserID(ID string, page int64) ([]*Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := db.Collection("tweet")

	var results []*Tweet

	query := bson.M{
		"userId": ID,
	}

	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, query, opts)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var registry Tweet
		err := cursor.Decode(&registry)
		if err != nil {
			return results, false
		}

		results = append(results, &registry)
	}
	return results, true
}

/*
CreateTweet function for add news tweets
*/
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
