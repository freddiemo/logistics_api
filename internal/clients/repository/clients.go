package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/clients/model"
)

type ClientRepository interface {
	Save(client model.Client) (model.Client, error)
	FindAll() ([]model.Client, error)
}

type clientRepo struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepo{
		db: db,
	}
}

func (clientRepo *clientRepo) Save(client model.Client) (model.Client, error) {
	result := clientRepo.db.Save(&client)

	if result.Error != nil {
		return model.Client{}, result.Error
	}

	return client, nil
}

func (clientRepo *clientRepo) FindAll() ([]model.Client, error) {
	var clients []model.Client
	result := clientRepo.db.Find(&clients)
	if result.Error != nil {
		return clients, result.Error
	}

	return clients, nil
}
