package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type userInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var users = []userInfo{
	{ID: "1", Name: "John Doe", Password: "12345"},
	{ID: "2", Name: "Jane Day", Password: "54321"},
}

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
	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Fprintln(w, "Logging in")
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
	var newUser userInfo
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = strconv.Itoa(len(users) + 1)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(&newUser)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, item := range users {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&userInfo{})
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

	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/register", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
