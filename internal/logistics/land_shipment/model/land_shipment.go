package model

import (
	"time"

	"gorm.io/gorm"
)

type LandShipment struct {
	gorm.Model
	RegisterDate    time.Time `json:"register_date" gorm:"default:CURRENT_TIMESTAMP"`
	DeliveryDate    time.Time `json:"delivery_date" binding:"required"`
	DeliveryPrice   float64   `json:"delivery_price" binding:"required"`
	ProductQuantity int       `json:"product_quantity" binding:"required" gorm:"not null"`
	VehiclePlate    string    `json:"vehicle_plate" binding:"required" gorm:"not null" validate:"is-vehicle-plate"`
	GuideNumber     string    `json:"guide_number" binding:"required" gorm:"not null" validate:"is-guide-number"`
	ClientId        int       `json:"client_id" binding:"required" gorm:"not null"`
	ProductTypeId   int       `json:"product_type_id" binding:"required" gorm:"not null"`
	StorageId       int       `json:"storage_id" binding:"required" gorm:"not null"`
}
