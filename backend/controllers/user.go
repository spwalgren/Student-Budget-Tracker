package controllers

import (
	"budget-tracker/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Origin", "http://localhost:4200")
// 	var users []models.UserInfo
// 	models.DB.Find(&users)
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(users)
// }

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var newUser models.UserInfo
	var users []models.UserInfo
	models.DB.Find(&users)
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	for _, entry := range users {
		if entry.Email == newUser.Email {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(password)
	models.DB.Create(&newUser)
	w.WriteHeader(http.StatusOK)
}

/*
 * Logout user by deleting the corresponding cookie
 */
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")
	http.SetCookie(w, &http.Cookie{
		Name:     "jtw",
		Expires:  time.Now().Add(-24),
		Domain:   "localhost",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
}

/*
*  Checks authentication by looking for user in database with matching email.
*	If no user found, returns message "email not found", otherwise checks if password matches
*	If password doesn't match, returns message "incorrect password"
*	If password matches, it creates token, sets cookies to that token, and returns "success"
 */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var userLoggingIn models.UserLoginInfo
	var info models.UserInfo

	_ = json.NewDecoder(r.Body).Decode(&userLoggingIn)

	searchResult := models.DB.Where("email = ?", userLoggingIn.Email).First(&info)

	// No user with matching email is not found
	if searchResult.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Password is incorrect
	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(userLoggingIn.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create Token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(info.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(models.SecretKey))

	// Error creating token
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set cookies to token if success
	cookie := http.Cookie{
		Name:     "jtw",
		Value:    token,
		Domain:   "localhost",
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Secure: true,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

/*
*	Gets jwt token from cookies
*	Gets claims from token
*	Claims issuer contains the logged in user ID
*	Returns user based on user ID
 */
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	cookie, err := r.Cookie("jtw")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	models.DB.Where("id = ?", claims.Issuer).First(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": user.ID, "email": user.Email, "firstName": user.FirstName, "lastName": user.LastName})
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var newTransaction models.Transaction
	var expenses models.FinancialInfo

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

	result := models.DB.Where("ID = ?", claims.Issuer).First(&expenses)

	// Transactions do not exist. Create one before moving forward
	if result.Error != nil {
		var user models.UserInfo
		models.DB.Where("ID = ?", claims.Issuer).First(&user)
		expenses.TransactionID = user.ID
		expenses.Transactions = append(expenses.Transactions, newTransaction)
	}

	json.NewEncoder(w).Encode(expenses)
}
