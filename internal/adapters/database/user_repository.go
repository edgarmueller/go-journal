package database

import (
	"github.com/asaskevich/EventBus"
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/edgarmueller/go-api-journal/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db       *gorm.DB
	eventBus EventBus.Bus
}

func NewGormUserRepository(db *gorm.DB, eventBus EventBus.Bus) repositories.UserRepository {
	return &GormUserRepository{db: db, eventBus: eventBus}
}

func (g *GormUserRepository) SaveUser(user domain.User) error {
	if err := g.db.Create(&user).Error; err != nil {
		return err
	}
	g.eventBus.Publish("user.created", user)
	return nil
}

func (g *GormUserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := g.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (g *GormUserRepository) GetUserByUUID(userId uuid.UUID) (domain.User, error) {
	var user domain.User
	if err := g.db.Where("uuid = ?", userId.String()).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
