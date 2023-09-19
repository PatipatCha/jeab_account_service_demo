package entities

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	// ID        ID        `json:"id"`
	Firstname string `json:"firstname" bson:"firstname"`
	Surname   string `json:"surname" bson:"surname"`
	// CreatedAt time.Time `json:"created_at" bson:"createdAt"`
	// UpdatedAt time.Time `json:"updated_at" bson:"updatedAt"`
	// DeletedAt time.Time `json:"deleted_at"`
}

func (Account) TableName() string {
	return "account"
}
