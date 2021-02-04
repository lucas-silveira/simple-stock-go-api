package dtos

// AuthCredentialsDto represents the user data request
type AuthCredentialsDto struct {
	Email    string `json:email"`
	Password string `json:password"`
}
