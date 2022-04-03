package models

import (
	"gorm.io/gorm"
)

type (
	Transaction struct {
		gorm.Model
		UserID           uint
		MerchantID       uint
		PaymentID        uint
		ShipmentID       uint
		TransactionItems []TransactionItem
		Username         string   `gorm:"not null"`
		MerchantName     string   `gorm:"not null"`
		PaymentInfo      string   `gorm:"not null"`
		ShipmentInfo     string   `gorm:"not null"`
		ShippingAddress  string   `gorm:"not null"`
		Total            uint     `gorm:"not null"`
		Status           string   `gorm:"not null"`
		User             User     `gorm:"OnDelete:SET NULL"`
		Shipment         Shipment `gorm:"OnDelete:SET NULL"`
		Payment          Payment  `gorm:"OnDelete:SET NULL"`
		Merchant         Merchant `gorm:"OnDelete:SET NULL"`
	}
)
