package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	//Get the instance of the datbase client.
	client, err := mongo.NewClient(options.Client().ApplyURI(GetMongoDBURI()))
	if err != nil {
		fmt.Println("Error in Connecting to Mongo db Database")
	}
	fmt.Println("Database getting connected")
	//Set the context to run the client in the background
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//Connect the client and run the background
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("productdb").Collection(collectionName)
	return collection
}
