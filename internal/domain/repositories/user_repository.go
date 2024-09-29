package repositories

import (
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	SaveUser(user domain.User) error
	GetUserByEmail(email string) (domain.User, error)
	GetUserByUUID(uuid uuid.UUID) (domain.User, error)
}
