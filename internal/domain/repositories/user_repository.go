package repositories

import "github.com/edgarmueller/go-api-journal/internal/domain"

type UserRepository interface {
	SaveUser(user domain.User) error
	GetUserByEmail(email string) (domain.User, error)
}
