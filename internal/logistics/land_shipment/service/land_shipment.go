package service

import (
	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/helpers"
	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/model"
	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/repository"

	storageServ "github.com/freddiemo/logistics-api/internal/register/storages/service"
)

type LandShipmentServiceInterface interface {
	Save(model.LandShipment) (model.LandShipment, error)
	FindAll(ctx *gin.Context) ([]model.LandShipment, error)
	FindById(id int64) (model.LandShipment, error)
	Update(model.LandShipment) (model.LandShipment, error)
	Delete(id int64) error
}

type landShipmentService struct {
	landShipmentRepository repository.LandShipmentRepository
	storageService         storageServ.StorageServiceInterface
}

func NewLandShipmentService(
	repository repository.LandShipmentRepository,
	storageService storageServ.StorageServiceInterface,
) LandShipmentServiceInterface {

	return &landShipmentService{
		landShipmentRepository: repository,
		storageService:         storageService,
	}
}

func (service *landShipmentService) Save(landShipment model.LandShipment) (model.LandShipment, error) {
	// validate only land storage type
	storage, err := service.storageService.FindById(int64(landShipment.StorageId))
	if err != nil {
		return model.LandShipment{}, err
	}
	if storage.TransportationType != 1 {
		return model.LandShipment{}, fmt.Errorf("err: Invalid Transportation Type")
	}

	// calculate discounts
	discount := helpers.GetDiscount(landShipment.ProductQuantity, landShipment.DeliveryPrice)
	if discount > 0 {
		landShipment.Discount = discount
	}

	// save
	landShipment, err = service.landShipmentRepository.Save(landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	return landShipment, nil
}

func (service *landShipmentService) FindAll(ctx *gin.Context) ([]model.LandShipment, error) {
	landShipments, err := service.landShipmentRepository.FindAll(ctx)
	if err != nil {
		return []model.LandShipment{}, err
	}

	return landShipments, nil
}

func (service *landShipmentService) FindById(id int64) (model.LandShipment, error) {
	landShipment, err := service.landShipmentRepository.FindById(id)
	if err != nil {
		return model.LandShipment{}, err
	}

	return landShipment, nil
}

func (service *landShipmentService) Update(landShipment model.LandShipment) (model.LandShipment, error) {
	// get landShipment from db
	landShipmentDb, err := service.landShipmentRepository.FindById(int64(landShipment.ID))
	if err != nil {
		return model.LandShipment{}, err
	}

	// validate only land storage type
	storage, err := service.storageService.FindById(int64(landShipment.StorageId))
	if err != nil {
		return model.LandShipment{}, err
	}
	if storage.TransportationType != 1 {
		return model.LandShipment{}, fmt.Errorf("err: Invalid Transportation Type")
	}

	// update discount
	if landShipment.DeliveryPrice != landShipmentDb.DeliveryPrice {
		discount := helpers.GetDiscount(landShipment.ProductQuantity, landShipment.DeliveryPrice)
		if discount > 0 {
			landShipment.Discount = discount
		}
	}

	// update
	landShipment, err = service.landShipmentRepository.Update(landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	return landShipment, nil
}

func (service *landShipmentService) Delete(id int64) error {
	if err := service.landShipmentRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
