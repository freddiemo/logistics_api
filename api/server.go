package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"io"
	"os"

	clients "github.com/freddiemo/logistics-api/internal/clients/controller"
	clientsRepository "github.com/freddiemo/logistics-api/internal/clients/repository"
	clientsService "github.com/freddiemo/logistics-api/internal/clients/service"
)

var server = gin.Default()

func Init(envs map[string]string, db *gorm.DB) {
	setupLogOutput(envs["APP_NAME"])

	var clientRepository clientsRepository.ClientRepository = clientsRepository.NewClientRepository(db)
	var clientService clientsService.ClientServiceInterface = clientsService.NewClientService(clientRepository)
	var clientsController clients.ClientController = clients.NewClientController(clientService)
	logisticsAPI := NewLogisticsAPI(clientsController)

	getRoutes(logisticsAPI)

	server.Run(":" + envs["APP_PORT"])
}

func setupLogOutput(app_name string) {
	f, _ := os.Create("./log/" + app_name + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
