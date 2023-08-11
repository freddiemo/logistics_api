package model

import (
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Code      string `json:"code"`
	Name      string `json:"name"`
	PhoneCode int    `json:"phone_code"`
}
