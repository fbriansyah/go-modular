package userRepository

import (
	"database/sql"
	"strings"

	"github.com/fbriansyah/go-modular/pkg/database"
	"github.com/lib/pq"
)

// handleError converts database errors to appropriate domain errors
func (r *UserRepository) handleError(operation string, err error) error {
	if err == nil {
		return nil
	}

	// Handle PostgreSQL specific errors
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23505": // unique_violation
			if strings.Contains(pqErr.Detail, "email") {
				return database.NewDatabaseErrorWithCode(operation, "users", "23505", "email already exists", database.ErrDuplicateKey)
			}
			return database.NewDatabaseErrorWithCode(operation, "users", "23505", "duplicate key violation", database.ErrDuplicateKey)
		case "23503": // foreign_key_violation
			return database.NewDatabaseErrorWithCode(operation, "users", "23503", "foreign key constraint violation", database.ErrForeignKeyViolation)
		case "23514": // check_violation
			return database.NewDatabaseErrorWithCode(operation, "users", "23514", "check constraint violation", database.ErrInvalidInput)
		}
	}

	// Handle common database errors
	if err == sql.ErrNoRows {
		return database.ErrNotFound
	}

	// Wrap other errors
	return database.NewDatabaseError(operation, "users", err)
}
