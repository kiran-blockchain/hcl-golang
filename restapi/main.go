package main

import (
	// "fmt"
	// "log"
	"context"
	"fmt"
	"multifileapp/config"
	"multifileapp/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "multifileapp/routes"
	// "net/http"
	// "github.com/gorilla/mux"
)

func main() {
	checkDBConnection()
	// fmt.Println(config.SetupPort())
	// router := mux.NewRouter() // create routes
	// // call the route configuration

	// // construct routing table
	// routes.Routes(router)

	// log.Fatal(http.ListenAndServe(config.SetupPort(), router))

}
func checkDBConnection() {
	var DB *mongo.Client = config.ConnectDB()
	myCollection := GetCollection(DB, "products")
	fmt.Println(myCollection)
}
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var products []models.Product
	productCollection := client.Database("productdb").Collection(collectionName)
	results, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
	defer results.Close(ctx)
	for results.Next(ctx) {
		var product models.Product
		if err = results.Decode(&product); err != nil {
			fmt.Println(err)
		}
		products = append(products, product)
	}
	fmt.Println(products)

	return productCollection
}
