package models

import (
	models "anya-day/models/web"
	"anya-day/utils"
	"fmt"
	"log"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		Name       string `gorm:"not null"`
		MerchantID uint   `gorm:"not null"`
		Price      uint   `gorm:"not null"`
		Desc       string
		Rating     uint
		Stock      uint     `gorm:"not null"`
		CategoryID uint     `gorm:"not null"`
		Merchant   Merchant `gorm:"constraint:OnDelete:CASCADE"`
		Category   Category
	}
)

func GetMerchantProducts(db *gorm.DB, data *[]models.ProductOutput, m *Merchant, filter *models.ProductFilter) error {
	db.First(m)

	poarr := &[]models.ProductOutput{}

	if err := db.
		Model(&Product{}).
		Select("id", "name", "price", "rating").
		Where("merchant_id = ?", m.ID).
		Find(poarr).Error; err != nil {
		return err
	}

	for _, po := range *poarr {
		if !strings.Contains(strings.ToLower(po.Name), strings.ToLower(*filter.Name)) {
			continue
		}
		if *filter.Price != "" {
			fPrice := utils.StringToIntIgnore(*filter.Price)
			fmt.Printf("fPrice: %v\n", fPrice)
			fmt.Printf("po.Price: %v\n", int(po.Price))
			if fPrice > int(po.Price) {
				log.Println("continued")
				continue
			}
		}
		if *filter.Rating != "" {
			fRating, err := strconv.ParseFloat(*filter.Rating, 32)
			if err != nil {
				return err
			}
			if float32(fRating) > m.Rating {
				continue
			}
		}
		*data = append(*data, po)
	}

	return nil
}

func GetDetailedProduct(db *gorm.DB, data *models.MerchantProductOutput, m *Merchant, p *Product) error {
	data.Name = m.Name
	data.ID = m.ID

	data.Product = models.ProductDetailOutput{
		Name:   p.Name,
		Desc:   p.Desc,
		Price:  p.Price,
		Rating: p.Rating,
		Stock:  p.Stock,
	}

	return nil
}

func DeleteProductById(db *gorm.DB, p *Product) error {
	return db.Unscoped().Where("merchant_id = ?", p.MerchantID).Delete(p).Error
}
