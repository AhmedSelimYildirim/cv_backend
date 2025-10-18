package repository

import (
	"cv_backend/model"
	"errors"

	"gorm.io/gorm"
)

type LanguageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{db: db}
}

func (r *LanguageRepository) GetAll() ([]model.Language, error) {
	var languages []model.Language
	err := r.db.Find(&languages).Error
	return languages, err
}

func (r *LanguageRepository) GetByID(id int64) (*model.Language, error) {
	var language model.Language
	err := r.db.First(&language, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &language, err
}

func (r *LanguageRepository) Delete(id int64) error {
	return r.db.Delete(&model.Language{}, id).Error
}
