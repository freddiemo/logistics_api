package api

import (
	"github.com/gin-gonic/gin"

	"net/http"

	clients "github.com/freddiemo/logistics-api/internal/clients/controller"
)

type LogisticsAPI struct {
	clientsController clients.ClientController
}

func NewLogisticsAPI(clientsController clients.ClientController) *LogisticsAPI {
	return &LogisticsAPI{
		clientsController: clientsController,
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
