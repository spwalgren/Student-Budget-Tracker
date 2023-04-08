package main

import (
	"budget-tracker/controllers"
	"budget-tracker/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	corsObj := cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:4200"},
		AllowedMethods:     []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token", "Origin"},
		OptionsPassthrough: true,
		AllowCredentials:   true,
	})

	r := mux.NewRouter()

	database.Initialize("budget_tracker")
	r.HandleFunc("/api/login", controllers.LoginHandler).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/users", controllers.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/signup", controllers.CreateUser).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/user", controllers.GetUser).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/delete-user", controllers.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/logout", controllers.LogoutHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/transaction", controllers.CreateTransaction).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/transaction", controllers.GetTransactions).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/transaction", controllers.UpdateTransaction).Methods(http.MethodOptions, http.MethodPut)
	r.HandleFunc("/api/transaction/{transactionId}", controllers.DeleteTransaction).Methods(http.MethodOptions, http.MethodDelete)
	r.HandleFunc("/api/budget", controllers.CreateBudget).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/budget", controllers.GetBudgets).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/deleted_budgets", controllers.GetDeletedBudgets).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/budget", controllers.UpdateBudget).Methods(http.MethodOptions, http.MethodPut)
	r.HandleFunc("/api/budget/categories", controllers.GetBudgetCategories).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/api/budget/{budgetId}", controllers.DeleteBudget).Methods(http.MethodOptions, http.MethodDelete)

	r.HandleFunc("/api/progress", controllers.CreateProgress).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/api/progress", controllers.GetProgress).Methods(http.MethodOptions, http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", corsObj.Handler(r)))
}
