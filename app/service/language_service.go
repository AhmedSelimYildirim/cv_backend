package service

import (
	"cv_backend/app/repository"
	"cv_backend/model"

	"gorm.io/gorm"
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
	language, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if language == nil {
		return gorm.ErrRecordNotFound
	}
	return s.repo.Delete(id)
}
