package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type model struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
