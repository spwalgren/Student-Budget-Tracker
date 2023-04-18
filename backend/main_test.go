package main

// use command "go test -v" to run tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	// "time"

	"budget-tracker/controllers"
	"budget-tracker/database"
	"budget-tracker/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/crypto/bcrypt"
)

var Router = mux.NewRouter()

func TestMain(m *testing.M) {

	database.Initialize("budget_tracker_test")
	Router.HandleFunc("/api/login", controllers.LoginHandler).Methods(http.MethodOptions, http.MethodPost)
	Router.HandleFunc("/api/users", controllers.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	Router.HandleFunc("/api/signup", controllers.CreateUser).Methods(http.MethodOptions, http.MethodPost)
	Router.HandleFunc("/api/user", controllers.GetUser).Methods(http.MethodGet, http.MethodOptions)
	Router.HandleFunc("/api/delete-user", controllers.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	Router.HandleFunc("/api/logout", controllers.LogoutHandler).Methods(http.MethodPost, http.MethodOptions)
	Router.HandleFunc("/api/transaction", controllers.CreateTransaction).Methods(http.MethodPost, http.MethodOptions)
	Router.HandleFunc("/api/transaction", controllers.GetTransactions).Methods(http.MethodGet, http.MethodOptions)
	Router.HandleFunc("/api/transaction", controllers.UpdateTransaction).Methods(http.MethodOptions, http.MethodPut)
	Router.HandleFunc("/api/transaction/{transactionId}", controllers.DeleteTransaction).Methods(http.MethodOptions, http.MethodDelete)
	Router.HandleFunc("/api/budget", controllers.CreateBudget).Methods(http.MethodOptions, http.MethodPost)
	Router.HandleFunc("/api/budget", controllers.GetBudgets).Methods(http.MethodGet, http.MethodOptions)
	Router.HandleFunc("/api/deleted_budgets", controllers.GetDeletedBudgets).Methods(http.MethodGet, http.MethodOptions)
	Router.HandleFunc("/api/budget", controllers.UpdateBudget).Methods(http.MethodOptions, http.MethodPut)
	Router.HandleFunc("/api/budget/categories", controllers.GetBudgetCategories).Methods(http.MethodOptions, http.MethodGet)
	Router.HandleFunc("/api/budget/cycle/{date}", controllers.GetCyclePeriod).Methods(http.MethodOptions, http.MethodGet)
	Router.HandleFunc("/api/budget/{budgetId}", controllers.DeleteBudget).Methods(http.MethodOptions, http.MethodDelete)
	Router.HandleFunc("/api/progress", controllers.GetProgress).Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	Router.HandleFunc("/api/progress/previous", controllers.GetPreviousProgress).Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	Router.HandleFunc("/api/budget/dates/{budgetId}/{date}", controllers.HelperGetStartEndDate).Methods(http.MethodOptions, http.MethodGet)
	Router.HandleFunc("/api/calendar/{month}", controllers.GetEvents).Methods(http.MethodGet, http.MethodOptions)

    code := m.Run()
    clearUserTable()
	clearTransactionTable()
	clearBudgetTable()
	clearBudgetTransactionTable()
    os.Exit(code)
}

// --------------------- Functions for tests ---------------------

func clearUserTable() {
    database.DB.Exec("DELETE FROM user_infos")
    database.DB.Exec("ALTER TABLE user_infos AUTO_INCREMENT = 1")
}

func clearTransactionTable() {
    database.DB.Exec("DELETE FROM transactions")
	database.DB.Exec("ALTER TABLE transactions AUTO_INCREMENT = 1")
}

func clearBudgetTable() {
    database.DB.Exec("DELETE FROM budgets")
	database.DB.Exec("ALTER TABLE budgets AUTO_INCREMENT = 1")
}

func clearBudgetTransactionTable() {
    database.DB.Exec("DELETE FROM budget_transactions")
	database.DB.Exec("ALTER TABLE budget_transactions AUTO_INCREMENT = 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    Router.ServeHTTP(rr, req)
    return rr
}

// --------------------- Tests ---------------------

func TestEmptyUsersTable(t *testing.T) {
    clearUserTable()
    req, _ := http.NewRequest("GET", "/api/users", nil)
    response := executeRequest(req)
    a := assert.New(t)
    a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")
    body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.UserInfo{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
    clearUserTable()

	payload := []byte(`{"firstName": "test-firstName1",
    "lastName": "test-lastName1",
    "email": "test-email1",
    "password": "test-password1"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

	payload = []byte(`{"firstName": "test-firstName2",
    "lastName": "test-lastName2",
    "email": "test-email2",
    "password": "test-password2"}`)
    req, _ = http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    req, _ = http.NewRequest("GET", "/api/users", nil)
    response := executeRequest(req)
    a := assert.New(t)
    a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")
    body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := []models.UserInfo{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	expected := []models.UserInfo{}
	if err := bcrypt.CompareHashAndPassword([]byte(actual[0].Password), []byte("test-password1")); err == nil {
		actual[0].Password = "test-password1"
	}
	if err := bcrypt.CompareHashAndPassword([]byte(actual[1].Password), []byte("test-password2")); err == nil {
		actual[1].Password = "test-password2"
	}
	expected = append(expected, models.UserInfo{
		ID:        1,
		FirstName: "test-firstName1",
		LastName:  "test-lastName1",
		Email:     "test-email1",
		Password:  "test-password1",
	})
	expected = append(expected, models.UserInfo{
		ID:        2,
		FirstName: "test-firstName2",
		LastName:  "test-lastName2",
		Email:     "test-email2",
		Password:  "test-password2",
	})
	a.Equal(expected, actual)
}

func TestSignup_OK(t *testing.T) {
    clearUserTable()

    payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    response := executeRequest(req)

    a := assert.New(t)
    a.Equal(http.MethodPost, req.Method, "HTTP request method error")
    a.Equal(http.StatusOK, response.Code, "HTTP request status code error")
}

func TestSignUp_DuplicateEmail(t *testing.T) {
    clearUserTable()

    payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    response := executeRequest(req)

    a := assert.New(t)
    a.Equal(http.MethodPost, req.Method, "HTTP request method error")
    a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

    payload = []byte(`{"firstName": "test-firstName2",
    "lastName": "test-lastName2",
    "email": "test-email",
    "password": "test-password2"}`)
    req, _ = http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    response = executeRequest(req)

    a = assert.New(t)
    a.Equal(http.MethodPost, req.Method, "HTTP request method error")
    a.Equal(http.StatusConflict, response.Code, "HTTP request status code error")
}

func TestLogin_OK(t *testing.T) {
    clearUserTable()
    payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

    a := assert.New(t)
    a.NotEmpty(response.Result().Cookies())
    a.Equal(http.MethodPost, req.Method, "HTTP request method error")
    a.Equal(http.StatusOK, response.Code, "HTTP request status code error")
}

func TestLogin_EmailError(t *testing.T) {
    clearUserTable()
    payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "incorrect", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

    a := assert.New(t)
    a.Equal(http.MethodPost, req.Method, "HTTP request method error")
    a.Equal(http.StatusNotFound, response.Code, "HTTP request status code error")
}

func TestLogin_PasswordError(t *testing.T) {
    clearUserTable()
    payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "incorrect"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

    a := assert.New(t)
    a.Equal(http.MethodPost, req.Method, "HTTP request method error")
    a.Equal(http.StatusUnauthorized, response.Code, "HTTP request status code error")
}

func TestGetUser_LoggedOut(t *testing.T) {
    clearUserTable()
    req, _ := http.NewRequest("GET", "/api/user", nil)
	req.AddCookie(&http.Cookie{
		Name:       "jtw",
		Value:      "",
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	})
    response := executeRequest(req)

    a := assert.New(t)
    a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusUnauthorized, response.Code, "HTTP request status code error")
}

func TestGetUser_OK(t *testing.T) {
    clearUserTable()
    payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	req, _ = http.NewRequest("GET", "/api/user", nil)
	req.AddCookie(response.Result().Cookies()[0])
    response = executeRequest(req)

    a := assert.New(t)
    a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.UserReturnInfo{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.UserReturnInfo{
		ID:        1,
		FirstName: "test-firstName",
		LastName:  "test-lastName",
		Email:     "test-email",
	}
	a.Equal(expected, actual)
}

func TestLogout(t *testing.T) {
	clearUserTable()
	req, _ := http.NewRequest("POST", "/api/logout", nil)
    response := executeRequest(req)

	a := assert.New(t)

	a.NotEmpty(response.Result().Cookies())
	a.Empty(response.Result().Cookies()[0].Value)
	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")
}

func TestCreateTransaction(t *testing.T) {
	clearTransactionTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)
	cookie := response.Result().Cookies()[0]
	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"data":{"amount": 100, "name": "test-name", "date": "test-date", "category": "test-category", "description": "test-description"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.CreateTransactionResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.CreateTransactionResponse{
		UserID:      1,
		TransactionID: 1,
	}
	a.Equal(expected, actual)
}

func TestCreateTransactionMissingBudget(t *testing.T) {
	clearTransactionTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"data":{"amount": 100, "name": "test-name", "date": "test-date", "category": "test-category", "description": "test-description"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(response.Result().Cookies()[0])
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, response.Code, "HTTP request status code error")
}
func TestGetTransaction(t *testing.T) {
	clearTransactionTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)
	cookie := response.Result().Cookies()[0]

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"data":{"amount": 100, "name": "test-name", "date": "test-date", "category": "test-category", "description": "test-description"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("GET", "/api/transaction", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.TransactionsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.TransactionsResponse{}
	expected.Data = append(expected.Data, models.Transaction{
		UserID: 1, 
		TransactionID: 1,
		Amount:      100,
		Name:        "test-name",
		Date:        "test-date",
		Category:    "test-category",
		Description: "test-description",})
	expected.Data = append(expected.Data, models.Transaction{
		UserID: 1,
		TransactionID: 2,
		Amount:      200,
		Name:        "test-name2",
		Date:        "test-date2",
		Category:    "test-category",
		Description: "test-description2",})
	a.Equal(expected, actual)
}
func TestUpdateTransaction_OK(t *testing.T) {
	clearTransactionTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)
	cookie := response.Result().Cookies()[0]

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"category": "test-category-updated", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"data":{"amount": 100, "name": "test-name", "date": "test-date", "category": "test-category", "description": "test-description"}}`)


	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category2", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"data":{"userId":1, "transactionId": 1,"amount": 300, "name": "test-name-updated", "date": "test-date-updated", "category": "test-category-updated", "description": "test-description-updated"}}`)

	req, _ = http.NewRequest("PUT", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)
	a.Equal(http.MethodPut, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	req, _ = http.NewRequest("GET", "/api/transaction", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.TransactionsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.TransactionsResponse{}
	expected.Data = append(expected.Data, models.Transaction{
		UserID: 1, 
		TransactionID: 1,
		Amount:      300,
		Name:        "test-name-updated",
		Date:        "test-date-updated",
		Category:    "test-category-updated",
		Description: "test-description-updated",})
	expected.Data = append(expected.Data, models.Transaction{
		UserID: 1,
		TransactionID: 2,
		Amount:      200,
		Name:        "test-name2",
		Date:        "test-date2",
		Category:    "test-category2",
		Description: "test-description2",})
	a.Equal(expected, actual)
}
func TestUpdateTransaction_WrongTransactionID(t *testing.T) {
	clearTransactionTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)
	cookie := response.Result().Cookies()[0]

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"category": "test-category-updated", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)
	
	payload = []byte(`{"data":{"amount": 100, "name": "test-name", "date": "test-date", "category": "test-category", "description": "test-description"}}`)


	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"data":{"userId":1, "transactionId": 4,"amount": 300, "name": "test-name-updated", "date": "test-date-updated", "category": "test-category-updated", "description": "test-description-updated"}}`)

	req, _ = http.NewRequest("PUT", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)
	a.Equal(http.MethodPut, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, response.Code, "HTTP request status code error")
}

func TestDeleteTransaction_OK(t *testing.T) {
	clearTransactionTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)
	cookie := response.Result().Cookies()[0]

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"data":{"amount": 100, "name": "test-name", "date": "test-date", "category": "test-category", "description": "test-description"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("DELETE", "/api/transaction/1", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response =executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	req, _ = http.NewRequest("GET", "/api/transaction", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.TransactionsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.TransactionsResponse{}
	expected.Data = append(expected.Data, models.Transaction{
		UserID: 1,
		TransactionID: 2,
		Amount:      200,
		Name:        "test-name2",
		Date:        "test-date2",
		Category:    "test-category",
		Description: "test-description2",})
	a.Equal(expected, actual)
}
func TestDeleteTransaction_WrongTransactionID(t *testing.T) {
	clearTransactionTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)
	cookie := response.Result().Cookies()[0]

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T04:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"data":{"amount": 100, "name": "test-name", "date": "test-date", "category": "test-category", "description": "test-description"}}`)


	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("DELETE", "/api/transaction/5", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response =executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, response.Code, "HTTP request status code error")
}

func TestCreateBudget(t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "Weekly", "duration": 2, "count": 1, "startDate": "3/27/2023"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(response.Result().Cookies()[0])
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.CreateBudgetResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.CreateBudgetResponse{
		UserID:      1,
		BudgetID: 1,
	}
	a.Equal(expected, actual)
}
func TestGetBudget(t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "Weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T00:00:00.000Z"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "Weekly", "duration": 1, "count": 2, "startDate": "2023-03-27T00:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("GET", "/api/budget", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.BudgetsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.BudgetsResponse{}
	expected.Budgets = append(expected.Budgets, models.Budget{
		Data:      models.BudgetContent{
			Category:      "test-category",
			AmountLimit:   100,
			Frequency:     "Weekly",
			CycleDuration: 2,
			CycleCount:    1,
			StartDate:     "2023-03-27T00:00:00.000Z",
		},
		UserID:    1,
		BudgetID:  1,
		IsDeleted: false,
	})
	expected.Budgets = append(expected.Budgets, models.Budget{
		Data:      models.BudgetContent{
			Category:      "test-category2",
			AmountLimit:   100,
			Frequency:     "Weekly",
			CycleDuration: 1,
			CycleCount:    2,
			StartDate:     "2023-03-27T00:00:00.000Z",
		},
		UserID:    1,
		BudgetID:  2,
		IsDeleted: false,
	})
	a.Equal(expected, actual)
}
func TestUpdateBudget_OK(t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "Weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T00:00:00.000Z"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "Weekly", "duration": 1, "count": 2, "startDate": "2023-03-27T00:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"newBudget": {
		"userId": 1,
		"budgetId": 1,
		"isDeleted": false,
		"data": {
			"category": "test-category-updated",
			"amountLimit": 50.5,
			"frequency": "Weekly",
			"duration": 2,
			"count":1,
			"startDate": "2023-03-28T00:00:00.000Z"
		}
	}
 }`)

	req, _ = http.NewRequest("PUT", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)
	a.Equal(http.MethodPut, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	req, _ = http.NewRequest("GET", "/api/budget", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.BudgetsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.BudgetsResponse{}
	expected.Budgets = append(expected.Budgets, models.Budget{
		Data:      models.BudgetContent{
			Category:      "test-category-updated",
			AmountLimit:   50.5,
			Frequency:     "Weekly",
			CycleDuration: 2,
			CycleCount:    1,
			StartDate:     "2023-03-28T00:00:00.000Z",
		},
		UserID:    1,
		BudgetID:  1,
		IsDeleted: false,
	})
	expected.Budgets = append(expected.Budgets, models.Budget{
		Data:      models.BudgetContent{
			Category:      "test-category2",
			AmountLimit:   100,
			Frequency:     "Weekly",
			CycleDuration: 1,
			CycleCount:    2,
			StartDate:     "2023-03-27T00:00:00.000Z",
		},
		UserID:    1,
		BudgetID:  2,
		IsDeleted: false,
	})
	a.Equal(expected, actual)
}

func TestUpdateBudget_WrongBudgetID(t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "Weekly", "duration": 2, "count": 1, "startDate": "3/27/2023"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "Weekly", "duration": 1, "count": 2, "startDate": "3/27/2023"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"newBudget": {

		"userId": 1,
		"budgetId": 12,
		"isDeleted": false,
		"data": {
			"category": "test-category-updated",
			"amountLimit": 50.5,
			"frequency": "Weekly",
			"duration": 2,
			"count":1,
			"startDate:": "3/28/2023"
		}
	}
 }`)

	req, _ = http.NewRequest("PUT", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)
	a.Equal(http.MethodPut, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, response.Code, "HTTP request status code error")
}

func TestDeleteBudget_OK(t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "Weekly", "duration": 2, "count": 1, "startDate": "3/27/2023"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "Weekly", "duration": 1, "count": 2, "startDate": "3/27/2023"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("DELETE", "/api/budget/1", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	req, _ = http.NewRequest("GET", "/api/budget", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.BudgetsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.BudgetsResponse{}
	expected.Budgets = append(expected.Budgets, models.Budget{
		Data:      models.BudgetContent{
			Category:      "test-category2",
			AmountLimit:   100,
			Frequency:     "Weekly",
			CycleDuration: 1,
			CycleCount:    2,
			StartDate:     "3/27/2023",
		},
		UserID:    1,
		BudgetID:  2,
		IsDeleted: false,
	})
	a.Equal(expected, actual)
}

func TestDeleteBudget_WrongBudgetID(t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "Weekly", "duration": 2, "count": 1, "startDate": "3/27/2023"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "Weekly", "duration": 1, "count": 2, "startDate": "3/27/2023"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("DELETE", "/api/budget/5", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, response.Code, "HTTP request status code error")
}

func TestGetBudgetCategories (t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "Weekly", "duration": 2, "count": 1, "startDate": "3/27/2023"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "Weekly", "duration": 1, "count": 2, "startDate": "3/27/2023"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "Weekly", "duration": 1, "count": 2, "startDate": "3/27/2023"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("GET", "/api/budget/categories", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.BudgetCategoriesResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.BudgetCategoriesResponse{}
	expected.Category = append(expected.Category, "test-category")
	expected.Category = append(expected.Category, "test-category2")
	a.Equal(expected, actual)
}

func TestGetCyclePeriod (t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T00:00:00.000Z"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "weekly", "duration": 1, "count": 2, "startDate": "2023-03-27T00:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("GET", "/api/budget/cycle/2023-03-27", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.CyclePeriodResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.CyclePeriodResponse{}
	expected.Data = append(expected.Data, models.Cycle{Start: "2023-03-27T00:00:00Z",End: "2023-04-09T23:59:59Z",Index: 0,BudgetID: 1})
	expected.Data = append(expected.Data, models.Cycle{Start: "2023-03-27T00:00:00Z",End: "2023-04-02T23:59:59Z",Index: 0,BudgetID: 2})
	a.Equal(expected, actual)
}

func TestGetEventsCurrMonth (t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T00:00:00.000Z"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "weekly", "duration": 1, "count": 0, "startDate": "2023-03-29T00:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("GET", "/api/calendar/0", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.EventsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.EventsResponse{}
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 0,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-03-27T00:00:00Z", EndDate: "2023-04-09T23:59:59Z", Category: "test-category"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 1,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-03-29T00:00:00Z", EndDate: "2023-04-04T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 2,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-04-05T00:00:00Z", EndDate: "2023-04-11T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 3,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-04-12T00:00:00Z", EndDate: "2023-04-18T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 4,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-04-19T00:00:00Z", EndDate: "2023-04-25T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 5,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-04-26T00:00:00Z", EndDate: "2023-05-02T23:59:59Z", Category: "test-category2"}})
	a.Equal(expected, actual)
}

func TestGetEventsNextMonth (t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T00:00:00.000Z"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "weekly", "duration": 1, "count": 0, "startDate": "2023-03-29T00:00:00.000Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("GET", "/api/calendar/1", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.EventsResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.EventsResponse{}
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 0,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-04-26T00:00:00Z", EndDate: "2023-05-02T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 1,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-05-03T00:00:00Z", EndDate: "2023-05-09T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 2,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-05-10T00:00:00Z", EndDate: "2023-05-16T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 3,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-05-17T00:00:00Z", EndDate: "2023-05-23T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 4,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-05-24T00:00:00Z", EndDate: "2023-05-30T23:59:59Z", Category: "test-category2"}})
	expected.Events = append(expected.Events, models.Event{UserID: 1,EventID: 5,Data: models.EventContent{Frequency: "weekly",AmountLimit: 100,TotalSpent: 0,StartDate: "2023-05-31T00:00:00Z", EndDate: "2023-06-06T23:59:59Z", Category: "test-category2"}})
	a.Equal(expected, actual)
}

func TestGetProgress (t *testing.T) {
	clearBudgetTable()
	clearBudgetTransactionTable()

	payload := []byte(`{"firstName": "test-firstName",
    "lastName": "test-lastName",
    "email": "test-email",
    "password": "test-password"}`)
    req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
    executeRequest(req)

    payload = []byte(`{"email": "test-email", "password": "test-password"}`)
    req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
    response := executeRequest(req)

	payload = []byte(`{"category": "test-category", "amountLimit": 100, "frequency": "weekly", "duration": 2, "count": 1, "startDate": "2023-03-27T00:00:00Z"}`)

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"category": "test-category2", "amountLimit": 100, "frequency": "weekly", "duration": 1, "count": 0, "startDate": "2023-03-29T00:00:00Z"}`)

	req, _ = http.NewRequest("POST", "/api/budget", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	payload = []byte(`{"data":{"amount": 30, "name": "test-name", "date": "2023-03-31T00:00:00Z", "category": "test-category", "description": "test-description"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "2023-04-03T00:00:00Z", "category": "test-category2", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("GET", "/api/progress", nil)
	req.AddCookie(cookie)
    response = executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, response.Code, "HTTP request status code error")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.Error(err)
	}
    actual := models.GetProgressResponse{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.GetProgressResponse{}
	expected.Data = append(expected.Data, models.Progress{UserID: 1,BudgetID: 1,Category: "test-category",Frequency: "weekly",BudgetGoal: 100,TotalSpent: 30,TransactionIDList: []uint{1}})
	a.Equal(expected, actual)
}


