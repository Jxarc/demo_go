package bookrepository_test

import (
	"testing"

	bookRepository "github.com/axel526/jikkosoft/src/data/repositories/book.repository"
)

func TestGetAllByUser(t *testing.T) {
	books, err := bookRepository.GetAllByUser("610f84c43c23706b80789962")

	if err != nil {
		t.Error("Error al obtener los usuarios:  " + err.Error())
		t.Fail()
		return
	}

	for _, book := range books {
		t.Log("Libro ->" + book.ID.Hex() + ": Titulo " + book.Title)
	}

}
