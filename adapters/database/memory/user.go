package memory

import (
	"errors"

	userApp "user-service/application/user"
)

var (
	errorUserNotFound = errors.New("user not found")
)

type UserMemoryDatabase struct {
	store map[string]userApp.User
}

func NewUserMemoryDatabase() *UserMemoryDatabase {
	return &UserMemoryDatabase{store: make(map[string]userApp.User)}
}

func (umdb *UserMemoryDatabase) Save(user userApp.User) error {
	umdb.store[user.ID] = user
	return nil
}

func (umdb *UserMemoryDatabase) FindById(id string) (userApp.User, error) {
	user, ok := umdb.store[id]
	if !ok {
		return user, errorUserNotFound
	}

	return user, nil
}
