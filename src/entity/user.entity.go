package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `json:"name"`
	Age       int                `json:"age"`
	Address   string             `json:"address"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
	Books     Books              `bson:"-" json:"books,omitempty"`
}

type Users []User
