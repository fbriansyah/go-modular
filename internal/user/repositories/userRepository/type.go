package userRepository

import (
	userModel "github.com/fbriansyah/go-modular/internal/model/user"
	"github.com/fbriansyah/go-modular/pkg/database"
	userPort "github.com/fbriansyah/go-modular/ports/user"
)

type UserRepository struct {
	database.BaseRepository[userModel.User, string]
}

var _ userPort.UserRepository = (*UserRepository)(nil)
