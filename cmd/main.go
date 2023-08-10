package main

import (
	"github.com/freddiemo/logistics-api/api"
	"github.com/freddiemo/logistics-api/config"
	"github.com/freddiemo/logistics-api/db"
)

func main() {
	envs := config.Init()
	db := db.Init(envs["DB_HOST"], envs["DB_USER"], envs["DB_PASSWORD"], envs["DB_NAME"], envs["DB_PORT"])

	api.Init(envs, db)
}
