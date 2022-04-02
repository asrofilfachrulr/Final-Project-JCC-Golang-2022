package models

import (
	wmodels "anya-day/models/web"
	"fmt"
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

func (u *User) GetCompleteUser(db *gorm.DB, cu *wmodels.UserCompleteDataResp) (err error) {
	defer func() {
		if s := recover(); s != nil {
			err = fmt.Errorf("%v", s)
		}
	}()

	addr := &UserAddress{}
	type Result struct {
		Name string
		Role string
	}

	res := &Result{}

	err = db.Table("user_addresses ua").
		Select("c.name, r.name role").
		Joins("left join countries c on c.id = ua.country_id").
		Joins("left join roles r on r.user_id = ua.user_id").
		Where("ua.user_id = ?", u.ID).
		Scan(res).Error

	if err != nil {
		return err
	}

	_u := &User{}
	_u.ID = u.ID
	db.First(_u)

	db.Where("user_id = ?", u.ID).Find(addr)

	cu.ID = _u.ID
	cu.Username = _u.Username
	cu.Email = _u.Email
	cu.Fullname = _u.FullName
	cu.Role = res.Role
	cu.Address = wmodels.AddressRespData{
		AddressLine: addr.AddressLine,
		City:        addr.City,
		Country:     res.Name,
		PhoneNumber: addr.PhoneNumber,
		PostalCode:  addr.PostalCode,
	}
	return nil
}

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
