package controllers

import (
	"budget-tracker/database"
	"budget-tracker/models"
	"encoding/json"
	"net/http"
	"strconv"
	//"fmt"
	//"github.com/gorilla/mux"
)


func CreateProgress(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}

	userID, _ := strconv.ParseInt(ReturnUserID(w, r), 10, 32)
	if userID == -1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var prog models.Progress
	_ = json.NewDecoder(r.Body).Decode(&prog)

	prog.UserID = uint(userID)

	database.DB.Create(&prog)
	json.NewEncoder(w).Encode(models.CreateProgressResponse{
		UserID: prog.UserID,
		ProgressID: prog.ProgressID,
	})
	w.WriteHeader(http.StatusOK)
}

func GetProgress(w http.ResponseWriter, r* http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}

	userID, _ := strconv.ParseInt(ReturnUserID(w, r), 10, 32)
	if userID == -1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var progResponse models.ProgressResponse

	database.DB.Where(map[string]interface{}{"user_id": userID}).Find(&progResponse.ProgressData)
	json.NewEncoder(w).Encode(progResponse)
	w.WriteHeader(http.StatusOK)
}
