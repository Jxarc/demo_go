package userrepository

import (
	e "github.com/axel526/jikkosoft/src/entity"
)

func GetAll() (e.Users, error) {

	var users = e.Users{
		{
			ID:      1,
			Name:    "Alexander",
			Age:     20,
			Address: "CR43 #12-32",
		},
		{
			ID:      2,
			Name:    "Johan",
			Age:     22,
			Address: "CR43 #12-32",
		},
		{
			ID:      3,
			Name:    "Johan",
			Age:     22,
			Address: "CR43 #12-32",
		},
	}

	return users, nil
}
