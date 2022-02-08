package main

import (
	// "fmt"
	// "log"
	"fmt"
	"multifileapp/config"

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
	collection := client.Database("productdb").Collection(collectionName)
	return collection
}
