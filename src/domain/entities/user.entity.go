package entities

// User is a entity that represents a user in ubiquitous language
type User struct {
	Entity
	Name, Email, Password string
}

// NewUser is a factory function to create a new user struct
func NewUser(name, email, password string) User {
	user := User{}
	user.SetName(name)
	user.SetEmail(email)
	user.SetPassword(password)

	return user
}

// SetName set name property
func (user *User) SetName(name string) {
	user.Name = name
}

// SetEmail set email property
func (user *User) SetEmail(email string) {
	user.Email = email
}

// SetPassword set password property
func (user *User) SetPassword(password string) {
	user.Password = password
}
