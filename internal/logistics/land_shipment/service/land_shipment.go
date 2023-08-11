package service

import (
	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/model"
	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/repository"
)

type LandShipmentServiceInterface interface {
	Save(model.LandShipment) (model.LandShipment, error)
}

type landShipmentService struct {
	landShipmentRepository repository.LandShipmentRepository
}

func NewLandShipmentService(repository repository.LandShipmentRepository) LandShipmentServiceInterface {
	return &landShipmentService{
		landShipmentRepository: repository,
	}
}

func (service *landShipmentService) Save(landShipment model.LandShipment) (model.LandShipment, error) {
	landShipment, err := service.landShipmentRepository.Save(landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	return landShipment, nil
}
