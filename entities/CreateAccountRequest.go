package entities

type CreateAccountRequest struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Surname   string `json:"surname" bson:"surname"`
	Email     string `json:"email" bson:"email"`
	Mobile    string `json:"mobile" bson:"mobile"`
	Role      string `json:"role" bson:"role"`
	Status    bool   `json:"is_active" bson:"is_active"`
}

func (CreateAccountRequest) TableName() string {
	return "account"
}
