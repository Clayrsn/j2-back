package models

type Role struct {
	model
	Name   string  `json:"name" bson:"name"`
	Rights []Right `json:"rights" bson:"rights"`
}
