package models

import (
	models "anya-day/models/web"

	"gorm.io/gorm"
)

type (
	ProductReview struct {
		UserID    uint `gorm:"not null"`
		ProductID uint `gorm:"not null"`
		Review    string
		Rating    float32 `gorm:"not null"`
		User      User    `gorm:"constraint:OnDelete:CASCADE"`
		Product   Product `gorm:"constraint:OnDelete:CASCADE"`
	}
)

func (r *ProductReview) PostReview(db *gorm.DB) error {
	return db.Create(r).Error
}

func GetReview(db *gorm.DB, data *[]models.Review, productId uint) error {
	return db.
		Table("product_reviews pr").
		Select("u.username", "pr.review", "pr.rating").
		Joins("left join users u on u.id = pr.user_id").
		Where("product_id = ?", productId).
		Find(data).Error
}
