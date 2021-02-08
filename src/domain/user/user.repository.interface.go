package user

// IRepository is a interface for user repository
type IRepository interface {
	Save(User) error
}
