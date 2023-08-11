package model

import (
	"gorm.io/gorm"

	registerModels "github.com/freddiemo/logistics-api/internal/register/model"
)

type Storage struct {
	gorm.Model
	CityId             int                               `json:"city_id" binding:"required" gorm:"not null"`
	Address            string                            `binding:"required" gorm:"type:varchar(50);not null"`
	TransportationType registerModels.TransportationType `json:"transportation_type" binding:"required" gorm:"type:transportation_type;column:transportation_type"`
	Code               string                            `binding:"required" gorm:"not null"`
}
