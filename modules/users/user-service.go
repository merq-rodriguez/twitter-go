package users

import (
	"context"
	"time"

	"github.com/merq-rodriguez/twitter-clone-backend-go/common/database"
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/crypt"
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var db = database.MongoCN.Database("twitter")

/*
CreateUser function: register a user
@Params u: models.User
*/
func CreateUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := db.Collection("users")
	u.ID = primitive.NewObjectID()

	u.Password, _ = crypt.HashPassword(u.Password)
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

/*
UserAlreadyExist function: find user by email
@Params email: string
*/
func UserAlreadyExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := db.Collection("users")
	query := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, query).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}

/*
Signin function for login user
**/
func Signin(email string, password string) (models.User, bool) {
	user, wanted, _ := UserAlreadyExist(email)
	if wanted == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
