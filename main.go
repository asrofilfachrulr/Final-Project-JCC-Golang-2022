package main

import (
	"anya-day/config"
	"anya-day/routes"
	"anya-day/utils"
	"os"

	"anya-day/docs"

	"github.com/joho/godotenv"
)

func main() {
	if env := utils.GetEnvWithFallback("ENVIRONMENT", "development"); env == "development" {
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
		docs.SwaggerInfo.Schemes = []string{"http"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	db := config.ConnectDataBase()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	dummiesMode := os.Getenv("INIT_DB")
	config.InitDB(db, dummiesMode)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Anya Day API"
	docs.SwaggerInfo.Description = "API provide backend service for your ecommerce app"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.GetEnvWithFallback("SWAGGER_HOST", "localhost:8080")

	r := routes.InitRoute(db)

	r.Run()
}
