package main

import (
	"j2-api/configs"
	"j2-api/controllers"

	"github.com/go-fuego/fuego"
)

func main() {
	configs.ConnectDB()
	s := fuego.NewServer()

	controllers.IngredientsResources{}.Routes(s)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
