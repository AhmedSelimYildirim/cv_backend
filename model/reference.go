package model

type Reference struct {
	BaseModel
	ReferenceName   string `json:"reference_name"`
	ReferenceNumber string `json:"reference_number"`
	PersonId        int64  `json:"person_id"`
	Person          Person `json:"person"`
}
