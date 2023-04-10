package models

type Transaction struct {
	UserID        uint    `json:"userId"`
	TransactionID uint    `json:"transactionId" gorm:"primaryKey;unique;column:transactionId"`
	Amount        float32 `json:"amount"`
	Name          string  `json:"name"`
	Date          string  `json:"date"`
	Category      string  `json:"category"`
	Description   string  `json:"description"`
	CycleIndex    int     `json:"cycleIndex"`
}

type CreateTransactionResponse struct {
	UserID        uint `json:"userId"`
	TransactionID uint `json:"transactionId"`
}

type UpdateTransactionRequest struct {
	Data Transaction `json:"data"`
}

type CreateTransactionRequest struct {
	Data Transaction `json:"data"`
}

type TransactionsResponse struct {
	Data []Transaction `json:"data"`
}