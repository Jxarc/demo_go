package bookservice_test

import (
	"testing"
	"time"

	e "github.com/axel526/jikkosoft/src/entity"
	bookservice "github.com/axel526/jikkosoft/src/services/book.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	userID, _ := primitive.ObjectIDFromHex("610f84c43c23706b80789962")

	var books e.Books = e.Books{
		{
			Title:     "Titulo 1",
			Pages:     12,
			CreatedAt: time.Now(),
			UserID:    userID,
		},

		{
			Title:     "Titulo 2",
			Pages:     22,
			CreatedAt: time.Now(),
			UserID:    userID,
		},
	}

	for _, book := range books {
		oid := primitive.NewObjectID()
		book.ID = oid

		err := bookservice.Create(book)

		if err != nil {
			t.Error("Ocurrio un error: " + err.Error())
			t.Fail()
		} else {
			t.Log("Book creado!")
		}
	}

}
