package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
User model struct for MongoDB
*/
type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name      string             `bson:"name"        json:"name, omitempty"`
	Lastname  string             `bson:"lastname"    json:"lastname, omitempty"`
	BirthDate time.Time          `bson:"birthDate"   json:"birthDate, omitempty"`
	Username  string             `bson:"username"    json:"username"`
	Email     string             `bson:"email"       json:"email"`
	Password  string             `bson:"password"    json:"password, omitempty"`
	Avatar    string             `bson:"avatar"      json:"avatar, omitempty"`
	Banner    string             `bson:"banner"      json:"banner, omitempty"`
	Location  string             `bson:"location"    json:"location, omitempty"`
	Website   string             `bson:"website"     json:"website, omitempty"`
	Biography string             `bson:"biography"   json:"biography, omitempty"`
}
