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
	CycleDuration uint    `json:"duration"` /* How many repeats of selected duration */
	CycleCount    uint    `json:"count"`    /* How many repeats of created cycle */
	StartDate     string  `json:"startDate"`
}

type Budget struct {
	Data      BudgetContent `json:"data" gorm:"embedded"`
	UserID    uint          `json:"userId"`
	BudgetID  uint          `json:"budgetId" gorm:"primaryKey;unique;column:budgetId"`
	IsDeleted bool          `json:"isDeleted" gorm:"column:isDeleted;type:boolean"`
	CurrentPeriodStart string			`json:"currentPeriodStart"`
	CurrentPeriodEnd		string			`json:"currentPeriodEnd"`
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

type BudgetCategoriesResponse struct {
	Category []string `json:"categories"`
}

type Cycle struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Index    int    `json:"index"`
	BudgetID uint   `json:"budgetId"`
}

type CyclePeriodResponse struct {
	Data []Cycle `json:"data"`
}
