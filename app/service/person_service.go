package service

import (
	"cv_backend/app/repository"
	"cv_backend/model"
	"github.com/google/uuid"
)

type PersonService struct {
	Repo *repository.PersonRepository
}

func NewPersonService(repo *repository.PersonRepository) *PersonService {
	return &PersonService{Repo: repo}
}

func (s *PersonService) CreatePerson(person *model.Person) error {
	person.UUID = uuid.New()
	return s.Repo.Create(person)
}

func (s *PersonService) GetPersonByID(id int64) (*model.Person, error) {
	return s.Repo.GetByID(id)
}

func (s *PersonService) UpdatePerson(person *model.Person) error {
	return s.Repo.Update(person)
}

func (s *PersonService) DeletePerson(id int64) error {
	return s.Repo.Delete(id)
}

// Pagination ve Status filtreleme
func (s *PersonService) GetPersonsPaginated(status string, page, limit int) ([]model.Person, int64, error) {
	var persons []model.Person
	var total int64
	offset := (page - 1) * limit

	query := s.Repo.DB.Preload("Positions").Preload("Languages").Preload("References")

	// Status filtreleme
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

	// Toplam kayıt
	if err := query.Model(&model.Person{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	if err := query.Offset(offset).Limit(limit).Find(&persons).Error; err != nil {
		return nil, 0, err
	}

	return persons, total, nil
}
