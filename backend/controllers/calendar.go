package controllers

import (
	"budget-tracker/database"
	"budget-tracker/models"
	"encoding/json"
	"log"
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
	log.Println("month: ", tempMonth)
	currentTime := time.Now().AddDate(0,tempMonth,0)
	firstOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
	lastOfMonth := firstOfMonth.AddDate(0,1,-1)
	log.Println("currentDate: ", currentTime.Format("01/02/2006"))
	log.Println("firstOfMonth: ", firstOfMonth.Format("01/02/2006"))
	log.Println("lastOfMonth: ", lastOfMonth.Format("01/02/2006"))

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
		log.Println("budget category: ", currBudget.Data.Category)
		log.Println("budget start time: ", budgetStartTime.Format("01/02/2006"))
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
			cycleRangeStart = (int(firstOfMonth.Sub(budgetStartTime).Hours() / 24 / float64(currBudget.Data.CycleDuration) / 7) - 1)
			cycleRangeEnd = (int(lastOfMonth.Sub(budgetStartTime).Hours() / 24 / float64(currBudget.Data.CycleDuration) / 7) + 1)
		} else if (currBudget.Data.Frequency == "monthly") {
			cycleRangeStart = (int(firstOfMonth.Sub(budgetStartTime).Hours() / 24 / float64(currBudget.Data.CycleDuration) / 31) - 1)
			cycleRangeEnd = (int(lastOfMonth.Sub(budgetStartTime).Hours() / 24 / float64(currBudget.Data.CycleDuration) / 28) + 1)
		} else {
			cycleRangeStart = (int(firstOfMonth.Sub(budgetStartTime).Hours() / 24 / float64(currBudget.Data.CycleDuration) / 366) - 1)
			cycleRangeEnd = (int(lastOfMonth.Sub(budgetStartTime).Hours() / 24 / float64(currBudget.Data.CycleDuration) / 365) + 1)
		}
		log.Println("cycle start range: ", cycleRangeStart)
		log.Println("cycle end range: ", cycleRangeEnd)
		for j := cycleRangeStart; j <= cycleRangeEnd; j++ {
			if (currBudget.Data.Frequency == "weekly") {
				var eventDate = budgetStartTime.AddDate(0,0,7*j*int(currBudget.Data.CycleDuration))
				if (eventDate.After(firstOfMonth) && eventDate.Before(lastOfMonth)) {
					var tempEvent models.Event
					tempEvent.UserID = uint(userID)
					tempEvent.EventID = uint(eventIdCount)
					eventIdCount++
					var tempEventContent models.EventContent
					tempEventContent.Frequency = currBudget.Data.Frequency
					tempEventContent.AmountLimit = currBudget.Data.AmountLimit
					tempEventContent.Category = currBudget.Data.Category
					tempEventContent.StartDate = eventDate.Format("01/02/2006")
					tempEventContent.EndDate = eventDate.AddDate(0,0,7*int(currBudget.Data.CycleDuration)).Format("01/02/2006")
					tempEventContent.TotalSpent = 0
					tempEvent.Data = tempEventContent
					eventsResponse.Events = append(eventsResponse.Events, tempEvent)
				}

			} else if (currBudget.Data.Frequency == "monthly") {
				var eventDate = budgetStartTime.AddDate(0,j*int(currBudget.Data.CycleDuration),0)
				if (eventDate.After(firstOfMonth) && eventDate.Before(lastOfMonth)) {
					var tempEvent models.Event
					tempEvent.UserID = uint(userID)
					tempEvent.EventID = uint(eventIdCount)
					eventIdCount++
					var tempEventContent models.EventContent
					tempEventContent.Frequency = currBudget.Data.Frequency
					tempEventContent.AmountLimit = currBudget.Data.AmountLimit
					tempEventContent.Category = currBudget.Data.Category
					tempEventContent.StartDate = eventDate.Format("01/02/2006")
					tempEventContent.EndDate = eventDate.AddDate(0,int(currBudget.Data.CycleDuration),0).Format("01/02/2006")
					tempEventContent.TotalSpent = 0
					tempEvent.Data = tempEventContent
					eventsResponse.Events = append(eventsResponse.Events, tempEvent)
				}
			} else {
				var eventDate = budgetStartTime.AddDate(j*int(currBudget.Data.CycleDuration),0,0)
				if (eventDate.After(firstOfMonth) && eventDate.Before(lastOfMonth)) {
					var tempEvent models.Event
					tempEvent.UserID = uint(userID)
					tempEvent.EventID = uint(eventIdCount)
					eventIdCount++
					var tempEventContent models.EventContent
					tempEventContent.Frequency = currBudget.Data.Frequency
					tempEventContent.AmountLimit = currBudget.Data.AmountLimit
					tempEventContent.Category = currBudget.Data.Category
					tempEventContent.StartDate = eventDate.Format("01/02/2006")
					tempEventContent.EndDate = eventDate.AddDate(int(currBudget.Data.CycleDuration),0,0).Format("01/02/2006")
					tempEventContent.TotalSpent = 0
					tempEvent.Data = tempEventContent
					eventsResponse.Events = append(eventsResponse.Events, tempEvent)
				}
			}
		}
	}
	json.NewEncoder(w).Encode(eventsResponse)
	w.WriteHeader(http.StatusOK)
}