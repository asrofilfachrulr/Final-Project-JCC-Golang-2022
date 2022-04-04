package models

type (
	ProductReview struct {
		UserID    uint `gorm:"not null"`
		ProductID uint `gorm:"not null"`
		Review    string
		Rating    uint    `gorm:"not null"`
		User      User    `gorm:"constraint:OnDelete:CASCADE"`
		Product   Product `gorm:"constraint:OnDelete:CASCADE"`
	}
)
