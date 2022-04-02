package config

import (
	models "anya-day/models/sql"
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
	db, err := gorm.Open(postgres.Open(dsn), config)

	if err != nil {
		panic(err.Error())
	}

	return db

}

// migration and seeding database
func Load(db *gorm.DB) {
	// table models
	users := &models.User{}
	userCredentials := &models.UserCredential{}
	roles := &models.Role{}
	countries := &models.Country{}
	uaddr := &models.UserAddress{}
	categories := &models.Category{}

	// DROP PREVIOUS TABLES
	isDropped := false
	if mode := utils.GetEnvWithFallback("DROP_TABLES", "false"); mode == "DROP" {
		log.Println("[MIGRATION] DROPPING TABLES")
		if db.Migrator().HasTable(users) {
			db.Migrator().DropTable(users)
		}
		if db.Migrator().HasTable(countries) {
			db.Migrator().DropTable(countries)
		}
		if db.Migrator().HasTable(categories) {
			db.Migrator().DropTable(categories)
		}
		if db.Migrator().HasTable(userCredentials) {
			db.Migrator().DropTable(userCredentials)
		}
		if db.Migrator().HasTable(roles) {
			db.Migrator().DropTable(roles)
		}
		if db.Migrator().HasTable(uaddr) {
			db.Migrator().DropTable(uaddr)
		}
		isDropped = true
	}

	log.Println("[MIGRATION] AUTO MIGRATING TABLES")
	// migrate new tables
	// static ordered first
	// independent ordered first
	db.AutoMigrate(
		users,
		countries,
		categories,
		userCredentials,
		roles,
		uaddr,
	)

	// init dynamic, dummy, for testing purposes data
	if isDropped {
		InitDynamicData(db)
		InitStaticData(db)
	}

}

// adding dummies, developer and admin account base INIT_DB env etc
func InitDynamicData(db *gorm.DB) {
	initUserDummy := func() {
		users := models.UsersDummy
		// batch insert
		db.Create(&users)

		for _, user := range users {
			// create record for user_credentials tab;e
			db.Create(&models.UserCredential{
				UserID: int(user.ID), Username: user.Username, Password: models.ConvToHash(user.Username),
			})

			// create record for roles table
			db.Create(&models.Role{
				UserID: user.ID, Name: "customer",
			})
		}
	}

	initDevAcc := func() {
		usr := &models.User{
			FullName: os.Getenv("DEV_FULLNM"),
			Username: os.Getenv("DEV_USRNM"),
			Email:    os.Getenv("DEV_MAIL"),
		}
		db.Create(usr)
		db.Create(&models.UserCredential{
			UserID:   int(usr.ID),
			Username: usr.Username,
			Password: models.ConvToHash(usr.Username),
		})
		db.Create(&models.Role{
			UserID: usr.ID, Name: "dev",
		})
	}

	mode := os.Getenv("INIT_DB")
	log.Printf("mode init: %s\n", mode)
	if mode == "USERS_DEV" {
		initUserDummy()
		initDevAcc()
	} else if mode == "USERS" {
		initUserDummy()
	}

}

// add static data which likely rarely be updated or deleted, so the table won't be dropped
func InitStaticData(db *gorm.DB) {
	initCountry := func() {
		db.Create(&models.AseanCountries)
	}

	initCategory := func() {
		db.Create(&models.Categories)
	}

	initCountry()
	initCategory()
}
