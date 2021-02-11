package database

import (
	"context"
	"log"

	"github.com/merq-rodriguez/twitter-go/common/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
MongoCN Object connection to database
*/
var MongoCN = connectionDB()

var viper, err = config.Settings()

func getClientOptions() *options.ClientOptions {
	uriMongo := viper.GetString("database.uri")
	return options.Client().ApplyURI(uriMongo)
}

/*
	connectionDB: Connect with database
*/
func connectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), getClientOptions())
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
