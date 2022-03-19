package ports

import "user-service/application/domains"

type Repository interface {
	Save(u domains.User) error
	FindById(id string) (domains.User, error)
}

type Usecase interface {
	Create(name, email, password string) (domains.User, error)
	GetById(id string) (domains.User, error)
}
