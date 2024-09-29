package usecases

import (
	"github.com/edgarmueller/go-api-journal/internal/app/services"
	"github.com/edgarmueller/go-api-journal/internal/domain/repositories"
)

type AuthUseCases struct {
	auth           services.AuthService
	userRepository repositories.UserRepository
}

func NewAuthUseCases(userRepository repositories.UserRepository, auth services.AuthService) *AuthUseCases {
	return &AuthUseCases{auth: auth, userRepository: userRepository}
}

func (t *AuthUseCases) GenerateToken(email, providedPassword string) (string, error) {
	user, err := t.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	credentialError := t.auth.CheckPassword(providedPassword, user.Password)
	if credentialError != nil {
		return "", credentialError
	}
	return services.GenerateJWT(user.UUID, user.Email, user.Username)
}
