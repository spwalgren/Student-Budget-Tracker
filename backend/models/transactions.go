package models

type Transaction struct {
	UserID      uint    `json:"userId"`
	TransactionID	uint	`json:"transactionId" gorm:"primaryKey;unique;column:transactionId"`
	Amount      float32 `json:"amount"`
	Name        string  `json:"name"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}
