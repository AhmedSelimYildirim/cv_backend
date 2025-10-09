package repository

import (
	"cv_backend/model"

	"gorm.io/gorm"
)

type LanguageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{db: db}
}

func (r *LanguageRepository) Create(language *model.Language) error {
	return r.db.Create(language).Error
}

func (r *LanguageRepository) GetAll() ([]model.Language, error) {
	var languages []model.Language
	err := r.db.Find(&languages).Error
	return languages, err
}

func (r *LanguageRepository) GetByID(id int64) (*model.Language, error) {
	var language model.Language
	err := r.db.First(&language, id).Error
	if err != nil {
		return nil, err
	}
	return &language, nil
}

func (r *LanguageRepository) Update(language *model.Language) error {
	return r.db.Save(language).Error
}

func (r *LanguageRepository) Delete(id int64) error {
	return r.db.Delete(&model.Language{}, id).Error
}
