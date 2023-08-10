package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/register/product_types/model"
)

type ProductTypeRepository interface {
	Save(productType model.ProductType) (model.ProductType, error)
	FindAll() ([]model.ProductType, error)
	FindById(id int64) (model.ProductType, error)
}

type productTypeRepo struct {
	db *gorm.DB
}

func NewProductTypeRepository(db *gorm.DB) ProductTypeRepository {
	return &productTypeRepo{
		db: db,
	}
}

func (productTypeRepo *productTypeRepo) Save(productType model.ProductType) (model.ProductType, error) {
	result := productTypeRepo.db.Save(&productType)

	if result.Error != nil {
		return model.ProductType{}, result.Error
	}

	return productType, nil
}

func (productTypeRepo *productTypeRepo) FindAll() ([]model.ProductType, error) {
	var productTypes []model.ProductType
	result := productTypeRepo.db.Find(&productTypes)
	if result.Error != nil {
		return productTypes, result.Error
	}

	return productTypes, nil
}

func (productTypeRepo *productTypeRepo) FindById(id int64) (model.ProductType, error) {
	var productType model.ProductType
	result := productTypeRepo.db.First(&productType, id)
	if result.Error != nil {
		return productType, result.Error
	}

	return productType, nil
}
