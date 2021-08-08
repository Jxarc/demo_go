package bookservice

import (
	bookRepository "github.com/axel526/jikkosoft/src/data/repositories/book.repository"
	e "github.com/axel526/jikkosoft/src/entity"
)

func Create(book e.Book) error {
	err := bookRepository.Create(book)

	if err != nil {
		return err
	}

	return nil
}
