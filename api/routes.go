package api

import (
	"github.com/gin-gonic/gin"
)

func addRoutes(rg *gin.RouterGroup, logisticsAPI *LogisticsAPI) {
	clients := rg.Group("/clients")
	{
		clients.GET("/", logisticsAPI.FindAllClients)
		clients.POST("/", logisticsAPI.SaveClient)
		clients.GET(":id", logisticsAPI.FindClientById)
	}
}

func getRoutes(logisticsAPI *LogisticsAPI) {
	v1 := server.Group("/v1")
	addRoutes(v1, logisticsAPI)
}
