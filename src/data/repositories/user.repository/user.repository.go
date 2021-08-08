package userrepository

import (
	"context"
	"errors"

	mongoDb "github.com/axel526/jikkosoft/src/data/databases/mongo"
	e "github.com/axel526/jikkosoft/src/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = mongoDb.GetCollection("user")
var ctx = context.Background()

func Create(user e.User) error {
	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
func GetAll() (e.Users, error) {

	var users e.Users

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user e.User
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetByUserId(userId string) (e.User, error) {

	var user e.User

	oid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": oid}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return user, err
	}

	if cur.Next(ctx) {
		err = cur.Decode(&user)
		return user, err
	}

	return user, errors.New("el usuario no existe")
}
