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

## Documentation on Functions Added in Sprint 4
### backend/progress.go

#### GetProgress()
**Description:**
Retrieves progress information for all budgets associated with a user from the database and sends it in the HTTP response body as a JSON object. It returns the progress of all budgets, including weekly, monthly, and yearly budgets. It also includes the transactions that have been made on each budget.

**Parameters:**
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.
r *http.Request - a pointer to a data structure that represents the client HTTP request.

**Behavior:**
Set the "Content-Type" header of the HTTP response to "\*".  
Check if the HTTP request method is "OPTIONS". If yes, it sets the HTTP status code of the response to 200 (OK) and returns from the function.  
Extract the user ID from the HTTP request by calling the ReturnUserID() function and converts it to an integer using the ParseInt function. If the user ID is -1, it sets the HTTP status code of the response to 401 (Unauthorized) and returns from the function.  
Define variables for the response structs: progResponse, weeklyProgResponse, monthlyProgResponse, yearlyProgResponse, weeklyBudgets, monthlyBudgets, yearlyBudgets.
Retrieve all weekly budgets from the database where user_id matches with the extracted user ID, frequency is "weekly", and isDeleted flag is false. Then, create a new progress entry for each budget. To do this, it retrieves transactions for the corresponding category using IsInBudget function and saves them in the progress struct along with the budget information. The weekly progress entries are added to weeklyProgResponse.  
Retrieve all monthly budgets from the database where user_id matches with the extracted user ID, frequency is "monthly", and isDeleted flag is false. Then, create a new progress entry for each budget by following the same steps used for weekly budgets. The monthly progress entries are added to monthlyProgResponse.  
Retrieve all yearly budgets from the database where user_id matches with the extracted user ID, frequency is "yearly", and isDeleted flag is false. Then, create a new progress entry for each budget by following the same steps used for weekly budgets. The yearly progress entries are added to yearlyProgResponse.  
Concatenate the progress entries for weekly, monthly, and yearly budgets into a single array called progResponse.  
Encode the progress response in JSON format and send it in the HTTP response body.  
Set the HTTP status code of the response to 200 (OK).  
