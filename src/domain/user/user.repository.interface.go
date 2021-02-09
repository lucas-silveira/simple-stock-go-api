package user

// IRepository is a interface for user repository
type IRepository interface {
	FindByEmail(email string) (*User, error)
	Save(User) error
}
