package models

type Ingredient struct {
	model
	Name     string `json:"name" bson:"name"`
	Proteins string `json:"proteins" bson:"proteins"`
	Lipids   string `json:"lipids" bson:"lipids"`
	Carbs    string `json:"carbs" bson:"carbs"`
	Calories string `json:"calories" bson:"calories"`
}

type IngredientCreate struct {
	Name     string `json:"name" bson:"name"`
	Proteins string `json:"proteins" bson:"proteins"`
	Lipids   string `json:"lipids" bson:"lipids"`
	Carbs    string `json:"carbs" bson:"carbs"`
	Calories string `json:"calories" bson:"calories"`
}

type IngredientUpdate struct {
	Name     string `json:"name" bson:"name"`
	Proteins string `json:"proteins" bson:"proteins"`
	Lipids   string `json:"lipids" bson:"lipids"`
	Carbs    string `json:"carbs" bson:"carbs"`
	Calories string `json:"calories" bson:"calories"`
}
