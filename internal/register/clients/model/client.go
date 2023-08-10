package model

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name       string `binding:"required" gorm:"type:varchar(32);not null"`
	LastName   string `binding:"required" gorm:"type:varchar(32);not null"`
	Email      string `binding:"required" gorm:"type:varchar(32);not null;UNIQUE" validate:"is-email"`
	DocumentID int    `binding:"required" gorm:"not null;UNIQUE"`
	Phone      int    `binding:"required" gorm:"not null"`
}
