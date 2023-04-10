package models

import (
	"github.com/lib/pq"
)

type Progress struct {
	UserID			  uint			`json:"userId"`
	ProgressID        uint          `json:"progressId" gorm:"unique;primaryKey"`
	TotalSpent        float32       `json:"totalSpent"`
	TransactionIDList pq.Int32Array `json:"transactionIdList" gorm:"type:integer"`
	BudgetIDList      pq.Int32Array `json:"budgetIdList" gorm:"type:integer"`
	Data              BudgetContent `json:"data" gorm:"embedded"`
}

type CreateProgressResponse struct {
	UserID	uint	`json:"userId"`
	ProgressID uint	`json:"progressId"`
}

type ProgressResponse struct {
	ProgressData []Progress `json:"progressData"`
}

/*

CategoryName
TotalSpent
BudgetGoal
MoneyLeft (or could just do BudgetGoal - TotalSpent)
StartDate
EndDate (calculated from budget's cycle duration, count, etc.)
DaysLeft (ditto)

*/
