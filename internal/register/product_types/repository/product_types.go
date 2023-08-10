package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/register/product_types/model"
)

type ProductTypeRepository interface {
	Save(productType model.ProductType) (model.ProductType, error)
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
