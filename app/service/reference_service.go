package service

import (
	"cv_backend/app/repository"
	"cv_backend/model"
)

type ReferenceService struct {
	repo *repository.ReferenceRepository
}

func NewReferenceService(repo *repository.ReferenceRepository) *ReferenceService {
	return &ReferenceService{repo: repo}
}

func (s *ReferenceService) CreateReference(reference *model.Reference) error {
	return s.repo.Create(reference)
}

func (s *ReferenceService) GetAllReferences() ([]model.Reference, error) {
	return s.repo.GetAll()
}

func (s *ReferenceService) GetReferenceByID(id int64) (*model.Reference, error) {
	return s.repo.GetByID(id)
}

func (s *ReferenceService) UpdateReference(reference *model.Reference) error {
	return s.repo.Update(reference)
}

func (s *ReferenceService) DeleteReference(id int64) error {
	return s.repo.Delete(id)
}
