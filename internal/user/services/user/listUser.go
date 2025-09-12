package userService

import (
	"context"
	"log/slog"

	userModel "github.com/fbriansyah/go-modular/internal/model/user"
)

func (s *UserService) ListUser(ctx context.Context, query *userModel.ListUserQuery) ([]*userModel.User, error) {
	slog.Info("ListUser", "query", query)

	filter := &userModel.User{
		FirstName: query.FirstName,
		LastName:  query.LastName,
		Email:     query.Email,
	}

	users, err := s.userRepository.List(ctx, filter, query.Limit, query.Offset)
	if err != nil {
		slog.Error("ListUser", "query", query, "error", err)
		return nil, err
	}

	return users, nil
}
