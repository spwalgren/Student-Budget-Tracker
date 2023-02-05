package controllers

import (
	"budget-tracker/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var users []models.UserInfo
	models.DB.Find(&users)

	for _, item := range users {
		if item.FirstName == param["firstName"] && item.LastName == param["lastName"] {
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
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "*")

	var newUser models.UserInfo
	var users []models.UserInfo
	models.DB.Find(&users)
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	for _, entry := range users {
		if entry.Email == newUser.Email {
			json.NewEncoder(w).Encode(models.ReturnInfo{ID: ""})
			return
		}
	}
	models.DB.Create(&newUser)
	json.NewEncoder(w).Encode(models.ReturnInfo{ID: strconv.FormatUint(uint64(newUser.ID), 10)})
}

/*  Checks authentication by looking for email/password combination in database.
 *  If that combo doesn't exist, checks if an email exists.
 *  Returns empty ID field for an account that doesn't exists under an email.
 *	Returns "-1" for a wrong password and returns the ID of the user that successfully
 *	logged in.
 */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var userLoggingIn models.UserLoginInfo
	var info models.UserInfo
	var returnInfo models.ReturnInfo

	_ = json.NewDecoder(r.Body).Decode(&userLoggingIn)
	searchResult := models.DB.First(&info, userLoggingIn)

	if searchResult.Error != nil {
		emailSearch := models.DB.First(&info, "email=?", userLoggingIn.Email)
		if emailSearch.Error != nil {
			json.NewEncoder(w).Encode(models.ReturnInfo{ID: ""})
			return
		}
		json.NewEncoder(w).Encode(models.ReturnInfo{ID: "-1"})
		return
	}

	// Successful login
	returnInfo.ID = strconv.FormatUint(uint64(info.ID), 10)
	json.NewEncoder(w).Encode(returnInfo)
}
