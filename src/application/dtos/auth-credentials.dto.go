package dtos

// AuthCredentialsDto represents the user data request
type AuthCredentialsDto struct {
	Email    string `json:Email"`
	Password string `json:Password"`
}
