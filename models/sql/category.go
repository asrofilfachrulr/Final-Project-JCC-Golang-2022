package models

type (
	Category struct {
		ID   uint   `gorm:"primary_key;autoIncrement"`
		Name string `gorm:"not null;unique"`
	}
)
