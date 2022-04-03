package models

import (
	wmodels "anya-day/models/web"
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
	}

	data.ID = mout.ID
	data.Name = mout.Name
	data.Rating = mout.Rating

	data.Address = *maddr

	return nil
}

func (m *Merchant) Delete(db *gorm.DB) error {
	return db.Unscoped().Delete(m).Error

}
