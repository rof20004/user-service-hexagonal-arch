package user

import (
	"errors"
	"sync"
	userApp "user-service/application/domains"
)

var (
	errorUserNotFound = errors.New("user not found")
)

type MemoryDatabase struct {
	lock *sync.RWMutex
	store map[string]userApp.User
}

func NewUserMemoryDatabase() *MemoryDatabase {
	return &MemoryDatabase{lock: &sync.RWMutex{}, store: make(map[string]userApp.User)}
}

func (umdb *MemoryDatabase) Save(user userApp.User) error {
	umdb.lock.Lock()
	defer umdb.lock.Unlock()

	umdb.store[user.ID] = user
	return nil
}

func (umdb *MemoryDatabase) FindById(id string) (userApp.User, error) {
	umdb.lock.RLock()
	defer umdb.lock.RUnlock()

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