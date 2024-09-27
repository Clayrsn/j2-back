package services

import (
	"context"
	"time"

	"j2-api/configs"
	"j2-api/models"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService interface {
	GetUser(id string) (models.User, error)
	GetUserByToken(token string) (models.User, error)
	CreateUser(userCreate models.UserCreate) (models.User, error)
	UpdateUser(id string, userUpdate models.UserUpdate) (models.User, error)
	UpdateUserToken(id string, token string) (any, error)
	DeleteUser(id string) (any, error)
}

var usersCollection = configs.GetCollection(configs.DB, "users")

func GetUser(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	err := usersCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	return user, err
}

func GetUserByToken(token string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	err := usersCollection.FindOne(ctx, bson.M{"current_token": token}).Decode(&user)

	return user, err
}

func CreateUser(userCreate models.UserCreate) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	err := copier.Copy(&user, &userCreate)
	if err != nil {
		return user, err
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = recipesCollection.InsertOne(ctx, user)

	return user, err
}

func UpdateUser(id string, userUpdate models.UserUpdate) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	err := copier.Copy(&user, &userUpdate)
	if err != nil {
		return user, err
	}

	user.UpdatedAt = time.Now()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	_, err = usersCollection.UpdateOne(ctx, filter, update)

	return user, err
}

func UpdateUserToken(id string, token string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"current_token": token}

	_, err := usersCollection.UpdateOne(ctx, filter, update)

	return true, err
}

func DeleteUser(id string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = usersCollection.DeleteOne(ctx, bson.M{"_id": objID})

	return true, err
}
