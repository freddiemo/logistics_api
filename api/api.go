package api

import (
	"github.com/gin-gonic/gin"

	"net/http"

	clients "github.com/freddiemo/logistics-api/internal/register/clients/controller"
	productTypes "github.com/freddiemo/logistics-api/internal/register/product_types/controller"
)

type LogisticsAPI struct {
	clientsController      clients.ClientController
	productTypesController productTypes.ProductTypeController
}

func NewLogisticsAPI(
	clientsController clients.ClientController,
	productTypesController productTypes.ProductTypeController) *LogisticsAPI {
	return &LogisticsAPI{
		clientsController:      clientsController,
		productTypesController: productTypesController,
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
