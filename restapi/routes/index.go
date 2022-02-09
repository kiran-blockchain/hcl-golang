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
	/* params
	route url  /users
	route handler  controller.GetUsers
	*/
	router.HandleFunc("/api/getProducts", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/addProduct", controllers.InsertProduct).Methods("POST")
	router.HandleFunc("/api/updateProduct", controllers.UpdateProduct).Methods("POST")
}
