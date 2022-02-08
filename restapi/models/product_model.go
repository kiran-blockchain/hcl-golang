package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `json:"_id,omitempty"`
	Price       int                `json:"price,omitempty" `
	Description string             `json:"description,omitempty" `
	Title       string             `json:"title,omitempty"`
	Name        string             `json:"name,omitempty"`
	Qty         int                `json:"qty,omitempty"`
}
