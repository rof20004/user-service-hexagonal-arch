package services

import (
	"errors"
	"user-service/application/domains"
	"user-service/ports"
)

type Services struct {
	userRepository ports.Repository
}

func NewUserService(userRepository ports.Repository) *Services {
	return &Services{userRepository}
}

func (s *Services) Create(name, email, password string) (domains.User, error) {
	user := domains.NewUser(name, email, password)

	if err := s.userRepository.Save(user); err != nil {
		return user, errors.New("user creation has failed")
	}

	return user, nil
}

func (s *Services) GetById(id string) (domains.User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return user, errors.New("error while fetching user by id")
	}

	return user, nil
}