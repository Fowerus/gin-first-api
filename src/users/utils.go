package users

import (
	"errors"
	"strconv"
)

func RetrieveUser(id string) (*User, error) {
	for _, user := range users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func getNewId(id string) string {
	intId, err := strconv.Atoi(id)
	if err == nil {
		intId++
		stringId := strconv.Itoa(intId)
		return stringId
	}
	return ""
}
