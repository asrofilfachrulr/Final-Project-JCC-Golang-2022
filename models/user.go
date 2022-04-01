package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		FullName string `json:"full_name" gorm:"not null;"`
		Username string `json:"username" gorm:"not null;unique"`
		Email    string `json:"email" gorm:"not null;unique"`
	}
)

func (u *User) SaveUser(db *gorm.DB) error {
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}
