package main

import (
	"fmt"

	"github.com/freddiemo/logistics-api/config"
	"github.com/freddiemo/logistics-api/db"
	clients "github.com/freddiemo/logistics-api/internal/register/clients/model"
	location "github.com/freddiemo/logistics-api/internal/register/location/model"
	productTypes "github.com/freddiemo/logistics-api/internal/register/product_types/model"
	storage "github.com/freddiemo/logistics-api/internal/register/storages/model"
)

func main() {
	envs := config.Init()
	db := db.Init(envs["DB_HOST"], envs["DB_USER"], envs["DB_PASSWORD"], envs["DB_NAME"], envs["DB_PORT"])

	var result interface{}
	db.Raw("CREATE TYPE transportation_type AS ENUM ('1', '2');").Scan(&result)
	if result != nil {
		fmt.Println(result)
	}

	if result != nil {
		fmt.Println(result)
	}

	db.AutoMigrate(&clients.Client{})
	db.AutoMigrate(&productTypes.ProductType{})
	db.AutoMigrate(&location.Country{})
	db.AutoMigrate(&location.State{})
	db.AutoMigrate(&location.City{})
	db.AutoMigrate(&storage.Storage{})
}
