package config

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetMongoDBURI()))
	if err!=nil{
		fmt.Println("Error in Connecting to Mongo db Database")
	}
	return client;
}
