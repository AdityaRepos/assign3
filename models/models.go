package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Branch    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Register struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Branch   string `json:"branch" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
}

type Question struct {
	ID     uint   `gorm:"primarykey"`
	UserID string `gorm:"not null"`
	Text   string `gorm:"not null"`
}
