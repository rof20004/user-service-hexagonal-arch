package user

// Usecase user domain business functionalities
type Usecase interface {
	// Create register user
	Create(dto CreateUserDto) (ViewUserDto, error)

	// GetById get user by id
	GetById(id string) (ViewUserDto, error)
}
