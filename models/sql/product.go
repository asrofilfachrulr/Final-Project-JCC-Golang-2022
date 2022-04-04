package models

import (
	models "anya-day/models/web"

	"gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		Name       string `gorm:"not null"`
		MerchantID uint   `gorm:"not null"`
		Price      uint   `gorm:"not null"`
		Desc       string
		Stock      uint     `gorm:"not null"`
		CategoryID uint     `gorm:"not null"`
		Merchant   Merchant `gorm:"constraint:OnDelete:CASCADE"`
		Category   Category
	}
)

func GetMerchantProducts(db *gorm.DB, data *[]models.ProductOutput, m *Merchant) error {
	db.First(m)

	if err := db.
		Model(&Product{}).
		Select("id", "name", "price").
		Where("merchant_id = ?", m.ID).
		Find(data).Error; err != nil {
		return err
	}

	return nil
}
