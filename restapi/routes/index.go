package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	//router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	// router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET")
	// router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT")
	// router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE")
	router.HandleFunc("/users", handleUser).Methods("GET")
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("came here")
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category")
}
