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

type RecipesService interface {
	GetRecipesWithFilters(bson.M, *options.FindOptions) ([]models.Recipe, error)
	GetRecipe(id string) (models.Recipe, error)
	CreateRecipe(recipeCreate models.RecipeCreate) (models.Recipe, error)
	UpdateRecipe(id string, recipeUpdate models.RecipeUpdate) (models.Recipe, error)
	DeleteRecipe(id string) (any, error)
}

var recipesCollection = configs.GetCollection(configs.DB, "recipes")

func GetRecipesWithFilters(filters bson.M, opts *options.FindOptions) ([]models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var recipes []models.Recipe
	defer cancel()

	cur, err := recipesCollection.Find(ctx, filters, opts)
	if err != nil {
		return recipes, err
	}

	for cur.Next(ctx) {
		var recipe models.Recipe
		err = cur.Decode(&recipe)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func GetRecipe(id string) (models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var recipe models.Recipe
	defer cancel()

	err := recipesCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&recipe)

	return recipe, err
}

func CreateRecipe(recipeCreate models.RecipeCreate) (models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var recipe models.Recipe
	defer cancel()

	err := copier.Copy(&recipe, &recipeCreate)
	if err != nil {
		return recipe, err
	}

	recipe.ID = primitive.NewObjectID()
	recipe.CreatedAt = time.Now()
	recipe.UpdatedAt = time.Now()

	_, err = recipesCollection.InsertOne(ctx, recipe)

	return recipe, err
}

func UpdateRecipe(id string, recipeUpdate models.RecipeUpdate) (models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var recipe models.Recipe
	defer cancel()

	err := copier.Copy(&recipe, &recipeUpdate)
	if err != nil {
		return recipe, err
	}

	recipe.UpdatedAt = time.Now()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": recipeUpdate}

	_, err = recipesCollection.UpdateOne(ctx, filter, update)

	return recipe, err
}

func DeleteRecipe(id string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var recipe models.Recipe
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return recipe, err
	}

	_, err = recipesCollection.DeleteOne(ctx, bson.M{"_id": objID})

	return recipe, err
}
