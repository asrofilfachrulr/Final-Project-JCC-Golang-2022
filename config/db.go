package config

import (
	"anya-day/models"
	"anya-day/utils"
	"fmt"
	"log"
	"os"
	"time"

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
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
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
	db, err := gorm.Open(postgres.Open(dsn), config)
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

// adding dummies, developer and admin account base INIT_DB env etc
func InitDB(db *gorm.DB, mode string) {
	if mode == "USERS" {
		// batch insert
		users := []models.User{
			{FullName: "John Doe", Username: "john", Email: "john@mail.com"},
			{FullName: "Mary Sue", Username: "mary", Email: "mary@mail.com"},
			{FullName: "Xi Ng", Username: "xi", Email: "nihaoma@mail.com"},
		}
		db.Create(&users)

		userCredentials := []models.UserCredential{
			{UserID: int(users[0].ID), Username: "john", Password: models.ConvToHash("john")},
			{UserID: int(users[1].ID), Username: "mary", Password: models.ConvToHash("mary")},
			{UserID: int(users[2].ID), Username: "xi", Password: models.ConvToHash("xi")},
		}
		db.Create(&userCredentials)
	}
}
