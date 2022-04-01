package models

type (
	Country struct {
		ID   uint   `gorm:"primary_key;autoIncrement"`
		Name string `gorm:"not null"`
	}
)
