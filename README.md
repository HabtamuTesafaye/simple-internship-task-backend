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
├── docs/               # Swagger/OpenAPI documentation
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
  - Serves Swagger UI at `/swagger/` in development mode.

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
    - `change_password.go`: Handles password reset requests.
    - `routes.go`: Registers public and protected routes.

### `internal/service/`
- **Purpose**: Application’s core business logic.
- **Responsibilities**:
  - Implements main use case flows.
  - Coordinates repository logic.
  - Handles password hashing and user creation.
  - **Files**:
    - `user.go`: User-related business logic.
    - `change_password.go`: Password reset logic.

### `internal/repository/`
- **Purpose**: Database access layer.
- **Responsibilities**:
  - Run SQL queries against PostgreSQL.
  - Isolate raw DB logic from business code.
  - Expose clean functions like `GetUserByEmail(email)`.
  - **Files**:
    - `user.go`: UserRepository implementation.
    - `change_password.go`: PasswordResetRepository implementation.

### `internal/model/`
- **Purpose**: Go structs for domain entities.
- **Responsibilities**:
  - Define data models (`User`, etc.).
  - Map database rows to Go structs.
  - **Files**:
    - `user.go`: User struct.
    - `change_password.go`: PasswordResetToken struct.

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
    - `context_with_claims.go`: Context helpers for JWT claims.
    - `email.go`: Email sending utility.

### `internal/db/`
- **Purpose**: Database connection handling.
- **Responsibilities**:
  - Load and pool PostgreSQL connection.
  - **Files**:
    - `postgres.go`: Connects to PostgreSQL.

### `internal/templates/`
- **Purpose**: HTML templates for emails.
- **Files**:
    - `reset_password.html`: Password reset email template.

### `migrations/`
- **Purpose**: SQL schema and migrations.
- **Responsibilities**:
  - Define and update database schema.
  - **Files**:
    - `schema.sql`: Table definitions for users, assignments, appointments.

### `docs/`
- **Purpose**: API documentation (Swagger/OpenAPI).
- **Files**:
    - `swagger.yaml` / `swagger.json`: OpenAPI definitions.
    - `docs.go`: Auto-generated Swagger docs for Go.

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

### 3. 🔑 Password Reset
- Request password reset via `/api/v1/forgot-password`.
- Reset password via `/api/v1/reset-password` using token sent to email.

### 4. 🗄️ Database
- PostgreSQL stores users, assignments, and appointments.

---

## 📖 API Documentation (Swagger)

- **Swagger UI** is available at [`/swagger/`](http://localhost:4000/swagger/) when running in development mode.
- OpenAPI specs are defined in [`docs/swagger.yaml`](docs/swagger.yaml) and [`docs/swagger.json`](docs/swagger.json).
- Endpoints include:
  - `POST /api/v1/login` – User login, returns JWT.
  - `POST /api/v1/register` – Create a new user.
  - `GET /api/v1/users` – List all users (JWT required).
  - `GET /api/v1/users/{id}` – Get user by ID (JWT required).
  - `PUT /api/v1/users/update/{id}` – Update user (JWT required).
  - `POST /api/v1/forgot-password` – Request password reset.
  - `POST /api/v1/reset-password` – Reset password with token.

---

## 📦 Tech Stack

- **Golang 1.23**
- **PostgreSQL**
- **JWT Authentication**
- **Swagger/OpenAPI Documentation**
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

### 5. View API docs
- Open [http://localhost:4000/swagger/](http://localhost:4000/swagger/) in your browser.

---

## 📝 Notes

- All protected endpoints require a valid JWT in the `Authorization: Bearer <token>` header.
- Password reset emails use the template in [`internal/templates/reset_password.html`](internal/templates/reset_password.html).
- See [`docs/swagger.yaml`](docs/swagger.yaml) for full endpoint specs and