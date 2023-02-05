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
		AllowedOrigins:     []string{"*"},   // All origins
		AllowedMethods:     []string{"GET"}, // Allowing only get, just an example
		AllowedHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		OptionsPassthrough: true,
	})

	r := mux.NewRouter()

	models.Connect()
	r.HandleFunc("/login", controllers.LoginHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions, http.MethodPost)
	r.HandleFunc("/user/{name}", controllers.GetUser).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions, http.MethodPost)
	r.HandleFunc("/users", controllers.GetUsers).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions, http.MethodPost)
	r.HandleFunc("/signup", controllers.CreateUser).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions, http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", corsObj.Handler(r)))
}
