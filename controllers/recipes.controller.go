package controllers

import (
	"strconv"

	"j2-api/models"
	"j2-api/services"

	"github.com/go-fuego/fuego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RecipesResources struct {
	RecipesService services.RecipesService
}

func (rs RecipesResources) Routes(s *fuego.Server) {
	recipesGroup := fuego.Group(s, "/recipes")

	fuego.Get(recipesGroup, "/", rs.getRecipeWithFilters)
	fuego.Get(recipesGroup, "/{id}", rs.getRecipes)
	fuego.Post(recipesGroup, "/", rs.postRecipes)
	fuego.Put(recipesGroup, "/{id}", rs.putRecipes)
	fuego.Delete(recipesGroup, "/{id}", rs.deleteRecipes)
}

func (rs RecipesResources) getRecipeWithFilters(c fuego.ContextNoBody) ([]models.Recipe, error) {
	filters := bson.M{}

	for key, values := range c.Request().URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	findOptions := options.Find()
	if limit := c.QueryParam("limit"); limit != "" {
		limitInt, err := strconv.ParseInt(limit, 10, 64)
		if err == nil {
			findOptions.SetLimit(limitInt)
		}
	}
	if skip := c.QueryParam("skip"); skip != "" {
		skipInt, err := strconv.ParseInt(skip, 10, 64)
		if err == nil {
			findOptions.SetSkip(skipInt)
		}
	}

	return rs.RecipesService.GetRecipesWithFilters(filters, findOptions)
}

func (rs RecipesResources) postRecipes(c *fuego.ContextWithBody[models.RecipeCreate]) (models.Recipe, error) {
	body, err := c.Body()
	if err != nil {
		return models.Recipe{}, err
	}

	createRecipe, err := rs.RecipesService.CreateRecipe(body)
	if err != nil {
		return models.Recipe{}, err
	}

	return createRecipe, nil
}

func (rs RecipesResources) getRecipes(c fuego.ContextNoBody) (models.Recipe, error) {
	id := c.PathParam("id")

	return rs.RecipesService.GetRecipe(id)
}

func (rs RecipesResources) putRecipes(c *fuego.ContextWithBody[models.RecipeUpdate]) (models.Recipe, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return models.Recipe{}, err
	}

	updateRecipe, err := rs.RecipesService.UpdateRecipe(id, body)
	if err != nil {
		return models.Recipe{}, err
	}

	return updateRecipe, nil
}

func (rs RecipesResources) deleteRecipes(c *fuego.ContextNoBody) (any, error) {
	return rs.RecipesService.DeleteRecipe(c.PathParam("id"))
}
