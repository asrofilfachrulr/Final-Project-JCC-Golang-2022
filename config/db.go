package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"anya-day/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
}

func ConfigByDebugMode(mode string) *gorm.Config {
	if mode == "DEBUG" {
		return &gorm.Config{
			Logger: NewLogger(),
		}
	}
	return &gorm.Config{}
}

func ConnectDataBase() *gorm.DB {
	var username, password, host, port, dbname, sslmode string

	username = os.Getenv("DATABASE_USERNAME")
	password = os.Getenv("DATABASE_PASSWORD")
	host = os.Getenv("DATABASE_HOST")
	port = os.Getenv("DATABASE_PORT")
	dbname = os.Getenv("DATABASE_NAME")
	sslmode = os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbname, port, sslmode)

	mode := os.Getenv("DEBUG_MODE")
	config := ConfigByDebugMode(mode)

	log.Println("Connecting to DB...")
	db, err := gorm.Open(postgres.Open(dsn), config)

	if err != nil {
		panic(err.Error())
	}

	log.Println("Auto migrating...")
	if err := migrations.Migrate(db); err != nil {
		log.Fatalln(err)
	}

	return db

}
