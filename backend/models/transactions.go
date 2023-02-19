package models

type Transaction struct {
	UserID      uint    `json:"userId"`
	Amount      float32 `json:"amount"`
	Name        string  `json:"name"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}
