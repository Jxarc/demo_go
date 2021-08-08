package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", DBUser, DBPass, DBHost, DBPort)
	fmt.Printf("Connecting: %s  \n", uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Connected: %s  \n", uri)
	return client.Database(DBName).Collection(collection)
}
