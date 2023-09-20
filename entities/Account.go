package entities

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model `json:"-"`
	ID         uint `db:"id" json:"id" bson:"id" gorm:"primarykey, AUTO_INCREMENT"`
	// Username  string          `json:"username" bson:"username"`
	Firstname string `db:"firstname" json:"firstname" bson:"firstname"`
	Surname   string `db:"surname" json:"surname" bson:"surname"`
	Email     string `db:"email" json:"email" bson:"email"`
	Mobile    string `db:"mobile" json:"mobile" bson:"mobile"`
	Role      string `db:"role" json:"role" bson:"Role"`
	IsActive  bool   `db:"is_active" json:"is_active" bson:"is_active"`
	// CreatedAt time.Time `json:"created_at" bson:"createdAt"`
	// UpdatedAt time.Time `json:"updated_at" bson:"updatedAt"`
	// DeletedAt time.Time `json:"deleted_at"`
}

func (Account) TableName() string {
	return "account"
}
