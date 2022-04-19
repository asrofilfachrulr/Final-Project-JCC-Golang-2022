package migrations

import (
	"os"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if os.Getenv("DROP") == "y" {
		db.Migrator().DropTable(Models...)
	}

	return db.AutoMigrate(Models...)
}
