package usecases

import (
	"github.com/edgarmueller/go-api-journal/internal/app/services"
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/edgarmueller/go-api-journal/internal/domain/repositories"
)

type AuthUseCases struct {
	auth           services.AuthService
	userRepository repositories.UserRepository
}

func NewAuthUseCases(userRepository repositories.UserRepository, auth services.AuthService) *AuthUseCases {
	return &AuthUseCases{auth: auth, userRepository: userRepository}
}

func (a *AuthUseCases) RegisterUser(registerUser domain.RegisterUser) (domain.User, error) {
	pw, err := a.auth.HashPassword(registerUser.Password)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		Email:    registerUser.Email,
		Username: registerUser.Username,
		Password: pw,
	}
	err = a.userRepository.SaveUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
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
	return services.GenerateJWT(user.Email, user.Username)
}
