package userRepository

import (
	"fmt"
	"strings"

	userModel "github.com/fbriansyah/go-modular/internal/model/user"
)

// buildListQuery constructs the SQL query for listing users with filters
func (r *UserRepository) buildListQuery(filter *userModel.User, limit, offset int) (string, []interface{}) {
	query := `
		SELECT id, email, password_hash as password, first_name, last_name, status, created_at, updated_at, version
		FROM users`

	whereClause, args := r.buildWhereClause(filter)
	if whereClause != "" {
		query += " WHERE " + whereClause
	}

	query += " ORDER BY created_at DESC"

	// Add pagination
	argIndex := len(args) + 1
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)

	return query, args
}

// buildCountQuery constructs the SQL query for counting users with filters
func (r *UserRepository) buildCountQuery(filter *userModel.User) (string, []interface{}) {
	query := "SELECT COUNT(*) FROM users"

	whereClause, args := r.buildWhereClause(filter)
	if whereClause != "" {
		query += " WHERE " + whereClause
	}

	return query, args
}

// buildWhereClause constructs the WHERE clause for filtering
func (r *UserRepository) buildWhereClause(filter *userModel.User) (string, []interface{}) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	if filter.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIndex))
		args = append(args, string(filter.Status))
		argIndex++
	}

	if filter.Email != "" {
		conditions = append(conditions, fmt.Sprintf("email ILIKE $%d", argIndex))
		args = append(args, "%"+filter.Email+"%")
		argIndex++
	}

	if filter.FirstName != "" {
		conditions = append(conditions, fmt.Sprintf("first_name ILIKE $%d", argIndex))
		args = append(args, "%"+filter.FirstName+"%")
		argIndex++
	}

	if filter.LastName != "" {
		conditions = append(conditions, fmt.Sprintf("last_name ILIKE $%d", argIndex))
		args = append(args, "%"+filter.LastName+"%")
		argIndex++
	}

	return strings.Join(conditions, " AND "), args
}
