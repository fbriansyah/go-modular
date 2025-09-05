package database

import (
	"context"
	"database/sql"
)

// Repository defines the base interface for all repositories
type Repository[T any, ID comparable] interface {
	// Create inserts a new entity and returns the created entity with ID
	Create(ctx context.Context, entity *T) error

	// GetByID retrieves an entity by its ID
	GetByID(ctx context.Context, id ID) (*T, error)

	// Update updates an existing entity
	Update(ctx context.Context, entity *T) error

	// Delete removes an entity by its ID
	Delete(ctx context.Context, id ID) error

	// List retrieves entities with pagination
	List(ctx context.Context, limit, offset int) ([]*T, error)

	// Count returns the total number of entities
	Count(ctx context.Context) (int64, error)

	// Exists checks if an entity exists by ID
	Exists(ctx context.Context, id ID) (bool, error)
}

// BaseRepository provides common repository functionality using SQLx
type BaseRepository[T any, ID comparable] struct {
	db        *DB
	tableName string
	idColumn  string
}

// NewBaseRepository creates a new base repository
func NewBaseRepository[T any, ID comparable](db *DB, tableName, idColumn string) *BaseRepository[T, ID] {
	return &BaseRepository[T, ID]{
		db:        db,
		tableName: tableName,
		idColumn:  idColumn,
	}
}

// Create inserts a new entity using the provided query
func (r *BaseRepository[T, ID]) Create(ctx context.Context, entity *T, query string, args ...interface{}) error {
	tx := GetTxFromContext(ctx)
	if tx != nil {
		_, err := tx.NamedExecContext(ctx, query, entity)
		return err
	}

	_, err := r.db.NamedExecContext(ctx, query, entity)
	return err
}

// GetByID retrieves an entity by ID using the provided query
func (r *BaseRepository[T, ID]) GetByID(ctx context.Context, id ID, query string) (*T, error) {
	var entity T

	tx := GetTxFromContext(ctx)
	if tx != nil {
		err := tx.GetContext(ctx, &entity, query, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, ErrNotFound
			}
			return nil, err
		}
		return &entity, nil
	}

	err := r.db.GetContext(ctx, &entity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &entity, nil
}

// Update updates an entity using the provided query
func (r *BaseRepository[T, ID]) Update(ctx context.Context, entity *T, query string, args ...interface{}) error {
	tx := GetTxFromContext(ctx)
	if tx != nil {
		result, err := tx.NamedExecContext(ctx, query, entity)
		if err != nil {
			return err
		}
		return r.checkRowsAffected(result)
	}

	result, err := r.db.NamedExecContext(ctx, query, entity)
	if err != nil {
		return err
	}
	return r.checkRowsAffected(result)
}

// Delete removes an entity by ID using the provided query
func (r *BaseRepository[T, ID]) Delete(ctx context.Context, id ID, query string) error {
	tx := GetTxFromContext(ctx)
	if tx != nil {
		result, err := tx.ExecContext(ctx, query, id)
		if err != nil {
			return err
		}
		return r.checkRowsAffected(result)
	}

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return r.checkRowsAffected(result)
}

// List retrieves entities with pagination using the provided query
func (r *BaseRepository[T, ID]) List(ctx context.Context, query string, limit, offset int) ([]*T, error) {
	var entities []*T

	tx := GetTxFromContext(ctx)
	if tx != nil {
		err := tx.SelectContext(ctx, &entities, query, limit, offset)
		return entities, err
	}

	err := r.db.SelectContext(ctx, &entities, query, limit, offset)
	return entities, err
}

// Count returns the total number of entities using the provided query
func (r *BaseRepository[T, ID]) Count(ctx context.Context, query string) (int64, error) {
	var count int64

	tx := GetTxFromContext(ctx)
	if tx != nil {
		err := tx.GetContext(ctx, &count, query)
		return count, err
	}

	err := r.db.GetContext(ctx, &count, query)
	return count, err
}

// Exists checks if an entity exists by ID using the provided query
func (r *BaseRepository[T, ID]) Exists(ctx context.Context, id ID, query string) (bool, error) {
	var exists bool

	tx := GetTxFromContext(ctx)
	if tx != nil {
		err := tx.GetContext(ctx, &exists, query, id)
		return exists, err
	}

	err := r.db.GetContext(ctx, &exists, query, id)
	return exists, err
}

// ExecuteInTransaction executes a function within a database transaction
func (r *BaseRepository[T, ID]) ExecuteInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return ExecuteInTransaction(ctx, r.db, fn)
}

// GetDB returns the underlying database connection
func (r *BaseRepository[T, ID]) GetDB() *DB {
	return r.db
}

// GetTableName returns the table name for this repository
func (r *BaseRepository[T, ID]) GetTableName() string {
	return r.tableName
}

// GetIDColumn returns the ID column name for this repository
func (r *BaseRepository[T, ID]) GetIDColumn() string {
	return r.idColumn
}

// checkRowsAffected checks if any rows were affected by the operation
func (r *BaseRepository[T, ID]) checkRowsAffected(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
