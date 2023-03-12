package models

type Period string

const (
	Weekly  Period = "weekly"
	Yearly  Period = "yearly"
	Monthly Period = "monthly"
)

type BudgetContent struct {
	Category	      string	 	`json:"category"`
	AmountLimit     float32  	`json:"amountLimit"`
	Frequency				Period		`json:"frequency"`
	CycleDuration		uint			`json:"duration"`
	CycleCount			uint			`json:"count"`
}

type Budget struct {
	Data 		 BudgetContent	`json:"data" gorm:"embedded"`
	UserID	 uint						`json:"userId"`
	BudgetID uint						`json:"budgetId" gorm:"primaryKey;unique;column:budgetId"`
}

type CreateBudgetResponse struct {
	UserID uint `json:"userId"`
	BudgetID uint `json:"budgetId"`
}

type BudgetsResponse struct {
	Budgets []Budget `json:"budgets"`
}

type UpdateBudgetRequest struct {
	NewBudget Budget `json:"newBudget"`
}
