# Work Completed in Sprint 3

**Main Points:**

- Added Budget Tracking functionality for users
  - Enabled user to track expired/deleted budgets

# Frontend Unit Tests

# Backend Unit Tests On Functions added in Sprint 3

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
