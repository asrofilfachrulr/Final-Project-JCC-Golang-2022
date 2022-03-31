package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	UserCredential struct {
		ID       int    `gorm:"primary_key;autoIncrement"`
		UserID   int    `json:"user_id" gorm:"not null"`
		Username string `json:"username" gorm:"not null"`
		Password string `json:"password" gorm:"not null"`
		User     User   `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	}
)

func (uc *UserCredential) ConvToHash() error {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(uc.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return errPassword
	}

	uc.Password = string(hashedPassword)

	return nil
}

func (uc *UserCredential) SaveCredential(u *User, db *gorm.DB) error {
	uc.UserID = int(u.ID)

	err := db.Create(uc).Error
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserCredential) AttemptLogin(db *gorm.DB) error {
	realUC := &UserCredential{}
	db.Where("username = ?", uc.Username).First(&realUC)

	if err := bcrypt.CompareHashAndPassword([]byte(realUC.Password), []byte(uc.Password)); err != nil {
		return err
	}

	uc.ID = realUC.ID
	return nil
}
