package userRepository

import (
	"context"
	"database/sql"

	userModel "github.com/fbriansyah/go-modular/internal/model/user"
	"github.com/fbriansyah/go-modular/pkg/database"
	"github.com/fbriansyah/go-modular/utils"
)

// Create implements userPort.UserRepository.
// Subtle: this method shadows the method (BaseRepository).Create of UserRepository.BaseRepository.
func (u *UserRepository) Create(ctx context.Context, entity *userModel.User) error {
	// Generate UUID if not provided
	if entity.ID == "" {
		entity.ID = utils.GenerateUUID()
	}

	query := `
		INSERT INTO users (id, email, password_hash, first_name, last_name, status, created_at, updated_at, version)
		VALUES (:id, :email, :password, :first_name, :last_name, :status, :created_at, :updated_at, :version)`

	err := u.BaseRepository.Create(ctx, entity, query)
	if err != nil {
		return u.handleError("Create", err)
	}

	return nil
}

// Delete implements userPort.UserRepository.
// Subtle: this method shadows the method (BaseRepository).Delete of UserRepository.BaseRepository.
func (u *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	err := u.BaseRepository.Delete(ctx, id, query)
	if err != nil {
		return u.handleError("Delete", err)
	}

	return nil
}

// Update implements userPort.UserRepository.
// Subtle: this method shadows the method (BaseRepository).Update of UserRepository.BaseRepository.
func (u *UserRepository) Update(ctx context.Context, entity *userModel.User) error {
	query := `
		UPDATE users 
		SET email = :email, 
		    password_hash = :password, 
		    first_name = :first_name, 
		    last_name = :last_name, 
		    status = :status, 
		    updated_at = :updated_at, 
		    version = :version
		WHERE id = :id AND version = :version - 1`

	tx := database.GetTxFromContext(ctx)
	var result sql.Result
	var err error

	if tx != nil {
		result, err = tx.NamedExecContext(ctx, query, entity)
	} else {
		result, err = u.GetDB().NamedExecContext(ctx, query, entity)
	}

	if err != nil {
		return u.handleError("Update", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return u.handleError("Update", err)
	}

	if rowsAffected == 0 {
		// Check if user exists to distinguish between not found and version conflict
		exists, existsErr := u.Exists(ctx, entity.ID)
		if existsErr != nil {
			return u.handleError("Update", existsErr)
		}
		if !exists {
			return database.ErrNotFound
		}
		return database.ErrOptimisticLock
	}

	return nil
}
