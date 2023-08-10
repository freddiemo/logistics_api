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
	}
}

func getRegisterRoutes(logisticsAPI *LogisticsAPI) {
	v1 := server.Group("/v1")
	register := v1.Group("/register")
	addRegisterRoutes(register, logisticsAPI)
}
