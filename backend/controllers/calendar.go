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

func GetEvents(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}

	var eventIdCount = 0
	vars := mux.Vars(r)
	tempMonth, _ := strconv.Atoi(vars["month"])
	currentTime := time.Now().AddDate(0,tempMonth,0)
	firstOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfMonth.AddDate(0,1,0)
	lastOfMonth = lastOfMonth.Add(-1 * time.Second)

	userID, _ := strconv.ParseInt(ReturnUserID(w, r), 10, 32)
	if userID == -1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var eventsResponse models.EventsResponse
	var budgets models.BudgetsResponse
	database.DB.Where(map[string]interface{}{"user_id": userID, "isDeleted": false}).Find(&budgets.Budgets)
	for i := 0; i < len(budgets.Budgets); i++ {
		var currBudget = budgets.Budgets[i]
		budgetStartTime, _ := time.Parse(time.RFC3339, currBudget.Data.StartDate)
		budgetStartTime = time.Date(budgetStartTime.Year(), budgetStartTime.Month(), budgetStartTime.Day(), 0, 0, 0, 0, budgetStartTime.Location())
		// First check if budget ends before the month selected
		if (currBudget.Data.CycleCount != 0) {
			budgetEndTime := time.Now()
			if (currBudget.Data.Frequency == "weekly") {
				budgetEndTime = budgetStartTime.AddDate(0,0,7*int(currBudget.Data.CycleDuration)*int(currBudget.Data.CycleCount))
			} else if (currBudget.Data.Frequency == "monthly") {
				budgetEndTime = budgetStartTime.AddDate(0,int(currBudget.Data.CycleDuration)*int(currBudget.Data.CycleCount),0)
			} else {
				budgetEndTime = budgetStartTime.AddDate(int(currBudget.Data.CycleDuration)*int(currBudget.Data.CycleCount),0,0)
			}
			budgetEndTime.Add(-1 * time.Second)
			if (budgetEndTime.Before(firstOfMonth)) {
				continue;
			}
		}
		// check if budget starts after month selected
		if (budgetStartTime.After(lastOfMonth)) {
			continue;
		}
		var cycleRangeStart = 0
		var cycleRangeEnd = 0
		// Getting number of cycles since budget created
		if (currBudget.Data.Frequency == "weekly") {
			cycleRangeStart = (int(firstOfMonth.Sub(budgetStartTime).Hours() / 24 / float64(currBudget.Data.CycleDuration) / 7))
			if (cycleRangeStart < 0) {
				cycleRangeStart = 0
			}
			if (int(currBudget.Data.CycleCount) != 0 && cycleRangeStart >= int(currBudget.Data.CycleCount)) {
				continue;
			}
			cycleRangeEnd = (int(lastOfMonth.Sub(budgetStartTime).Hours() / 24 / float64(currBudget.Data.CycleDuration) / 7))
			if (int(currBudget.Data.CycleCount) != 0 && cycleRangeEnd >= int(currBudget.Data.CycleCount)) {
				cycleRangeEnd = int(currBudget.Data.CycleCount) - 1
			}
		} else if (currBudget.Data.Frequency == "monthly") {
			year, month, _, _, _, _ := diff(firstOfMonth, budgetStartTime)
			cycleRangeStart = int(float64(year*12.0 + month) / float64(currBudget.Data.CycleDuration))
			if (cycleRangeStart < 0) {
				cycleRangeStart = 0
			}
			if (int(currBudget.Data.CycleCount) != 0 && cycleRangeStart >= int(currBudget.Data.CycleCount)) {
				continue;
			}
			year, month, _, _, _, _ = diff(lastOfMonth, budgetStartTime)
			cycleRangeEnd = int(float64(year*12.0 + month) / float64(currBudget.Data.CycleDuration))
			if (int(currBudget.Data.CycleCount) != 0 && cycleRangeEnd >= int(currBudget.Data.CycleCount)) {
				cycleRangeEnd = int(currBudget.Data.CycleCount) - 1
			}
		} else {
			year, _, _, _, _, _ := diff(firstOfMonth, budgetStartTime)
			cycleRangeStart = int(float64(year) / float64(currBudget.Data.CycleDuration))
			if (cycleRangeStart < 0) {
				cycleRangeStart = 0
			}
			if (int(currBudget.Data.CycleCount) != 0 && cycleRangeStart >= int(currBudget.Data.CycleCount)) {
				continue;
			}
			year, _, _, _, _, _ = diff(lastOfMonth, budgetStartTime)
			cycleRangeEnd = int(float64(year) / float64(currBudget.Data.CycleDuration))
			if (int(currBudget.Data.CycleCount) != 0 && cycleRangeEnd >= int(currBudget.Data.CycleCount)) {
				cycleRangeEnd = int(currBudget.Data.CycleCount) - 1
			}
		}
		for j := cycleRangeStart; j <= cycleRangeEnd; j++ {
			if (currBudget.Data.Frequency == "weekly") {
				var eventStartDate = budgetStartTime.AddDate(0,0,7*j*int(currBudget.Data.CycleDuration))
				var eventEndDate = budgetStartTime.AddDate(0,0,7*(j+1)*int(currBudget.Data.CycleDuration)).Add(-1 * time.Second)
				if (eventStartDate.Equal(firstOfMonth) || eventStartDate.Equal(lastOfMonth) || (eventStartDate.After(firstOfMonth) && eventStartDate.Before(lastOfMonth)) || (eventEndDate.After(firstOfMonth) && eventEndDate.Before(lastOfMonth))) {
					var tempEvent models.Event
					tempEvent.UserID = uint(userID)
					tempEvent.EventID = uint(eventIdCount)
					eventIdCount++
					var tempEventContent models.EventContent
					tempEventContent.Frequency = currBudget.Data.Frequency
					tempEventContent.AmountLimit = currBudget.Data.AmountLimit
					tempEventContent.Category = currBudget.Data.Category
					tempEventContent.StartDate = eventStartDate.Format(time.RFC3339)
					tempEventContent.EndDate = eventEndDate.Format(time.RFC3339)
					var transactions models.TransactionsResponse
					database.DB.Where(map[string]interface{}{"user_id": userID, "category": currBudget.Data.Category}).Find(&transactions.Data)
					tempSpent := float32(0.0)
					for k := 0; k < len(transactions.Data); k++ {
						var cycle []models.BudgetTransaction
						database.DB.Where(map[string]interface{}{"budget_id": currBudget.BudgetID, "transaction_id": transactions.Data[k].TransactionID, "cycle_index":j}).Find(&cycle)
						if (len(cycle) != 0) {
							tempSpent += transactions.Data[k].Amount
						}
					}
					tempEventContent.TotalSpent = tempSpent
					tempEvent.Data = tempEventContent
					eventsResponse.Events = append(eventsResponse.Events, tempEvent)
				}

			} else if (currBudget.Data.Frequency == "monthly") {
				var eventStartDate = budgetStartTime.AddDate(0,j*int(currBudget.Data.CycleDuration),0)
				eventEndDate := budgetStartTime.AddDate(0,(j+1)*int(currBudget.Data.CycleDuration),0).Add(-1*time.Second)
				if (eventStartDate.Equal(firstOfMonth) || eventStartDate.Equal(lastOfMonth) || (eventStartDate.After(firstOfMonth) && eventStartDate.Before(lastOfMonth)) || (eventEndDate.After(firstOfMonth) && eventEndDate.Before(lastOfMonth))) {
					var tempEvent models.Event
					tempEvent.UserID = uint(userID)
					tempEvent.EventID = uint(eventIdCount)
					eventIdCount++
					var tempEventContent models.EventContent
					tempEventContent.Frequency = currBudget.Data.Frequency
					tempEventContent.AmountLimit = currBudget.Data.AmountLimit
					tempEventContent.Category = currBudget.Data.Category
					tempEventContent.StartDate = eventStartDate.Format(time.RFC3339)
					tempEventContent.EndDate = eventEndDate.Format(time.RFC3339)
					var transactions models.TransactionsResponse
					database.DB.Where(map[string]interface{}{"user_id": userID, "category": currBudget.Data.Category}).Find(&transactions.Data)
					tempSpent := float32(0.0)
					for k := 0; k < len(transactions.Data); k++ {
						var cycle []models.BudgetTransaction
						database.DB.Where(map[string]interface{}{"budget_id": currBudget.BudgetID, "transaction_id": transactions.Data[k].TransactionID, "cycle_index":j}).Find(&cycle)
						if (len(cycle) != 0) {
							tempSpent += transactions.Data[k].Amount
						}
					}
					tempEventContent.TotalSpent = tempSpent
					tempEvent.Data = tempEventContent
					eventsResponse.Events = append(eventsResponse.Events, tempEvent)
				}
			} else {
				var eventStartDate = budgetStartTime.AddDate(j*int(currBudget.Data.CycleDuration),0,0)
				eventEndDate := budgetStartTime.AddDate((j+1)*int(currBudget.Data.CycleDuration),0,0).Add(-1 * time.Second)
				if (eventStartDate.Equal(firstOfMonth) || eventStartDate.Equal(lastOfMonth) || (eventStartDate.After(firstOfMonth) && eventStartDate.Before(lastOfMonth)) || (eventEndDate.After(firstOfMonth) && eventEndDate.Before(lastOfMonth))) {
					var tempEvent models.Event
					tempEvent.UserID = uint(userID)
					tempEvent.EventID = uint(eventIdCount)
					eventIdCount++
					var tempEventContent models.EventContent
					tempEventContent.Frequency = currBudget.Data.Frequency
					tempEventContent.AmountLimit = currBudget.Data.AmountLimit
					tempEventContent.Category = currBudget.Data.Category
					tempEventContent.StartDate = eventStartDate.Format(time.RFC3339)
					tempEventContent.EndDate = eventEndDate.Format(time.RFC3339)
					var transactions models.TransactionsResponse
					database.DB.Where(map[string]interface{}{"user_id": userID, "category": currBudget.Data.Category}).Find(&transactions.Data)
					tempSpent := float32(0.0)
					for k := 0; k < len(transactions.Data); k++ {
						var cycle []models.BudgetTransaction
						database.DB.Where(map[string]interface{}{"budget_id": currBudget.BudgetID, "transaction_id": transactions.Data[k].TransactionID, "cycle_index":j}).Find(&cycle)
						if (len(cycle) != 0) {
							tempSpent += transactions.Data[k].Amount
						}
					}
					tempEventContent.TotalSpent = tempSpent
					tempEvent.Data = tempEventContent
					eventsResponse.Events = append(eventsResponse.Events, tempEvent)
				}
			}
		}
	}
	json.NewEncoder(w).Encode(eventsResponse)
}