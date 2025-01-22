package user

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (s *serv) Create(ctx context.Context, dto user.CreateDTO) (user.User, error) {
	// check user exists

	return s.userRepository.Create(ctx, dto)
}
