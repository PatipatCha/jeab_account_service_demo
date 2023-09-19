package entities

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	// ID        ID        `json:"id"`
	// CreatedAt time.Time `json:"created_at" bson:"createdAt"`
	// UpdatedAt time.Time `json:"updated_at" bson:"updatedAt"`
	// DeletedAt time.Time `json:"deleted_at"`
	// Username  string          `json:"username" bson:"username"`
	Firstname string `json:"firstname" bson:"firstname"`
	Surname   string `json:"surname" bson:"surname"`
	Email     string `json:"email" bson:"email"`
	Mobile    string `json:"mobile" bson:"mobile"`
	Role      string `json:"role" bson:"role"`
	Status    bool   `gorm:is_active json:"is_active" bson:"is_active"`
}

func (Account) TableName() string {
	return "account"
}
