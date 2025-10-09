package service

import (
	"cv_backend/app/repository"
	"cv_backend/model"
)

type PositionService struct {
	repo *repository.PositionRepository
}

func NewPositionService(repo *repository.PositionRepository) *PositionService {
	return &PositionService{repo: repo}
}

func (s *PositionService) CreatePosition(position *model.Position) error {
	return s.repo.Create(position)
}

func (s *PositionService) GetAllPositions() ([]model.Position, error) {
	return s.repo.GetAll()
}

func (s *PositionService) GetPositionByID(id int64) (*model.Position, error) {
	return s.repo.GetByID(id)
}

func (s *PositionService) UpdatePosition(position *model.Position) error {
	return s.repo.Update(position)
}

func (s *PositionService) DeletePosition(id int64) error {
	return s.repo.Delete(id)
}
