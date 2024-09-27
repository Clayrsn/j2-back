package services

import (
	"context"
	"time"

	"j2-api/configs"
	"j2-api/models"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IngredientsService interface {
	GetIngredientsWithFilters(bson.M, *options.FindOptions) ([]models.Ingredient, error)
	GetIngredient(id string) (models.Ingredient, error)
	CreateIngredient(ingredientsCreate models.IngredientCreate) (models.Ingredient, error)
	UpdateIngredient(id string, ingredientsUpdate models.IngredientUpdate) (models.Ingredient, error)
	DeleteIngredient(id string) (any, error)
}

var ingredientsCollection = configs.GetCollection(configs.DB, "ingredients")

func GetIngredientsWithFilter(filters bson.M, opts *options.FindOptions) ([]models.Ingredient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var ingredients []models.Ingredient
	defer cancel()

	cur, err := ingredientsCollection.Find(ctx, filters, opts)
	if err != nil {
		return ingredients, err
	}

	for cur.Next(ctx) {
		var ingredient models.Ingredient
		err = cur.Decode(&ingredient)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func GetIngredient(id string) (models.Ingredient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var ingredient models.Ingredient
	defer cancel()

	err := ingredientsCollection.FindOne(ctx, bson.M{"ingredients.id": id}).Decode(&ingredient)

	return ingredient, err
}

func CreateIngredient(ingredientsCreate models.IngredientCreate) (models.Ingredient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var ingredient models.Ingredient
	defer cancel()

	err := copier.Copy(&ingredient, &ingredientsCreate)
	if err != nil {
		return ingredient, err
	}

	ingredient.ID = primitive.NewObjectID()
	ingredient.CreatedAt = time.Now()
	ingredient.UpdatedAt = time.Now()

	_, err = ingredientsCollection.InsertOne(ctx, ingredient)

	return ingredient, err
}

func UpdateIngredient(id string, ingredientsUpdate models.IngredientUpdate) (models.Ingredient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var ingredient models.Ingredient
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ingredient, err
	}

	err = copier.Copy(&ingredient, &ingredientsUpdate)
	if err != nil {
		return ingredient, err
	}

	ingredient.UpdatedAt = time.Now()

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": ingredient}

	_, err = ingredientsCollection.UpdateOne(ctx, filter, update)

	return ingredient, err
}

func DeleteIngredient(id string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = ingredientsCollection.DeleteOne(ctx, bson.M{"_id": objID})

	return true, err
}
