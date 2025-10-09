package model

import "github.com/google/uuid"

type PersonStatusType int

const (
	PersonDurumBeklemede PersonStatusType = iota
	PersonDurumOnaylandi
	PersonDurumReddedildi
	PersonDurumIlgileniliyor
)

type Person struct {
	BaseModel
	Name                      string           `json:"name"`
	Surname                   string           `json:"surname"`
	PhoneNumber               string           `json:"phone_number"`
	Email                     string           `json:"email"`
	EducationStatus           string           `json:"education_status"`
	StudyingDepartmen         string           `json:"studying_departmen"`
	UUID                      uuid.UUID        `json:"uuid"`
	Positions                 []Position       `json:"positions"`
	Languages                 []Language       `json:"languages"`
	References                []Reference      `json:"references"`
	StatusType                PersonStatusType `json:"status_type"`
	Reviewer                  User             `json:"reviewer"`
	ReviewerID                *int64           `json:"reviewer_id" `
	ReasonForRejection        string           `json:"reason_for_rejection"`
	ReasonForRejectionSummary string           `json:"reason_for_rejection_summary"`
}
