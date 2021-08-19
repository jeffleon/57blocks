package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Email    string `gorm:"unique" json:"email" validate:"required,email,min=6,max=32"`
	Password []byte `json:"-" validate:"min=8"`
}
