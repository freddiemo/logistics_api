package service

import (
	"github.com/freddiemo/logistics-api/internal/register/product_types/model"
	"github.com/freddiemo/logistics-api/internal/register/product_types/repository"
)

type ProductTypeServiceInterface interface {
	Save(model.ProductType) (model.ProductType, error)
	FindAll() ([]model.ProductType, error)
	FindById(id int64) (model.ProductType, error)
	Update(model.ProductType) (model.ProductType, error)
	Delete(id int64) error
}

type productTypeService struct {
	productTypeRepository repository.ProductTypeRepository
}

func NewProductTypeService(repository repository.ProductTypeRepository) ProductTypeServiceInterface {
	return &productTypeService{
		productTypeRepository: repository,
	}
}

func (service *productTypeService) Save(productType model.ProductType) (model.ProductType, error) {
	productType, err := service.productTypeRepository.Save(productType)
	if err != nil {
		return model.ProductType{}, err
	}

	return productType, nil
}

func (service *productTypeService) FindAll() ([]model.ProductType, error) {
	productTypes, err := service.productTypeRepository.FindAll()
	if err != nil {
		return []model.ProductType{}, err
	}

	return productTypes, nil
}

func (service *productTypeService) FindById(id int64) (model.ProductType, error) {
	productType, err := service.productTypeRepository.FindById(id)
	if err != nil {
		return model.ProductType{}, err
	}

	return productType, nil
}

func (service *productTypeService) Update(productType model.ProductType) (model.ProductType, error) {
	productType, err := service.productTypeRepository.Update(productType)
	if err != nil {
		return model.ProductType{}, err
	}

	return productType, nil
}

func (service *productTypeService) Delete(id int64) error {
	if err := service.productTypeRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
