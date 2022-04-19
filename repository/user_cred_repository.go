package repository

import (
	modelSql "anya-day/models/sql"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(d *gorm.DB) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) Create(uc *modelSql.UserCredential) error {
	return r.db.Create(uc).Error
}

func (r *UserRepository) UpdateById(id uint, uc *modelSql.UserCredential) (*modelSql.UserCredential, error) {
	_uc := &modelSql.UserCredential{CredentialID: id}
	res := r.db.Model(_uc).Updates(*uc)

	return _uc, res.Error
}

func (r *UserRepository) FindById(id uint) (*modelSql.UserCredential, error) {
	_uc := &modelSql.UserCredential{CredentialID: id}
	res := r.db.First(_uc)

	return _uc, res.Error
}

func (r *UserRepository) DeleteById(id uint) error {
	_uc := &modelSql.UserCredential{CredentialID: id}

	return r.db.Delete(_uc).Error
}
