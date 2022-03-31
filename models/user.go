package models

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint      `json:"id" gorm:"primary_key;autoIncrement"`
		FullName  string    `json:"full_name" gorm:"not null;"`
		Username  string    `json:"username" gorm:"not null;unique"`
		Email     string    `json:"email" gorm:"not null;unique"`
		CreatedAt time.Time `json:"joined_at"`
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
