package model

import (
	"time"

	"gorm.io/gorm"
)

type MaritimeShipment struct {
	gorm.Model
	RegisterDate    time.Time `json:"register_date" gorm:"default:CURRENT_TIMESTAMP"`
	DeliveryDate    time.Time `json:"delivery_date" binding:"required"`
	DeliveryPrice   float64   `json:"delivery_price" binding:"required"`
	Discount        float64   `json:"discount"`
	ProductQuantity int       `json:"product_quantity" binding:"required" gorm:"not null"`
	FleetNumber     string    `json:"fleet_number" binding:"required" gorm:"not null" validate:"is-fleet-number"`
	GuideNumber     string    `json:"guide_number" binding:"required" gorm:"not null; uniqueIndex" validate:"is-guide-number" filter:"searchable;filterable"`
	ClientId        int       `json:"client_id" binding:"required" gorm:"not null"`
	ProductTypeId   int       `json:"product_type_id" binding:"required" gorm:"not null"`
	StorageId       int       `json:"storage_id" binding:"required" gorm:"not null"`
}
