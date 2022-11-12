package main

import (
	"log"
	"os"

	"github.com/alexyslozada/ecommerce/infrastructure/handler"
	"github.com/alexyslozada/ecommerce/infrastructure/handler/response"
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

	port := os.Getenv("SERVER_PORT")
	if os.Getenv("IS_HTTPS") == "true" {
		err = e.StartTLS(":"+port, os.Getenv("CERT_PEM_FILE"), os.Getenv("KEY_PEM"))
	} else {
		err = e.Start(":" + port)
	}
	if err != nil {
		log.Fatal(err)
	}

}
