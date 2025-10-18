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

func (r *PersonRepository) GetAll() ([]model.Person, error) {
	var persons []model.Person
	err := r.DB.Preload("Positions").Preload("Languages").Preload("References").Find(&persons).Error
	return persons, err
}

func (r *PersonRepository) GetByID(id int64) (*model.Person, error) {
	var person model.Person
	err := r.DB.Preload("Positions").Preload("Languages").Preload("References").First(&person, id).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &person, err
}

func (r *PersonRepository) Update(person *model.Person) error {
	return r.DB.Save(person).Error
}

func (r *PersonRepository) Delete(id int64) error {
	return r.DB.Delete(&model.Person{}, id).Error
}
