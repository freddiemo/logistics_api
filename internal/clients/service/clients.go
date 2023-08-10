package service

import (
	"github.com/freddiemo/logistics-api/internal/clients/model"
	"github.com/freddiemo/logistics-api/internal/clients/repository"
)

type ClientServiceInterface interface {
	Save(model.Client) (model.Client, error)
	FindAll() ([]model.Client, error)
	FindById(id int64) (model.Client, error)
	Update(model.Client) (model.Client, error)
	Delete(id int64) error
}

type clientService struct {
	clientRepository repository.ClientRepository
}

func NewClientService(repository repository.ClientRepository) ClientServiceInterface {
	return &clientService{
		clientRepository: repository,
	}
}

func (service *clientService) Save(client model.Client) (model.Client, error) {
	client, err := service.clientRepository.Save(client)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (service *clientService) FindAll() ([]model.Client, error) {
	clients, err := service.clientRepository.FindAll()
	if err != nil {
		return []model.Client{}, err
	}

	return clients, nil
}

func (service *clientService) FindById(id int64) (model.Client, error) {
	client, err := service.clientRepository.FindById(id)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (service *clientService) Update(client model.Client) (model.Client, error) {
	client, err := service.clientRepository.Update(client)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (service *clientService) Delete(id int64) error {
	if err := service.clientRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
