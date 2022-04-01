package models

import "gorm.io/gorm"

type (
	UserAddress struct {
		gorm.Model
		User        User    `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
		Country     Country `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
		UserID      uint    `gorm:"not null"`
		AddressLine string  `gorm:"not null"`
		City        string  `gorm:"not null"`
		CountryID   uint    `gorm:"not null"`
		PostalCode  uint    `gorm:"not null"`
		PhoneNumber uint    `gorm:"not null"`
	}
)
