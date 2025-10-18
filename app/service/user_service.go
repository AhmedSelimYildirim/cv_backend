package service

import (
	"cv_backend/app/repository"
	"cv_backend/model"
	"cv_backend/utils"
	"errors"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *model.User) error {
	exists, _ := s.repo.GetByEmail(user.Email)
	if exists != nil && exists.ID != 0 {
		return errors.New("email already registered")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func (s *UserService) GetByID(id int64) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}
