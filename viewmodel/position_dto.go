package viewmodel

import "cv_backend/model"

type PositionDTO struct {
	BaseDTO
	Position string `json:"position"`
	PersonID int64  `json:"person_id"`
}

func ToPositionDTO(p *model.Position) *PositionDTO {
	if p == nil {
		return nil
	}

	return &PositionDTO{
		BaseDTO: BaseDTO{
			ID:        p.ID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
		Position: p.Position,
		PersonID: p.PersonID,
	}
}
func (d *PositionDTO) ToModel() *model.Position {
	if d == nil {
		return nil
	}

	return &model.Position{
		BaseModel: model.BaseModel{ID: d.ID},
		Position:  d.Position,
		PersonID:  d.PersonID,
	}
}
