package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"io"
	"os"

	clients "github.com/freddiemo/logistics-api/internal/register/clients/controller"
	clientsRepository "github.com/freddiemo/logistics-api/internal/register/clients/repository"
	clientsService "github.com/freddiemo/logistics-api/internal/register/clients/service"

	productTypes "github.com/freddiemo/logistics-api/internal/register/product_types/controller"
	productTypesRepository "github.com/freddiemo/logistics-api/internal/register/product_types/repository"
	productTypesService "github.com/freddiemo/logistics-api/internal/register/product_types/service"
)

var server = gin.Default()

func Init(envs map[string]string, db *gorm.DB) {
	setupLogOutput(envs["APP_NAME"])

	var clientRepository clientsRepository.ClientRepository = clientsRepository.NewClientRepository(db)
	var clientService clientsService.ClientServiceInterface = clientsService.NewClientService(clientRepository)
	var clientsController clients.ClientController = clients.NewClientController(clientService)

	var productTypesRepository productTypesRepository.ProductTypeRepository = productTypesRepository.NewProductTypeRepository(db)
	var productTypesService productTypesService.ProductTypeServiceInterface = productTypesService.NewProductTypeService(productTypesRepository)
	var productTypesController productTypes.ProductTypeController = productTypes.NewProductTypeController(productTypesService)

	logisticsAPI := NewLogisticsAPI(clientsController, productTypesController)

	getRegisterRoutes(logisticsAPI)

	server.Run(":" + envs["APP_PORT"])
}

func setupLogOutput(app_name string) {
	f, _ := os.Create("./log/" + app_name + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
