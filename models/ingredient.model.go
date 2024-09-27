package models

type Ingredient struct {
	model
	Name       string  `json:"name" bson:"name"`
	Proteins   float64 `json:"proteins" bson:"proteins"`
	Lipids     float64 `json:"lipids" bson:"lipids"`
	Carbs      float64 `json:"carbs" bson:"carbs"`
	Calories   float64 `json:"calories" bson:"calories"`
	IsPersonal bool    `json:"isPersonal" bson:"isPersonal" default:"false"`
	Creator    User    `json:"creator" bson:"creator" default:"NULL"`
}

type IngredientCreate struct {
	Name       string  `json:"name" bson:"name"`
	Proteins   float64 `json:"proteins" bson:"proteins"`
	Lipids     float64 `json:"lipids" bson:"lipids"`
	Carbs      float64 `json:"carbs" bson:"carbs"`
	Calories   float64 `json:"calories" bson:"calories"`
	IsPersonal bool    `json:"isPersonal" bson:"isPersonal" default:"false"`
	Creator    User    `json:"creator" bson:"creator" default:"NULL"`
}

type IngredientUpdate struct {
	Name       string  `json:"name" bson:"name"`
	Proteins   float64 `json:"proteins" bson:"proteins"`
	Lipids     float64 `json:"lipids" bson:"lipids"`
	Carbs      float64 `json:"carbs" bson:"carbs"`
	Calories   float64 `json:"calories" bson:"calories"`
	IsPersonal bool    `json:"isPersonal" bson:"isPersonal" default:"false"`
	Creator    User    `json:"creator" bson:"creator" default:"NULL"`
}
