package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
MongoCN Object connection to database
*/
var MongoCN = connectionDB()
var url = "mongodb://localhost:27017/twitter"

var clientOptions = options.Client().ApplyURI(url)

/*
	connectionDB: Connect with database
*/
func connectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connected successful with database")
	return client
}

/*
CheckConnection: ping to database
*/
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
