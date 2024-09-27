package models

type User struct {
	model
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
