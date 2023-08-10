package api

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func addRoutes(rg *gin.RouterGroup, logisticsAPI *LogisticsAPI) {
	clients := rg.Group("/clients")
	{
		clients.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "clients")
		})

		clients.POST("/", logisticsAPI.SaveClient)
	}
}

func getRoutes(logisticsAPI *LogisticsAPI) {
	v1 := server.Group("/v1")
	addRoutes(v1, logisticsAPI)
}
