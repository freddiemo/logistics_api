package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/freddiemo/logistics-api/internal/register/storages/model"
	"github.com/freddiemo/logistics-api/internal/register/storages/service"
)

type StorageController interface {
	Save(ctx *gin.Context) (model.Storage, error)
	FindAll(ctx *gin.Context) ([]model.Storage, error)
	FindById(ctx *gin.Context) (model.Storage, error)
	Update(ctx *gin.Context) (model.Storage, error)
	Delete(ctx *gin.Context) error
}

type storageController struct {
	storageService service.StorageServiceInterface
}

func NewStorageController(storageService service.StorageServiceInterface) StorageController {
	return &storageController{
		storageService: storageService,
	}
}

func (controller *storageController) Save(ctx *gin.Context) (model.Storage, error) {
	var storage model.Storage
	err := ctx.ShouldBind(&storage)
	if err != nil {
		return model.Storage{}, err
	}

	storage, err = controller.storageService.Save(storage)
	if err != nil {
		return model.Storage{}, err
	}

	return storage, nil
}

func (controller *storageController) FindAll(ctx *gin.Context) ([]model.Storage, error) {
	storages, err := controller.storageService.FindAll()
	if err != nil {
		return []model.Storage{}, err
	}

	return storages, nil
}

func (controller *storageController) FindById(ctx *gin.Context) (model.Storage, error) {
	var storage model.Storage

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.Storage{}, err
	}

	storage, err = controller.storageService.FindById(id)
	if err != nil {
		return model.Storage{}, err
	}

	return storage, nil
}

func (controller *storageController) Update(ctx *gin.Context) (model.Storage, error) {
	var storage model.Storage
	err := ctx.ShouldBind(&storage)
	if err != nil {
		return model.Storage{}, err
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.Storage{}, err
	}

	storage.ID = uint(id)
	storage, err = controller.storageService.Update(storage)
	if err != nil {
		return model.Storage{}, err
	}

	return storage, nil
}

func (controller *storageController) Delete(ctx *gin.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	if err = controller.storageService.Delete(id); err != nil {
		return err
	}

	return nil
}
