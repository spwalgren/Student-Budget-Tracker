package models

type BudgetTransaction struct {
	BudgetTransactionID uint `json:"budgetTransactionId" gorm:"primaryKey;unique"`
	BudgetID            uint `json:"budgetId"`
	TransactionID       uint `json:"transactionId"`
	CycleIndex          int  `json:"cycleIndex"`
}

type Transaction struct {
	UserID        uint    `json:"userId"`
	TransactionID uint    `json:"transactionId" gorm:"primaryKey;unique;column:transactionId"`
	Amount        float32 `json:"amount"`
	Name          string  `json:"name"`
	Date          string  `json:"date"`
	Category      string  `json:"category"`
	Description   string  `json:"description"`
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