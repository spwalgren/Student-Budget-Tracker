package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserInfo struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []UserInfo{} // temporary database for testing purposes

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser UserInfo
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	for _, entry := range users {
		log.Println(entry)
		if entry.Email == newUser.Email {
			fmt.Println("Email already associated with an account")
			return
		}
	}

	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, item := range users {
		if item.FirstName == param["firstname"] && item.LastName == param["lastname"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&UserInfo{})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, item := range users {
		json.NewEncoder(w).Encode(item)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Check username and password
	var userLoggingIn UserLoginInfo
	_ = json.NewDecoder(r.Body).Decode(&userLoggingIn)
	for _, entry := range users {
		if entry.Email == userLoggingIn.Email && entry.Password == userLoggingIn.Password {

			fmt.Fprintln(w, "Logging in")

			// Redirect to user homepage on success

			return
		}
	}

	// Populate error for email/password combo not matching
	fmt.Fprintln(w, "That is not an email-password combination associated with a registered account")
	loginHandler(w, r)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", loginHandler).Methods("GET")
	r.HandleFunc("/user/{name}", getUser).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/signup", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
