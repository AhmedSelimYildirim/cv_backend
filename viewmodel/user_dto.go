package viewmodel

import "cv_backend/model"

type UserDTO struct {
	BaseDTO
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

func ToUserDTO(u *model.User) *UserDTO {
	if u == nil {
		return nil
	}

	return &UserDTO{
		BaseDTO: BaseDTO{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Name:    u.Name,
		Surname: u.Surname,
		Email:   u.Email,
	}
}
func (d *UserDTO) ToModel() *model.User {
	if d == nil {
		return nil
	}

	return &model.User{
		BaseModel: model.BaseModel{ID: d.ID},
		Name:      d.Name,
		Surname:   d.Surname,
		Email:     d.Email,
	}
}
