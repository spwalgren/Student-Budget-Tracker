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

	"budget-tracker/controllers"
	"budget-tracker/database"
	"budget-tracker/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var Router = mux.NewRouter()

func TestMain(m *testing.M) {

	database.Initialize("budget_tracker_test")
	Router.HandleFunc("/api/login", controllers.LoginHandler).Methods(http.MethodOptions, http.MethodPost)
	Router.HandleFunc("/api/users", controllers.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	Router.HandleFunc("/api/signup", controllers.CreateUser).Methods(http.MethodOptions, http.MethodPost)
	Router.HandleFunc("/api/user", controllers.GetUser).Methods(http.MethodGet, http.MethodOptions)
	Router.HandleFunc("/api/logout", controllers.LogoutHandler).Methods(http.MethodPost, http.MethodOptions)
	Router.HandleFunc("/api/transaction", controllers.CreateTransaction).Methods(http.MethodPost, http.MethodOptions)
	Router.HandleFunc("/api/transaction", controllers.GetTransactions).Methods(http.MethodGet, http.MethodOptions)
	Router.HandleFunc("/api/transaction", controllers.UpdateTransaction).Methods(http.MethodOptions, http.MethodPut)
	Router.HandleFunc("/api/transaction/{userId}/{transactionId}", controllers.DeleteTransaction).Methods(http.MethodOptions, http.MethodDelete)

    code := m.Run()
    clearUserTable()
	clearTransactionTable()
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
func TestGetTransaction(t *testing.T) {
	clearTransactionTable()

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

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category2", "description": "test-description2"}}`)

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
    actual := []models.Transaction{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := []models.Transaction{}
	expected = append(expected, models.Transaction{
		UserID: 1, 
		TransactionID: 1,
		Amount:      100,
		Name:        "test-name",
		Date:        "test-date",
		Category:    "test-category",
		Description: "test-description",})
	expected = append(expected, models.Transaction{
		UserID: 1,
		TransactionID: 2,
		Amount:      200,
		Name:        "test-name2",
		Date:        "test-date2",
		Category:    "test-category2",
		Description: "test-description2",})
	a.Equal(expected, actual)
}
func TestUpdateTransaction_OK(t *testing.T) {
	clearTransactionTable()

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

	cookie := response.Result().Cookies()[0]

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
    actual := []models.Transaction{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := []models.Transaction{}
	expected = append(expected, models.Transaction{
		UserID: 1, 
		TransactionID: 1,
		Amount:      300,
		Name:        "test-name-updated",
		Date:        "test-date-updated",
		Category:    "test-category-updated",
		Description: "test-description-updated",})
	expected = append(expected, models.Transaction{
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

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category2", "description": "test-description2"}}`)

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

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category2", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("DELETE", "/api/transaction/1/1", bytes.NewBuffer(payload))
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
    actual := []models.Transaction{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := []models.Transaction{}
	expected = append(expected, models.Transaction{
		UserID: 1,
		TransactionID: 2,
		Amount:      200,
		Name:        "test-name2",
		Date:        "test-date2",
		Category:    "test-category2",
		Description: "test-description2",})
	a.Equal(expected, actual)
}
func TestDeleteTransaction_WrongTransactionID(t *testing.T) {
	clearTransactionTable()

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

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category2", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("DELETE", "/api/transaction/1/5", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response =executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, response.Code, "HTTP request status code error")
}
func TestDeleteTransaction_WrongUserID(t *testing.T) {
	clearTransactionTable()

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

	cookie := response.Result().Cookies()[0]

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
  	executeRequest(req)

	payload = []byte(`{"data":{"amount": 200, "name": "test-name2", "date": "test-date2", "category": "test-category2", "description": "test-description2"}}`)

	req, _ = http.NewRequest("POST", "/api/transaction", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    executeRequest(req)

	req, _ = http.NewRequest("DELETE", "/api/transaction/3/1", bytes.NewBuffer(payload))
	req.AddCookie(cookie)
    response =executeRequest(req)

	a := assert.New(t)

	a.Equal(http.MethodDelete, req.Method, "HTTP request method error")
	a.Equal(http.StatusForbidden, response.Code, "HTTP request status code error")
}