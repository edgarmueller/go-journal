package dto

import (
	"github.com/edgarmueller/go-api-journal/internal/domain"
)

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:       user.UUID.String(),
		Username: user.Username,
		Email:    user.Email,
	}
}
