package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint      `gorm:"unique;primaryKey;autoIncrement"`
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `json:"username" gorm:"unique"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
}

type RegisterUser struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func CreateUser(username, email, password string) User {
	return User{
		UUID:     uuid.New(),
		Username: username,
		Email:    email,
		Password: password,
	}
}
