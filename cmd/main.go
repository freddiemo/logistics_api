package main

import (
	"fmt"

	"github.com/freddiemo/logistics-api/config"
	"github.com/freddiemo/logistics-api/db"
)

func main() {
	envs := config.Init()
	db := db.Init(envs["DB_HOST"], envs["DB_USER"], envs["DB_PASSWORD"], envs["DB_NAME"], envs["DB_PORT"])

	fmt.Println(db)
	fmt.Println("logistics api")
}
