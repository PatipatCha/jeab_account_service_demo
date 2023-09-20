package entities

type UpdateAccountRequest struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname" bson:"firstname"`
	Surname   string `json:"surname" bson:"surname"`
	Email     string `json:"email" bson:"email"`
	Mobile    string `json:"mobile" bson:"mobile"`
	Role      string `json:"role" bson:"role"`
	IsActive  bool   `json:"is_active" bson:"is_active"`
}

// func (UpdateAccountRequest) TableName() string {
// 	return "account"
// }
