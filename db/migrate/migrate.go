package main

import (
	"github.com/freddiemo/logistics-api/config"
	"github.com/freddiemo/logistics-api/db"
	clients "github.com/freddiemo/logistics-api/internal/clients/model"
)

func main() {
	envs := config.Init()
	db := db.Init(envs["DB_HOST"], envs["DB_USER"], envs["DB_PASSWORD"], envs["DB_NAME"], envs["DB_PORT"])

	db.AutoMigrate(&clients.Client{})
}
