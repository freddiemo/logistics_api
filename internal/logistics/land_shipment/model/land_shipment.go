package model

import (
	"time"

	"gorm.io/gorm"
)

type LandShipment struct {
	gorm.Model
	RegisterDate    time.Time `json:"register_date" gorm:"default:CURRENT_TIMESTAMP"`
	DeliveryDate    time.Time `json:"delivery_date" binding:"required" gorm:"not null"`
	DeliveryPrice   float64   `json:"delivery_price" binding:"required" gorm:"not null;column:delivery_price" filter:"param:delivery_price;filterable"`
	Discount        float64   `json:"discount" gorm:"column:discount" filter:"filterable"`
	ProductQuantity int       `json:"product_quantity" binding:"required" gorm:"not null" filter:"param:product_quantity;filterable"`
	VehiclePlate    string    `json:"vehicle_plate" binding:"required" gorm:"not null;column:vehicle_plate" validate:"is-vehicle-plate" filter:"param:vehicle_plate;searchable;filterable"`
	GuideNumber     string    `json:"guide_number" binding:"required" gorm:"not null;column:guide_number" validate:"is-guide-number" filter:"param:guide_number;searchable;filterable"`
	ClientId        int       `json:"client_id" binding:"required" gorm:"not null;column:client_id" filter:"param:client_id;filterable"`
	ProductTypeId   int       `json:"product_type_id" binding:"required" gorm:"not null;column:product_type_id" filter:"param:product_type_id;filterable"`
	StorageId       int       `json:"storage_id" binding:"required" gorm:"not null;column:storage_id" filter:"param:storage_id;filterable"`
}
