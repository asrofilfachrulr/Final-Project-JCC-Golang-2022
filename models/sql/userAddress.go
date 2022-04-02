package models

import (
	wmodels "anya-day/models/web"
	"anya-day/utils"

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

	err := utils.ValidateDigitInt(int(data.PhoneNumber), 10, 15, "phone number")
	if err != nil {
		return err
	}
	ua.PhoneNumber = data.PhoneNumber

	err = utils.ValidateDigitInt(int(data.PostalCode), 4, 5, "postal code")
	if err != nil {
		return err
	}
	ua.PostalCode = data.PostalCode

	if err := db.Save(ua).Error; err != nil {
		return err
	}

	return nil
}

func (ua *UserAddress) UpdateAddress(db *gorm.DB, data *wmodels.AddressInputNotBinding) error {
	isChanged := false
	_ua := &UserAddress{}
	if data.AddressLine != "" {
		_ua.AddressLine = data.AddressLine
		isChanged = true
	}
	if data.City != "" {
		_ua.City = data.City
		isChanged = true
	}
	if data.Country != "" {
		country := &Country{}
		err := db.Where(&Country{Name: data.Country}).Find(country).Error
		if err != nil {
			return err
		}
		_ua.CountryID = country.ID
		isChanged = true
	}
	if data.PhoneNumber != 0 {
		err := utils.ValidateDigitInt(int(data.PhoneNumber), 10, 15, "phone number")
		if err != nil {
			return err
		}
		_ua.PhoneNumber = data.PhoneNumber
		isChanged = true
	}
	if data.PostalCode != 0 {
		err := utils.ValidateDigitInt(int(data.PostalCode), 4, 5, "postal code")
		if err != nil {
			return err
		}
		_ua.PostalCode = data.PostalCode
		isChanged = true
	}
	if isChanged {
		db.Model(&UserAddress{}).Where("user_id = ?", ua.UserID).Updates(_ua)
	}
	return nil
}
