package users

import (
	. "github.com/merq-rodriguez/twitter-go/common/context"
	"github.com/merq-rodriguez/twitter-go/common/database"
	. "github.com/merq-rodriguez/twitter-go/helpers"
	"github.com/merq-rodriguez/twitter-go/modules/crypt"
	. "github.com/merq-rodriguez/twitter-go/modules/users/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = database.MongoCN.Database("twitter")

func SearchUsers(s SearchUser) ([]*User, error) {
	ctx, cancel := GetContext()
	defer cancel()

	var results []*User

	col := db.Collection("users")
	opts := options.Find()
	opts.SetSkip((s.Page - 1) * 20)
	opts.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + s.TextSearch},
	}

	cursor, err := col.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var us User
		err := cursor.Decode(us)
		if err != nil {
			return nil, err
		}

		results = append(results, &us)
	}
	return results, nil
}

/*
FindUserByEmail function: find user by email
@Params email: string
*/
func FindUserByEmail(email string) (User, error) {
	ctx, cancel := GetContext()
	defer cancel()

	var user User
	col := db.Collection("users")
	query := bson.M{"email": email}

	err := col.FindOne(ctx, query).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

/*
CreateUser function: register a user
@Params u: models.User
*/
func CreateUser(user User) (User, error) {
	ctx, cancel := GetContext()
	defer cancel()

	col := db.Collection("users")

	user.Password, _ = crypt.HashPassword(user.Password)
	user.ID = primitive.NewObjectID()
	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return User{}, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

/*
UpdateUser function: update user profile
@Params u: models.User
*/
func UpdateUser(ID string, u User) (User, error) {
	ctx, cancel := GetContext()
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
		return User{}, err
	}
	return u, nil
}
