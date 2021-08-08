package userservice_test

import (
	"testing"
	"time"

	"strconv"

	e "github.com/axel526/jikkosoft/src/entity"
	userService "github.com/axel526/jikkosoft/src/services/user.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {

	var users e.Users = e.Users{
		{
			Name:      "Johan",
			Address:   "CR23",
			Age:       23,
			CreatedAt: time.Now(),
		},
		{
			Name:      "Alexander",
			Address:   "CR43",
			Age:       32,
			CreatedAt: time.Now(),
		},
	}

	for _, user := range users {
		oid := primitive.NewObjectID()
		user.ID = oid
		err := userService.Create(user)

		if err != nil {
			t.Error("Error al insertar usuario: " + err.Error())
			t.Fail()
		} else {
			t.Log("Usuario Creado!")
		}
	}

}

func TestGetAll(t *testing.T) {
	users, err := userService.GetAll()

	if err != nil {
		t.Error("Ocurrio un error al obtener la información: " + err.Error())
		t.Fail()
		return
	}

	for i, user := range users {
		t.Log(strconv.Itoa(i) + ": Nombre ->" + user.Name)
		t.Log(strconv.Itoa(i) + ": ID ->" + user.ID.Hex())
	}
}
