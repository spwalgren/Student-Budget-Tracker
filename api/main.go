package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type UserInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []UserInfo{} // temporary database for testing purposes

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key, nil)
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	fmt.Println(session.Values)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome!")
}

func ForbiddenHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	session, _ := store.Get(r, "cookie-name")

	// check username and password or authentication

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "YOU SHALL NOT PASS", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "Here lies my deepest secret")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check username and password
	var userLoggingIn UserLoginInfo
	_ = json.NewDecoder(r.Body).Decode(&userLoggingIn)
	for _, entry := range users {
		if entry.Email == userLoggingIn.Email && entry.Password == userLoggingIn.Password {
			session.Values["authenticated"] = true
			session.Save(r, w)
			fmt.Fprintln(w, "Logging in")
			return
		}
	}

	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintln(w, "That is not an email-password combination associated with a registered account")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintln(w, "Logging out")
}

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
		if item.Name == param["name"] {
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

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("GET")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")
	r.HandleFunc("/forbidden", ForbiddenHandler).Methods("GET")

	r.HandleFunc("/user/{name}", getUser).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/register", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
