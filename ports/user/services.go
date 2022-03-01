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
	var view userApp.ViewUserDto

	if err := dto.Validate(); err != nil {
		return view, err
	}

	var user = dto.ToDomain()

	if err := s.repo.Save(user); err != nil {
		log.Println(err)
		return view, errorSavingUser
	}

	return view.FromDomain(user), nil
}

func (s Service) GetById(id string) (userApp.ViewUserDto, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return user.ToViewUserDto(), err
	}

	return user.ToViewUserDto(), nil
}

func (s Service) Update(id string, dto userApp.UpdateUserDto) (userApp.ViewUserDto, error) {
	var view userApp.ViewUserDto

	if err := dto.Validate(); err != nil {
		return view, err
	}

	user, err := s.repo.FindById(id)
	if err != nil {
		return view, err
	}

	dto.ApplyNewValues(&user)

	if err := s.repo.Update(user); err != nil {
		return view, err
	}

	return view.FromDomain(user), err
}
