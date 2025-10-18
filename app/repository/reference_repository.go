package repository

import (
	"cv_backend/model"
	"errors"

	"gorm.io/gorm"
)

type ReferenceRepository struct {
	db *gorm.DB
}

func NewReferenceRepository(db *gorm.DB) *ReferenceRepository {
	return &ReferenceRepository{db: db}
}

func (r *ReferenceRepository) GetAll() ([]model.Reference, error) {
	var references []model.Reference
	err := r.db.Preload("Person").Find(&references).Error
	return references, err
}

func (r *ReferenceRepository) GetByID(id int64) (*model.Reference, error) {
	var reference model.Reference
	err := r.db.Preload("Person").First(&reference, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &reference, err
}

func (r *ReferenceRepository) Delete(id int64) error {
	return r.db.Delete(&model.Reference{}, id).Error
}
