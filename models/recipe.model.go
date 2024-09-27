package models

type Recipe struct {
	model
	Name        string       `json:"name" bson:"name"`
	Ingredients []Ingredient `json:"ingredients" bson:"ingredients"`
}

type RecipeCreate struct {
	Name        string       `json:"name" bson:"name"`
	Ingredients []Ingredient `json:"ingredients" bson:"ingredients"`
}

type RecipeUpdate struct {
	Name        string       `json:"name" bson:"name"`
	Ingredients []Ingredient `json:"ingredients" bson:"ingredients"`
}
