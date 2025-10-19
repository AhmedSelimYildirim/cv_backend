package repository

import (
	"cv_backend/model"

	"gorm.io/gorm"
)

type PersonRepository struct {
	DB *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{DB: db}
}

func (r *PersonRepository) GetAllPaginated(status string, offset, limit int) ([]model.Person, int64, error) {
	var persons []model.Person
	var total int64

	query := r.DB.Model(&model.Person{}).
		Preload("Positions").
		Preload("Languages").
		Preload("References")

	if status != "" {
		query = query.Where("status_type = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Find(&persons).Error; err != nil {
		return nil, 0, err
	}

	return persons, total, nil
}

func (r *PersonRepository) GetByID(id int64) (*model.Person, error) {
	var person model.Person
	err := r.DB.Preload("Positions").
		Preload("Languages").
		Preload("References").
		First(&person, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &person, err
}

func (r *PersonRepository) Create(person *model.Person) error {
	return r.DB.Create(person).Error
}

func (r *PersonRepository) Update(person *model.Person) error {
	return r.DB.Save(person).Error
}

func (r *PersonRepository) DeleteWithRelations(id int64) error {
	tx := r.DB.Begin()

	if err := tx.Where("person_id = ?", id).Delete(&model.Language{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("person_id = ?", id).Delete(&model.Position{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("person_id = ?", id).Delete(&model.Reference{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&model.Person{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
