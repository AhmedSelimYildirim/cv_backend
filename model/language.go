package model

type Language struct {
	BaseModel
	LanguageName  string `json:"language_name"`
	LanguageLevel string `json:"language_level"`
	PersonId      int64  `json:"person_id"`
	Person        Person `json:"person"`
}
