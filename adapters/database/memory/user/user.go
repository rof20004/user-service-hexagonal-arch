package user

import (
	"errors"

	userApp "user-service/application/user"
)

var (
	errorUserNotFound = errors.New("user not found")
)

type MemoryDatabase struct {
	store map[string]userApp.User
}

func NewUserMemoryDatabase() *MemoryDatabase {
	return &MemoryDatabase{store: make(map[string]userApp.User)}
}

func (umdb *MemoryDatabase) Save(user userApp.User) error {
	umdb.store[user.ID] = user
	return nil
}

func (umdb *MemoryDatabase) FindById(id string) (userApp.User, error) {
	user, ok := umdb.store[id]
	if !ok {
		return user, errorUserNotFound
	}

	return user, nil
}

func (umdb *MemoryDatabase) Update(user userApp.User) error {
	umdb.store[user.ID] = user
	return nil
}