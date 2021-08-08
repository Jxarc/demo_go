package bookrepository

import (
	"context"

	mongodb "github.com/axel526/jikkosoft/src/data/databases/mongo"
	e "github.com/axel526/jikkosoft/src/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = mongodb.GetCollection("book")
var ctx = context.Background()

func Create(book e.Book) error {
	_, err := collection.InsertOne(ctx, book)

	if err != nil {
		return err
	}

	return nil
}

func GetAllByUser(userId string) (e.Books, error) {
	var books e.Books

	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"user_id": oid}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var book e.Book
		err = cur.Decode(&book)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
