package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"budget-tracker/models"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var users []models.UserInfo
	models.DB.Find(&users)

	for _, item := range users {
		if item.FirstName == param["firstname"] && item.LastName == param["lastname"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.UserInfo{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Origin", "*")
	var users []models.UserInfo
	models.DB.Find(&users)

	json.NewEncoder(w).Encode(users)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser models.UserInfo
	var users []models.UserInfo
	models.DB.Find(&users)
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	for _, entry := range users {
		log.Println(entry)
		if entry.Email == newUser.Email {
			fmt.Println("Email already associated with an account")
			return
		}
	}
	models.DB.Create(&newUser)
	json.NewEncoder(w).Encode(newUser)
}
