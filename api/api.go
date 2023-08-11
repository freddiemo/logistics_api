package api

import (
	"github.com/gin-gonic/gin"

	"net/http"

	clients "github.com/freddiemo/logistics-api/internal/register/clients/controller"
	productTypes "github.com/freddiemo/logistics-api/internal/register/product_types/controller"
	storages "github.com/freddiemo/logistics-api/internal/register/storages/controller"
)

type LogisticsAPI struct {
	clientsController      clients.ClientController
	productTypesController productTypes.ProductTypeController
	storagesController     storages.StorageController
}

func NewLogisticsAPI(
	clientsController clients.ClientController,
	productTypesController productTypes.ProductTypeController,
	storagesController storages.StorageController) *LogisticsAPI {
	return &LogisticsAPI{
		clientsController:      clientsController,
		productTypesController: productTypesController,
		storagesController:     storagesController,
	}
}

func (api *LogisticsAPI) SaveClient(ctx *gin.Context) {
	client, err := api.clientsController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, client)
	}
}

func (api *LogisticsAPI) FindAllClients(ctx *gin.Context) {
	clients, err := api.clientsController.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, clients)
	}
}

func (api *LogisticsAPI) FindClientById(ctx *gin.Context) {
	client, err := api.clientsController.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, client)
	}
}

func (api *LogisticsAPI) UpdateClient(ctx *gin.Context) {
	client, err := api.clientsController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, client)
	}
}

func (api *LogisticsAPI) DeleteClient(ctx *gin.Context) {
	err := api.clientsController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (api *LogisticsAPI) SaveProductType(ctx *gin.Context) {
	productType, err := api.productTypesController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, productType)
	}
}

func (api *LogisticsAPI) FindAllProductTypes(ctx *gin.Context) {
	productTypes, err := api.productTypesController.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, productTypes)
	}
}

func (api *LogisticsAPI) FindProductTypeById(ctx *gin.Context) {
	productType, err := api.productTypesController.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, productType)
	}
}

func (api *LogisticsAPI) UpdateProductType(ctx *gin.Context) {
	productType, err := api.productTypesController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, productType)
	}
}

func (api *LogisticsAPI) DeleteProductType(ctx *gin.Context) {
	err := api.productTypesController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (api *LogisticsAPI) SaveStorage(ctx *gin.Context) {
	storage, err := api.storagesController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, storage)
	}
}

func (api *LogisticsAPI) FindAllStorages(ctx *gin.Context) {
	storages, err := api.storagesController.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, storages)
	}
}

func (api *LogisticsAPI) FindStorageById(ctx *gin.Context) {
	storage, err := api.storagesController.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, storage)
	}
}

func (api *LogisticsAPI) UpdateStorage(ctx *gin.Context) {
	storage, err := api.storagesController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, storage)
	}
}

func (api *LogisticsAPI) DeleteStorage(ctx *gin.Context) {
	err := api.storagesController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusNoContent, nil)
	}
}
