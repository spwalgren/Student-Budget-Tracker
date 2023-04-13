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
)


func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID := ReturnUserID(w,r)
		if userID == "-1" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

	var newTransactionData models.CreateTransactionRequest
	_ = json.NewDecoder(r.Body).Decode(&newTransactionData)


	var budgets models.BudgetsResponse
	database.DB.Where(map[string]interface{}{"user_id": userID, "isDeleted": false, "category": newTransactionData.Data.Category}).Find(&budgets.Budgets)

	// If the transaction category doesn't match a budget category
	if (len(budgets.Budgets) == 0 && newTransactionData.Data.Category != "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTransaction := models.Transaction{
		UserID:        0,
		TransactionID: 0,
		Amount:        newTransactionData.Data.Amount,
		Name:          newTransactionData.Data.Name,
		Date:          newTransactionData.Data.Date,
		Category:      newTransactionData.Data.Category,
		Description:   newTransactionData.Data.Description,
	}
	var user models.UserInfo
	database.DB.First(&user, userID)
	newTransaction.UserID = user.ID
	fmt.Println(newTransaction)
	if (newTransactionData.Data.Category == "") {
		newTransactionData.Data.Category = "[None]"
	}
	database.DB.Create(&newTransaction)
	

	// Getting cycle index for current transaction based on matching transaction category with budget category
	transactionDate, err := time.Parse(time.RFC3339, newTransactionData.Data.Date)
	if err != nil {
		transactionDate, _ = time.Parse("2006-01-02", newTransactionData.Data.Date)
		transactionDate = time.Date(transactionDate.Year(), transactionDate.Month(), transactionDate.Day(), 4, 0, 0, 0, transactionDate.Location())
	}
	for i := 0; i < len(budgets.Budgets); i++ {
		var budgetCycle models.BudgetTransaction
		budgetStartDate, err := time.Parse(time.RFC3339, budgets.Budgets[i].Data.StartDate)
		if err != nil {
			budgetStartDate, _ = time.Parse("2006-01-02", budgets.Budgets[i].Data.StartDate)
			budgetStartDate = time.Date(budgetStartDate.Year(), budgetStartDate.Month(), budgetStartDate.Day(), 4, 0, 0, 0, budgetStartDate.Location())
		}
		var cycleIndex = 0
		if (budgets.Budgets[i].Data.Frequency == "weekly") {
			cycleIndex = int(transactionDate.Sub(budgetStartDate).Hours() / 24 / 7 / float64(budgets.Budgets[i].Data.CycleDuration))
		} else if (budgets.Budgets[i].Data.Frequency == "monthly") {
			year, month, _, _, _, _ := diff(transactionDate, budgetStartDate)
			cycleIndex = int(float64(year*12.0 + month) / float64(budgets.Budgets[i].Data.CycleDuration))
		} else {
			year, _, _, _, _, _ := diff(transactionDate, budgetStartDate)
			cycleIndex = int(float64(year) / float64(budgets.Budgets[i].Data.CycleDuration))
		}
		budgetCycle.BudgetID = budgets.Budgets[i].BudgetID
		budgetCycle.CycleIndex = cycleIndex
		budgetCycle.TransactionID = newTransaction.TransactionID
		database.DB.Create(&budgetCycle)
	}


	json.NewEncoder(w).Encode(models.CreateTransactionResponse{
		UserID:newTransaction.UserID,
		TransactionID: newTransaction.TransactionID,
	})
	
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var user models.UserInfo
	userID := ReturnUserID(w,r)

	if userID == "-1" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	database.DB.First(&user, userID)

	var expenses []models.Transaction
	var data models.TransactionsResponse
	database.DB.Where(map[string]interface{}{"user_id": user.ID}).Find(&expenses)
	data.Data = expenses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var updateTransaction models.Transaction
	var updateTransactionData models.UpdateTransactionRequest
	_ = json.NewDecoder(r.Body).Decode(&updateTransactionData)
	updateTransaction = updateTransactionData.Data


	// get userID to get the list of transactions from current user

	var expenses models.Transaction
	var user models.UserInfo
	userID := ReturnUserID(w,r)

	if userID == "-1" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If the user ID's don't match, the intruder shouldn't be in here anyways
	database.DB.First(&user, userID)
	if user.ID != updateTransaction.UserID {
		w.WriteHeader(http.StatusForbidden)
	}

	// Now using the unique transactionID, get the specific transaction that needs to be updated.
	if err := database.DB.First(&expenses, updateTransaction.TransactionID).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expenses = updateTransaction

	var budgets models.BudgetsResponse
	database.DB.Where(map[string]interface{}{"user_id": userID, "isDeleted": false, "category": expenses.Category}).Find(&budgets.Budgets)
	
	// If the transaction category doesn't match a budget category
	if (len(budgets.Budgets) == 0 && expenses.Category != "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	// Gets cycle index for current transaction based on matching transaction category with budget category
	transactionDate, err := time.Parse(time.RFC3339, expenses.Date)
	if err != nil {
		transactionDate, _ = time.Parse("2006-01-02", expenses.Date)
		transactionDate = time.Date(transactionDate.Year(), transactionDate.Month(), transactionDate.Day(), 4, 0, 0, 0, transactionDate.Location())
	}
	for i := 0; i < len(budgets.Budgets); i++ {
		var budgetCycle models.BudgetTransaction
		database.DB.Where(map[string]interface{}{"transaction_id": expenses.TransactionID, "budget_id": budgets.Budgets[i].BudgetID}).Find(&budgetCycle)
		budgetStartDate, err := time.Parse(time.RFC3339, budgets.Budgets[i].Data.StartDate)
		if err != nil {
			budgetStartDate, _ = time.Parse("2006-01-02", budgets.Budgets[i].Data.StartDate)
			budgetStartDate = time.Date(budgetStartDate.Year(), budgetStartDate.Month(), budgetStartDate.Day(), 4, 0, 0, 0, budgetStartDate.Location())
		}
		var cycleIndex = 0
		if (budgets.Budgets[i].Data.Frequency == "weekly") {
			cycleIndex = int(transactionDate.Sub(budgetStartDate).Hours() / 24 / 7 / float64(budgets.Budgets[i].Data.CycleDuration))
		} else if (budgets.Budgets[i].Data.Frequency == "monthly") {
			year, month, _, _, _, _ := diff(transactionDate, budgetStartDate)
			cycleIndex = int(float64(year*12.0 + month) / float64(budgets.Budgets[i].Data.CycleDuration))
		} else {
			year, _, _, _, _, _ := diff(transactionDate, budgetStartDate)
			cycleIndex = int(float64(year) / float64(budgets.Budgets[i].Data.CycleDuration))
		}
		budgetCycle.BudgetID = budgets.Budgets[i].BudgetID
		budgetCycle.CycleIndex = cycleIndex
		budgetCycle.TransactionID = expenses.TransactionID
		database.DB.Save(&budgetCycle)
	}
	if (expenses.Category == "") {
		expenses.Category = "[None]"
	}
	database.DB.Save(expenses)
	w.WriteHeader(http.StatusOK)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	vars := mux.Vars(r)

	var toDelete models.Transaction

	// UserID and TransactionID will be in the request. Can setup a check to make sure the
	// requesting user matches with the UserID


	var user models.UserInfo
	userID := ReturnUserID(w,r)

	if userID == "-1" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If the user ID's don't match, the intruder shouldn't be in here anyways
	deletingUser, _ := strconv.Atoi(userID)
	deletingUserId := uint(deletingUser)
	temp, _ := strconv.Atoi(vars["transactionId"])
	deletingTransactionId := uint(temp)
	database.DB.First(&user, userID)
	if user.ID != deletingUserId {
		w.WriteHeader(http.StatusForbidden)
	}

	// deletes entry based on the userID and the transactionID
	err := database.DB.Where(map[string]interface{}{"user_id": deletingUserId, "transactionId": deletingTransactionId}).First(&toDelete).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	database.DB.Delete(&toDelete)
	var budgetCycles models.BudgetTransaction
	database.DB.Where(map[string]interface{}{"transaction_id": deletingTransactionId}).Delete(&budgetCycles)
}
