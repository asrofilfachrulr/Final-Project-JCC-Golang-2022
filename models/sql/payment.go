package models

type (
	Payment struct {
		ID     uint   `gorm:"primary key,autoIncrement,not null"`
		Name   string `gorm:"not null"`
		Method string `gorm:"not null"`
	}
)
