package models

type (
	Category struct {
		ID        uint      `json:"id" gorm:"primary_key;autoIncrement"`
		Name      string    `json:"name" gorm:"not null;unique"`
		Countries []Country `json:"-" gorm:"many2many:prohibit_categories"`
	}
)
