package usecases

import (
	"github.com/edgarmueller/go-api-journal/internal/app/services"
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/edgarmueller/go-api-journal/internal/domain/repositories"
)

type UserUseCases struct {
	auth           services.AuthService
	userRepository repositories.UserRepository
}

func NewUserUseCases(userRepository repositories.UserRepository, auth services.AuthService) *UserUseCases {
	return &UserUseCases{auth: auth, userRepository: userRepository}
}

func (a *UserUseCases) RegisterUser(registerUser domain.RegisterUser) (domain.User, error) {
	pw, err := a.auth.HashPassword(registerUser.Password)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.CreateUser(registerUser.Username, registerUser.Email, pw)
	err = a.userRepository.SaveUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
