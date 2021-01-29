package types

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Clain struct type for processing JWT
*/
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id, omitempty" json:"_id"`
	jwt.StandardClaims
}
