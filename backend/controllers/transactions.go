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
	if (len(budgets.Budgets) == 0) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Getting cycle index for current transaction based on matching transaction category with budget category
	var budget = budgets.Budgets[0]
	transactionDate, _ := time.Parse(time.RFC3339, newTransactionData.Data.Date)
	budgetStartDate, _ := time.Parse(time.RFC3339, budget.Data.StartDate)
	var cycleIndex = 0
	if (budget.Data.Frequency == "weekly") {
		cycleIndex = int(transactionDate.Sub(budgetStartDate).Hours() / 24 / 7 / float64(budget.Data.CycleDuration))
	} else if (budget.Data.Frequency == "monthly") {
		year, month, _, _, _, _ := diff(transactionDate, budgetStartDate)
		cycleIndex = int(float64(year*12.0 + month) / float64(budget.Data.CycleDuration))
	} else {
		year, _, _, _, _, _ := diff(transactionDate, budgetStartDate)
		cycleIndex = int(float64(year) / float64(budget.Data.CycleDuration))
	}


	newTransaction := models.Transaction{
		UserID:        0,
		TransactionID: 0,
		Amount:        newTransactionData.Data.Amount,
		Name:          newTransactionData.Data.Name,
		Date:          newTransactionData.Data.Date,
		Category:      newTransactionData.Data.Category,
		Description:   newTransactionData.Data.Description,
		CycleIndex:    cycleIndex,
	}
	var user models.UserInfo
	database.DB.First(&user, userID)
	newTransaction.UserID = user.ID
	fmt.Println(newTransaction)
	database.DB.Create(&newTransaction)
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
	if (len(budgets.Budgets) == 0) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	// Gets cycle index for current transaction based on matching transaction category with budget category
	var budget = budgets.Budgets[0]
	transactionDate, _ := time.Parse(time.RFC3339, expenses.Date)
	budgetStartDate, _ := time.Parse(time.RFC3339, budget.Data.StartDate)
	var cycleIndex = 0
	if (budget.Data.Frequency == "weekly") {
		cycleIndex = int(transactionDate.Sub(budgetStartDate).Hours() / 24 / 7 / float64(budget.Data.CycleDuration))
	} else if (budget.Data.Frequency == "monthly") {
		year, month, _, _, _, _ := diff(transactionDate, budgetStartDate)
		cycleIndex = int(float64(year*12.0 + month) / float64(budget.Data.CycleDuration))
	} else {
		year, _, _, _, _, _ := diff(transactionDate, budgetStartDate)
		cycleIndex = int(float64(year) / float64(budget.Data.CycleDuration))
	}


	expenses.CycleIndex = cycleIndex
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
}
