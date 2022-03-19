package domains

import "github.com/google/uuid"

// User domain
type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) User {
	return User{
		ID:       uuid.NewString(),
		Name:     name,
		Email:    email,
		Password: password,
	}
}