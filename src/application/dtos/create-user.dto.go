package dtos

// CreateUserDto represents the user create request
type CreateUserDto struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
