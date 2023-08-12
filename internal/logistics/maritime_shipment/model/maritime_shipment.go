package model

import (
	"time"

	"gorm.io/gorm"
)

type MaritimeShipment struct {
	gorm.Model
	RegisterDate    time.Time `json:"register_date" gorm:"default:CURRENT_TIMESTAMP"`
	DeliveryDate    time.Time `json:"delivery_date" binding:"required"`
	DeliveryPrice   float64   `json:"delivery_price" binding:"required" gorm:"not null;column:delivery_price" filter:"param:delivery_price;filterable"`
	Discount        float64   `json:"discount" gorm:"column:discount" filter:"filterable"`
	ProductQuantity int       `json:"product_quantity" binding:"required" gorm:"not null"`
	FleetNumber     string    `json:"fleet_number" binding:"required" gorm:"not null;column:fleet_number" validate:"is-fleet-number" filter:"param:fleet_number;searchable;filterable"`
	GuideNumber     string    `json:"guide_number" binding:"required" gorm:"not null;column:guide_number" validate:"is-guide-number" filter:"param:guide_number;searchable;filterable"`
	ClientId        int       `json:"client_id" binding:"required" gorm:"not null;column:client_id" filter:"param:client_id;filterable"`
	ProductTypeId   int       `json:"product_type_id" binding:"required" gorm:"not null;column:product_type_id" filter:"param:product_type_id;filterable"`
	StorageId       int       `json:"storage_id" binding:"required" gorm:"not null;column:storage_id" filter:"param:storage_id;filterable"`
}
