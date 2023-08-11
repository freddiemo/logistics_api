package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"io"
	"os"

	// register
	clients "github.com/freddiemo/logistics-api/internal/register/clients/controller"
	clientsRepository "github.com/freddiemo/logistics-api/internal/register/clients/repository"
	clientsService "github.com/freddiemo/logistics-api/internal/register/clients/service"

	productTypes "github.com/freddiemo/logistics-api/internal/register/product_types/controller"
	productTypesRepository "github.com/freddiemo/logistics-api/internal/register/product_types/repository"
	productTypesService "github.com/freddiemo/logistics-api/internal/register/product_types/service"

	storages "github.com/freddiemo/logistics-api/internal/register/storages/controller"
	storagesRepository "github.com/freddiemo/logistics-api/internal/register/storages/repository"
	storagesService "github.com/freddiemo/logistics-api/internal/register/storages/service"

	// logistics
	landShipment "github.com/freddiemo/logistics-api/internal/logistics/land_shipment/controller"
	landShipmentRepository "github.com/freddiemo/logistics-api/internal/logistics/land_shipment/repository"
	landShipmentService "github.com/freddiemo/logistics-api/internal/logistics/land_shipment/service"
)

var server = gin.Default()

func Init(envs map[string]string, db *gorm.DB) {
	setupLogOutput(envs["APP_NAME"])

	// register
	var clientRepository clientsRepository.ClientRepository = clientsRepository.NewClientRepository(db)
	var clientService clientsService.ClientServiceInterface = clientsService.NewClientService(clientRepository)
	var clientsController clients.ClientController = clients.NewClientController(clientService)

	var productTypesRepository productTypesRepository.ProductTypeRepository = productTypesRepository.NewProductTypeRepository(db)
	var productTypesService productTypesService.ProductTypeServiceInterface = productTypesService.NewProductTypeService(productTypesRepository)
	var productTypesController productTypes.ProductTypeController = productTypes.NewProductTypeController(productTypesService)

	var storagesRepository storagesRepository.StorageRepository = storagesRepository.NewStorageRepository(db)
	var storagesService storagesService.StorageServiceInterface = storagesService.NewStorageService(storagesRepository)
	var storagesController storages.StorageController = storages.NewStorageController(storagesService)

	// logistics
	var landShipmentRepository landShipmentRepository.LandShipmentRepository = landShipmentRepository.NewLandShipmentRepository(db)
	var landShipmentService landShipmentService.LandShipmentServiceInterface = landShipmentService.NewLandShipmentService(landShipmentRepository)
	var landShipmentController landShipment.LandShipmentController = landShipment.NewLandShipmentController(landShipmentService)

	logisticsAPI := NewLogisticsAPI(
		// register
		clientsController, productTypesController, storagesController,
		// logistics
		landShipmentController,
	)

	getRegisterRoutes(logisticsAPI)

	server.Run(":" + envs["APP_PORT"])
}

func setupLogOutput(app_name string) {
	f, _ := os.Create("./log/" + app_name + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
