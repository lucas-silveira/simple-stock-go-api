package user

// IRepository is a interface for user repository
type IRepository interface {
	save(User) error
}
