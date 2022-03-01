package user

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

var (
	errorNameRequired     = errors.New("name is required")
	errorEmailRequired    = errors.New("e-mail is required")
	errorPasswordRequired = errors.New("password is required")
)

// User domain
type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

// CreateUserDto create user request data
type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ViewUserDto user data to send to clients
type ViewUserDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UpdateUserDto update user request data
type UpdateUserDto struct {
	Name string `json:"name"`
}

// Validate check create user data request
func (dto CreateUserDto) Validate() error {
	if strings.TrimSpace(dto.Name) == "" {
		return errorNameRequired
	}

	if strings.TrimSpace(dto.Email) == "" {
		return errorEmailRequired
	}

	if strings.TrimSpace(dto.Password) == "" {
		return errorPasswordRequired
	}

	return nil
}

// Validate check update user data request
func (dto UpdateUserDto) Validate() error {
	if strings.TrimSpace(dto.Name) == "" {
		return errorNameRequired
	}

	return nil
}

// ToDomain converts CreateUserDto into User domain
func (dto CreateUserDto) ToDomain() User {
	return User{
		ID:       uuid.NewString(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

// ApplyNewValues populate User domain with UpdateUserDto values
func (dto UpdateUserDto) ApplyNewValues(loadedUser *User) {
	loadedUser.Name = dto.Name
}

// FromDomain build ViewUserDto from User domain data
func (dto ViewUserDto) FromDomain(user User) ViewUserDto {
	return ViewUserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

// ToViewUserDto converts User domain into ViewUserDto
func (u User) ToViewUserDto() ViewUserDto {
	return ViewUserDto{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
