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
	offset := (page - 1) * limit
	return s.repo.GetAllPaginated(status, offset, limit)
}

func (s *PersonService) CreatePerson(person *model.Person) (*model.Person, error) {
	if err := s.repo.Create(person); err != nil {
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
	return s.repo.DeleteWithRelations(id)
}
