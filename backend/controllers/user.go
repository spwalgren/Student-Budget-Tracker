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
	password, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(password)
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
	// var returnInfo models.ReturnInfo

	_ = json.NewDecoder(r.Body).Decode(&userLoggingIn)

	searchResult := models.DB.Where("email = ?", userLoggingIn.Email).First(&info)

	// No user with matching email is not found
	if searchResult.Error != nil {
		// json.NewEncoder(w).Encode(models.ReturnInfo{ID: ""})
		json.NewEncoder(w).Encode(models.Error{Message: "Email not found"})
		return
	}

	// Password is incorrect
	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(userLoggingIn.Password)); err != nil {
		// json.NewEncoder(w).Encode(models.ReturnInfo{ID: "-1"})
		json.NewEncoder(w).Encode(models.Error{Message:string([]byte(info.Password)) + " " + string([]byte(userLoggingIn.Password))})
		return
	}

	// Create Token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(info.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(models.SecretKey))

	// Error creating token
	if err != nil {
		json.NewEncoder(w).Encode(models.Error{Message: "Could not login"})
		return
	}

	// Set cookies to token if success
	cookie := http.Cookie{
		Name: "jtw",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	
	json.NewEncoder(w).Encode(models.Error{Message:"success"})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jtw")
	if err != nil {
        json.NewEncoder(w).Encode(models.Error{Message:"Error getting cookies"})
		return
    }
	tempClaims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(cookie.Value, &tempClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.SecretKey), nil
	})

	if err != nil {
		json.NewEncoder(w).Encode(models.Error{Message:"unauthenticated"})
		return
	}
	
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.UserInfo

	models.DB.Where("id = ?", claims.Issuer).First(&user)

	json.NewEncoder(w).Encode(map[string]interface{}{"ID": user.ID, "email": user.Email, "firstName": user.FirstName, "lastName": user.LastName})
}
