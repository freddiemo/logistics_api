package service

import (
	"github.com/freddiemo/logistics-api/internal/register/storages/model"
	"github.com/freddiemo/logistics-api/internal/register/storages/repository"
)

type StorageServiceInterface interface {
	Save(model.Storage) (model.Storage, error)
	FindAll() ([]model.Storage, error)
	FindById(id int64) (model.Storage, error)
}

type storageService struct {
	storageRepository repository.StorageRepository
}

func NewStorageService(repository repository.StorageRepository) StorageServiceInterface {
	return &storageService{
		storageRepository: repository,
	}
}

func (service *storageService) Save(storage model.Storage) (model.Storage, error) {
	storage, err := service.storageRepository.Save(storage)
	if err != nil {
		return model.Storage{}, err
	}

	return storage, nil
}

func (service *storageService) FindAll() ([]model.Storage, error) {
	storages, err := service.storageRepository.FindAll()
	if err != nil {
		return []model.Storage{}, err
	}

	return storages, nil
}

func (service *storageService) FindById(id int64) (model.Storage, error) {
	storage, err := service.storageRepository.FindById(id)
	if err != nil {
		return model.Storage{}, err
	}

	return storage, nil
}
