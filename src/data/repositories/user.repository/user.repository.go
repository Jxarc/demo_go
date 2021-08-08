package userrepository

import (
	"context"

	mongoDb "github.com/axel526/jikkosoft/src/data/databases/mongo"
	e "github.com/axel526/jikkosoft/src/entity"
	"go.mongodb.org/mongo-driver/bson"
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
