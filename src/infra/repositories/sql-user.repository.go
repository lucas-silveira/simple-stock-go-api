package repositories

import (
	"database/sql"
	"errors"
	"main/src/domain/user"
)

// User is a user repository
type User struct {
	db *sql.DB
}

// NewUserRepository generate a new repository for user
func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

// Save persist user object into database
func (userRepo User) Save(user user.User) error {
	_, err := userRepo.db.Exec("INSERT INTO users (name, email, password) VALUES(?,?,?)", user.Name, user.Email, user.Password)

	if err != nil {
		return errors.New("Houve um erro ao persistir o usuário no banco")
	}

	return nil
}
