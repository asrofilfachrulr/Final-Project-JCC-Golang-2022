package main

import (
	"anya-day/config"
	"anya-day/helper"
	"flag"
	"log"
	"os"

	"anya-day/docs"

	"github.com/joho/godotenv"
)

func init() {
	env := helper.GetEnvWithFallback("ENVIRONMENT", "development")
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			// exit program in fail load .env
			log.Fatalln(err.Error())
		}

		// set schemes only http if run in local
		docs.SwaggerInfo.Schemes = []string{"http"}

		// set debug mode by given flag
		debugMode := flag.String("d", "", "")
		// drop all table before migration
		drop := flag.String("drop", "", "")
		flag.Parse()

		// set to this runtime env
		os.Setenv("DEBUG_MODE", *debugMode)
		os.Setenv("DROP", *drop)
	} else {
		// set schemes only https if run in production
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	// get host url from env
	docs.SwaggerInfo.Host = helper.GetEnvWithFallback("SWAGGER_HOST", "localhost:8080")
}

// @title Anya Day API
// @version 2.0
// @description API which provide you backend service for your minimalist ecommerce app

// @contact.name Developer
// @contact.email riidloa@gmail.com

// @BasePath /api/v1
func main() {
	server := config.NewServer()
	server.Init()
	server.Run()
}
