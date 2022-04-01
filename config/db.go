package config

import (
	"anya-day/models"
	"anya-day/utils"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	var username, password, host, port, dbname, sslmode string

	username = os.Getenv("DATABASE_USERNAME")
	password = os.Getenv("DATABASE_PASSWORD")
	host = os.Getenv("DATABASE_HOST")
	port = os.Getenv("DATABASE_PORT")
	dbname = os.Getenv("DATABASE_NAME")
	sslmode = os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// table models
	users := &models.User{}
	userCredentials := &models.UserCredential{}

	// DROP PREVIOUS TABLES
	if mode := utils.GetEnvWithFallback("DROP_TABLES", "false"); mode == "DROP" {
		log.Println("[MIGRATION] DROPPING TABLES")
		if db.Migrator().HasTable(users) {
			db.Migrator().DropTable(users)
		}
		if db.Migrator().HasTable(userCredentials) {
			db.Migrator().DropTable(userCredentials)
		}
	}

	log.Println("[MIGRATION] AUTO MIGRATING TABLES")
	// migrate new tables
	db.AutoMigrate(
		users,
		userCredentials,
	)

	return db

}
