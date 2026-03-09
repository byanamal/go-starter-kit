# Go Starter Kit 🚀

A clean, production-ready Go REST API starter kit built with performance, scalability, and developer experience in mind.

## ✨ Features

- **Standard Library First**: Built on top of Go's `net/http` (leveraging Go 1.22+ features like enhanced routing).
- **Clean Architecture**: Separation of concerns using Handlers, Services, Repositories, and DTOs.
- **Database Support**: PostgreSQL integration with `sqlx` for powerful and type-safe data access.
- **Authentication & RBAC**:
  - JWT-based Authentication.
  - Role-Based Access Control (RBAC) with granular Permission checks.
- **Developer Experience (DX)**:
  - **Pretty Logging**: Colored console output for requests and errors (INFO, WARN, ERROR, DEBUG).
  - **Automated Seeding**: Ready-to-use seeders for Users, Roles, and Permissions.
  - **Migration System**: Simple SQL-based migrations.
- **Lean & Fast**: No heavy frameworks, zero unnecessary dependencies.

## 🏗 Project Structure

```text
├── cmd/
│   ├── api/            # API entry point
│   ├── migrate/        # Database migration runner
│   └── seeder/         # Data seeder runner
├── internal/
│   ├── dto/            # Data Transfer Objects
│   ├── handler/        # HTTP Handlers & Routers
│   ├── model/          # Database entities
│   ├── repository/     # Database access layer
│   ├── service/        # Business logic layer
│   ├── pkg/
│   │   ├── config/     # Configuration management
│   │   ├── db/         # Database connection
│   │   ├── helper/     # Shared utilities
│   │   ├── logger/     # Pretty colored logger
│   │   └── middleware/ # Auth, RBAC, Logger, CORS
│   └── server/         # Server initialization
├── migrations/         # SQL migration files
└── Makefile            # Common development commands
```

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.22 or higher
- [PostgreSQL](https://www.postgresql.org/)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/byanamal/base-api.git
   cd base-api
   ```

2. Setup environment variables:

   ```bash
   cp .env.example .env
   ```

   _Edit `.env` and fill in your database credentials._

3. Download dependencies:

   ```bash
   make tidy
   ```

4. Run migrations:

   ```bash
   make migrate-up
   ```

5. Seed initial data (optional):
   ```bash
   make seed
   ```

### Running the Server

Start the API server:

```bash
make run
```

The server will be available at `http://localhost:8080`.

## 🛠 Available Commands

| Command                         | Description                                       |
| ------------------------------- | ------------------------------------------------- |
| `make run`                      | Starts the API server                             |
| `make build`                    | Builds the binary to `bin/api`                    |
| `make migrate-create name=name` | Creates a new migration file                      |
| `make migrate-up`               | Runs all pending database migrations              |
| `make migrate-down`             | Reverts the last database migration               |
| `make seed`                     | Runs database seeders (Roles, Permissions, Users) |
| `make tidy`                     | Cleans up `go.mod` and `go.sum`                   |
| `make docker-up`                | Starts services using Docker Compose              |

## 🔒 Security

- **JWT Auth**: Tokens are required for protected endpoints.
- **RBAC**: Use `@Permission("action:resource")` to guard your routes.

## 📝 License

This project is licensed under the MIT License.
