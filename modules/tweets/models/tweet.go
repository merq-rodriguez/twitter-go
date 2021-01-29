package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Tweet struct model
*/
type Tweet struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	UserID    string             `bson:"userId" json:"userId, omitempty"`
	Message   string             `bson:"message" json:"message, omitempty"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp, omitempty"`
}
