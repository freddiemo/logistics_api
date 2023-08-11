package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/model"
)

type LandShipmentRepository interface {
	Save(storage model.LandShipment) (model.LandShipment, error)
	FindAll() ([]model.LandShipment, error)
}

type landShipmentRepo struct {
	db *gorm.DB
}

func NewLandShipmentRepository(db *gorm.DB) LandShipmentRepository {
	return &landShipmentRepo{
		db: db,
	}
}

func (landShipmentRepo *landShipmentRepo) Save(landShipment model.LandShipment) (model.LandShipment, error) {
	result := landShipmentRepo.db.Save(&landShipment)

	if result.Error != nil {
		return model.LandShipment{}, result.Error
	}

	return landShipment, nil
}

func (landShipmentRepo *landShipmentRepo) FindAll() ([]model.LandShipment, error) {
	var landShipments []model.LandShipment
	result := landShipmentRepo.db.Find(&landShipments)
	if result.Error != nil {
		return landShipments, result.Error
	}

	return landShipments, nil
}
