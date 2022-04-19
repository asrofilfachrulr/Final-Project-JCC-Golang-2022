package repository

import (
	modelSql "anya-day/models/sql"

	"gorm.io/gorm"
)

type UserCredentialRepository struct {
	db *gorm.DB
}

func NewUserCredRepo(d *gorm.DB) *UserCredentialRepository {
	return &UserCredentialRepository{
		db: d,
	}
}

func (r *UserCredentialRepository) Create(uc *modelSql.UserCredential) error {
	return r.db.Create(uc).Error
}

func (r *UserCredentialRepository) UpdateById(id uint, uc *modelSql.UserCredential) (*modelSql.UserCredential, error) {
	_uc := &modelSql.UserCredential{CredentialID: id}
	res := r.db.Model(_uc).Updates(*uc)

	return _uc, res.Error
}

func (r *UserCredentialRepository) FindById(id uint) (*modelSql.UserCredential, error) {
	_uc := &modelSql.UserCredential{CredentialID: id}
	res := r.db.First(_uc)

	return _uc, res.Error
}

func (r *UserCredentialRepository) DeleteById(id uint) error {
	_uc := &modelSql.UserCredential{CredentialID: id}

	return r.db.Delete(_uc).Error
}
