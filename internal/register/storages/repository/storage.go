package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/register/storages/model"
)

type StorageRepository interface {
	Save(storage model.Storage) (model.Storage, error)
	FindAll() ([]model.Storage, error)
	FindById(id int64) (model.Storage, error)
	Update(model.Storage) (model.Storage, error)
	Delete(id int64) error
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

func (storageRepo *storageRepo) FindById(id int64) (model.Storage, error) {
	var storage model.Storage
	result := storageRepo.db.First(&storage, id)
	if result.Error != nil {
		return storage, result.Error
	}

	return storage, nil
}

func (storageRepo *storageRepo) Update(storage model.Storage) (model.Storage, error) {
	result := storageRepo.db.Where("id = ?", storage.ID).Updates(storage)
	if result.Error != nil {
		return storage, result.Error
	}
	if result.RowsAffected == 0 {
		return storage, gorm.ErrRecordNotFound
	}

	return storage, nil
}

func (storageRepo *storageRepo) Delete(id int64) error {
	var storage model.Storage
	result := storageRepo.db.Delete(&storage, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
