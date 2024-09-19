package database

import (
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/edgarmueller/go-api-journal/internal/domain/repositories"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GormUserRepository{db: db}
}

func (g *GormUserRepository) SaveUser(user domain.User) error {
	if err := g.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormUserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := g.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
