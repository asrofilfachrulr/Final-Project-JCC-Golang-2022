package models

import "gorm.io/gorm"

type (
	Role struct {
		gorm.Model
		UserID uint   `json:"-" gorm:"not null"`
		Name   string `json:"-" gorm:"not null"`
		User   User   `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	}
)
