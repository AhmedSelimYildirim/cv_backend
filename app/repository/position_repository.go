package repository

import (
	"cv_backend/model"
	"errors"

	"gorm.io/gorm"
)

type PositionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) *PositionRepository {
	return &PositionRepository{db: db}
}

func (r *PositionRepository) GetAll() ([]model.Position, error) {
	var positions []model.Position
	err := r.db.Preload("Person").Find(&positions).Error
	return positions, err
}

func (r *PositionRepository) GetByID(id int64) (*model.Position, error) {
	var position model.Position
	err := r.db.Preload("Person").First(&position, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &position, err
}

func (r *PositionRepository) Delete(id int64) error {
	result := r.db.Delete(&model.Position{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
