package user

// Repository user domain database interface
type Repository interface {
	// Save register user into the database
	Save(user User) error

	// FindById load a user from database
	FindById(id string) (User, error)
}
