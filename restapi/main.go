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
	//Create the context for fetching the data
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//incase of timeout or error  cancel the context
	defer cancel()
	//Variable for the list of products
	var products []models.Product
	//Get the data from the product db collection name
	productCollection := client.Database("productdb").Collection(collectionName)
	//Fetch the array of searched items
	results, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
	//close the context incase of error
	defer results.Close(ctx)
	//loop through the result set and convert into the the product strucutre
	for results.Next(ctx) {
		var product models.Product
		//Decode the results
		if err = results.Decode(&product); err != nil {
			fmt.Println(err)
		}
		//append the product to the products collection.
		products = append(products, product)
	}
	fmt.Println(products)

	return productCollection
}
