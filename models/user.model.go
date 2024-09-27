package models

type User struct {
	model
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Role         Role   `json:"role" bson:"role"`
	CurrentToken string `json:"current_token" bson:"current_token"`
}

type UserCreate struct {
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Role         Role   `json:"role" bson:"role"`
	CurrentToken string `json:"current_token" bson:"current_token"`
}

type UserUpdate struct {
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Role         Role   `json:"role" bson:"role"`
	CurrentToken string `json:"current_token" bson:"current_token"`
}
