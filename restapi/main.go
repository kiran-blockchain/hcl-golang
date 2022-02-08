package main

import (
	"log"
	"multifileapp/config"

	"multifileapp/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//checkDBConnection()
	// fmt.Println(config.SetupPort())
	router := mux.NewRouter() // create routes
	// // call the route configuration
	config.ConnectDB()
	// // construct routing table
	routes.Routes(router)

	log.Fatal(http.ListenAndServe(config.SetupPort(), router))

}

// func checkDBConnection() {
// 	var DB *mongo.Client = config.ConnectDB()
// 	myCollection := GetCollection(DB, "employees")
// 	fmt.Println(myCollection)
// }
