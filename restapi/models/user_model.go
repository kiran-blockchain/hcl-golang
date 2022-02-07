package models

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Title    string `json:"title"`
}
