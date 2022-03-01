package user

import (
	"errors"
	"log"

	userApp "user-service/application/user"
)

var (
	errorSavingUser = errors.New("error while saving user, try again later or verify if user data is correct")
)

type Service struct {
	repo userApp.Repository
}

func NewUserService(r userApp.Repository) Service {
	return Service{r}
}

func (s Service) Create(dto userApp.CreateUserDto) (userApp.ViewUserDto, error) {
	if err := dto.Validate(); err != nil {
		return userApp.ViewUserDto{}, err
	}

	var user userApp.User
	user.SetValuesFromCreateUserDto(dto)

	if err := s.repo.Save(user); err != nil {
		log.Println(err)
		return userApp.ViewUserDto{}, errorSavingUser
	}

	return user.ToViewUserDto(), nil
}

func (s Service) GetById(id string) (userApp.ViewUserDto, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return userApp.ViewUserDto{}, err
	}

	return user.ToViewUserDto(), nil
}

func (s Service) Update(id string, dto userApp.UpdateUserDto) (userApp.ViewUserDto, error) {
	if err := dto.Validate(); err != nil {
		return userApp.ViewUserDto{}, err
	}

	user, err := s.repo.FindById(id)
	if err != nil {
		return userApp.ViewUserDto{}, err
	}

	user.SetValuesFromUpdateUserDto(dto)

	if err := s.repo.Update(user); err != nil {
		return userApp.ViewUserDto{}, err
	}

	return user.ToViewUserDto(), err
}
