package main

import (
	"anya-day/config"
	"anya-day/routes"
	"anya-day/utils"

	"anya-day/docs"

	"github.com/joho/godotenv"
)

func main() {
	if env := utils.Getenv("ENVIRONMENT", "development"); env == "development" {
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
	}

	db := config.ConnectDataBase()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Anya Day API"
	docs.SwaggerInfo.Description = "API provide backend service for your ecommerce app"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := routes.InitRoute(db)

	r.Run()
}
