package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

type RegisterUser struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
