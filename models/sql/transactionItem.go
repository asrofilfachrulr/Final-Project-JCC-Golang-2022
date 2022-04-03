package models

type (
	TransactionItem struct {
		TransactionID uint `gorm:"not null"`
		ProductID     uint `gorm:"not null"`
	}
)
