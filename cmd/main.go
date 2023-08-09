package main

import (
	"fmt"

	"github.com/freddiemo/logistics-api/config"
)

func main() {
	envs := config.Init()
	fmt.Println(envs["APP_NAME"])
	fmt.Println("logistics api")
}
