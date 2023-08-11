package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/register/storages/model"
)

type StorageRepository interface {
	Save(storage model.Storage) (model.Storage, error)
	FindAll() ([]model.Storage, error)
}

type storageRepo struct {
	db *gorm.DB
}

func NewStorageRepository(db *gorm.DB) StorageRepository {
	return &storageRepo{
		db: db,
	}
}

func (storageRepo *storageRepo) Save(storage model.Storage) (model.Storage, error) {
	result := storageRepo.db.Save(&storage)

	if result.Error != nil {
		return model.Storage{}, result.Error
	}

	return storage, nil
}

func (storageRepo *storageRepo) FindAll() ([]model.Storage, error) {
	var storages []model.Storage
	result := storageRepo.db.Find(&storages)
	if result.Error != nil {
		return storages, result.Error
	}

	return storages, nil
}
