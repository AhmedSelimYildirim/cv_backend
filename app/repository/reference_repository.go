package repository

import (
	"cv_backend/model"

	"gorm.io/gorm"
)

type ReferenceRepository struct {
	db *gorm.DB
}

func NewReferenceRepository(db *gorm.DB) *ReferenceRepository {
	return &ReferenceRepository{db: db}
}

func (r *ReferenceRepository) Create(reference *model.Reference) error {
	return r.db.Create(reference).Error
}

func (r *ReferenceRepository) GetAll() ([]model.Reference, error) {
	var references []model.Reference
	err := r.db.Preload("Person").Find(&references).Error
	return references, err
}

func (r *ReferenceRepository) GetByID(id int64) (*model.Reference, error) {
	var reference model.Reference
	err := r.db.Preload("Person").First(&reference, id).Error
	if err != nil {
		return nil, err
	}
	return &reference, nil
}

func (r *ReferenceRepository) Update(reference *model.Reference) error {
	return r.db.Save(reference).Error
}

func (r *ReferenceRepository) Delete(id int64) error {
	return r.db.Delete(&model.Reference{}, id).Error
}
