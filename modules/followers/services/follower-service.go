package services

import (
	"context"
	"time"

	"github.com/merq-rodriguez/twitter-go/common/database"
	. "github.com/merq-rodriguez/twitter-go/modules/followers/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = database.MongoCN.Database("twitter")

func AddFollower(follower UserFollower) (UserFollower, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := db.Collection("follower")
	follower.CreatedAt = time.Now()
	_, err := col.InsertOne(ctx, follower)

	if err != nil {
		return UserFollower{}, err
	}
	return follower, nil
}

func GetFollowersByUser(ID string, page int64) ([]*UserFollower, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []*UserFollower
	col := db.Collection("follower")

	query := bson.M{"userId": ID}

	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, query, opts)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var follower UserFollower
		err := cursor.Decode(&follower)
		if err != nil {
			return nil, err
		}
		results = append(results, &follower)

	}
	return results, nil
}
