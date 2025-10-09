package viewmodel

import "cv_backend/model"

type LanguageDTO struct {
	BaseDTO
	LanguageName  string `json:"language_name"`
	LanguageLevel string `json:"language_level"`
	PersonID      int64  `json:"person_id"`
}

func ToLanguageDTO(l *model.Language) *LanguageDTO {
	if l == nil {
		return nil
	}

	return &LanguageDTO{
		BaseDTO: BaseDTO{
			ID:        l.ID,
			CreatedAt: l.CreatedAt,
			UpdatedAt: l.UpdatedAt,
		},
		LanguageName:  l.LanguageName,
		LanguageLevel: l.LanguageLevel,
		PersonID:      l.PersonId,
	}
}

func (d *LanguageDTO) ToModel() *model.Language {
	if d == nil {
		return nil
	}

	return &model.Language{
		BaseModel:     model.BaseModel{ID: d.ID},
		LanguageName:  d.LanguageName,
		LanguageLevel: d.LanguageLevel,
		PersonId:      d.PersonID,
	}
}
