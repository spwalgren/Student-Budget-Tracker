package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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


/*  Checks authentication by looking for email/password combination in database.
	 *  If that combo doesn't exist, checks if an email exists.
	 *  Returns empty ID field for an account that doesn't exists under an email.
	 *	Returns "-1" for a wrong password and returns the ID of the user that successfully
	 *	logged in.
	 */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Origin", "*")
	var userLoggingIn models.UserLoginInfo
	var info models.UserInfo
	var returnInfo models.ReturnLoginInfo

	_ = json.NewDecoder(r.Body).Decode(&userLoggingIn)
	searchResult := models.DB.First(&info, userLoggingIn)

	if searchResult.Error != nil {
		emailSearch := models.DB.First(&info, "email=?", userLoggingIn.Email)
		if emailSearch.Error != nil {
			json.NewEncoder(w).Encode(models.ReturnLoginInfo{ID: ""})
			return
		}
		json.NewEncoder(w).Encode(models.ReturnLoginInfo{ID: "-1"})
		return
	}

	// Successful login
	returnInfo.ID = strconv.FormatUint(uint64(info.ID), 10)
	json.NewEncoder(w).Encode(returnInfo)
}
