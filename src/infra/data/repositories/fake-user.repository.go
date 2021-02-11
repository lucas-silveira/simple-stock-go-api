package repositories

import (
	"main/src/domain/user"
)

// FakeUser is a fake user repository
type FakeUser struct {
	data map[int]user.User
}

// NewFakeUser is a factory function
func NewFakeUser() *FakeUser {
	data := make(map[int]user.User)
	return &FakeUser{data}
}

// FindByEmail find a user object from database by email
func (userRepo FakeUser) FindByEmail(email string) (*user.User, error) {
	var user user.User

	for _, anUser := range userRepo.data {
		if anUser.Email == email {
			return &anUser, nil
		}
	}

	return &user, nil
}

// Save user object into database
func (userRepo FakeUser) Save(user user.User) error {
	user.ID = len(userRepo.data) + 1
	userRepo.data[user.ID] = user

	return nil
}
