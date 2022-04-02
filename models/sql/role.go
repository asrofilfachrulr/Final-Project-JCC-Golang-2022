package models

import (
	"log"

	"gorm.io/gorm"
)

type (
	Role struct {
		gorm.Model
		UserID uint   `json:"-" gorm:"not null"`
		Name   string `json:"-" gorm:"not null"`
		User   User   `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	}
)

func (r *Role) ChangeRole(db *gorm.DB) error {
	_r := &Role{}
	_u := &User{}
	_u.ID = r.UserID
	db.Where(r).Find(_r)
	db.First(_u)

	log.Printf("role nested struct: %v\n", *_r)

	if _r.Name == "customer" {
		_r.Name = "merchant"
	} else if _r.Name == "merchant" {
		_r.Name = "customer"

		//TODO: wipe out merchant data
	}

	err := db.Save(_r).Error
	if err != nil {
		return err
	}

	r.User = *_u
	r.Name = _r.Name

	return nil
}
