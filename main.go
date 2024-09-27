package main

import (
	"os"

	"j2-api/controllers"
	"j2-api/middlewares"

	"github.com/go-fuego/fuego"
)

func main() {
	client, err := middlewares.InitializeFirebaseApp(os.Getenv("FIREBASE_CONFIG_PATH"))
	if err != nil {
		panic(err)
	}

	s := fuego.NewServer()
	api := fuego.Group(s, "/api")

	fuego.Use(api, middlewares.MiddlewareResources{}.FirebaseAuthMiddleware(client))

	controllers.IngredientsResources{}.Routes(api)
	controllers.RecipesResources{}.Routes(api)

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
