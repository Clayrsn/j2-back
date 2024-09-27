package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	var (
		client *mongo.Client
	)

	err := godotenv.Load()
	if err != nil {
		panic("No .env file found")
	}

	ctx := context.TODO()
	//host := os.Getenv("MONGO_INITDB_HOST")
	//username := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	//password := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	connection := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(ctx, connection)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("mongo connection established")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	database := os.Getenv("MONGO_INITDB_DATABASE")
	collection := client.Database(database).Collection(collectionName)
	return collection
}
