package controllers

import (
	"budget-tracker/database"
	"budget-tracker/models"
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	//"github.com/dgrijalva/jwt-go"
)

func GetProgress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID, _ := strconv.ParseInt(ReturnUserID(w, r), 10, 32)
	if userID == -1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var progResponse models.GetProgressResponse
	var weeklyProgResponse models.GetProgressResponse
	var monthlyProgResponse models.GetProgressResponse
	var yearlyProgResponse models.GetProgressResponse
	var weeklyBudgets models.BudgetsResponse
	var monthlyBudgets models.BudgetsResponse
	var yearlyBudgets models.BudgetsResponse

	// GET ALL WEEKLY BUDGETS
	database.DB.Where(map[string]interface{}{"user_id": userID, "frequency": "weekly", "isDeleted": false}).Find(&weeklyBudgets.Budgets)
	// CREATE A NEW PROGRESS ENTRY FOR EACH BUDGET
	for i := 0; i < len(weeklyBudgets.Budgets); i++ {
		temp := weeklyBudgets.Budgets[i]

		// add all transactions to progress tab ONLY PULLS TRANSACTIONS WITH MATCHING CATEGORY
		var transactionResp models.TransactionsResponse
		database.DB.Where(map[string]interface{}{"user_id": userID, "category": temp.Data.Category}).Find(&transactionResp.Data)
		var idList []uint
		var totalSpent float32 = 0
		budgetTransactions , error:= IsInBudget(transactionResp.Data, temp, r)
		if error != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for j := 0; j < len(budgetTransactions.Data); j++ {
			idList = append(idList, budgetTransactions.Data[j].TransactionID)
			totalSpent += budgetTransactions.Data[j].Amount
		}
		weeklyProgResponse.Data = append(weeklyProgResponse.Data, models.Progress{UserID: temp.UserID, Frequency: temp.Data.Frequency, Category: temp.Data.Category, BudgetGoal: temp.Data.AmountLimit, BudgetID: temp.BudgetID, TransactionIDList: idList, TotalSpent: totalSpent})
	}

	// ADD ALL MONTHLY PROGRESS
	database.DB.Where(map[string]interface{}{"user_id": userID, "frequency": "monthly", "isDeleted": false}).Find(&monthlyBudgets.Budgets)
	for i := 0; i < len(monthlyBudgets.Budgets); i++ {
		tempBudget := monthlyBudgets.Budgets[i]

		var transactionResp models.TransactionsResponse
		database.DB.Where(map[string]interface{}{"user_id": userID, "category": tempBudget.Data.Category}).Find(&transactionResp.Data)
		var idList []uint
		var totalSpent float32 = 0
		budgetTransactions , error:= IsInBudget(transactionResp.Data, tempBudget, r)
		if error != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for j := 0; j < len(budgetTransactions.Data); j++ {
			idList = append(idList, budgetTransactions.Data[j].TransactionID)
			totalSpent += budgetTransactions.Data[j].Amount
		}

		monthlyProgResponse.Data = append(monthlyProgResponse.Data, models.Progress{UserID: tempBudget.UserID, Frequency: tempBudget.Data.Frequency, Category: tempBudget.Data.Category, BudgetGoal: tempBudget.Data.AmountLimit, BudgetID: tempBudget.BudgetID, TransactionIDList: idList, TotalSpent: totalSpent})
	}

	// ADD ALL YEARLY PROGRESS
	database.DB.Where(map[string]interface{}{"user_id": userID, "frequency": "yearly", "isDeleted": false}).Find(&yearlyBudgets.Budgets)
	for i := 0; i < len(yearlyBudgets.Budgets); i++ {
		tempBudget := yearlyBudgets.Budgets[i]

		var transactionResp models.TransactionsResponse
		database.DB.Where(map[string]interface{}{"user_id": userID, "category": tempBudget.Data.Category}).Find(&transactionResp.Data)
		var idList []uint
		var totalSpent float32 = 0
		budgetTransactions , error:= IsInBudget(transactionResp.Data, tempBudget, r)
		if error != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for j := 0; j < len(budgetTransactions.Data); j++ {
			idList = append(idList, budgetTransactions.Data[j].TransactionID)
			totalSpent += budgetTransactions.Data[j].Amount
		}

		yearlyProgResponse.Data = append(yearlyProgResponse.Data, models.Progress{UserID: tempBudget.UserID, Frequency: tempBudget.Data.Frequency, Category: tempBudget.Data.Category, BudgetGoal: tempBudget.Data.AmountLimit, BudgetID: tempBudget.BudgetID, TotalSpent: totalSpent, TransactionIDList: idList})
	}

	temp1 := append(weeklyProgResponse.Data, monthlyProgResponse.Data...)
	progResponse.Data = append(temp1, yearlyProgResponse.Data...)

	json.NewEncoder(w).Encode(progResponse)
	w.WriteHeader(http.StatusOK)
}

func GetPreviousProgress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	userID, _ := strconv.ParseInt(ReturnUserID(w, r), 10, 32)
	if userID == -1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var previousProgResponse models.GetProgressResponse
	var weeklyProgResponse models.GetProgressResponse
	var monthlyProgResponse models.GetProgressResponse
	var yearlyProgResponse models.GetProgressResponse
	var weeklyBudgets models.BudgetsResponse
	var monthlyBudgets models.BudgetsResponse
	var yearlyBudgets models.BudgetsResponse

	// GET ALL WEEKLY BUDGETS
	database.DB.Where(map[string]interface{}{"user_id": userID, "frequency": "weekly", "isDeleted": false}).Find(&weeklyBudgets.Budgets)
	// CREATE A NEW PROGRESS ENTRY FOR EACH BUDGET
	for i := 0; i < len(weeklyBudgets.Budgets); i++ {
		temp := weeklyBudgets.Budgets[i]

		// add all transactions to progress tab ONLY PULLS TRANSACTIONS WITH MATCHING CATEGORY
		var transactionResp models.TransactionsResponse
		database.DB.Where(map[string]interface{}{"user_id": userID, "category": temp.Data.Category}).Find(&transactionResp.Data)
		var idList []uint
		var totalSpent float32 = 0
		budgetTransactions, error:= IsInPreviousBudget(transactionResp.Data, temp, r)
		if error != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for j := 0; j < len(budgetTransactions.Data); j++ {
			idList = append(idList, budgetTransactions.Data[j].TransactionID)
			totalSpent += budgetTransactions.Data[j].Amount
		}
		weeklyProgResponse.Data = append(weeklyProgResponse.Data, models.Progress{UserID: temp.UserID, Frequency: temp.Data.Frequency, Category: temp.Data.Category, BudgetGoal: temp.Data.AmountLimit, BudgetID: temp.BudgetID, TransactionIDList: idList, TotalSpent: totalSpent})
	}

	// ADD ALL MONTHLY PROGRESS
	database.DB.Where(map[string]interface{}{"user_id": userID, "frequency": "monthly", "isDeleted": false}).Find(&monthlyBudgets.Budgets)
	for i := 0; i < len(monthlyBudgets.Budgets); i++ {
		tempBudget := monthlyBudgets.Budgets[i]

		var transactionResp models.TransactionsResponse
		database.DB.Where(map[string]interface{}{"user_id": userID, "category": tempBudget.Data.Category}).Find(&transactionResp.Data)
		var idList []uint
		var totalSpent float32 = 0
		budgetTransactions, error := IsInPreviousBudget(transactionResp.Data, tempBudget, r)
		if error != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for j := 0; j < len(budgetTransactions.Data); j++ {
			idList = append(idList, budgetTransactions.Data[j].TransactionID)
			totalSpent += budgetTransactions.Data[j].Amount
		}

		monthlyProgResponse.Data = append(monthlyProgResponse.Data, models.Progress{UserID: tempBudget.UserID, Frequency: tempBudget.Data.Frequency, Category: tempBudget.Data.Category, BudgetGoal: tempBudget.Data.AmountLimit, BudgetID: tempBudget.BudgetID, TransactionIDList: idList, TotalSpent: totalSpent})
	}

	// ADD ALL YEARLY PROGRESS
	database.DB.Where(map[string]interface{}{"user_id": userID, "frequency": "yearly", "isDeleted": false}).Find(&yearlyBudgets.Budgets)
	for i := 0; i < len(yearlyBudgets.Budgets); i++ {
		tempBudget := yearlyBudgets.Budgets[i]

		var transactionResp models.TransactionsResponse
		database.DB.Where(map[string]interface{}{"user_id": userID, "category": tempBudget.Data.Category}).Find(&transactionResp.Data)
		var idList []uint
		var totalSpent float32 = 0
		budgetTransactions , error:= IsInPreviousBudget(transactionResp.Data, tempBudget, r)
		if error != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for j := 0; j < len(budgetTransactions.Data); j++ {
			idList = append(idList, budgetTransactions.Data[j].TransactionID)
			totalSpent += budgetTransactions.Data[j].Amount
		}

		yearlyProgResponse.Data = append(yearlyProgResponse.Data, models.Progress{UserID: tempBudget.UserID, Frequency: tempBudget.Data.Frequency, Category: tempBudget.Data.Category, BudgetGoal: tempBudget.Data.AmountLimit, BudgetID: tempBudget.BudgetID, TransactionIDList: idList, TotalSpent: totalSpent})
	}

	temp1 := append(weeklyProgResponse.Data, monthlyProgResponse.Data...)
	previousProgResponse.Data = append(temp1, yearlyProgResponse.Data...)

	json.NewEncoder(w).Encode(previousProgResponse)
	w.WriteHeader(http.StatusOK)
}

func IsInPreviousBudget(transactions []models.Transaction, budget models.Budget, r *http.Request) (models.TransactionsResponse, error) {
	// setup backend request
	// Get dates of current cycle then set "current" date to be day before the start date of current cycle. Requires two calls
	reqURL := "http://localhost:8080/api/budget/dates/" + strconv.Itoa(int(budget.BudgetID)) + "/" + time.Now().Format(time.RFC3339)[:10]

	req, _ := http.NewRequest("GET", reqURL, nil)

	// Set cookie to do backend get call to retrieve start and end date
	cookie, _ := r.Cookie("jtw")
	req.AddCookie(cookie)
	resp, error := http.DefaultClient.Do(req)
	if error != nil {
		return models.TransactionsResponse{}, error
	}
	var cycleResp models.Cycle
	json.NewDecoder(resp.Body).Decode(&cycleResp)


	newTempDate, _ := time.Parse(time.RFC3339, cycleResp.Start)
	newStartDate:= newTempDate.AddDate(0, 0, -1)
	reqURL2 := "http://localhost:8080/api/budget/dates/" + strconv.Itoa(int(budget.BudgetID)) + "/" + newStartDate.String()[:10]

	req2, _ := http.NewRequest("GET", reqURL2, nil)

	// Set cookie to do backend get call to retrieve start and end date
	cookie2, _ := r.Cookie("jtw")
	req2.AddCookie(cookie2)
	resp2, error2 := http.DefaultClient.Do(req2)
	if error2 != nil {
		return models.TransactionsResponse{}, error
	}
	var cycleResp2 models.Cycle
	json.NewDecoder(resp2.Body).Decode(&cycleResp2)
	fmt.Println(cycleResp2)

	var returnTransactions models.TransactionsResponse
	for _, element := range transactions {
		transactionDate, _ := time.Parse(time.RFC3339, element.Date)
		endDate, _ := time.Parse(time.RFC3339, cycleResp2.End)
		startDate, _ := time.Parse(time.RFC3339, cycleResp2.Start)

		if (transactionDate.Before(endDate) || DateEqual(transactionDate, endDate))&& (transactionDate.After(startDate) || DateEqual(transactionDate, startDate)) {
			returnTransactions.Data = append(returnTransactions.Data, element)
		}
	}
	return returnTransactions, nil
}

// Take in an array of transactions that match the category of the budget
// Loop through the input transactions, check if transaction date is within the date of the budget
// Return transactions that fall within range
func IsInBudget(transactions []models.Transaction, budget models.Budget, r *http.Request) (models.TransactionsResponse, error){
	// setup backend request
	reqURL := "http://localhost:8080/api/budget/dates/" + strconv.Itoa(int(budget.BudgetID)) + "/" + time.Now().Format(time.RFC3339)[:10]
	req, _ := http.NewRequest("GET", reqURL, nil)

	// Set cookie to do backend get call to retrieve start and end date
	cookie, _ := r.Cookie("jtw")
	req.AddCookie(cookie)
	resp, error := http.DefaultClient.Do(req)
	if error != nil {
		return models.TransactionsResponse{}, error
	}
	var cycleResp models.Cycle
	json.NewDecoder(resp.Body).Decode(&cycleResp)

	var returnTransactions models.TransactionsResponse
	for _, element := range transactions {
		transactionDate, _ := time.Parse(time.RFC3339, element.Date)
		endDate, _ := time.Parse(time.RFC3339, cycleResp.End)
		startDate, _ := time.Parse(time.RFC3339, cycleResp.Start)

		if (transactionDate.Before(endDate) || DateEqual(transactionDate, endDate)) && (transactionDate.After(startDate) || DateEqual(transactionDate, startDate)) {
			returnTransactions.Data = append(returnTransactions.Data, element)
		}
	}
	return returnTransactions, nil
}

// Copy of GetCyclePeriod()
func HelperGetStartEndDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID, _ := strconv.ParseInt(ReturnUserID(w, r), 10, 32)
	if userID == -1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Getting budgetId and date params
	vars := mux.Vars(r)
	dateTemp := vars["date"]
	tempBudgetId, _ := strconv.Atoi(vars["budgetId"])
	budgetId := uint(tempBudgetId)
	date, _ := time.Parse("2006-01-02", dateTemp)

	// Gets budget and checks if budgetId is valid
	var budgets models.BudgetsResponse
	err := database.DB.Where(map[string]interface{}{"user_id": userID, "isDeleted": false, "budgetId": budgetId}).Find(&budgets.Budgets).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var budget = budgets.Budgets[0]
	budgetStartDate, _ := time.Parse(time.RFC3339, budget.Data.StartDate)

	// Gets cycle index and start/end
	var cycleIndex = 0
	var cycleRangeStart = time.Now()
	var cycleRangeEnd = time.Now()
	if budget.Data.Frequency == "weekly" {
		cycleIndex = int(date.Sub(budgetStartDate).Hours() / 24 / 7 / float64(budget.Data.CycleDuration))
		cycleRangeStart = budgetStartDate.AddDate(0, 0, cycleIndex*7*int(budget.Data.CycleDuration))
		cycleRangeEnd = cycleRangeStart.AddDate(0, 0, 7*int(budget.Data.CycleDuration)).Add(-1 * time.Second)
	} else if budget.Data.Frequency == "monthly" {
		year, month, _, _, _, _ := diff(date, budgetStartDate)
		cycleIndex = int(float64(year*12.0+month) / float64(budget.Data.CycleDuration))
		cycleRangeStart = budgetStartDate.AddDate(0, cycleIndex, 0)
		cycleRangeEnd = cycleRangeStart.AddDate(0, int(budget.Data.CycleDuration), 0).Add(-1 * time.Second)
	} else {
		year, _, _, _, _, _ := diff(date, budgetStartDate)
		cycleIndex = int(float64(year) / float64(budget.Data.CycleDuration))
		cycleRangeStart = budgetStartDate.AddDate(cycleIndex, 0, 0)
		cycleRangeEnd = cycleRangeStart.AddDate(int(budget.Data.CycleDuration), 0, 0).Add(-1 * time.Second)
	}

	var cycleResponse models.Cycle
	cycleResponse.Index = cycleIndex
	cycleResponse.Start = cycleRangeStart.Format(time.RFC3339)
	cycleResponse.End = cycleRangeEnd.Format(time.RFC3339)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cycleResponse)
}

func DateEqual(time1 time.Time, time2 time.Time) bool {
	return (time1.Year() == time2.Year()) && (time1.Month() == time2.Month()) && (time1.Day() == time2.Day())
}
