package repository

import (
	filter "github.com/ActiveChooN/gin-gorm-filter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/model"
)

type LandShipmentRepository interface {
	Save(storage model.LandShipment) (model.LandShipment, error)
	FindAll(ctx *gin.Context) ([]model.LandShipment, error)
	FindById(id int64) (model.LandShipment, error)
	Update(model.LandShipment) (model.LandShipment, error)
	Delete(id int64) error
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

func (landShipmentRepo *landShipmentRepo) FindAll(ctx *gin.Context) ([]model.LandShipment, error) {
	var landShipments []model.LandShipment
	result := landShipmentRepo.db.Scopes(
		filter.FilterByQuery(ctx, filter.ALL),
	).Find(&landShipments)
	if result.Error != nil {
		return landShipments, result.Error
	}

	return landShipments, nil
}

func (landShipmentRepo *landShipmentRepo) FindById(id int64) (model.LandShipment, error) {
	var landShipment model.LandShipment
	result := landShipmentRepo.db.First(&landShipment, id)
	if result.Error != nil {
		return landShipment, result.Error
	}

	return landShipment, nil
}

func (landShipmentRepo *landShipmentRepo) Update(landShipment model.LandShipment) (model.LandShipment, error) {
	result := landShipmentRepo.db.Where("id = ?", landShipment.ID).Updates(landShipment)
	if result.Error != nil {
		return landShipment, result.Error
	}
	if result.RowsAffected == 0 {
		return landShipment, gorm.ErrRecordNotFound
	}

	return landShipment, nil
}

func (landShipmentRepo *landShipmentRepo) Delete(id int64) error {
	var landShipment model.LandShipment
	result := landShipmentRepo.db.Delete(&landShipment, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
