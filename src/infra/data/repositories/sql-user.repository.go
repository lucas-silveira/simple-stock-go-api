package repositories

import (
	"database/sql"
	"fmt"
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
	_, err := userRepo.db.Query("INSERT INTO users (name, email, password) VALUES($1, $2, $3)", user.Name, user.Email, user.Password)

	if err != nil {
		return fmt.Errorf("There was an error persisting the user in the database, err: %s", err)
	}

	return nil
}
