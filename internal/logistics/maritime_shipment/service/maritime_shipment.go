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
	FindAll() ([]model.MaritimeShipment, error)
	FindById(id int64) (model.MaritimeShipment, error)
	Update(model.MaritimeShipment) (model.MaritimeShipment, error)
	Delete(id int64) error
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

func (service *maritimeShipmentService) FindAll() ([]model.MaritimeShipment, error) {
	maritimeShipments, err := service.maritimeShipmentRepository.FindAll()
	if err != nil {
		return []model.MaritimeShipment{}, err
	}

	return maritimeShipments, nil
}

func (service *maritimeShipmentService) FindById(id int64) (model.MaritimeShipment, error) {
	maritimeShipment, err := service.maritimeShipmentRepository.FindById(id)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	return maritimeShipment, nil
}

func (service *maritimeShipmentService) Update(maritimeShipment model.MaritimeShipment) (model.MaritimeShipment, error) {
	// get maritimeShipment from db
	maritimeShipmentDb, err := service.maritimeShipmentRepository.FindById(int64(maritimeShipment.ID))
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	// validate only maritime storage type
	storage, err := service.storageService.FindById(int64(maritimeShipment.StorageId))
	if err != nil {
		return model.MaritimeShipment{}, err
	}
	if storage.TransportationType != 2 {
		return model.MaritimeShipment{}, fmt.Errorf("err: Invalid Transportation Type")
	}

	// update discount
	if maritimeShipment.DeliveryPrice != maritimeShipmentDb.DeliveryPrice {
		maritimeShipment.Discount = helpers.GetDiscount(maritimeShipment.ProductQuantity, maritimeShipment.DeliveryPrice)
	}

	// update
	maritimeShipment, err = service.maritimeShipmentRepository.Update(maritimeShipment)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	return maritimeShipment, nil
}

func (service *maritimeShipmentService) Delete(id int64) error {
	if err := service.maritimeShipmentRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
