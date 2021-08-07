package userservice

import (
	userRepository "github.com/axel526/jikkosoft/src/data/repositories/user.repository"
	e "github.com/axel526/jikkosoft/src/entity"
)

func Create() {

}

func GetAll() (e.Users, error) {
	return userRepository.GetAll()
}

func Delete() {

}

func Update() {

}
