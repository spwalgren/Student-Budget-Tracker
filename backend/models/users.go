package models

type UserInfo struct {
	ID        uint   `json:"ID" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReturnInfo struct {
	ID string `json:"id"`
}
