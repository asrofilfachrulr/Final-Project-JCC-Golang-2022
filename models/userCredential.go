package models

import (
	"anya-day/token"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	UserCredential struct {
		ID        int       `gorm:"primary_key;autoIncrement"`
		UserID    int       `json:"user_id" gorm:"not null"`
		Username  string    `json:"username" gorm:"not null"`
		Password  string    `json:"password" gorm:"not null"`
		CreatedAt time.Time `gorm:"not null"`
		UpdatedAt time.Time `gorm:"not null"`
		User      User      `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
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
	// realUC for real credential saved in db
	realUC := &UserCredential{}

	// check if the login attempt instance using username or user_id as identity
	if uc.Username == "" {
		db.Where("user_id = ?", uc.UserID).First(&realUC)
	} else {
		db.Where("username = ?", uc.Username).First(&realUC)
	}

	// compare real password (in hash) and instance password
	if err := bcrypt.CompareHashAndPassword([]byte(realUC.Password), []byte(uc.Password)); err != nil {
		return err
	}

	// attach user id
	uc.UserID = realUC.UserID
	return nil
}

// Check whether given password is identical as former
func (uc *UserCredential) CheckPasswordTwin(db *gorm.DB) bool {
	dumm := &UserCredential{
		UserID:   uc.UserID,
		Password: uc.Password,
	}

	// simulating login, if success.. given password is identical
	if err := dumm.AttemptLogin(db); err != nil {
		return false
	}
	return true
}

func (uc *UserCredential) AttemptChangePw(newPw string, db *gorm.DB, c *gin.Context) error {
	// get user_id from jwt payload
	uid, err := token.ExtractUID(c)
	if err != nil {
		return err
	}

	uc.UserID = int(uid)

	// verify given old password by attempting login with user_id
	err = uc.AttemptLogin(db)
	if err != nil {
		return err
	}

	// check whether given new password is identical as former
	uc.Password = newPw
	if uc.CheckPasswordTwin(db) {
		return fmt.Errorf("password is identical as previous, create the new one instead")
	}

	err = uc.ConvToHash()
	if err != nil {
		return err
	}

	// update password
	db.Model(uc).Where("user_id = ?", uc.UserID).Update("password", uc.Password)

	return nil
}
