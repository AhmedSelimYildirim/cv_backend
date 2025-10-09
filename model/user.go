package model

type User struct {
	BaseModel
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email" gorm:"index"`
	Password string `json:"password"`
}

func (User) ModelName() string {
	return "user"
}

func (k User) String() string {
	return k.Name + " " + k.Surname
}
