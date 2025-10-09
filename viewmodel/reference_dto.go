package viewmodel

import "cv_backend/model"

type ReferenceDTO struct {
	BaseDTO
	ReferenceName   string `json:"reference_name"`
	ReferenceNumber string `json:"reference_number"`
	PersonID        int64  `json:"person_id"`
}

func ToReferenceDTO(r *model.Reference) *ReferenceDTO {
	if r == nil {
		return nil
	}

	return &ReferenceDTO{
		BaseDTO: BaseDTO{
			ID:        r.ID,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		},
		ReferenceName:   r.ReferenceName,
		ReferenceNumber: r.ReferenceNumber,
		PersonID:        r.PersonId,
	}
}
func (d *ReferenceDTO) ToModel() *model.Reference {
	if d == nil {
		return nil
	}

	return &model.Reference{
		BaseModel:       model.BaseModel{ID: d.ID},
		ReferenceName:   d.ReferenceName,
		ReferenceNumber: d.ReferenceNumber,
		PersonId:        d.PersonID,
	}
}
