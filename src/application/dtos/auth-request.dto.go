package dtos

// AuthRequestDto represents the user data request
type AuthRequestDto struct {
	Email    string `json:Email"`
	Password string `json:Password"`
}
