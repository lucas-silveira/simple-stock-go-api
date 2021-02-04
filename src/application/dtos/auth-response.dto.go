package dtos

// AuthResponseDto ...
type AuthResponseDto struct {
	UserId int    `json:"user_id"`
	token  string `json:"token"`
}

// NewAuthResponseDto is a factory function that generate a new AuthResponseDto
func NewAuthResponseDto(userID int, token string) *AuthResponseDto {
	return &AuthResponseDto{
		UserId: userID,
		token:  token,
	}
}
