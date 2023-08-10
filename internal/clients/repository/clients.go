package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/logistics-api/internal/clients/model"
)

type ClientRepository interface {
	Save(client model.Client) (model.Client, error)
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
