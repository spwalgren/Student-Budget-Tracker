package controllers

import (
	"budget-tracker/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	var newTransaction models.Transaction

	_ = json.NewDecoder(r.Body).Decode(&newTransaction)

	// Retrieve transactions for user. If none exist, create one. Retrieve the current cookie to get the user info
	cookie, err := r.Cookie("jtw")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("oops")
		return
	}
	tempClaims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(cookie.Value, &tempClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.SecretKey), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// Transactions do not exist. Create one before moving forward
	var user models.UserInfo
	models.DB.First(&user, claims.Issuer)
	newTransaction.UserID = user.ID
	models.DB.Create(&newTransaction)
	fmt.Println("error")
	json.NewEncoder(w).Encode(newTransaction)

	//  else {
	// 	fmt.Println("nil error")
	// 	// Transactions already exist. Update.
	// 	var temp models.FinancialInfo
	// 	models.DB.First(&temp, claims.Issuer)
	// 	temp.Transactions = append(temp.Transactions, newTransaction)
	// 	models.DB.Model(&temp).First(&temp, claims.Issuer).Update("transactions", temp.Transactions)
	// }
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")
	cookie, err := r.Cookie("jtw")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("oops")
		return
	}
	tempClaims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(cookie.Value, &tempClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.SecretKey), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.UserInfo
	models.DB.First(&user, claims.Issuer)


	var expenses []models.Transaction
	models.DB.Find(&expenses, claims.Issuer)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}
