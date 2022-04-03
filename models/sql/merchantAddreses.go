package models

type (
	MerchantAddress struct {
		MerchantID          uint
		City                string `gorm:"not null"`
		OfflineStoreAddress string
		CountryID           uint
		Country             Country
		Merchant            Merchant `gorm:"constraint:OnDelete:CASCADE"`
	}
)
