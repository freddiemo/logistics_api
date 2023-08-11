package model

import (
	"gorm.io/gorm"

	registerModels "github.com/freddiemo/logistics-api/internal/register/model"
)

type ProductType struct {
	gorm.Model
	Name               string                            `binding:"required" gorm:"type:varchar(32);not null;UNIQUE"`
	TransportationType registerModels.TransportationType `json:"transportation_type" binding:"required" gorm:"type:transportation_type;column:transportation_type"`
	Code               string                            `binding:"required" gorm:"not null"`
}
