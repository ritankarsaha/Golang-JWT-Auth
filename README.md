# Golang JWT Auth

## Overview

This project implements a JWT authentication system using Go and MongoDB. It features user registration, login, and data retrieval, employing JWTs for secure access management.

## Features

- User signup and login
- JWT-based authentication with access and refresh tokens
- User data management and validation
- Error handling and input validation

## Getting Started

### Prerequisites

- Go 1.18 or later
- MongoDB instance
- Required Go modules:
  - `github.com/gin-gonic/gin`
  - `github.com/go-playground/validator/v10`
  - `golang.org/x/crypto/bcrypt`
  - `github.com/golang-jwt/jwt`
  - `github.com/joho/godotenv`




### Project Structure:- 

```
|--controllers
| |--userController.go     # Contains handlers for user-related operations
|--database
| |--databaseConnection.go # Handles MongoDB connection 
|--helpers
| |--authHelper.go         # Contains helper functions for authentication
| |--tokenHelper.go        # Contains helper functions for JWT token management
|--middleware
| |--authMiddleware.go     # Middleware for JWT authentication
|--models
| |--userModel.go          # Defines the user model and schema
|--routes
| |--authRouter.go         # Routes for authentication-related endpoints
| |--userRouter.go         # Routes for user-related endpoints
|--.env                    
|--go.mod                  # Go module file
|--go.sum                  # Go module checksum file
|--main.go                 # Entry point for the application
|--README.md               # Project documentation

```

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/Golang-JWT-Auth.git
   cd Golang-JWT-Auth
   ```
2. **Install Dependencies**

   ```bash
   go mod tidy
   ```
3. **Setup Environment Variables**

   ```bash
   SECRET_KEY=your_secret_key_here
   MONGODB_URL=your_mongodb_url
   PORT=your_desired_port
   ```
4. **Configure MongoDB**
Ensure your MongoDB instance is running and update connection details in the database package if necessary.

5. **Run the Application**
    ```bash
    go run main.go
     ```

API Endpoints

   •   Signup

	•	Endpoint: POST /signup
	•	Description: Registers a new user.
	•	Request Body:

```
{
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "password": "yourpassword",
  "phone": "1234567890",
  "user_type": "USER"
}
```

	•	Response:

```
{
  "message": "User created successfully"
}
```

Login

	•	Endpoint: POST /login
	•	Description: Authenticates a user and returns JWT tokens.
	•	Request Body:

```
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

	•	Response:

    ```
    {
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "user_type": "USER",
    "token": "your_access_token",
    "refresh_token": "your_refresh_token"
    }
```

Get User

	•	Endpoint: GET /user/:user_id
	•	Description: Retrieves user data by user ID.
	•	Request Header: token: <access_token>



	•	Response:

    ```
    {
  "user": {
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "phone": "1234567890",
    "user_type": "USER",
    "created_at": "2024-08-02T12:00:00Z",
    "updated_at": "2024-08-02T12:00:00Z",
    "user_id": "your_user_id"
  }
}
```
Error Handling

Errors are managed with appropriate HTTP status codes and messages. Common errors include invalid requests, validation errors, and internal server issues.

Contributing

Contributions are welcome. Please submit issues or pull requests following the community guidelines.



  
   