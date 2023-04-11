package controllers

import (
	"budget-tracker/database"
	"budget-tracker/models"
	"encoding/json"

	//"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//"github.com/gorilla/mux"
	//"time"
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
	var yearlyBudgets	models.BudgetsResponse

	// GET ALL WEEKLY BUDGETS
	database.DB.Where(map[string]interface{}{"user_id": userID,"frequency": "weekly", "isDeleted": false}).Find(&weeklyBudgets.Budgets)
	// CREATE A NEW PROGRESS ENTRY FOR EACH BUDGET
	for i := 0; i < len(weeklyBudgets.Budgets); i++ {
		temp := weeklyBudgets.Budgets[i]
		weeklyProgResponse.Data = append(weeklyProgResponse.Data, models.Progress{UserID: temp.UserID, Frequency: temp.Data.Frequency, Category: temp.Data.Category, BudgetGoal: temp.Data.AmountLimit, BudgetID: temp.BudgetID})
	}

	// ADD ALL MONTHLY PROGRESS
	database.DB.Where(map[string]interface{}{"user_id": userID,"frequency": "monthly", "isDeleted": false}).Find(&monthlyBudgets.Budgets)
	for i := 0; i < len(monthlyBudgets.Budgets); i++ {
		fmt.Println("monthly")
		temp := monthlyBudgets.Budgets[i]
		monthlyProgResponse.Data = append(monthlyProgResponse.Data, models.Progress{UserID: temp.UserID, Frequency: temp.Data.Frequency, Category: temp.Data.Category, BudgetGoal: temp.Data.AmountLimit, BudgetID: temp.BudgetID})
	}

	// ADD ALL YEARLY PROGRESS
	database.DB.Where(map[string]interface{}{"user_id": userID,"frequency": "yearly", "isDeleted": false}).Find(&yearlyBudgets.Budgets)
	for i := 0; i < len(yearlyBudgets.Budgets); i++ {
		fmt.Println("yearly")
		temp := yearlyBudgets.Budgets[i]
		yearlyProgResponse.Data = append(yearlyProgResponse.Data, models.Progress{UserID: temp.UserID, Frequency: temp.Data.Frequency, Category: temp.Data.Category, BudgetGoal: temp.Data.AmountLimit, BudgetID: temp.BudgetID})
	}

	temp1 := append(weeklyProgResponse.Data, monthlyProgResponse.Data...)
	progResponse.Data = append(temp1, yearlyProgResponse.Data...)

	json.NewEncoder(w).Encode(progResponse)
	w.WriteHeader(http.StatusOK)
}
