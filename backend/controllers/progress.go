package controllers

import (
	"budget-tracker/database"
	"budget-tracker/models"
	"encoding/json"
	//"fmt"
	"net/http"
	"reflect"
	"strconv"
	//"github.com/gorilla/mux"
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

	var progRequest models.GetProgressRequest
	_ = json.NewDecoder(r.Body).Decode(&progRequest)

	// Sets up response and initializes the response.Frequency field
	var progResponse models.GetProgressResponse
	progResponse.Data.Frequency = progRequest.Frequency
	progResponse.Data.UserID = uint(userID)

	var transactionData []models.Transaction
	var budgetData []models.Budget
	//var categories []string
	 categoryMap := make(map[string]bool)

	// Gets all budgets for specified period
	database.DB.Where(map[string]interface{}{"user_id": userID, "frequency": progRequest.Frequency, "isDeleted": false}).Find(&budgetData)

	// Gets all the categories from the budgets made with the requested frequency
	var budgetGoal float32 = 0.0
	for _ , element := range budgetData {
		categoryMap[element.Data.Category] = true
		progResponse.Data.BudgetIDList = append(progResponse.Data.BudgetIDList, element.BudgetID)
		budgetGoal += element.Data.AmountLimit
	}

	progResponse.Data.BudgetGoal = budgetGoal

	// Gets all transactions from categories previously collected and sums total spent
	for _ , element := range reflect.ValueOf(categoryMap).MapKeys() {
		category := element.Interface().(string)
		database.DB.Where(map[string]interface{}{"user_id": userID, "category": category}).Find(&transactionData)
	}

	var totalSpent float32 = 0.0
	for _ , element := range transactionData {
		totalSpent += element.Amount
		progResponse.Data.TransactionIDList = append(progResponse.Data.TransactionIDList, element.TransactionID)
	}

	progResponse.Data.TotalSpent = totalSpent




	json.NewEncoder(w).Encode(progResponse)
	w.WriteHeader(http.StatusOK)
}
