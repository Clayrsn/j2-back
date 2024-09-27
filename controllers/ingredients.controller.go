package controllers

import (
	"strconv"

	"j2-api/models"
	"j2-api/services"

	"github.com/go-fuego/fuego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IngredientsResources struct {
	IngredientsService services.IngredientsService
}

func (rs IngredientsResources) Routes(s *fuego.Server) {
	ingredientsGroup := fuego.Group(s, "/ingredients")

	fuego.Get(ingredientsGroup, "/", rs.getIngredientsWithFilters)
	fuego.Get(ingredientsGroup, "/{id}", rs.getIngredient)
	fuego.Post(ingredientsGroup, "/", rs.postIngredient)
	fuego.Put(ingredientsGroup, "/{id}", rs.putIngredient)
	fuego.Delete(ingredientsGroup, "/{id}", rs.deleteIngredient)
}

func (rs IngredientsResources) getIngredientsWithFilters(c fuego.ContextNoBody) ([]models.Ingredient, error) {
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

	return rs.IngredientsService.GetIngredientsWithFilters(filters, findOptions)
}

func (rs IngredientsResources) postIngredient(c *fuego.ContextWithBody[models.IngredientCreate]) (models.Ingredient, error) {
	body, err := c.Body()
	if err != nil {
		return models.Ingredient{}, err
	}

	createIngredient, err := rs.IngredientsService.CreateIngredient(body)
	if err != nil {
		return models.Ingredient{}, err
	}

	return createIngredient, nil
}

func (rs IngredientsResources) getIngredient(c fuego.ContextNoBody) (models.Ingredient, error) {
	id := c.PathParam("id")

	return rs.IngredientsService.GetIngredient(id)
}

func (rs IngredientsResources) putIngredient(c *fuego.ContextWithBody[models.IngredientUpdate]) (models.Ingredient, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return models.Ingredient{}, err
	}

	updateIngredient, err := rs.IngredientsService.UpdateIngredient(id, body)
	if err != nil {
		return models.Ingredient{}, err
	}

	return updateIngredient, nil
}

func (rs IngredientsResources) deleteIngredient(c *fuego.ContextNoBody) (any, error) {
	return rs.IngredientsService.DeleteIngredient(c.PathParam("id"))
}
