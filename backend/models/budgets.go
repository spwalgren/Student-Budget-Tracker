package models

type Period string

const (
	Weekly  Period = "weekly"
	Yearly  Period = "yearly"
	Monthly Period = "monthly"
)

type BudgetContent struct {
	Category      string  `json:"category"`
	AmountLimit   float32 `json:"amountLimit"`
	Frequency     Period  `json:"frequency"`
	CycleDuration uint    `json:"duration"`
	CycleCount    uint    `json:"count"`
	StartDate     string  `json:"startDate"`
}

type Budget struct {
	Data      BudgetContent `json:"data" gorm:"embedded"`
	UserID    uint          `json:"userId"`
	BudgetID  uint          `json:"budgetId" gorm:"primaryKey;unique;column:budgetId"`
	IsDeleted bool					`json:"isDeleted" gorm:"column:isDeleted;type:boolean"`
}

type CreateBudgetResponse struct {
	UserID   uint `json:"userId"`
	BudgetID uint `json:"budgetId"`
}

type BudgetsResponse struct {
	Budgets []Budget `json:"budgets"`
}

type DeletedBudgetsResponse struct {
	DeletedBudgets []Budget `json:"deletedBudgets"`
}

type UpdateBudgetRequest struct {
	NewBudget Budget `json:"newBudget"`
}
