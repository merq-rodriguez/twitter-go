package users

import (
	"context"
	"time"

	"github.com/merq-rodriguez/twitter-clone-backend-go/common/database"
	. "github.com/merq-rodriguez/twitter-clone-backend-go/helpers"
	"github.com/merq-rodriguez/twitter-clone-backend-go/modules/crypt"
	. "github.com/merq-rodriguez/twitter-clone-backend-go/modules/users/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = database.MongoCN.Database("twitter")

/*
GetProfile function for obtain profile data of user
@Params: ID string
*/
func GetProfile(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var user User
	col := db.Collection("users")

	query := bson.M{"username": username}

	err := col.FindOne(ctx, query).Decode(&user)
	user.Password = ""

	if err != nil {
		return user, err
	}
	return user, nil
}

/*
CreateUser function: register a user
@Params u: models.User
*/
func CreateUser(u User) (string, bool, error) {
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
UpdateUser function: update user profile
@Params u: models.User
*/
func UpdateUser(ID string, u User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := db.Collection("users")

	register := make(map[string]interface{})
	if !IsEmpty(u.Name) {
		register["name"] = u.Name
	}

	if !IsEmpty(u.Lastname) {
		register["lastname"] = u.Lastname
	}

	register["birthDate"] = u.BirthDate

	if !IsEmpty(u.Banner) {
		register["banner"] = u.Banner
	}

	if !IsEmpty(u.Biography) {
		register["biography"] = u.Biography
	}

	if !IsEmpty(u.Location) {
		register["location"] = u.Location
	}

	queryUpdate := bson.M{
		"$set": register,
	}

	objectID, _ := primitive.ObjectIDFromHex(ID)

	queryFilter := bson.M{
		"_id": bson.M{"$eq": objectID},
	}

	_, err := col.UpdateOne(ctx, queryFilter, queryUpdate)

	if err != nil {
		return false, err
	}

	return true, nil
}

/*
UserAlreadyExist function: find user by email
@Params email: string
*/
func UserAlreadyExist(email string) (User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	col := db.Collection("users")
	query := bson.M{"email": email}

	var result User

	err := col.FindOne(ctx, query).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
