package services

import (
	sqlModel "anya-day/models/sql"
	repo "anya-day/repository"
)

type UserServices struct {
	UserRepo     *repo.UserRepository
	UserCredRepo *repo.UserCredentialRepository
}

func NewUserServices(urep *repo.UserRepository, ucrep *repo.UserCredentialRepository) *UserServices {
	return &UserServices{
		UserRepo:     urep,
		UserCredRepo: ucrep,
	}
}

func (s *UserServices) CreateUser(user *sqlModel.User, userCred *sqlModel.UserCredential) error {
	if err := s.UserRepo.Create(user); err != nil {
		return err
	}

	userCred.UserID = user.ID

	if err := s.UserCredRepo.Create(userCred); err != nil {
		return err
	}

	return nil
}
