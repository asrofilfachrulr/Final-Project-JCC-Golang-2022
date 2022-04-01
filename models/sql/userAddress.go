package models

import (
	wmodels "anya-day/models/web"

	"gorm.io/gorm"
)

type (
	UserAddress struct {
		gorm.Model
		User        User `gorm:"constraint:OnDelete:CASCADE;"`
		Country     Country
		UserID      uint   `gorm:"not null"`
		AddressLine string `gorm:"not null"`
		City        string `gorm:"not null"`
		CountryID   uint   `gorm:"not null"`
		PhoneNumber uint   `gorm:"not null"`
		PostalCode  uint
	}
)

func (ua *UserAddress) PostAddress(db *gorm.DB, data *wmodels.AddressInput) error {
	country := &Country{}
	db.Where(&Country{Name: data.Country}).Find(country)

	ua.CountryID = country.ID
	ua.AddressLine = data.AddressLine
	ua.City = data.City
	ua.PhoneNumber = data.PhoneNumber
	ua.PostalCode = data.PostalCode

	if err := db.Save(ua).Error; err != nil {
		return err
	}

	return nil
}
