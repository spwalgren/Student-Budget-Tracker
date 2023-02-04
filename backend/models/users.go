package models

type UserInfo struct {
	ID		  uint   `json:"ID" gorm:"primary_key"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReturnLoginInfo struct {
	ID	string `json:"id"`
}
