package service

import (
	"cv_backend/app/repository"
	"cv_backend/model"
)

type PersonService struct {
	repo *repository.PersonRepository
}

func NewPersonService(repo *repository.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) GetPersonByID(id int64) (*model.Person, error) {
	return s.repo.GetByID(id)
}

func (s *PersonService) GetPersonsPaginated(status string, page, limit int) ([]model.Person, int64, error) {
	var persons []model.Person
	var total int64
	offset := (page - 1) * limit

	query := s.repo.DB.Preload("Positions").Preload("Languages").Preload("References")

	if status != "" {
		switch status {
		case "beklemede":
			query = query.Where("status_type = ?", model.PersonDurumBeklemede)
		case "onaylandi":
			query = query.Where("status_type = ?", model.PersonDurumOnaylandi)
		case "reddedildi":
			query = query.Where("status_type = ?", model.PersonDurumReddedildi)
		case "ilgileniliyor":
			query = query.Where("status_type = ?", model.PersonDurumIlgileniliyor)
		default:
			return []model.Person{}, 0, nil
		}
	}

	if err := query.Model(&model.Person{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Find(&persons).Error; err != nil {
		return nil, 0, err
	}

	return persons, total, nil
}

func (s *PersonService) CreatePerson(person *model.Person) (*model.Person, error) {
	err := s.repo.Create(person)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (s *PersonService) UpdatePersonStatus(person *model.Person) error {
	return s.repo.Update(person)
}

func (s *PersonService) DeletePerson(id int64) error {
	person, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if person == nil {
		return nil
	}

	tx := s.repo.DB.Begin()

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
