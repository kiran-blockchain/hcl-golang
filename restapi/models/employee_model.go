package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	Id     primitive.ObjectID `json:"_id,omitempty"`
	Name   string             `json:"name,omitempty" `
	Salary int                `json:"salary,omitempty" `
	age    int                `json:"age,omitempty"`
}
