package services

import (
	sqlModel "anya-day/models/sql"
	repo "anya-day/repository"
	"fmt"
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

	if len(userCred.Password) < 8 {
		return fmt.Errorf("password too short (at least 8 characters)")
	}

	if err := s.UserCredRepo.Create(userCred); err != nil {
		return err
	}

	return nil
}
