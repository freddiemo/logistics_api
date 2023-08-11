package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/register/storages/model"
)

type StorageRepository interface {
	Save(storage model.Storage) (model.Storage, error)
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
