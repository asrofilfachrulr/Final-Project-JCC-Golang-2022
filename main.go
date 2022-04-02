package main

import (
	"anya-day/config"
	"anya-day/routes"
	"anya-day/utils"
	"flag"
	"os"

	"anya-day/docs"

	"github.com/joho/godotenv"
)

// @title Anya Day API
// @version beta
// @description API which provide you backend service for your minimalist ecommerce app

// @contact.name Developer
// @contact.email riidloa@gmail.com

// @BasePath /api/v1
func main() {
	env := utils.GetEnvWithFallback("ENVIRONMENT", "development")
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
		docs.SwaggerInfo.Schemes = []string{"http"}

		dropMode := flag.String("drop", "nope", "drop all tables and seeding the db if set to DROP, for debugging purpose")
		debugMode := flag.String("d", "nope", "logging every sql operation if set to DEBUG")
		initDBMode := flag.String("initdb", "nope", "seeding db whether with dummies only (USERS) or dev users too (USERS_DEV)")
		flag.Parse()

		os.Setenv("DROP_TABLES", *dropMode)
		os.Setenv("DEBUG_MODE", *debugMode)
		os.Setenv("INIT_DB", *initDBMode)
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	db := config.ConnectDataBase()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	config.Load(db)

	docs.SwaggerInfo.Host = utils.GetEnvWithFallback("SWAGGER_HOST", "localhost:8080")

	r := routes.InitRoute(db)

	r.Run()
}
