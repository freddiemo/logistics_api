package service

import (
	"fmt"

	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/helpers"
	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/model"
	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/repository"

	storageServ "github.com/freddiemo/logistics-api/internal/register/storages/service"
)

type MaritimeShipmentServiceInterface interface {
	Save(model.MaritimeShipment) (model.MaritimeShipment, error)
}

type maritimeShipmentService struct {
	maritimeShipmentRepository repository.MaritimeShipmentRepository
	storageService             storageServ.StorageServiceInterface
}

func NewMaritimeShipmentService(
	repository repository.MaritimeShipmentRepository,
	storageService storageServ.StorageServiceInterface,
) MaritimeShipmentServiceInterface {
	return &maritimeShipmentService{
		maritimeShipmentRepository: repository,
		storageService:             storageService,
	}
}

func (service *maritimeShipmentService) Save(maritimeShipment model.MaritimeShipment) (model.MaritimeShipment, error) {
	// validate only maritime storage type
	storage, err := service.storageService.FindById(int64(maritimeShipment.StorageId))
	if err != nil {
		return model.MaritimeShipment{}, err
	}
	if storage.TransportationType != 2 {
		return model.MaritimeShipment{}, fmt.Errorf("err: Invalid Transportation Type")
	}

	// get discount
	if maritimeShipment.ProductQuantity > 10 {
		maritimeShipment.Discount -= maritimeShipment.DeliveryPrice * 0.03
	}
	discount := helpers.GetDiscount(maritimeShipment.ProductQuantity, maritimeShipment.DeliveryPrice)
	if discount > 0 {
		maritimeShipment.Discount = discount
	}

	// save
	maritimeShipment, err = service.maritimeShipmentRepository.Save(maritimeShipment)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	return maritimeShipment, nil
}
