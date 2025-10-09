package viewmodel

import (
	"cv_backend/model"
)

type PersonDTO struct {
	BaseDTO
	Name                      string                 `json:"name"`
	Surname                   string                 `json:"surname"`
	PhoneNumber               string                 `json:"phone_number"`
	Email                     string                 `json:"email"`
	EducationStatus           string                 `json:"education_status"`
	StudyingDepartmen         string                 `json:"studying_departmen"`
	StatusType                model.PersonStatusType `json:"status_type"`
	ReasonForRejection        string                 `json:"reason_for_rejection"`
	ReasonForRejectionSummary string                 `json:"reason_for_rejection_summary"`
	Reviewer                  *UserDTO               `json:"reviewer,omitempty"`

	Positions  []PositionDTO  `json:"positions,omitempty"`
	Languages  []LanguageDTO  `json:"languages,omitempty"`
	References []ReferenceDTO `json:"references,omitempty"`
}

// Model -> DTO
func ToPersonDTO(p *model.Person) *PersonDTO {
	if p == nil {
		return nil
	}

	dto := &PersonDTO{
		BaseDTO: BaseDTO{
			ID:        p.ID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
		Name:                      p.Name,
		Surname:                   p.Surname,
		PhoneNumber:               p.PhoneNumber,
		Email:                     p.Email,
		EducationStatus:           p.EducationStatus,
		StudyingDepartmen:         p.StudyingDepartmen,
		StatusType:                p.StatusType,
		ReasonForRejection:        p.ReasonForRejection,
		ReasonForRejectionSummary: p.ReasonForRejectionSummary,
		Reviewer:                  ToUserDTO(&p.Reviewer),
	}

	for _, pos := range p.Positions {
		dto.Positions = append(dto.Positions, *ToPositionDTO(&pos))
	}
	for _, lang := range p.Languages {
		dto.Languages = append(dto.Languages, *ToLanguageDTO(&lang))
	}
	for _, ref := range p.References {
		dto.References = append(dto.References, *ToReferenceDTO(&ref))
	}

	return dto
}

// DTO -> Model (handler'larda kullanılıyor)
func (d *PersonDTO) ToModel() *model.Person {
	if d == nil {
		return nil
	}

	person := &model.Person{
		BaseModel: model.BaseModel{
			ID: d.ID,
		},
		Name:                      d.Name,
		Surname:                   d.Surname,
		PhoneNumber:               d.PhoneNumber,
		Email:                     d.Email,
		EducationStatus:           d.EducationStatus,
		StudyingDepartmen:         d.StudyingDepartmen,
		StatusType:                d.StatusType,
		ReasonForRejection:        d.ReasonForRejection,
		ReasonForRejectionSummary: d.ReasonForRejectionSummary,
	}

	for _, pos := range d.Positions {
		person.Positions = append(person.Positions, *pos.ToModel())
	}
	for _, lang := range d.Languages {
		person.Languages = append(person.Languages, *lang.ToModel())
	}
	for _, ref := range d.References {
		person.References = append(person.References, *ref.ToModel())
	}

	if d.Reviewer != nil {
		person.Reviewer = *d.Reviewer.ToModel()
	}

	return person
}
