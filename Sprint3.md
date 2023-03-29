# Work Completed in Sprint 3

**Main Points:**

- Added Budget Tracking functionality for users
  - Users can add, view, edit, and delete their budgets
  - Budget tables show detailed information about...
    - How they are defined
    - Whether or not they repeat a finite number of times
    - What the timespan is for the current period
    - How much time is left in the current period
  - Enabled user to track expired budgets
- Added Button to allow user to delete their user information on the database

# Frontend Unit Tests

For unit testing, we opted to use Cypress.
The code for writing component tests in Cypress is less complex than Jasmine and is easier to understand.
Additionally, it offers useful tools such as running the tests in a browser and reviewing how Cypress sees the individual components.
Seeing that the Cypress component tests accomplish the same tasks as the Jasmine tests, we believed it to be an acceptable substitute.

## New/Changed for Sprint 3

**components/dash-budgets** - Tests if it can add tables, add to existing tables, delete data from tables, delete tables, and edit table entries.
Note: This test seemingly passes in the Cypress GUI but fails in the Cypress CLI. This issue is planned to be addressed in a future sprint.

**components/dash-settings** - Tests if the delete user button behaves as it should and calls a function.

## Unchanged from Sprint 2

**components/alert** - Tests how it displays colors, text, and icons.

**components/dash-transactions** - Tests if it has a table, has a button that opens a modal, can display entries from the mock service, and has a detail row that pops out.

**components/page-not-found** - Tests if it has a button and if that button calls a function when clicked.

**components/dash-home** - Tests if it can mount

**routes/landing** - Tests if it can mount

**routes/dashboard** - Tests if it has visible buttons, if it calls a function immediately, and if it calls a function upon clicking log out.

**routes/login** - Tests if it automatically attempts to login the user and tests various inputs to see what will permit the user to log in.

**routes/signup** - Tests various inputs to see what will allow the user to login

# Frontend Cypress E2E Tests

## New/Changed for Sprint 3

**budgets**

- Tests if the user can access the budget page and budget modal.
- Tests if the user can add data to the table.
- Tests if the user will still have their data after logging out and logging back in.
- Tests if old data does not appear when the user recreates their account after deletion.

**transactions**

- Tests if the user can access the transaction page and transaction modal.
- Tests if the user can add data to the table.
- Tests if the user will still have their data after logging out and logging back in.
- Tests if old data does not appear when the user recreates their account after deletion.

**login**

- Runs several tests to see what permits the user to log in.
- Logs the user in and tests if the user is directed to the dashboard.
- Tests if going to the login page when already logged in will redirect the user.
- Tests user registration and user deletion.

## Unchanged from Sprint 2

**dashboard**

- Tests that it opens when the user is logged in.
- Tests if the buttons link to different parts of the app.
- Tests if the 404 page works.
- Tests the logout button and checks that the user cannot easily log back in.

# Backend Unit Tests On Functions added in Sprint 3

**TestCreateBudget()  
TestGetBudget()  
TestUpdateBudget_OK()  
TestUpdateBudget_WrongBudgetID()  
TestDeleteBudget_OK()  
TestDeleteBudget_WrongBudgetID()
TestGetBudgetCategories()**

# Documentation On Functions added in Sprint 3

## controllers/user.go

### DeleteUser()

Deletes the user and all the associated transactions and budgets for that user.

**Parameters**  
w (http.ResponseWriter): An interface that allows the server to construct an HTTP response  
r (http.Request): A data structure that represents the client's HTTP request

**Behavior**  
Sets the content type of the response to "\*"  
Checks if the HTTP request method is "OPTIONS". If it is, the function sends an HTTP response with status code 200 and returns.  
Retrieves the user ID of the user making the request.  
If the user ID is "-1" (i.e., the user is not authenticated), the function sends an HTTP response with status code 401 (Unauthorized) and returns.  
Retrieves the user information from the database.  
Deletes the user information from the database.  
Deletes all the transactions associated with the user.  
Deletes all the budgets associated with the user.  
Deletes the jwt token cookies to sign the user out.
Sends an HTTP response with status code 200.

## controllers/budgets.go

### CreateBudget()

Creates a new budget for a user in the database based on the request body and sends a JSON response with the newly created budget's details.

**Parameters**  
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.  
r \*http.Request - a pointer to a data structure that represents the client HTTP request.

**Function Behavior**  
It sets the "Content-Type" header of the HTTP response to "\*".  
If the HTTP request method is "OPTIONS", it sets the HTTP status code of the response to 200 (OK) and returns from the function.  
It extracts the user ID from the HTTP request by calling the ReturnUserID() function and converts it to an integer using the ParseInt() function. If the user ID is -1, it sets the HTTP status code of the response to 401 (Unauthorized) and returns from the function.  
It creates a new empty BudgetContent model and decodes the request body to populate it.  
It creates a new Budget model with the user ID, a budget ID of 0, a false value for the IsDeleted field, and the BudgetContent model from the request body.  
It saves the new Budget model to the database using the Create() function.  
It sets the HTTP status code of the response to 200 (OK).  
It encodes a CreateBudgetResponse model with the user ID and budget ID of the newly created budget as JSON and sends it in the HTTP response body.

### GetBudgets()

Retrieves all the budgets for a user from the database and sends them in the HTTP response body as a JSON object.

**Parameters**  
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.  
r \*http.Request - a pointer to a data structure that represents the client HTTP request.

**Function Behavior**  
It sets the "Content-Type" header of the HTTP response to "\*".  
If the HTTP request method is "OPTIONS", it sets the HTTP status code of the response to 200 (OK) and returns from the function.  
It extracts the user ID from the HTTP request by calling the ReturnUserID() function and converts it to an integer using the ParseInt() function. If the user ID is -1, it sets the HTTP status code of the response to 401 (Unauthorized) and returns from the function.  
It creates a new empty BudgetsResponse model.  
It retrieves all the budgets from the database that belong to the user with the extracted user ID and stores them in the Budgets field of the BudgetsResponse model using the Where() and Find() functions.  
It prints the BudgetsResponse model to the console for debugging purposes.  
It sets the HTTP status code of the response to 200 (OK).  
It encodes the BudgetsResponse model as JSON and sends it in the HTTP response body.

### GetDeletedBudgets()

Retrieves all the deleted budgets for a user from the database and sends them in the HTTP response body as a JSON object.

**Parameters**
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.  
r \*http.Request - a pointer to a data structure that represents the client HTTP request.

**Function Behavior**
It sets the "Content-Type" header of the HTTP response to "\*".  
If the HTTP request method is "OPTIONS", it sets the HTTP status code of the response to 200 (OK) and returns from the function.  
It extracts the user ID from the HTTP request by calling the ReturnUserID() function and converts it to an integer using the ParseInt() function. If the user ID is -1, it sets the HTTP status code of the response to 401 (Unauthorized) and returns from the function.  
It creates a new empty BudgetsResponse model.  
It retrieves all the deleted budgets from the database that belong to the user with the extracted user ID and stores them in the Budgets field of the BudgetsResponse model using the Where() and Find() functions.  
It sets the HTTP status code of the response to 200 (OK).  
It encodes the BudgetsResponse model as JSON and sends it in the HTTP response body.

### UpdateBudget()

Updates a budget with new data provided in the HTTP request body.

**Parameters**
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.  
r \*http.Request - a pointer to a data structure that represents the client HTTP request.

**Function Behavior**
It sets the "Content-Type" header of the HTTP response to "\*".  
If the HTTP request method is "OPTIONS", it sets the HTTP status code of the response to 200 (OK) and returns from the function.  
It creates an empty UpdateBudgetRequest model and decodes the new budget data from the HTTP request body using the json.NewDecoder() and Decode() functions.  
It extracts the user ID from the HTTP request by calling the ReturnUserID() function and converts it to an integer using the ParseInt() function. If the user ID is -1, it sets the HTTP status code of the response to 401 (Unauthorized) and returns from the function.  
It compares the extracted user ID with the user ID of the new budget data. If they are not the same, it sets the HTTP status code of the response to 403 (Forbidden) and returns from the function.  
It retrieves the old budget data from the database using the First() function of the DB object and stores it in an empty Budget model. If an error occurs, it sets the HTTP status code of the response to 400 (Bad Request) and returns from the function.  
It updates the old budget data with the new budget data and saves it to the database using the Save() function of the DB object.  
It sets the HTTP status code of the response to 200 (OK).

### DeleteBudget()

Deletes the specified budget or moves it to the deleted state by setting the isDeleted flag to true in the database.

**Parameters**
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.
r \*http.Request - a pointer to a data structure that represents the client HTTP request.

**Function Behavior**
The function sets the Content-Type header to "\*".  
If the HTTP request method is OPTIONS, the function sets the HTTP status code of the response to 200 (OK) and returns.  
The function retrieves the user ID from the HTTP request header and sets the HTTP status code of the response to 401 (Unauthorized) if the user ID is not valid.  
The function retrieves the budget ID from the URL parameters.  
The function retrieves the budget from the database that matches the user ID and budget ID.  
If the budget does not exist, the function sets the HTTP status code of the response to 400 (Bad Request) and returns.  
If the budget is not already deleted, the function sets the isDeleted flag to true and saves the changes to the database. Otherwise, the function deletes the budget from the database.

### GetBudgetCategories()

Gets a list of all of the budget categories of the logged in user.

**Parameters**
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.
r \*http.Request - a pointer to a data structure that represents the client HTTP request.

**Function Behavior**
The function sets the Content-Type header to "\*".  
If the HTTP request method is OPTIONS, the function sets the HTTP status code of the response to 200 (OK) and returns.  
The function retrieves the user ID from the HTTP request header and sets the HTTP status code of the response to 401 (Unauthorized) if the user ID is not valid.  
The function retrieves the budget ID from the URL parameters.  
The function retrieves all budgets from the database that matches the user ID.  
The function returns a list of all of the categories from each retrieved budget.
