package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `json:"title"`
	Pages     int                `json:"pages"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type Books []Book
