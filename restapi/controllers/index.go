package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"multifileapp/config"
	"multifileapp/models"
	"net/http"
	"time"

	//"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//var employeeCollection *mongo.Collection = config.GetCollection(config.DB, "employees")
var productCollection *mongo.Collection = config.GetCollection(config.DB, "products")

func GetProducts(w http.ResponseWriter, r *http.Request) {

	// users := []models.User{
	// 	//user := models.User
	// 	{Id: "1", Name: "Kiran", Location: "Hyd", Title: "Trainer"},
	// 	{Id: "2", Name: "Sanjeev", Location: "Everywhere", Title: "Everything"},
	// }
	products := getProducts(config.DB, "products")
	fmt.Println("came here")
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}}
	json.NewEncoder(w).Encode(products)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println(product)
		err := insertProduct(config.DB, "products", product)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			response := models.UserResponse{Status: http.StatusOK,
				Message: "Unable to Save the data",
				Data:    map[string]interface{}{"data": err}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//stroe the data to the database
		w.WriteHeader(http.StatusOK)
		response := models.UserResponse{Status: http.StatusOK,
			Message: "Saved Successfuly",
			Data:    map[string]interface{}{"data": ""}}
		json.NewEncoder(w).Encode(response)
		return
	}

}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		fmt.Println(product)
		err := updateProduct(config.DB, "products", product)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			response := models.UserResponse{Status: http.StatusOK,
				Message: "Unable to Save the data",
				Data:    map[string]interface{}{"data": err}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//stroe the data to the database
		w.WriteHeader(http.StatusOK)
		response := models.UserResponse{Status: http.StatusOK,
			Message: "Saved Successfuly",
			Data:    map[string]interface{}{"data": ""}}
		json.NewEncoder(w).Encode(response)
		return
	}

}

func getProducts(client *mongo.Client, collectionName string) []models.Product {
	//Create the context for fetching the data
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//incase of timeout or error  cancel the context
	defer cancel()
	//Variable for the list of products
	var products []models.Product
	//var employees []models.Employee
	//Get the data from the product db collection name
	productCollection = client.Database("productdb").Collection(collectionName)
	//employeeCollection := client.Database("productdb").Collection(collectionName)
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
		//employees = append(employees, employee)
	}
	fmt.Println(products)

	return products
}

func insertProduct(client *mongo.Client, collectionName string, product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//incase of timeout or error  cancel the context
	defer cancel()
	produtToInsert := models.Product{
		Id:          primitive.NewObjectID(),
		Name:        product.Name,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Qty:         product.Qty,
	}
	result, err := productCollection.InsertOne(ctx, produtToInsert)
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)
	return nil

}

func updateProduct(client *mongo.Client, collectionName string, product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//incase of timeout or error  cancel the context
	defer cancel()
	productToUpdate := models.Product{
		Id:          product.Id,
		Name:        product.Name,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Qty:         product.Qty,
	}

	update := bson.M{
		"qty":         productToUpdate.Qty,
		"name":        productToUpdate.Name,
		"title":       productToUpdate.Title,
		"price":       productToUpdate.Price,
		"description": productToUpdate.Description,
	}

	result, err := productCollection.UpdateOne(ctx, bson.M{"id": productToUpdate.Id}, bson.M{"$set": update})

	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
