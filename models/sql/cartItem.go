package models

import "gorm.io/gorm"

type (
	CartItem struct {
		gorm.Model
		CartID    uint    `gorm:"not null"`
		ProductID uint    `gorm:"not null"`
		Price     uint    `gorm:"not null"`
		Qty       uint    `gorm:"not null"`
		SubTotal  uint    `gorm:"not null"`
		Cart      Cart    `gorm:"constraint:OnDelete:CASCADE"`
		Product   Product `gorm:"constraint:OnDelete:CASCADE"`
	}
)
