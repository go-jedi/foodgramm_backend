package user

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (s *serv) GetByID(ctx context.Context, userID int64) (user.User, error) {
	s.logger.Debug("[get user by id] execute service")

	return s.userRepository.GetByID(ctx, userID)
}
