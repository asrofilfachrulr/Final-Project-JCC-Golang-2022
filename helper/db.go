package helper

import "gorm.io/gorm"

func RollbackIfErr(tx *gorm.DB) {
	if s := recover(); s != nil {
		tx.Rollback()
	}
}
