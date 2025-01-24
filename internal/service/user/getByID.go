package user

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (s *serv) GetByID(ctx context.Context, userID int64) (user.User, error) {
	return s.userRepository.GetByID(ctx, userID)
}
