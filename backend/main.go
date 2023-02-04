package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"budget-tracker/models"
	"budget-tracker/controllers"
	"github.com/gorilla/handlers"
)

func main() {

	corsObj:=handlers.AllowedOrigins([]string{"*"})
	r := mux.NewRouter()

	models.Connect()

	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET")
	r.HandleFunc("/user/{name}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/signup", controllers.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(r)))
}
