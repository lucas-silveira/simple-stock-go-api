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

// FindByEmail find a user object from database by email
func (userRepo User) FindByEmail(email string) (*user.User, error) {
	var user user.User
	selectError :=
		userRepo.db.QueryRow(
			"SELECT id, name, email, password FROM users WHERE email = $1",
			email,
		).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if selectError != nil {
		if selectError != sql.ErrNoRows {
			return nil, fmt.Errorf("There was an error to select user from id, err: %s", selectError)
		}
	}

	return &user, nil
}

// Save user object into database
func (userRepo User) Save(user user.User) error {
	var id string
	selectError := userRepo.db.QueryRow("SELECT id FROM users WHERE id = $1", user.ID).Scan(&id)

	if selectError != nil && id == "" {
		if selectError != sql.ErrNoRows {
			return fmt.Errorf("There was an error to select user from id, err: %s", selectError)
		}

		_, insertError := userRepo.db.Query("INSERT INTO users (name, email, password) VALUES($1, $2, $3)", user.Name, user.Email, user.Password)

		if insertError != nil {
			return fmt.Errorf("There was an error creating the user in the database, err: %s", insertError)
		}
	}

	_, updateError := userRepo.db.Query("UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4", user.Name, user.Email, user.Password, user.ID)

	if updateError != nil {
		return fmt.Errorf("There was an error updating the user in the database, err: %s", updateError)
	}

	return nil
}
