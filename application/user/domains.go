package user

import (
	"errors"
	"github.com/google/uuid"
	"strings"
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

// SetValuesFromCreateUserDto converts CreateUserDto into User domain
func (u *User) SetValuesFromCreateUserDto(dto CreateUserDto) {
	u.ID = uuid.NewString()
	u.Name = dto.Name
	u.Email = dto.Email
	u.Password = dto.Password
}

// SetValuesFromUpdateUserDto populate User domain from UpdateUserDto values
func (u *User) SetValuesFromUpdateUserDto(dto UpdateUserDto) {
	u.Name = dto.Name
}

// ToViewUserDto build ViewUserDto from User domain data
func (u *User) ToViewUserDto() ViewUserDto {
	return ViewUserDto{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
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
