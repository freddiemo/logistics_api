package api

import (
	"github.com/gin-gonic/gin"
)

func addRegisterRoutes(rg *gin.RouterGroup, logisticsAPI *LogisticsAPI) {
	clients := rg.Group("/clients")
	{
		clients.GET("/", logisticsAPI.FindAllClients)
		clients.POST("/", logisticsAPI.SaveClient)
		clients.GET(":id", logisticsAPI.FindClientById)
		clients.PUT(":id", logisticsAPI.UpdateClient)
		clients.DELETE(":id", logisticsAPI.DeleteClient)
	}

	productTypes := rg.Group("/product_types")
	{
		productTypes.GET("/", logisticsAPI.FindAllProductTypes)
		productTypes.POST("/", logisticsAPI.SaveProductType)
		productTypes.GET(":id", logisticsAPI.FindProductTypeById)
		productTypes.PUT(":id", logisticsAPI.UpdateProductType)
		productTypes.DELETE(":id", logisticsAPI.DeleteProductType)
	}

	storages := rg.Group("/storages")
	{
		storages.GET("/", logisticsAPI.FindAllStorages)
		storages.POST("/", logisticsAPI.SaveStorage)
		storages.GET(":id", logisticsAPI.FindStorageById)
		storages.PUT(":id", logisticsAPI.UpdateStorage)
		storages.DELETE(":id", logisticsAPI.DeleteStorage)
	}
}

func addLogisticsRoutes(rg *gin.RouterGroup, logisticsAPI *LogisticsAPI) {
	landShipments := rg.Group("/land_shipments")
	{
		landShipments.GET("/", logisticsAPI.FindAllLandShipments)
		landShipments.POST("/", logisticsAPI.SaveLandShipment)
		landShipments.GET(":id", logisticsAPI.FindLandShipmentById)
		landShipments.PUT(":id", logisticsAPI.UpdateLandShipment)
		landShipments.DELETE(":id", logisticsAPI.DeleteLandShipment)
	}
}

func getRegisterRoutes(logisticsAPI *LogisticsAPI) {
	v1 := server.Group("/v1")

	// register
	register := v1.Group("/register")
	addRegisterRoutes(register, logisticsAPI)

	// logistics
	logistics := v1.Group("/logistics")
	addLogisticsRoutes(logistics, logisticsAPI)
}
