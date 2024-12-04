# User CRUD API with Authentication

This project provides a **User CRUD API** built using **Golang**, **Gin**, **GORM**, and **JWT** (JSON Web Tokens) for secure authentication. It includes registration, login, and protected routes with role-based access control.

## Features
- **User Registration & Login** with JWT authentication
- **Role-based access** (Admin role seeded on server startup)
- **CRUD Operations**: Get, Update, Delete user data
- **GORM ORM** for interacting with relational databases
- **Custom JWT Middleware** for protected routes

## Technologies Used
- **Golang**
- **Gin**
- **GORM**
- **JWT** (JSON Web Tokens)
- **PostgreSQL** or **MySQL**

## Project Structure
```
├── controllers/ # API endpoint controllers 
├── middlewares/ # JWT authentication middleware
├── seed/ # Seed data for master admin creation
├── services/ # Business logic layer
├── dao/ # Database interactions
├── models/ # Data models (e.g., User)
├── utils/ # Utility functions (e.g., JWT generation)
├── config/ # Configuration files
├── main.go # Main entry point 
```

## Installation & Setup

### 1. Clone the Repository
```
git clone https://github.com/KoKhant02/user-crud-api.git
cd user-crud-api
```

### 2. Install Dependencies
```
go mod tidy
```

### 3. Run the Server
```
go run main.go
```
OR if docker is installed
```
docker compose up --build
```

### 5. API Endpoints
```
POST /api/register: Register a new user.
POST /api/login: Login and receive a JWT token.
GET /api/user/{id}: Get user details by ID (Protected route).
PUT /api/user/{id}: Update user details by ID (Protected route).
DELETE /api/user/{id}: Delete user by ID (Protected route).
```

Contributing
Feel free to fork the repository, open issues, and submit pull requests.
