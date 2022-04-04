package models

import (
	wmodels "anya-day/models/web"
	"fmt"
	"log"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type (
	Merchant struct {
		gorm.Model
		Name    string `gorm:"not null"`
		Rating  float32
		AdminId uint `gorm:"not null"`
		User    User `gorm:"foreignKey:AdminId;constraint:OnDelete:CASCADE"`
	}
)

func CreateMerchant(db *gorm.DB, m *Merchant, ma *MerchantAddress) error {
	r := Role{}

	db.Where(&Role{UserID: m.AdminId}).First(&r)

	if r.Name == "customer" {
		return fmt.Errorf("you must have merchant role first before creating merchant, please update your profile first with [PATCH] to /user/role")
	}

	if err := db.Create(m).Error; err != nil {
		return err
	}

	ma.MerchantID = m.ID
	if err := db.Create(ma).Error; err != nil {
		return err
	}
	return nil
}

func GetAll(db *gorm.DB, filter *wmodels.MerchantFilter, m []Merchant, mout *[]wmodels.MerchantOutput) error {
	var mo []wmodels.MerchantOutput

	db.
		Table("merchants m").
		Select("m.id, m.name, m.rating, ma.city").
		Joins("left join merchant_addresses ma on ma.merchant_id = m.id").
		Find(&mo)

	for _, m := range mo {
		if !strings.Contains(strings.ToLower(m.Name), strings.ToLower(*filter.Name)) {
			continue
		}
		if *filter.City != "" && !strings.EqualFold(*filter.City, m.City) {
			continue
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
		*mout = append(*mout, m)
	}
	return nil
}

func (m *Merchant) GetMy(db *gorm.DB, data *wmodels.MerchantDetailsOutput) error {
	if err := db.
		Model(&Merchant{}).
		Where("admin_id = ?", m.AdminId).
		Find(m).
		Error; err != nil {
		return fmt.Errorf("you have not a merchant yet")
	}

	log.Printf("get merchant info: %v\n", *m)

	if err := m.GetById(db, data); err != nil {
		return err
	}

	return nil
}

func (m *Merchant) GetById(db *gorm.DB, data *wmodels.MerchantDetailsOutput) error {
	row, err := db.
		Table("merchants m").
		Select("m.id, m.name, m.rating, ma.city, ma.offline_store_address, c.name").
		Joins("left join merchant_addresses ma on ma.merchant_id = m.id left join countries c on ma.country_id = c.id").
		Where("m.id = ?", m.ID).
		Rows()

	if err != nil {
		return err
	}

	mout := &wmodels.MerchantDetailsOutput{}
	maddr := &wmodels.MerchantAddrOutput{}
	if row.Next() {
		err := row.Scan(&mout.ID, &mout.Name, &mout.Rating, &maddr.City, &maddr.AddressLine, &maddr.Country)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("kami tidak memiliki data merchant anda")
	}

	data.ID = mout.ID
	data.Name = mout.Name
	data.Rating = mout.Rating

	data.Address = *maddr

	return nil
}

func (m *Merchant) Delete(db *gorm.DB) error {
	if err := db.Unscoped().Delete(m).Error; err != nil {
		return err
	}

	if err := db.
		Model(&Role{}).
		Where("user_id = ?", m.AdminId).
		Updates(&Role{Name: "customer"}).
		Error; err != nil {
		return err
	}

	return nil
}

func (m *Merchant) Put(db *gorm.DB, addr *MerchantAddress) error {

	if err := db.Model(m).Updates(m).Error; err != nil {
		return err
	}

	err := db.
		Model(&MerchantAddress{}).
		Where("merchant_id = ?", addr.MerchantID).
		Updates(addr).
		Error

	if err != nil {
		return err
	}

	return nil
}
