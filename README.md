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
│   ├── api/            # API handlers
│   ├── db/             # Database connections
│   ├── model/          # Data models
│   ├── repository/     # Data access layer
│   └── service/        # Business logic
```


---

## 📌 Detailed Folder & File Overview

### `cmd/server/main.go`
- **Purpose**: Main entrypoint of the backend.
- **Responsibilities**:
  - Loads environment/config values.
  - Connects to PostgreSQL.
  - Sets up router and middleware.
  - Starts the HTTP server.

---

### `config/`
- **Purpose**: Centralized configuration logic.
- **Responsibilities**:
  - Loads environment variables (via `.env`).
  - Validates config values (e.g., `DATABASE_URL`).
  - Makes config values accessible throughout the app.

---

### `internal/api/`
- **Purpose**: HTTP route handlers.
- **Responsibilities**:
  - Parse/validate incoming requests.
  - Call service layer functions.
  - Return appropriate JSON HTTP responses.

> **Example**: `reading.go` handles user progress on reading assignments.

---

### `internal/service/`
- **Purpose**: Application’s core business logic.
- **Responsibilities**:
  - Implements main use case flows.
  - Coordinates repository and email logic.
  - Enforces rules like reading must be done before appointments.

---

### `internal/repository/`
- **Purpose**: Database access layer.
- **Responsibilities**:
  - Run SQL queries against PostgreSQL.
  - Isolate raw DB logic from business code.
  - Expose clean functions like `GetUserByEmail(email)`.

---

### `internal/model/`
- **Purpose**: Go structs for domain entities.
- **Responsibilities**:
  - Define data models (`User`, `Assignment`, `Appointment`, etc.).
  - Map database rows to Go structs.

---

### `internal/email/`
- **Purpose**: Email integration module.
- **Responsibilities**:
  - Send notification emails via SMTP or other providers.
  - Define templates and formatting.

---

### `internal/middleware/`
- **Purpose**: HTTP middleware components.
- **Responsibilities**:
  - Add JWT-based authentication.
  - Handle request logging.
  - Attach context values like current user ID.

---

### `internal/util/`
- **Purpose**: Reusable utility functions.
- **Responsibilities**:
  - Handle common logic like hashing, string manipulation, or timestamp formatting.

---

### `internal/db/`
- **Purpose**: Database connection handling.
- **Responsibilities**:
  - Load and pool PostgreSQL connection.
  - Handle migrations or setup.

---

### `scripts/`
- **Purpose**: Dev & DB helper scripts.
- **Examples**:
  - SQL file to create tables.
  - Seeder to populate initial data.

---

### `Dockerfile` & `docker-compose.yml`
- **Purpose**: Containerization and service orchestration.
- **Responsibilities**:
  - Package the Go app into a Docker image.
  - Spin up services like the backend + PostgreSQL using `docker-compose`.

---

## 🧠 System Workflow

### 1. 🔐 User Registration & Reading Progress
- New users register via `/register`.
- Reading topics are tracked via an API.
- Each topic completion is recorded.

### 2. 🧠 Business Logic
- Once **all reading topics** are completed:
  - User status updates automatically.
  - An email notification is triggered.

### 3. 💻 Practical Project
- After reading is complete, users may request the **practical project**.

### 4. 📅 Appointment Scheduling
- Users schedule an appointment to present the project.
- Appointments are tracked in the `appointments` table.

### 5. 📬 Notification Emails
- Automatic emails are sent on:
  - Registration success.
  - Reading completion.
  - Appointment confirmation.

---

## 📦 Tech Stack

- **Golang 1.22**
- **PostgreSQL**
- **Docker + Docker Compose**
- **JWT Authentication**
- **SMTP Email Notifications**
- **Clean Architecture Principles**

---

## 🚀 Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/minab/internship-backend.git
cd internship-backend
```
## Local Development Setup

1. Install dependencies:
```bash
go mod download
```

2. Set up environment variables:
   - Copy .env.example to .env
   - Update the values accordingly

3. Run the server:
```bash
go run cmd/server/main.go
```
