# Project Overview

This project is a Go application with a modular architecture. It appears to be a web service or API, given the use of libraries like `sqlx` for database interaction and `pq` as a PostgreSQL driver. The project is structured with a clear separation of concerns, with distinct packages for configuration, database management, and different business logic modules (like `auth` and `user`).

The project's dependencies, as defined in `go.mod`, include:
- `github.com/golang-migrate/migrate/v4` for database migrations.
- `github.com/google/uuid` for generating UUIDs.
- `github.com/jmoiron/sqlx` for database interactions.
- `github.com/lib/pq` as the PostgreSQL database driver.
- `golang.org/x/crypto` for cryptographic operations.

# Building and Running

**TODO:** The entry points for the application, `cmd/http/main.go` and `cmd/migration/main.go`, are currently empty. To make the project runnable, you will need to implement the main application logic in these files.

A typical `main.go` for the HTTP server might look like this:

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fbriansyah/go-modular/config"
	"github.com/fbriansyah/go-modular/pkg/database"
)

func main() {
	// Load configuration
	dbConfig := &config.DatabaseConfig{
		URL:      os.Getenv("DATABASE_URL"),
		// Populate other fields from environment variables or a config file
	}

	// Initialize database manager
	dbManager, err := database.NewManager(dbConfig, "./migrations")
	if err != nil {
		log.Fatalf("Failed to initialize database manager: %v", err)
	}
	defer dbManager.Close()

	// Initialize database
	if err := dbManager.Initialize(context.Background(), true); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// TODO: Initialize your HTTP server and routes here

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// TODO: Start your server in a goroutine

	<-stop
	log.Println("Shutting down server...")

	// TODO: Add server shutdown logic
}
```

# Development Conventions

- **Modular Architecture:** The project follows a modular design, with features like `auth` and `user` separated into their own modules under the `internal` directory. Each module should be self-contained, with its own handlers, services, and repositories.
- **Database Migrations:** Database schema changes should be managed through migration files, likely located in a `migrations` directory at the project root. The `golang-migrate` library is used to apply these migrations.
- **Configuration:** Application configuration is handled in the `config` package. Database configuration is defined in `config/config.go`. Sensitive information should be managed separately, potentially in `config/secret.go` and loaded from environment variables.
- **Database Interaction:** The `pkg/database` package provides a robust database manager that handles connections, health checks, and migrations. Use this manager for all database operations.
