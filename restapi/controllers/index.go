package controllers

import (
	"encoding/json"
	"fmt"
	"multifileapp/models"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		//user := models.User
		{Id: "1", Name: "Kiran", Location: "Hyd", Title: "Trainer"},
		{Id: "2", Name: "Sanjeev", Location: "Everywhere", Title: "Everything"},
	}

	fmt.Println("came here")
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
