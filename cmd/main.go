package main

import (
	"github.com/alexyslozada/ecommerce/infrastructure/handler"
	"github.com/alexyslozada/ecommerce/infrastructure/handler/response"
	"log"
	"os"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	e := newHTTP(response.HTTPErrorHandler)

	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(e, dbPool)

	err = e.Start(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
