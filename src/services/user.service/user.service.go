package userservice

import (
	userRepository "github.com/axel526/jikkosoft/src/data/repositories/user.repository"
	e "github.com/axel526/jikkosoft/src/entity"
)

func Create(user e.User) error {
	return userRepository.Create(user)
}

func GetAll() (e.Users, error) {
	return userRepository.GetAll()
}

func Delete() {

}

func Update() {

}
