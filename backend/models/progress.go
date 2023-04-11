package models

type Progress struct {
	UserID            uint    `json:"userId"`
	TotalSpent        float32 `json:"totalSpent"`
	TransactionIDList []uint  `json:"transactionIdList"`
	BudgetID      		uint  `json:"budgetId"`
	Category          string  `json:"category"`
	BudgetGoal        float32 `json:"budgetGoal"`
	Frequency         Period  `json:"frequency"`
}

type GetProgressResponse struct {
	Data []Progress `json:"data"`
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
