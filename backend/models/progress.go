package models

type ProgressByPeriod struct {
	UserID						uint					`json:"userId"`
	TotalSpent        float32       `json:"totalSpent"`
	TransactionIDList []uint 			`json:"transactionIdList"`
	BudgetIDList      []uint 			`json:"budgetIdList"`
	Category					string				`json:"category"`
	BudgetGoal				float32					`json:"budgetGoal"`
	Frequency					Period				`json:"frequency"`
}

type GetProgressRequest struct {
	Frequency	Period	`json:"frequency"`
}

type GetProgressResponse struct {
	Data ProgressByPeriod `json:"data"`
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
