package main

import (
	"budget-tracker/controllers"
	"budget-tracker/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	corsObj := cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:4200"},
		AllowedMethods:     []string{"GET, OPTIONS, POST"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		OptionsPassthrough: true,
		AllowCredentials:   true,
	})

	r := mux.NewRouter()

	models.Connect()
	r.HandleFunc("/api/login", controllers.LoginHandler).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/users", controllers.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/signup", controllers.CreateUser).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/user", controllers.GetUser).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/logout", controllers.LogoutHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", corsObj.Handler(r)))
}
