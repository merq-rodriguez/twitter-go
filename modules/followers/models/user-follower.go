package models

import "time"

type UserFollower struct {
	UserID         string    `bson:"userId" json:"userId"`
	UserFollowerID string    `bson:"followerId" json:"followerId"`
	CreatedAt      time.Time `bson:"timestamp" json:"timestamp"`
}
