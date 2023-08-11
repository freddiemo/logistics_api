package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/model"
)

type MaritimeShipmentRepository interface {
	Save(storage model.MaritimeShipment) (model.MaritimeShipment, error)
}

type maritimeShipmentRepo struct {
	db *gorm.DB
}

func NewMaritimeShipmentRepository(db *gorm.DB) MaritimeShipmentRepository {
	return &maritimeShipmentRepo{
		db: db,
	}
}

func (maritimeShipmentRepo *maritimeShipmentRepo) Save(maritimeShipment model.MaritimeShipment) (model.MaritimeShipment, error) {
	result := maritimeShipmentRepo.db.Save(&maritimeShipment)

	if result.Error != nil {
		return model.MaritimeShipment{}, result.Error
	}

	return maritimeShipment, nil
}
