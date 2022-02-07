package main

import (
	"fmt"
	"log"
	"multifileapp/config"
	"multifileapp/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(config.SetupPort())
	router := mux.NewRouter() // create routes
	// call the route configuration
	routes.Routes(router)
	log.Fatal(http.ListenAndServe(config.SetupPort(), router))

}
