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

func (r *UserRepository) Create(u *modelSql.User) error {
	return r.db.Create(u).Error
}

func (r *UserRepository) UpdateById(id uint, u *modelSql.User) (*modelSql.User, error) {
	_u := &modelSql.User{ID: id}
	//TODO check whether Updates accept pointer or not
	res := r.db.Model(_u).Updates(*u)

	return _u, res.Error
}

func (r *UserRepository) FindById(id uint) (*modelSql.User, error) {
	_u := &modelSql.User{ID: id}
	res := r.db.First(_u)

	return _u, res.Error
}

func (r *UserRepository) FindAll() ([]modelSql.User, error) {
	var _us []modelSql.User
	res := r.db.Find(&_us)

	return _us, res.Error
}

func (r *UserRepository) DeleteById(id uint) error {
	_u := &modelSql.User{ID: id}

	return r.db.Delete(_u).Error
}
