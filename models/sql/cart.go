package models

import "gorm.io/gorm"

type (
	Cart struct {
		gorm.Model
		UserID     uint     `gorm:"not null"`
		MerchantID uint     `gorm:"not null"`
		Total      uint     `gorm:"not null"`
		User       User     `gorm:"constraint:OnDelete:CASCADE"`
		Merchant   Merchant `gorm:"constraint:OnDelete:CASCADE"`
	}
)
