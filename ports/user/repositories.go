package userPort

import (
	"context"

	userModel "github.com/fbriansyah/go-modular/internal/model/user"
	"github.com/fbriansyah/go-modular/pkg/database"
)

type UserRepository interface {
	database.Repository[userModel.User, string]
	GetByEmail(ctx context.Context, email string) (*userModel.User, error)
}
