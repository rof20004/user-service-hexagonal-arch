package user

// Usecase user domain business functions
type Usecase interface {
	// Create register user
	Create(dto CreateUserDto) (ViewUserDto, error)

	// GetById get user by id
	GetById(id string) (ViewUserDto, error)

	// Update update some user data
	Update(id string, dto UpdateUserDto) (ViewUserDto, error)
}
