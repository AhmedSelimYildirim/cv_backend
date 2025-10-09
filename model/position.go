package model

type Position struct {
	BaseModel
	Position string `json:"position"`
	PersonID int64  `json:"person_id"`
	Person   Person `json:"person"`
}
