package models

import (
	wmodels "anya-day/models/web"

	"gorm.io/gorm"
)

type (
	Merchant struct {
		gorm.Model
		Name     string `gorm:"not null"`
		Rating   string
		AdminId  uint `gorm:"not null"`
		User     User `gorm:"foreignKey:AdminId;constraint:OnDelete:CASCADE"`
		Products []Product
	}
)

func GetMerchants(db *gorm.DB, m []Merchant, mout *[]wmodels.MerchantOutput) error {
	var users []User
	db.Find(&users)
	db.Find(&m)

	usernameLookup := func(id uint) string {
		for _, u := range users {
			if u.ID == id {
				return u.Username
			}
		}
		return ""
	}

	for _, merchant := range m {
		username := usernameLookup(merchant.AdminId)

		*mout = append(*mout, wmodels.MerchantOutput{
			ID:     merchant.ID,
			Name:   merchant.Name,
			Rating: merchant.Rating,
			Owner:  username,
		})
	}

	return nil
}
