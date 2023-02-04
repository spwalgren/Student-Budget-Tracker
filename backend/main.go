package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"budget-tracker/models"
	"budget-tracker/controllers"
)

// type UserInfo struct {
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// 	Email     string `json:"email"`
// 	Password  string `json:"password"`
// }

type UserLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// var users = []UserInfo{} // temporary database for testing purposes

/*func loginHandler(w http.ResponseWriter, r *http.Request) {
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
}*/

func main() {
	r := mux.NewRouter()

	models.Connect()
	
	//r.HandleFunc("/login", loginHandler).Methods("GET")
	r.HandleFunc("/user/{name}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/signup", controllers.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
