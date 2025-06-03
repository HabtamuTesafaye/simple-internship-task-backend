# Internship Backend

A Go-based backend service for internship management with user API endpoints and PostgreSQL database integration.

## Project Structure

The project follows a clean architecture approach:

```
internship_backend/
├── cmd/                # Application entry points
│   └── server/         # Main server application
├── config/             # Configuration management
├── internal/           # Private application code
│   ├── api/            # API handlers (auth, user, routes)
│   ├── db/             # Database connection helpers
│   ├── model/          # Data models (User, etc.)
│   ├── repository/     # Data access layer (UserRepository)
│   ├── service/        # Business logic (UserService)
│   ├── middleware/     # HTTP middleware (JWT auth)
│   └── util/           # Utility functions (JWT, hashing, context)
├── migrations/         # SQL schema and migrations
├── .env                # Environment variables (not committed)
├── .gitignore          # Git ignore rules
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

---

## 📌 Folder & File Overview

### `cmd/server/main.go`
- **Purpose**: Main entrypoint of the backend.
- **Responsibilities**:
  - Loads environment/config values.
  - Connects to PostgreSQL.
  - Sets up router and middleware.
  - Starts the HTTP server.

### `config/`
- **Purpose**: Centralized configuration logic.
- **Responsibilities**:
  - Loads environment variables (via `.env`).
  - Validates config values (e.g., `DATABASE_URL`).
  - Makes config values accessible throughout the app.

### `internal/api/`
- **Purpose**: HTTP route handlers.
- **Responsibilities**:
  - Parse/validate incoming requests.
  - Call service layer functions.
  - Return appropriate JSON HTTP responses.
  - **Files**:
    - `auth.go`: Handles login and JWT issuance.
    - `user.go`: Handles user CRUD endpoints.
    - `routes.go`: Registers public and protected routes.

### `internal/service/`
- **Purpose**: Application’s core business logic.
- **Responsibilities**:
  - Implements main use case flows.
  - Coordinates repository logic.
  - Handles password hashing and user creation.
  - **Files**:
    - `user.go`: User-related business logic.

### `internal/repository/`
- **Purpose**: Database access layer.
- **Responsibilities**:
  - Run SQL queries against PostgreSQL.
  - Isolate raw DB logic from business code.
  - Expose clean functions like `GetUserByEmail(email)`.
  - **Files**:
    - `user.go`: UserRepository implementation.

### `internal/model/`
- **Purpose**: Go structs for domain entities.
- **Responsibilities**:
  - Define data models (`User`, etc.).
  - Map database rows to Go structs.
  - **Files**:
    - `user.go`: User struct.

### `internal/middleware/`
- **Purpose**: HTTP middleware components.
- **Responsibilities**:
  - Add JWT-based authentication.
  - Attach context values like current user claims.
  - **Files**:
    - `auth.go`: JWT authentication middleware.

### `internal/util/`
- **Purpose**: Reusable utility functions.
- **Responsibilities**:
  - Handle JWT creation/parsing, password hashing, and context helpers.
  - **Files**:
    - `jwt.go`: JWT generation and parsing.
    - `encrypt.go`: Password hashing and verification.
    - `ContextWithClaims.go`: Context helpers for JWT claims.

### `internal/db/`
- **Purpose**: Database connection handling.
- **Responsibilities**:
  - Load and pool PostgreSQL connection.
  - **Files**:
    - `postgres.go`: Connects to PostgreSQL.

### `migrations/`
- **Purpose**: SQL schema and migrations.
- **Responsibilities**:
  - Define and update database schema.
  - **Files**:
    - `schema.sql`: Table definitions for users, assignments, appointments.

### `.env`
- **Purpose**: Environment variables for local development.

### `.gitignore`
- **Purpose**: Ignore sensitive files like `.env`.

### `go.mod`
- **Purpose**: Go module and dependency management.

---

## 🧠 System Workflow

### 1. 🔐 User Login
- Users authenticate via `/api/v1/login`.
- JWT is issued on successful login.

### 2. 👤 User Management
- CRUD endpoints for users under `/api/v1/users` (protected by JWT).

### 3. 🗄️ Database
- PostgreSQL stores users, assignments, and appointments.

---

## 📦 Tech Stack

- **Golang 1.23**
- **PostgreSQL**
- **JWT Authentication**
- **Clean Architecture Principles**

---

## 🚀 Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/minab/internship-backend.git
cd internship-backend
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Set up environment variables
- Copy `.env.example` to `.env` (if available) and update values.

### 4. Run the server
```bash
go run cmd/server/main.go
```