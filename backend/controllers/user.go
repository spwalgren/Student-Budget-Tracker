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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Origin", "http://localhost:4200")
	var users []models.UserInfo
	models.DB.Find(&users)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "*")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Content-Type", "*")
	w.WriteHeader(http.StatusOK)

	var newUser models.UserInfo
	var users []models.UserInfo
	models.DB.Find(&users)
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	for _, entry := range users {
		if entry.Email == newUser.Email {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(models.Error{Message: "duplicate email"})
			return
		}
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(password)
	models.DB.Create(&newUser)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ReturnInfo{ID: strconv.FormatUint(uint64(newUser.ID), 10)})
}

/*
*  Checks authentication by looking for user in database with matching email.
*	If no user found, returns message "email not found", otherwise checks if password matches
*	If password doesn't match, returns message "incorrect password"
*	If password matches, it creates token, sets cookies to that token, and returns "success"
 */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
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
		json.NewEncoder(w).Encode(models.Error{Message: "email not found"})
		return
	}

	// Password is incorrect
	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(userLoggingIn.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.Error{Message: "incorrect password"})
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
		json.NewEncoder(w).Encode(models.Error{Message: "could not login"})
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
	json.NewEncoder(w).Encode(models.Error{Message: "success"})
}

/*
*	Gets jwt token from cookies
*	Gets claims from token
*	Claims issuer contains the logged in user ID
*	Returns user based on user ID
 */
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "*")
	w.WriteHeader(http.StatusOK)
	cookie, err := r.Cookie("jtw")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Error{Message: "error getting cookies"})
		return
	}
	tempClaims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(cookie.Value, &tempClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.SecretKey), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.Error{Message: "unauthenticated"})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.UserInfo

	models.DB.Where("id = ?", claims.Issuer).First(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"ID": user.ID, "email": user.Email, "firstName": user.FirstName, "lastName": user.LastName})
}
