package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/model"
)

type MaritimeShipmentRepository interface {
	Save(storage model.MaritimeShipment) (model.MaritimeShipment, error)
	FindAll() ([]model.MaritimeShipment, error)
	FindById(id int64) (model.MaritimeShipment, error)
	Update(model.MaritimeShipment) (model.MaritimeShipment, error)
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

func (maritimeShipmentRepo *maritimeShipmentRepo) FindAll() ([]model.MaritimeShipment, error) {
	var maritimeShipments []model.MaritimeShipment
	result := maritimeShipmentRepo.db.Find(&maritimeShipments)
	if result.Error != nil {
		return maritimeShipments, result.Error
	}

	return maritimeShipments, nil
}

func (maritimeShipmentRepo *maritimeShipmentRepo) FindById(id int64) (model.MaritimeShipment, error) {
	var maritimeShipment model.MaritimeShipment
	result := maritimeShipmentRepo.db.First(&maritimeShipment, id)
	if result.Error != nil {
		return maritimeShipment, result.Error
	}

	return maritimeShipment, nil
}

func (maritimeShipmentRepo *maritimeShipmentRepo) Update(maritimeShipment model.MaritimeShipment) (model.MaritimeShipment, error) {
	result := maritimeShipmentRepo.db.Where("id = ?", maritimeShipment.ID).Updates(maritimeShipment)
	if result.Error != nil {
		return maritimeShipment, result.Error
	}
	if result.RowsAffected == 0 {
		return maritimeShipment, gorm.ErrRecordNotFound
	}

	return maritimeShipment, nil
}
