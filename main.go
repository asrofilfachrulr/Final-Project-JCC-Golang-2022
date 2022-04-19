package main

import (
	"anya-day/config"
	"anya-day/routes"
	"anya-day/utils"
	"flag"
	"log"
	"os"

	"anya-day/docs"

	"github.com/joho/godotenv"
)

func init() {
	env := utils.GetEnvWithFallback("ENVIRONMENT", "development")
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
	docs.SwaggerInfo.Host = utils.GetEnvWithFallback("SWAGGER_HOST", "localhost:8080")
}

// @title Anya Day API
// @version 2.0
// @description API which provide you backend service for your minimalist ecommerce app

// @contact.name Developer
// @contact.email riidloa@gmail.com

// @BasePath /api/v1
func main() {
	db := config.ConnectDataBase()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.Init(db)

	routes.Attach(r, map[string]any{
		"db": db,
	})

	r.Run()
}
