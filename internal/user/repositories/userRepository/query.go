package userRepository

import (
	"context"

	userModel "github.com/fbriansyah/go-modular/internal/model/user"
	"github.com/fbriansyah/go-modular/pkg/database"
)

// Count implements userPort.UserRepository.
// Subtle: this method shadows the method (BaseRepository).Count of UserRepository.BaseRepository.
func (u *UserRepository) Count(ctx context.Context, filter *userModel.User) (int64, error) {
	query, args := u.buildCountQuery(filter)

	var count int64
	tx := database.GetTxFromContext(ctx)

	var err error
	if tx != nil {
		err = tx.GetContext(ctx, &count, query, args...)
	} else {
		err = u.GetDB().GetContext(ctx, &count, query, args...)
	}

	if err != nil {
		return 0, u.handleError("Count", err)
	}

	return count, nil
}

// GetByID implements userPort.UserRepository.
// Subtle: this method shadows the method (BaseRepository).GetByID of UserRepository.BaseRepository.
func (u *UserRepository) GetByID(ctx context.Context, id string) (*userModel.User, error) {
	query := `
		SELECT id, email, password_hash as password, first_name, last_name, status, created_at, updated_at, version
		FROM users 
		WHERE id = $1`

	user, err := u.BaseRepository.GetByID(ctx, id, query)
	if err != nil {
		return nil, u.handleError("GetByID", err)
	}

	return user, nil
}

// List implements userPort.UserRepository.
// Subtle: this method shadows the method (BaseRepository).List of UserRepository.BaseRepository.
func (u *UserRepository) List(ctx context.Context, filter *userModel.User, limit int, offset int) ([]*userModel.User, error) {
	query, args := u.buildListQuery(filter, limit, offset)

	var users []*userModel.User
	tx := database.GetTxFromContext(ctx)

	var err error
	if tx != nil {
		err = tx.SelectContext(ctx, &users, query, args...)
	} else {
		err = u.GetDB().SelectContext(ctx, &users, query, args...)
	}

	if err != nil {
		return nil, u.handleError("List", err)
	}

	return users, nil
}

// Exists implements userPort.UserRepository.
// Subtle: this method shadows the method (BaseRepository).Exists of UserRepository.BaseRepository.
func (u *UserRepository) Exists(ctx context.Context, id string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`

	var exists bool
	tx := database.GetTxFromContext(ctx)

	var err error
	if tx != nil {
		err = tx.GetContext(ctx, &exists, query, id)
	} else {
		err = u.GetDB().GetContext(ctx, &exists, query, id)
	}

	if err != nil {
		return false, u.handleError("Exists", err)
	}

	return exists, nil
}
