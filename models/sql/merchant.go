package models

import "gorm.io/gorm"

type (
	Merchant struct {
		gorm.Model
		Name    string `gorm:"not null"`
		Rating  uint8
		AdminId uint `gorm:"not null"`
		User    User `gorm:"foreignKey:AdminId;constraint:OnDelete:CASCADE"`
	}
)
