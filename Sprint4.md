# Work Completed in Sprint 4

**Main Points:**
- Added a Progress and Calendar System so the user can have a visual tracking system
- Updated the current Budget System so that a Transaction can only be added to an active Budget.

# Frontend Unit Tests

For unit testing, we opted to use Cypress.
The code for writing component tests in Cypress is less complex than Jasmine tests and is easier to understand.
Additionally, it offers useful tools such as running the tests in a browser and reviewing how Cypress sees the individual components.
Seeing that the Cypress component tests accomplish the same tasks as the Jasmine tests, we believed it to be an acceptable substitute.

## New/Changed for Sprint 4

**components/dash-budgets** - Tests if it can add tables, add to existing tables, delete data from tables, delete tables, and edit table entries.
Changed to be compatible with the current version that uses an API to calculate period information.

**components/dash-transactions** - Tests if it has a table, has a button that opens a modal, can display entries from the mock service, has a detail row that pops out, can add transactions, can edit transactions, and can delete transactions.
Changed to test the component more thoroughly with add, edit, and delete functions.

**components/event-card** - Tests how it displays information in each of its states, including "On Track", "Upcoming", "Overspent", and "Completed".

**components/dash-calendar** - Tests if it can display the current month and can navigate between months.

## Unchanged from Sprint 3

**components/alert** - Tests how it displays colors, text, and icons.

**components/dash-home** - Tests if it can mount

**components/dash-settings** - Tests if the delete user button behaves as it should and calls a function.

**components/page-not-found** - Tests if it has a button and if that button calls a function when clicked.

**routes/landing** - Tests if it can mount

**routes/dashboard** - Tests if it has visible buttons, it calls a function immediately, and it calls a function upon clicking log out.

**routes/login** - Tests if it automatically attempts to log in and tests various inputs to see what will permit the user to log in.

**routes/signup** - Tests various inputs to see what will allow the user to log in.

# Frontend Cypress E2E Tests

## New/Changed for Sprint 4

**dashboard**

- Tests that it opens when the user is logged in.
- Tests if the buttons link to different parts of the app.
- Tests if the 404 page works.
- Tests the logout button and checks that the user cannot easily log back in.
- Changed to account for new button for progress page.

**transactions**

- Tests if the user can access the transaction page and transaction modal.
- Tests if the user can add data to the table.
- Tests if the user will still have their data after logging out and logging back in.
- Tests if old data does not appear when the user recreates their account after deletion.
- Changed to account for the new method for selecting transaction categories.

**budgets**

- Tests if the user can access the budget page and budget modal.
- Tests if the user can add data to the table.
- Tests if the user will still have their data after logging out and logging back in.
- Tests if old data does not appear when the user recreates their account after deletion.
- Changed to now test how adding budget categories influences the category options in transactions.
- Now tests how renaming/deleting budget categories influence existing transactions.

**calendar**

- Tests if the user can access the calendar page.
- Tests if the user can add budgets and have them appear as calendar events.
- Tests if the user can add transactions and have them count toward calendar events.

## Unchanged from Sprint 3

**login**

- Runs several tests to see what permits the user to log in.
- Logs the user in and tests if the user is directed to the dashboard.
- Tests if going to the login page when already logged in will redirect the user.
- Tests user registration and user deletion.

## Backend Unit Tests on Functions Added in Sprint 4
**
TestGetCyclePeriod()
TestGetEventsCurrMonth()
TestGetEventsNextMonth()
TestGetProgress()
TestGetPreviousProgress()
**
## Documentation on Functions Added in Sprint 4
### backend/budgets.go
#### GetCyclePeriod()
Retrieves the cycle period for a given date and budget associated with a user from the database and sends it in the HTTP response body as a JSON object. It returns the cycle period, including start and end dates, for a given budget, including weekly, monthly, and yearly budgets.

**Parameters:**
- w http.ResponseWriter: an HTTP response writer used to send a response back to the client.
- r *http.Request: a pointer to an HTTP request object that contains the client's request information.
**Behavior:**
- Sets the Content-Type of the response header to "*".
- If the request method is OPTIONS, sets the HTTP status code to 200 and returns.
- Extracts the user ID from the request header by calling the ReturnUserID() function.
- If the user ID is -1, sets the HTTP status code to 401 (Unauthorized) and returns.
- Extracts the date and budget ID from the URL parameters using the Gorilla mux package.
- Retrieves the budget associated with the user ID and budget ID from the database by calling the Find() function on the corresponding models.
- Determines the cycle index and cycle start and end dates based on the budget frequency, duration, and start date.
- Creates a Cycle object for the given budget, including its index, start date, end date, and budget ID.
- Encodes the cycle period information as a JSON object and sends it in the response body using the http.ResponseWriter object.
- Sets the HTTP status code to 200.

#### diff()
The diff function takes two time.Time objects, a and b, and returns the difference between them as integers representing the number of years, months, days, hours, minutes, and seconds.

**Parameters:**
- a time.Time: the first time object to compare.
- b time.Time: the second time object to compare.
**Returns:**
- year int: the number of years between the two time objects.
- month int: the number of months between the two time objects.
- day int: the number of days between the two time objects.
- hour int: the number of hours between the two time objects.
- min int: the number of minutes between the two time objects.
- sec int: the number of seconds between the two time objects.
**Behavior:**
- If the locations of the two time objects are different, sets the location of the second time object to the location of the first time object.
- If the first time object is after the second time object, swaps the two time objects.
- Extracts the year, month, and day components of each time object.
- Extracts the hour, minute, and second components of each time object.
- Calculates the difference between the corresponding components of the two time objects.
- Normalizes negative values by adjusting the higher-level components accordingly.
- Returns the calculated differences as integers.

#### GetBudgetCategories()
Retrieves a list of unique budget categories associated with a user from the database and sends it in the HTTP response body as a JSON object.

**Parameters:**
- w http.ResponseWriter: an HTTP response writer used to send a response back to the client.
- r *http.Request: a pointer to an HTTP request object that contains the client's request information.

**Behavior:**
- Sets the Content-Type of the response header to "*".
- If the request method is OPTIONS, sets the HTTP status code to 200 and returns.
- Extracts the user ID from the request header by calling the ReturnUserID() function.
- If the user ID is -1, sets the HTTP status code to 401 (Unauthorized) and returns.
- Retrieves all budgets associated with the user ID from the database by calling the Find() function on the BudgetsResponse model.
- Creates a list of unique budget categories by iterating through each budget and adding its category to a map. If the category is not already in the map, it is added to the list of categories.
- Encodes the list of categories as a JSON object and sends it in the response body using the http.ResponseWriter object.
- Sets the HTTP status code to 200.

### backend/calendar.go
#### GetEvents()
Handler function that processes GET requests to retrieve events and budgets from the database for a particular month.

**Parameters**
- w http.ResponseWriter - an interface that allows the server to write the response headers and body for the client.
- r *http.Request - a pointer to a struct that represents the client's HTTP request.

**Function Behavior**
- Sets the response content type to *
- If the request method is "OPTIONS", writes an HTTP status of 200 and returns.
- Retrieves the month parameter from the request URL using mux.Vars().
- Calculates the current time and the first and last day of the month for the month specified in the request.
- Retrieves the user ID from the request header using the ReturnUserID() function.
- If the user ID is -1 (meaning the user is not authorized), sets the HTTP status to 401 and returns.
- Retrieves the budgets from the database that match the user ID and have not been deleted.
- Iterates through each budget retrieved, and for each budget:
  - Parses the start date of the budget and sets the start time to midnight of the same day.
  - If the budget ends before the selected month, continues to the next budget.
  - If the budget starts after the selected month, continues to the next budget.
  - Calculates the range of cycles for the budget that fall within the selected month.
  - For each cycle in the range:
    - If the budget frequency is weekly, calculates the start and end dates of the cycle.
    - If the budget frequency is monthly, calculates the start and end months of the cycle.
    - If the budget frequency is yearly, calculates the start and end years of the cycle.
    - Retrieves the events from the database that fall within the cycle, and adds them to the eventsResponse object.
- Writes the eventsResponse object to the response body.

### backend/progress.go
#### GetProgress()
**Description:**
- Retrieves progress information for all budgets associated with a user from the database and sends it in the HTTP response body as a JSON object. It returns the progress of all budgets, including weekly, monthly, and yearly budgets. It also includes the transactions that have been made on each budget.

**Parameters:**
- w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.
- r *http.Request - a pointer to a data structure that represents the client HTTP request.

**Behavior:**  
- Set the "Content-Type" header of the HTTP response to "\*".  
- Check if the HTTP request method is "OPTIONS". If yes, it sets the HTTP status code of the response to 200 (OK) and returns from the function.  
- Extract the user ID from the HTTP request by calling the ReturnUserID() function and converts it to an integer using the ParseInt function. If the user ID is -1, it sets the HTTP status code of the response to 401 (Unauthorized) and returns from the function.  
- Define variables for the response structs: progResponse, weeklyProgResponse, monthlyProgResponse, yearlyProgResponse, weeklyBudgets, monthlyBudgets, yearlyBudgets.
- Retrieve all weekly budgets from the database where user_id matches with the extracted user ID, frequency is "weekly", and isDeleted flag is false. Then, create a new progress entry for each budget. To do this, it retrieves transactions for the corresponding category using IsInBudget function and saves them in the progress struct along with the budget information. The weekly progress entries are added to weeklyProgResponse.  
- Retrieve all monthly budgets from the database where user_id matches with the extracted user ID, frequency is "monthly", and isDeleted flag is false. Then, create a new progress entry for each budget by following the same steps used for weekly budgets. The monthly progress entries are added to monthlyProgResponse.  
- Retrieve all yearly budgets from the database where user_id matches with the extracted user ID, frequency is "yearly", and isDeleted flag is false. Then, create a new progress entry for each budget by following the same steps used for weekly budgets. The yearly progress entries are added to yearlyProgResponse.  
- Concatenate the progress entries for weekly, monthly, and yearly budgets into a single array called progResponse.  
- Encode the progress response in JSON format and send it in the HTTP response body.  
- Set the HTTP status code of the response to 200 (OK).  

#### GetPreviousProgress()
Retrieves progress information for all budgets associated with a user from a previous period from the database and sends it in the HTTP response body as a JSON object. It returns the progress of all budgets, including weekly, monthly, and yearly budgets. It also includes the transactions that have been made on each budget.

**Parameters:**

- w http.ResponseWriter: an HTTP response writer used to send a response back to the client.
- r *http.Request: a pointer to an HTTP request object that contains the client's request information.

**Behavior:**
- Sets the Content-Type of the response header to "*".
- If the request method is OPTIONS, sets the HTTP status code to 200 and returns.
- Extracts the user ID from the request header by calling the ReturnUserID() function.
- If the user ID is -1, sets the HTTP status code to 401 (Unauthorized) and returns.
- Retrieves all weekly, monthly, and yearly budgets associated with the user ID from the database by calling the Find() function on the corresponding models.
- For each weekly, monthly, and yearly budget, retrieves all transactions that match the budget category by calling the Find() function on the TransactionsResponse model.
- Calls the IsInPreviousBudget() function to check if the transactions belong to a previous budget cycle.
- Creates a Progress object for each budget, including its budget ID, category, budget goal, transaction ID list, and total spent amount.
- Encodes the progress information as a JSON object and sends it in the response body using the http.ResponseWriter object.
- Sets the HTTP status code to 200.

#### IsInPreviousBudget() 
Retrieves all transactions that occurred within the previous budget cycle. It takes in three parameters - an array of transactions, a budget object, and an HTTP request. The function sends two HTTP requests to the backend to retrieve the start and end dates of the previous budget cycle. It then compares the date of each transaction in the array with the start and end dates of the previous budget cycle and returns an array of transactions that occurred within that cycle.

**Parameters:**
- transactions []models.Transaction: an array of transactions to filter.
- budget models.Budget: a budget object used to retrieve the budget cycle dates.
- r *http.Request: an HTTP request object used to send requests to the backend.

**Return Values:**
- models.TransactionsResponse: a struct containing an array of transactions that occurred within the previous budget cycle.
- error: an error object indicating whether there was an error sending HTTP requests or decoding JSON responses.
**Function Steps:**
- Send an HTTP GET request to the backend to retrieve the start and end dates of the current budget cycle using the budget object passed in as a parameter.
- Send a second HTTP GET request to retrieve the start and end dates of the previous budget cycle by subtracting one day from the start date of the current cycle.
- Decode the JSON response from both requests into structs.
- Loop through each transaction in the array of transactions.
- Parse the date of the transaction as a time.Time object.
- Parse the start and end dates of the previous budget cycle as time.Time objects.
- Compare the transaction date with the start and end dates of the previous budget cycle.
- If the transaction date is within the previous budget cycle, append it to a new array of transactions.
- Return the new array of transactions as a models.TransactionsResponse struct.

#### IsInBudget() 
Function that takes a slice of transactions, a budget model, and a pointer to an http.Request object as inputs. It returns a TransactionsResponse model and an error. The function filters the transactions to only include those that fall within the current budget cycle.

**Parameters:**
- transactions: A slice of Transaction models
- budget: A Budget model
- r: A pointer to an http.Request object
**Returns:**
- TransactionsResponse: A model containing a slice of Transaction models that fall within the current budget cycle
- error: An error, if any occurred during the function's execution

**Function Behavior**
- Set up a backend request using the budget model and the current date.
- Set a cookie for the backend request to retrieve the start and end dates of the budget cycle.
- Send the backend request using the http.DefaultClient and handle any errors that occur during the process.
- Decode the response body to get the start and end dates of the budget cycle.
- Iterate through the transactions slice.
- For each transaction, parse the date and check if it falls within the current budget cycle using the start and end dates.
- If the transaction falls within the current budget cycle, append it to the TransactionsResponse model.
- Return the TransactionsResponse model and nil error.

#### HelperGetStartEndDate()
Handler function that takes in an http.ResponseWriter and an http.Request as its parameters. It retrieves the budget ID and date parameters from the request URL, and uses them to calculate the start and end dates of the budget cycle. It then returns the start and end dates in JSON format in the response body.

**Parameters:**
- w http.ResponseWriter: The response writer used to write HTTP response headers and body.
- r *http.Request: The HTTP request received by the server.

**Errors:**
This function returns an HTTP error status code if there is an error retrieving the budget ID and date parameters, or if the budget ID is invalid.

**Function Behavior**
- Set the Content-Type header of the response to "*".
- Check if the HTTP method of the request is OPTIONS. If it is, set the status code of the response to 200 and return.
- Parse the user ID from the request using the ReturnUserID() function and convert it to an int64.
- If the user ID is -1, set the status code of the response to 401 (Unauthorized) and return.
- Retrieve the budget ID and date parameters from the request using the mux.Vars() function.
- Convert the budget ID parameter to an int and the date parameter to a time.Time object.
- Retrieve the budget model from the database using the user ID and budget ID. If the query fails, set the status code of the response to 400 (Bad Request) and return.
- Calculate the start and end dates of the budget cycle based on the budget model and the provided date.
- Create a Cycle model containing the cycle index, start date, and end date.
- Set the status code of the response to 200 (OK) and encode the Cycle model to JSON in the response body.

#### DateEqual()
Function that takes two time.Time objects as input and returns a boolean value indicating if they represent the same date.

**Parameters:**
- time1: A time.Time object representing a date
- time2: A time.Time object representing a date
**Returns:**
- A boolean value indicating whether the dates represented by time1 and time2 are equal.
Function Behavior:
- Compare the year component of time1 and time2. If they are not equal, return false.
- Compare the month component of time1 and time2. If they are not equal, return false.
- Compare the day component of time1 and time2. If they are not equal, return false.
- Return true.

## Backend Functions Updated
### backend/budgets.go
#### CreateBudget()
- Functionality added to incorporate a transaction index and cycle index to better track which cycle of the budget it is currently on
#### UpdateBudget()
- Functionality added to update the category, cycle index, and transaction index added to the budget
### backend/transactions.go
#### CreateTransaction()
- Functionality added to create a budget cycle tracking system along with the created transaction
#### UpdateTransaction()
- Functionality added to update the cycle period associated with a transaction.
