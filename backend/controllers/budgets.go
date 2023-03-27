package controllers

import (
	"budget-tracker/database"
	"budget-tracker/models"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func CreateBudget(w http.ResponseWriter, r *http.Request) {
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

	var budgetData models.BudgetContent
	_ = json.NewDecoder(r.Body).Decode(&budgetData)

	newBudget := models.Budget{
		UserID:    uint(userID),
		BudgetID:  0,
		IsDeleted: false,
		Data:      budgetData,
	}

	database.DB.Create(&newBudget)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.CreateBudgetResponse{
		UserID:   newBudget.UserID,
		BudgetID: newBudget.BudgetID,
	})

}

func GetBudgets(w http.ResponseWriter, r *http.Request) {
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

	var budgets models.BudgetsResponse
	database.DB.Where(map[string]interface{}{"user_id": userID, "isDeleted": false}).Find(&budgets.Budgets)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budgets)
}

func GetDeletedBudgets(w http.ResponseWriter, r *http.Request) {
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

	var budgets models.BudgetsResponse
	database.DB.Where(map[string]interface{}{"user_id": userID, "isDeleted": true}).Find(&budgets.Budgets)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(budgets)
}

func UpdateBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var updateBudget models.UpdateBudgetRequest
	_ = json.NewDecoder(r.Body).Decode(&updateBudget.NewBudget)

	userID, _ := strconv.ParseInt(ReturnUserID(w, r), 10, 32)
	if userID == -1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userID != int64(updateBudget.NewBudget.UserID) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var oldBudget models.Budget
	if err := database.DB.First(&oldBudget, updateBudget.NewBudget.BudgetID).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oldBudget = updateBudget.NewBudget
	database.DB.Save(oldBudget)
	w.WriteHeader(http.StatusOK)
}

func DeleteBudget(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	deletingUserId := uint(userID)
	tempBudgetId, _ := strconv.Atoi(vars["budgetId"])
	deletingBudgetId := uint(tempBudgetId)

	var toDelete models.Budget
	err := database.DB.Where(map[string]interface{}{"user_id": deletingUserId, "budgetId": deletingBudgetId}).First(&toDelete).Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !toDelete.IsDeleted {
		toDelete.IsDeleted = true
		database.DB.Save(toDelete)
	} else {
		database.DB.Delete(&toDelete)
	}
}
