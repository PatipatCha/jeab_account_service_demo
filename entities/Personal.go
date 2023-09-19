package entities

import (
	"gorm.io/gorm"
)

type Personal struct {
	gorm.Model
	IDCard    string `json:"id_card" bson:"id_card"`
	Firstname string `json:"firstname" bson:"firstname"`
	Surname   string `json:"surname" bson:"surname"`
	Email     string `json:"email" bson:"email"`
	Mobile    string `json:"mobile" bson:"mobile"`
}

func (Personal) TableName() string {
	return "personal"
}
