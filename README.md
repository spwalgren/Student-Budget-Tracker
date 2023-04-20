
# Student-Budget-Tracker

A mobile app that helps college students track their expenses and income, set budgets, and get insights into their spending habits.

Frontend: User interface and interaction

Backend: Handling financial data and providing analytics

Features: Budget planner, savings calculator, expense tracker

Members:

- Emily Jiji (Front-end)
- Grayson Kornberg (Back-end)
- Brian Magnuson (Front-end)
- Spencer Walgren (Back-end)

## Setting up the database

Our app uses a MySQL database for the backend. You must use a cloud hosting service (AWS, Azure, etc.) to create the database.

Once your database is set up, navigate to 

``
Student-Budget-Tracker/backend/database/setup.go
``

On line 13, replace models.Username, models.Password, and models.Url with the username, password, and url respectively for your database:

``
var  dsn  =  fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",models.Username,models.Password,models.Url, dbname)
``

Finally, navigate to 

``
Student-Budget-Tracker/backend/controllers/user.go
``

On lines 115 and 174, replace models.SecretKey with any string of your choice.

## Running the app

To run the app, make sure Go and Node.js are both installed.

Install all necessary npm packages
```
npm install
```

Install necessary Go modules
```
cd backend
go mod tidy
```

Start the app (cd back to the root if you haven't already)
```
npm start
```

## Frontend Documentation

View our frontend documentation on our wiki:
[Student Budget Tracker Wiki](https://github.com/spwalgren/Student-Budget-Tracker/wiki)
