package models

type Right struct {
	model
	Name string `json:"name" bson:"name"`
}
