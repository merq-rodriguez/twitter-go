package models

import "time"

type Tweet struct {
	UserID    string    `bson:"userId" json:"userId, omitempty"`
	Message   string    `bson:"message" json:"message, omitempty"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp, omitempty"`
}
