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
	// UUID oluştur
	person.UUID = uuid.New()
	return s.Repo.Create(person)
}

func (s *PersonService) GetAllPersons() ([]model.Person, error) {
	return s.Repo.GetAll()
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
