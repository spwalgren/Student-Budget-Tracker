package controllers

import (
	"budget-tracker/database"
	"budget-tracker/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
	_ = json.NewDecoder(r.Body).Decode(&updateBudget)

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

func GetBudgetCategories(w http.ResponseWriter, r *http.Request) {
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
	var categories models.BudgetCategoriesResponse
	for i := range budgets.Budgets {
		categories.Category = append(categories.Category, budgets.Budgets[i].Data.Category)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
    if a.Location() != b.Location() {
        b = b.In(a.Location())
    }
    if a.After(b) {
        a, b = b, a
    }
    y1, M1, d1 := a.Date()
    y2, M2, d2 := b.Date()

    h1, m1, s1 := a.Clock()
    h2, m2, s2 := b.Clock()

    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)
    hour = int(h2 - h1)
    min = int(m2 - m1)
    sec = int(s2 - s1)

    // Normalize negative values
    if sec < 0 {
        sec += 60
        min--
    }
    if min < 0 {
        min += 60
        hour--
    }
    if hour < 0 {
        hour += 24
        day--
    }
    if day < 0 {
        // days in month:
        t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
        day += 32 - t.Day()
        month--
    }
    if month < 0 {
        month += 12
        year--
    }

	return
}

func GetCyclePeriod(w http.ResponseWriter, r *http.Request) {
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
	dateTemp := vars["date"]
	tempBudgetId, _ := strconv.Atoi(vars["budgetId"])
	budgetId := uint(tempBudgetId)
	date, _ := time.Parse(time.RFC3339, dateTemp)

	var budgets models.BudgetsResponse
	database.DB.Where(map[string]interface{}{"user_id": userID, "isDeleted": false, "budgetId":budgetId}).Find(&budgets.Budgets)
	var budget = budgets.Budgets[0]
	budgetStartDate, _ := time.Parse(time.RFC3339, budget.Data.StartDate)
	var cycleIndex = 0
	var cycleRangeStart = time.Now()
	var cycleRangeEnd = time.Now()
	if (budget.Data.Frequency == "weekly") {
		cycleIndex = int(date.Sub(budgetStartDate).Hours() / 24 / 7 / float64(budget.Data.CycleDuration))
		cycleRangeStart = budgetStartDate.AddDate(0,0,cycleIndex*7)
		cycleRangeEnd = cycleRangeStart.AddDate(0,0,7*int(budget.Data.CycleDuration)).Add(-1 * time.Second)
	} else if (budget.Data.Frequency == "monthly") {
		year, month, _, _, _, _ := diff(date, budgetStartDate)
		cycleIndex = int(float64(year*12.0 + month) / float64(budget.Data.CycleDuration))
		cycleRangeStart = budgetStartDate.AddDate(0,cycleIndex,0)
		cycleRangeEnd = cycleRangeStart.AddDate(0,int(budget.Data.CycleDuration),0).Add(-1 * time.Second)
	} else {
		year, _, _, _, _, _ := diff(date, budgetStartDate)
		cycleIndex = int(float64(year) / float64(budget.Data.CycleDuration))
		cycleRangeStart = budgetStartDate.AddDate(cycleIndex,0,0)
		cycleRangeEnd = cycleRangeStart.AddDate(int(budget.Data.CycleDuration),0,0).Add(-1 * time.Second)
	}
	var cycleResponse models.CyclePeriodResponse
	cycleResponse.Index = cycleIndex
	cycleResponse.Start = cycleRangeStart.Format(time.RFC3339)
	cycleResponse.End = cycleRangeEnd.Format(time.RFC3339)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cycleResponse)
}