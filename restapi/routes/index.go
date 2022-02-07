package routes

import (
	"multifileapp/controllers"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	//router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	// router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET")
	// router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT")
	// router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
}
