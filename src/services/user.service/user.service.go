package userservice

import (
	bookRepository "github.com/axel526/jikkosoft/src/data/repositories/book.repository"
	userRepository "github.com/axel526/jikkosoft/src/data/repositories/user.repository"
	e "github.com/axel526/jikkosoft/src/entity"
)

func Create(user e.User) error {
	return userRepository.Create(user)
}

func GetAll() (e.Users, error) {
	return userRepository.GetAll()
}

func GetUser(userId string) (e.User, error) {

	user, err := userRepository.GetByUserId(userId)

	if err != nil {
		return user, err
	}

	books, err := bookRepository.GetAllByUser(userId)

	if err != nil {
		return user, err
	}

	user.Books = books

	return user, nil
}

func Delete() {

}

func Update() {

}
