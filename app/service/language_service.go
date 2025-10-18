package service

import (
	"cv_backend/app/repository"
	"cv_backend/model"
)

type LanguageService struct {
	repo *repository.LanguageRepository
}

func NewLanguageService(repo *repository.LanguageRepository) *LanguageService {
	return &LanguageService{repo: repo}
}

func (s *LanguageService) GetAll() ([]model.Language, error) {
	return s.repo.GetAll()
}

func (s *LanguageService) GetByID(id int64) (*model.Language, error) {
	return s.repo.GetByID(id)
}

func (s *LanguageService) Delete(id int64) error {
	return s.repo.Delete(id)
}
