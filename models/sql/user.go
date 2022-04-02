package models

import (
	wmodels "anya-day/models/web"
	"html"
	"net/mail"
	"strings"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		FullName string `gorm:"not null;"`
		Username string `gorm:"not null;unique"`
		Email    string `gorm:"not null;unique"`
	}
)

func (u *User) SaveUser(db *gorm.DB) error {
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Username = strings.ToLower(u.Username)

	var err error = db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUserProfile(db *gorm.DB, data *wmodels.UpdateProfileInput) error {
	isChanged := false
	if data.Email != "" {
		addr, err := mail.ParseAddress(data.Email)
		if err != nil {
			return err
		}
		u.Email = addr.Address
		isChanged = true
	}
	if data.FullName != "" {
		u.FullName = data.FullName
		isChanged = true
	}
	if data.Username != "" {
		u.Username = data.Username
		isChanged = true
	}

	if isChanged {
		// commiting change(s)
		if err := db.Save(u).Error; err != nil {
			return err
		}
	}
	return nil
}

func (u *User) AttemptDeleteUser(db *gorm.DB) error {
	// retrieve all data first for later response
	db.First(u)

	// delete user
	err := db.Unscoped().Delete(u).Error
	if err != nil {
		return err
	}
	return nil
}
