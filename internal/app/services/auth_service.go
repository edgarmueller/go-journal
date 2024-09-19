package services

import (
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
}

func NewAuthService() AuthService {
	return AuthService{}
}

func (a *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (a *AuthService) CheckPassword(providedPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
}
