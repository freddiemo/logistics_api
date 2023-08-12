package repository

import (
	filter "github.com/ActiveChooN/gin-gorm-filter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/model"
)

type MaritimeShipmentRepository interface {
	Save(storage model.MaritimeShipment) (model.MaritimeShipment, error)
	FindAll(ctx *gin.Context) ([]model.MaritimeShipment, error)
	FindById(id int64) (model.MaritimeShipment, error)
	Update(model.MaritimeShipment) (model.MaritimeShipment, error)
	Delete(id int64) error
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

func (maritimeShipmentRepo *maritimeShipmentRepo) FindAll(ctx *gin.Context) ([]model.MaritimeShipment, error) {
	var maritimeShipments []model.MaritimeShipment
	result := maritimeShipmentRepo.db.Model(&model.MaritimeShipment{}).Scopes(
		filter.FilterByQuery(ctx, filter.ALL),
	).Find(&maritimeShipments)
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

func (maritimeShipmentRepo *maritimeShipmentRepo) Delete(id int64) error {
	var maritimeShipment model.MaritimeShipment
	result := maritimeShipmentRepo.db.Delete(&maritimeShipment, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
