# Documentation
## controllers/user.go
### GetUsers()
**Description**  
This is a handler function that retrieves all the user information from the database and sends it in the HTTP response body as a JSON array.

**Parameters**  
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.  
r *http.Request - a pointer to a data structure that represents the client HTTP request.  
  
**Function Behavior**  
It initializes an empty slice of UserInfo models.  
It retrieves all the user information from the database and stores it in the users slice.  
It sets the HTTP status code of the response to 200 (OK).  
It encodes the users slice as a JSON object and sends it in the HTTP response body.  
  
### CreateUser()  
**Description**  
This is a handler function that creates a new user record in the database. It retrieves the user details from the request body, checks if the email is already registered in the database, encrypts the password, and saves the new user record to the database.  

**Parameters**  
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.  
r *http.Request - a pointer to a data structure that represents the client HTTP request.  

**Function Behavior**   
If the HTTP request method is OPTIONS, it sets the HTTP status code of the response to 200 (OK) and returns.  
It initializes an empty slice of UserInfo models.  
It retrieves all the user information from the database and stores it in the users slice.  
It decodes the request body into a new instance of the UserInfo model.  
It checks if the email is already registered in the database. If the email is already registered, it sets the HTTP status code of the response to 409 (Conflict) and returns.  
It encrypts the password using bcrypt.  
It saves the new user record to the database.  
It sets the HTTP status code of the response to 200 (OK).  

### LogoutHandler() 
**Description**
This is a handler function that deletes the JWT token cookie from the client browser to logout the user.

**Parameters**  
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.  
r *http.Request - a pointer to a data structure that represents the client HTTP request.  
  
**Function Behavior**
It creates a new HTTP cookie with the name "jtw" and an expiration time of 24 hours ago.  
It sets the domain, path, HTTP-only, and same-site attributes of the cookie.  
It sets the HTTP status code of the response to 200 (OK).  

### LoginHandler()
**Description**
This is a handler function that checks the user's email and password, and creates a new JWT token for the authenticated user.

**Parameters**  
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.  
r *http.Request - a pointer to a data structure that represents the client HTTP request.  
  
**Function Flow**
If the HTTP request method is OPTIONS, it sets the HTTP status code of the response to 200 (OK) and returns.  
It decodes the request body into a new instance of the UserLoginInfo model.  
It searches the database for a user record with the same email as the one provided in the request.  
If no user record is found, it sets the HTTP status code of the response to 404 (Not Found) and returns.    
  
## controllers/transactions.go
### CreateTransaction()
This is a handler function that creates a new transaction record in the database. It retrieves the transaction details from the request body, and associates the transaction with the current user. The user information is obtained from the JWT token present in the request cookie.
 
**Parameters:**  
w http.ResponseWriter - constructs the response.   
r *http.Request - pointer to the request.   

**Function Behavior**   
It creates a new instance of the Transaction model.  
It decodes the request body into the newTransaction object.  
It retrieves the JWT token from the request cookie.  
It parses the JWT token to extract the user information.  
It retrieves the user record from the database using the user information obtained from the JWT token.  
It sets the user ID for the new transaction record.  
It creates the new transaction record in the database.  
It encodes the new transaction record as a JSON object and sends it in the HTTP response body.  
  
### GetTransactions()
This is a handler function that retrieves all transaction records associated with the current user. The user information is obtained from the JWT token present in the request cookie.

**Parameters**  
w http.ResponseWriter - constructs the response.  
r *http.Request - pointer to the request.  

**Function Flow**   
It retrieves the JWT token from the request cookie.  
It parses the JWT token to extract the user information.  
It retrieves the user record from the database using the user information obtained from the JWT token.  
It retrieves all transaction records associated with the user from the database.  
It encodes the transaction records as a JSON array and sends it in the HTTP response body.  


### UpdateTransaction()
This is a handler function that updates a transaction in the database with the new information provided in the request body. It requires a valid JSON Web Token (JWT) in the HTTP request header to identify the user who is authorized to update the transaction.

**Parameters**  
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.
r *http.Request - a pointer to a data structure that represents the client HTTP request.

**Function Behavior**
It extracts the transaction information to be updated from the request body and stores it in an updateTransaction variable of type models.Transaction.
It extracts the JWT from the HTTP request cookie and parses it to obtain the user ID from the token's standard claims.
It retrieves the user information from the database using the user ID from the token.
It compares the user ID in the updateTransaction with the user ID from the token to verify that the requesting user is authorized to update the transaction.
It retrieves the transaction information to be updated from the database using the transaction ID in the updateTransaction.
It updates the transaction information in the expenses variable with the information from the updateTransaction variable.
It saves the updated transaction information in the database.
It sets the HTTP status code of the response to 200 (OK).

### DeleteTransaction()
This is a handler function that deletes a transaction from the database. It requires a valid JSON Web Token (JWT) in the HTTP request header to identify the user who is authorized to delete the transaction.

**Parameters**
w http.ResponseWriter - an interface that allows the handler to construct an HTTP response.
r *http.Request - a pointer to a data structure that represents the client HTTP request.

**Function Behavior**
It extracts the transaction information to be deleted from the request body and stores it in a toDelete variable of type models.Transaction.
It extracts the JWT from the HTTP request cookie and parses it to obtain the user ID from the token's standard claims.
It retrieves the user information from the database using the user ID from the token.
It compares the user ID in the toDelete with the user ID from the token to verify that the requesting user is authorized to delete the transaction.
It deletes the transaction information from the database that matches the user ID and transaction ID in the toDelete variable.
It sets the HTTP status code of the response to 200 (OK).
