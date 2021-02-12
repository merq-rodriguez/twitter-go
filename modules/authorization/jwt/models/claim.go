package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Clain struct type for processing JWT
*/
type JWTCustomClaim struct {
	Email    string             `json:"email"`
	Username string             `json:"username"`
	ID       primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	jwt.StandardClaims
}
