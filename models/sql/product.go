package models

import "gorm.io/gorm"

type (
	Product struct {
		gorm.Model
		Name       string   `gorm:"not null"`
		MerchantID uint     `gorm:"not null"`
		Price      uint     `gorm:"not null"`
		CategoryID uint     `gorm:"not null"`
		Merchant   Merchant `gorm:"constraint:OnDelete:CASCADE"`
		Category   Category
	}
)
