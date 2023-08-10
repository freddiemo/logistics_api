package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/register/product_types/model"
)

type ProductTypeRepository interface {
	Save(productType model.ProductType) (model.ProductType, error)
	FindAll() ([]model.ProductType, error)
	FindById(id int64) (model.ProductType, error)
	Update(model.ProductType) (model.ProductType, error)
	Delete(id int64) error
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

func (productTypeRepo *productTypeRepo) Update(productType model.ProductType) (model.ProductType, error) {
	result := productTypeRepo.db.Where("id = ?", productType.ID).Updates(productType)
	if result.Error != nil {
		return productType, result.Error
	}
	if result.RowsAffected == 0 {
		return productType, gorm.ErrRecordNotFound
	}

	return productType, nil
}

func (productTypeRepo *productTypeRepo) Delete(id int64) error {
	var productType model.ProductType
	result := productTypeRepo.db.Delete(&productType, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
